package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var (
	Links []string
	visited map[string]bool = map[string]bool{}
)

func main() {
	visitLink("https://www.github.com")

	fmt.Println(len(Links))
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
		panic(fmt.Errorf("Status != de 200 %v", resp.StatusCode))
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
			Links = append(Links, link.String())
			visitLink(link.String())

		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ExtractLinks(c)
	}
}
