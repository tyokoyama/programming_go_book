package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinks1: %v\n", err)
		os.Exit(1)
	}

	visit(doc)
}

func visit(n *html.Node) {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for _, link := range links {
		fmt.Println(link)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}

}