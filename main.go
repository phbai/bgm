package main

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
  "os"
  "regexp"
  "github.com/phbai/bgm/util"
	// "strings"
)

type Anime struct {
	Name string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type Post struct {
  URL string `json:"url"`
  Title string `json:"title"`
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

func getLink(keyword string) string {
	// reg := regexp.MustCompile("alt=\"" + keyword + "\"></div></a><div class=\"item-info\"><a href=\"(.*)\">" + keyword + "</a>")
	reg := regexp.MustCompile(".*<a href=\"(.*)\">" + keyword + "</a>")
	html, ok := util.GetHTML(URL)
  
  if ok {
    params := reg.FindStringSubmatch(html)
    url := URL + params[1]
    return url;
  }
  return ""
}

func getInfo(link string) []Post {
  posts := []Post{}
  body, ok := util.GetBody(link)
  defer body.Close()
  if !ok {
    fmt.Print("error\n");
  }
  doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".player_list a").Each(func(i int, s *goquery.Selection) {
    name := s.Text()
    link, _ := s.Attr("href")
    post := Post{Title: name, URL: URL + link}
    posts = append(posts, post)
    // fmt.Printf("%-20s%-20s\n", name, URL + link);
  });
  // fmt.Printf("%-20s%-20s\n", name, URL + link);
  return posts;
}

func getVideoURL(link string) string {
  html, ok := util.GetHTML(link)
  if !ok {
    fmt.Println("error")
  }
  reg := regexp.MustCompile("\"url\":\"(.*?)\"")

  params := reg.FindStringSubmatch(html)
  videoURL := params[1]
  // url := URL + params[1]
  // return url;
  return videoURL
}

func getMultiVideoURL(posts []Post) {
  dic := make(map[string](chan string))
  // chs := make(chan string, len(posts))


  for _, v := range posts {
    dic[v.Title] = make(chan string, 1)

    go func(v Post) {
      url := getVideoURL(v.URL)
      dic[v.Title] <- url
    }(v)
  }

  for _, v := range posts {
    url := <- dic[v.Title]
    fmt.Printf("%-20s%-20s\n", v.Title, url);
  }
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
      link := getLink(name)
      fmt.Printf("%-15s%-20s\n", "解析成功地址为:", link)
      posts := getInfo(link)
      getMultiVideoURL(posts)
    case "parse":
      link := args[1]
      getVideoURL(link)
		case "list":
			list()
		case "test":
			fmt.Print("")
	}
}