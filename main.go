package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoRafa19/webcrawler/db"
	"golang.org/x/net/html"
)

var (
	visited map[string]bool = map[string]bool{}
)

type VisitedLinks struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

func main() {
	visitLink("https://www.github.com")

}

func visitLink(link string) {
	fmt.Println(link)

	if ok := visited[link]; ok {
		return
	}
	visited[link] = true
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("status != de 200 %v", resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}
	ExtractLinks(doc)

}

func ExtractLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attrr := range node.Attr {
			if attrr.Key != "href" {
				continue
			}
			link, err := url.Parse(attrr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}
			// Links = append(Links, link.String())
			visitedLink := VisitedLinks{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}
			db.Insert("links", visitedLink)

			visitLink(link.String())

		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractLinks(c)
	}
}
