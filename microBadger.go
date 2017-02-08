package main

import (
	"errors"
	"flag"
	"fmt"
	website "github.com/allentechnology/website"
	"github.com/vharitonsky/iniflags"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type mbSlice []*microBadge

func (m mbSlice) TrimWhiteSpace(text string) string {
	newString := strings.Replace(text, " ", "-", -1)
	newString = strings.Replace(newString, "(", "", -1)
	newString = strings.Replace(newString, ")", "", -1)
	return newString
}

var (
	slotMap          = map[string]*slot{}
	categoryMap      = map[string]mbSlice{}
	microBadgeMap    = map[string]*microBadge{}
	tmpMicroBadgeMap = map[string]*microBadge{}
	client           *http.Client
)

var (
	loginReady = make(chan bool)
)

var (
	username = flag.String("username", "", "The boardgamegeek.com username used to log into the site")
	password = flag.String("password", "", "The boardgamegeek.com password associated with the given username")
	version  = flag.Bool("version", false, "Print the executable version to the screen")
	interval = flag.Int("interval", 1, "The interval between randomizations in minutes")
)

var (
	appDir    = ""
	runtimeOS = ""
)

func init() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	runtimeOS = runtime.GOOS
	switch runtime.GOOS {
	case "linux", "darwin":
		appDir = currentUser.HomeDir + ".microBadger"
	case "windows":
		appDir = currentUser.HomeDir + "microBadger"
	}
}

func main() {
	iniflags.Parse()
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	go webServer()

	var err error
	switch runtimeOS {
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:6060/").Start()
	case "darwin":
		err = exec.Command("open", "http://localhost:6060/").Start()
	case "windows":
		err = exec.Command("cmd", "/C", "start", "http://localhost:6060/").Start()
	default:
		err = fmt.Errorf("unsupported platform")
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("Failed to open browser. Navigate to http://localhost/ on your preferred web browser.")
	}
	fmt.Println("MicroBadger version ", VERSION)
	fmt.Println("To use microBadger, navigate to http://localhost:6060 in any web browser.")
	<-loginReady

	//	client = logIntoBGG()
	//	loggedIn := true
	for {
		// if !loggedIn {
		// 	client = logIntoBGG()
		// }

		fmt.Print(time.Now().Format("2006-01-02 15:04:05 "))
		fmt.Print("Attempting to randomize badges: ")
		err := getMicroBadges(client)
		if err != nil {
			fmt.Println("Failed")
			fmt.Println(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		randomizeBadges()
		time.Sleep(time.Duration(*interval) * time.Minute)
	}
}

func randomizeBadges() {
	badgeList := getRandomBadges()
	updateSuccess := make([]bool, len(badgeList))
	var err error
	for i, v := range badgeList {
		err = assignSlot(v.Id, fmt.Sprintf("%d", i+1), client)
		if err != nil {
			// fmt.Println("Error assigning slot ", i+1, ": ", err.Error())
			updateSuccess[i] = false
			//	loggedIn = false
		} else {
			updateSuccess[i] = true
		}

	}
	fmt.Println()
	fmt.Print("Slots ")
	slotUpdated := false
	for i, v := range updateSuccess {
		if v {
			fmt.Printf("%d ", i+1)
			slotUpdated = true
		}
	}
	if slotUpdated {
		fmt.Println("updated successfully")
	} else {
		fmt.Println("not updated")
	}

}

func webServer() {
	//Web server here
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/slot/", slotHandler)
	http.HandleFunc("/slotSubmit", slotSubmitHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/setInterval", setIntervalHandler)
	http.HandleFunc("/randomize", randomizeHandler)
	//	http.HandleFunc("/test", testHandler)
	serverErr := http.ListenAndServe(":6060", nil)

	if serverErr != nil {
		os.Exit(0)
	}

}

func randomizeHandler(w http.ResponseWriter, r *http.Request) {
	randomizeBadges()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in test handler")
	fmt.Fprintln(w, "in test handler")
}

func setIntervalHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intervalSlice := r.Form["interval"]

	if len(intervalSlice) > 0 {
		formInterval, err := strconv.Atoi(intervalSlice[0])
		if err != nil {
			return
		}
		*interval = formInterval
	}
	return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("webpage.html")

	if err != nil {
		fmt.Fprintf(w, "error: "+err.Error())
	}
	err = tmpl.Execute(w, categoryMap)
	if err != nil {
		fmt.Fprintf(w, "error: "+err.Error())
	}
}

func slotSubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formSlots := make(map[string][]string)
	formSlots["1"] = r.Form["slot1"]
	formSlots["2"] = r.Form["slot2"]
	formSlots["3"] = r.Form["slot3"]
	formSlots["4"] = r.Form["slot4"]
	formSlots["5"] = r.Form["slot5"]
	for i := 1; i < 6; i++ {
		slotID := fmt.Sprintf("%d", i)
		if s, ok := slotMap[slotID]; ok {
			s.AvailableBadges = map[string]*microBadge{}
			for _, v := range formSlots[slotID] {
				if mb, ok := microBadgeMap[v]; ok {
					s.AvailableBadges[v] = mb
				}
			}
		} else {
			newMap := make(map[string]*microBadge)
			slotMap[slotID] = &slot{Id: slotID, AvailableBadges: newMap}
			for _, v := range formSlots[slotID] {
				if mb, ok := microBadgeMap[v]; ok {
					slotMap[slotID].AvailableBadges[v] = mb
				}
			}

		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usernameSlice := r.Form["username"]
	passwordSlice := r.Form["password"]
	if len(usernameSlice) < 1 || len(passwordSlice) < 1 {
		//Print error to notification area
		return
	}
	*username = usernameSlice[0]
	*password = passwordSlice[0]

	for {
		var err error
		client, err = website.Login("https://boardgamegeek.com/login", *username, *password, 30*time.Second)

		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "Login failed" {
				fmt.Println("Exiting microBadger")
				os.Exit(1)
			}
		} else {
			break
		}
		time.Sleep(10 * time.Minute)

	}

	loginReady <- true
}

func slotHandler(w http.ResponseWriter, r *http.Request) {
	slotNumber := r.URL.Path[6:]
	if mb, ok := microBadgeMap[slotMap[string(slotNumber)].AssignedBadge]; ok {
		fmt.Fprintf(w, "<html><head><meta http-equiv='refresh' content='0; url=http:%s' /></head></html>", mb.ImgURL)
	}

}

func logIntoBGG() (client *http.Client) {
	for {
		var err error
		client, err = website.Login("https://boardgamegeek.com/login", *username, *password, 30*time.Second)

		if err != nil {
			fmt.Println("BGG currently unavailable")
			//			log.Fatal(err)
		} else {
			return
		}
		time.Sleep(10 * time.Second)
	}
}

func getRandomBadges() []microBadge {

	badgeList := []microBadge{}
	for i := 1; i < 6; i++ {
		slotID := fmt.Sprintf("%d", i)
		if currentSlot, ok := slotMap[slotID]; ok {
			for _, mb := range currentSlot.AvailableBadges {
				badgeList = append(badgeList, *mb)
				break
			}
		}
	}

	return badgeList
}

func assignSlot(id, slotNumber string, client *http.Client) error {
	resp, err := client.PostForm("https://boardgamegeek.com/geekmicrobadge.php", url.Values{
		"badgeid": {id},
		"slot":    {slotNumber},
		"ajax":    {"1"},
		"action":  {"setslot"},
	})
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	//Response if not logged in is 85 bytes long
	if len(data) < 86 {
		return errors.New("Invalid username or password. Restart microBadger and attempt to log in again.")
	}
	if givenSlot, ok := slotMap[slotNumber]; ok {
		givenSlot.AssignedBadge = id
	} else {
		slotMap[slotNumber] = &slot{Id: slotNumber, AssignedBadge: id}
	}
	return nil
}
