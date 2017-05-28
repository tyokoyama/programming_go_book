package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var mapping map[string]int

func main() {
	mapping = make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinks1: %v\n", err)
		os.Exit(1)
	}

	visit(doc)

	fmt.Println(mapping)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		mapping[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}