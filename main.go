package main

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"io/ioutil"
	"regexp"
	// "strings"
)

type Anime struct {
	Name string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

const URL = "http://www.halihali.cc"

func list() {
	resp, err := http.Get(URL)
	
	weekends := [7]string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
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
	doc.Find(".bangumi-bar .tab-cont .item").Each(func(i int, s *goquery.Selection) {
		animes := []string{}
		// For each item found, get the band and title
		// title := s.Find("a").Text()
		// fmt.Println(s.Text())
		// sel := s.Parent().Parent();
		fmt.Print(weekends[i], "\n");
		// category := s.Find("span").Text()
		
		s.Find(".item-info").Each(func(i int, animeSel *goquery.Selection) {
			name := animeSel.Find("a").Text()
			// href, _ := animeSel.Find("a").Attr("href")
			// thumbnail, _ := animeSel.Find("img").Attr("data-src")
			// anime := Anime{Name: name, Thumbnail: URL + thumbnail}
			animes = append(animes, name)
		})

		for i, name := range animes {
			printName(name)
			if (i % 4 == 3) {
				fmt.Println()
			}
		}
		fmt.Println();
	})
}

func get(keyword string) {
	reg := regexp.MustCompile(".*<a href=\"(.*)\">" + keyword + "</a>")
	resp, err := http.Get(URL)
	
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			// handle error
	}

	html := string(body)
	
	params := reg.FindStringSubmatch(html)
	
	fmt.Println(params)
}
func printName(s string) {
	length := 20;
	trueLength := len([]rune(s))
	fmt.Print(s)
	if (trueLength < length) {
		for i := 0; i < length - trueLength; i++ {
			fmt.Print("　")
		}
	}
}
func main() {
	// fmt.Printf("%-15s%-15s%-15s%-15s\n", "邪神与厨二病少女", "妖怪旅馆营业中", "京都寺町三条商店街的福尔摩斯", "进击的巨人 第三季")
	// fmt.Printf("%-15s%-15s%-15s%-15s\n", "付丧神出租中", "轻羽飞扬", "刃牙 死囚篇", "黄金神威")

	args := os.Args[1:]
	switch args[0] {
		case "get":
			name := args[1]
			get(name)
		case "list":
			list()
		case "test":
			fmt.Print("")
	}
}