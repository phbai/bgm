package main

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	// "strings"
)

type Anime struct {
	Name string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

const URL = "https://mikanani.me"

func main() {
	resp, err := http.Get(URL)
	animes := []string{}
	// excludeTexts := []string{"pv", "台版", "PV"}
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".monday").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// title := s.Find("a").Text()
		fmt.Println("---------------------------------------------------");
		fmt.Println(s.Text())
		sel := s.Parent().Parent();

		// category := s.Find("span").Text()
		
		sel.Find(".m-week-square").Each(func(i int, animeSel *goquery.Selection) {
			name := animeSel.Find(".small-title").Text()
			// thumbnail, _ := animeSel.Find("img").Attr("data-src")
			// anime := Anime{Name: name, Thumbnail: URL + thumbnail}
			animes = append(animes, name)
		})

		for i, name := range animes {
			fmt.Print(name, "    ")
			if (i % 4 == 0) {
				fmt.Println()
			}
		}
	})
}