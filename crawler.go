package main
import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)
// const url = "http://www.dilidili.wang/anime/ysjmwzhsn/"
const url = "http://www.dilidili.wang/anime/steinsgate0/"
func getInfoByName(name string) string {
	return "test"
}

func main() {
	resp, err := http.Get(url)
	excluseTexts := []string{"pv", "台版", "PV"}
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".xfswiper1 li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		// title := s.Find("a").Text()
		name := s.Find("span").Text()
		href, _ := s.Find("a").Attr("href")
		for _, keyword := range excluseTexts {
			if (strings.Contains(name, keyword)) {
				return;
			}
		}
		
		// res := strings.Split(title, " ")
		// title := s.Find("i").Text()
		fmt.Println(name, href)
	})
}
