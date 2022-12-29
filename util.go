package nu

import (
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"strings"
)

func queryFirstChildDataOrDefault(n *html.Node, queryStr string, defaultStr string) string {

	node, _ := query(n, queryStr)
	nodeData := defaultStr

	if node != nil {
		nodeData = node.FirstChild.Data
	}

	return nodeData
}

func queryAttrOrDefault(n *html.Node, queryStr string, attrStr string, defaultStr string) string {

	node, _ := query(n, queryStr)
	attrData := defaultStr

	if node != nil {
		attrValue := attr(node, attrStr)
		if attrValue != "" {
			attrData = attrValue
		}
	}

	return attrData
}

func query(n *html.Node, query string) (*html.Node, error) {

	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}, err
	}

	return cascadia.Query(n, sel), nil
}

func queryAll(n *html.Node, query string) ([]*html.Node, error) {

	sel, err := cascadia.Parse(query)
	if err != nil {
		return []*html.Node{}, err
	}

	return cascadia.QueryAll(n, sel), nil
}

func attr(n *html.Node, attrName string) string {

	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}

	return ""
}

var (
	slugReplacer    = strings.NewReplacer(" ", "-", "/", "-slash-", "'", "")
	bracketReplacer = strings.NewReplacer("(", "", ")", "")
	newlineReplacer = strings.NewReplacer("\n", "", "\r", "")
)

func normalisedSlug(n string) string {
	return slugReplacer.Replace(strings.ToLower(n))
}

func normalisedDescription(n string) string {
	return newlineReplacer.Replace(n)
}
