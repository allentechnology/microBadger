package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"strings"
)

func getMicroBadges(client *http.Client) error {
	resp, err := client.Get("https://boardgamegeek.com/user/" + *username + "/microbadges")
	if err != nil {
		return err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	profileTitleBlocks := scrape.FindAllNested(root, scrape.ByClass("profile_title"))
	for _, v := range profileTitleBlocks {
		if strings.Contains(v.FirstChild.Data, "Microbadges") {
			parseMicroBadgeData(v.NextSibling.NextSibling.FirstChild)
		}
	}

	microBadgeMap = tmpMicroBadgeMap
	categoryMap = make(map[string]mbSlice)
	for _, v := range microBadgeMap {
		categoryMap[v.Category] = append(categoryMap[v.Category], v)
	}

	return nil
}

func parseMicroBadgeData(node *html.Node) {

	badgeRows := scrape.FindAllNested(node, scrape.ByTag(atom.Tr))
	categories := []string{}
	for _, v := range badgeRows {
		categoryTdNode := v.FirstChild.NextSibling
		currentCategory := categoryTdNode.FirstChild.NextSibling.FirstChild.Data
		categories = append(categories, currentCategory)

		getBadgesInCategory(currentCategory, v)
	}

	return
}

func getBadgesInCategory(category string, node *html.Node) {
	microbadges := scrape.FindAllNested(node, scrape.ByTag(atom.A))
	for _, mbAnchor := range microbadges {
		mb := mbAnchor.Attr[0].Val
		if strings.Contains(mb, "/microbadge/") {
			mb = strings.Trim(mb, "/")
			mb = strings.Split(mb, "/")[1]
			if existingMB, ok := tmpMicroBadgeMap[mb]; ok {
				existingMB.Category = category
			} else {
				tmpMicroBadgeMap[mb] = &microBadge{Id: mb, Category: category}
			}
		}
	}

	mbImages := scrape.FindAllNested(node, scrape.ByTag(atom.Img))
	for _, mbImg := range mbImages {
		for _, v := range mbImg.Attr {
			switch v.Key {
			case "class":
				if v.Val == "tilebadge" {
					//TODO get tilebadge img url here
				}
			case "data-frz-src":
				imgLink := v.Val
				mb := strings.Split(imgLink, "_")[1]
				if existingMB, ok := tmpMicroBadgeMap[mb]; ok {
					existingMB.ImgURL = imgLink
				} else {
					tmpMicroBadgeMap[mb] = &microBadge{Id: mb, ImgURL: imgLink, Category: category}
				}
			case "onmouseover":
				mbDescription := v.Val
				mbDescription = strings.TrimPrefix(mbDescription, "return overlib('")
				mbDescription = strings.TrimSuffix(mbDescription, "', WRAP );")
				mbDescription = strings.Replace(mbDescription, "\\'", "'", -1)

				mb := mbImg.Parent.Attr[0].Val
				if strings.Contains(mb, "/microbadge/") {
					mb = strings.Trim(mb, "/")
					mb = strings.Split(mb, "/")[1]
					if existingMB, ok := tmpMicroBadgeMap[mb]; ok {
						existingMB.Description = mbDescription
					} else {
						tmpMicroBadgeMap[mb] = &microBadge{Id: mb, Name: mbDescription}
					}

				}
			}

		}

	}

}
