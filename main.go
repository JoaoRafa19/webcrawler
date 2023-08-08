package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoRafa19/webcrawler/db"
	"golang.org/x/net/html"
)

type VisitedLinks struct {
	Website     string    `bson:"website"`
	Link        string    `bson:"link"`
	VisitedDate time.Time `bson:"visited_date"`
}

var link string

func init() {
	flag.StringVar(&link, "url", "https://www.github.com", "url para iniciar visitas")
}

func main() {
	flag.Parse()
	done := make(chan bool)
	go visitLink(link)
	<-done
}

func visitLink(link string) {
	fmt.Printf("visitando: %v\n", link)

	resp, err := http.Get(link)
	if err != nil {
		err := fmt.Errorf("unsuported protocol or status != 200 : %v\n error: %v", resp.Status, err)
		if err != nil {
			return
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("unsuported protocol or status != 200 : %v\n error: %v", resp.Status, err)
		if err != nil {
			return
		}
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		err := fmt.Errorf("an error ocourred\n\t%v", err)
		if err != nil {
			return
		}
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
			if err != nil || link.Scheme == "" || link.Scheme == "mailto" {
				continue
			}
			if db.VisitedLink(link.String()) {
				fmt.Printf("link ja visitado: %v\n", link)
				continue
			}

			// Links = append(Links, link.String())
			visitedLink := VisitedLinks{
				Website:     link.Host,
				Link:        link.String(),
				VisitedDate: time.Now(),
			}
			db.Insert("links", visitedLink)

			go visitLink(link.String())
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractLinks(c)
	}
}
