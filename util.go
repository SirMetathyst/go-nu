package nu

import (
	"fmt"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"strings"
)

func query(n *html.Node, query string) (*html.Node, error) {

	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}, err
	}

	node := cascadia.Query(n, sel)
	if node == nil {
		return nil, fmt.Errorf("query: selector resulted in nil node")
	}

	return node, nil
}

func queryAll(n *html.Node, query string) ([]*html.Node, error) {

	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}, err
	}

	nodes := cascadia.QueryAll(n, sel)
	if nodes == nil {
		return nil, fmt.Errorf("query: selector resulted in nil nodes")
	}

	return nodes, nil
}

func attr(n *html.Node, attrName string) string {

	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}

	return ""
}

func normalisedSlug(n string) string {

	n = strings.ToLower(n)
	n = strings.Replace(n, " ", "-", -1)
	n = strings.Replace(n, "/", "-slash-", -1)
	n = strings.Replace(n, "'", "", -1)

	return n
}
