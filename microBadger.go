package main

import (
	"errors"
	"flag"
	"fmt"
	website "github.com/allentechnology/website"
	"github.com/vharitonsky/iniflags"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/user"
	"runtime"
	"time"
)

var (
	slotMap          = map[int]slot{}
	microBadgeMap    = map[string]*microBadge{}
	tmpMicroBadgeMap = map[string]*microBadge{}
)

var (
	username = flag.String("username", "", "The boardgamegeek.com username used to log into the site")
	password = flag.String("password", "", "The boardgamegeek.com password associated with the given username")
)

var (
	appDir = ""
)

func init() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		appDir = currentUser.HomeDir + ".microBadger"
	case "windows":
		appDir = currentUser.HomeDir + "microBadger"
	}
}

func main() {
	iniflags.Parse()
	for *username == "" || *password == "" {
		getUsername()
		if *username == "" {
			fmt.Println("Invalid username")
		} else if *password == "" {
			fmt.Println("Invalid password")
		}
	}
	var client *http.Client
	for {
		var err error
		client, err = website.Login("https://boardgamegeek.com/login", *username, *password, 30*time.Second)

		if err != nil {
			fmt.Println("BGG currently unavailable")
		} else {
			break
		}
		//		time.Sleep(10 * time.Second)
		time.Sleep(10 * time.Minute)
	}

	//	client = logIntoBGG()
	//	loggedIn := true
	for {
		// if !loggedIn {
		// 	client = logIntoBGG()
		// }
		fmt.Print("Attempting to randomize badges: ")
		err := getMicroBadges(client)
		if err != nil {
			fmt.Println("Failed")
			fmt.Println(err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		badgeList := getRandomBadges(5)

		updateSuccess := make([]bool, len(badgeList))
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
		for i, v := range updateSuccess {
			if v {
				fmt.Printf("%d ", i+1)
			}
		}
		fmt.Println("updated successfully")
		time.Sleep(1 * time.Minute)
		//		time.Sleep(10 * time.Second)
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

func getRandomBadges(numBadges int) []microBadge {
	i := 0
	badgeList := []microBadge{}
	for _, v := range microBadgeMap {
		if v.Category != "Contest (Official)" {
			badgeList = append(badgeList, *v)
			i++
		}
		if i >= numBadges {
			break
		}
	}
	return badgeList
}

func assignSlot(id, slot string, client *http.Client) error {
	resp, err := client.PostForm("https://boardgamegeek.com/geekmicrobadge.php", url.Values{
		"badgeid": {id},
		"slot":    {slot},
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
	return nil
}