package nu

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TagResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) Tags() (results []TagResult, err error) {

	response, err := s.client.Get("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("tags: %w", err)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("tags: %w", err)
	}

	tagOptionNodes, err := queryAll(doc, "#tags_include option")
	if err != nil {
		return nil, fmt.Errorf("tags (#tags_include option): %w", err)
	}

	for _, option := range tagOptionNodes {
		results = append(results, TagResult{
			Slug:  normalisedSlug(option.FirstChild.Data),
			Name:  cases.Title(language.English).String(option.FirstChild.Data),
			Value: attr(option, "value"),
		})
	}

	return results, nil
}
