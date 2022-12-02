package nu

import (
	"fmt"
	"golang.org/x/net/html"
)

type TagResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) Tags() (tags []TagResult, err error) {

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
		tags = append(tags, TagResult{
			Slug:  normalisedSlug(option.FirstChild.Data),
			Name:  option.FirstChild.Data,
			Value: attr(option, "value"),
		})
	}

	return tags, nil
}
