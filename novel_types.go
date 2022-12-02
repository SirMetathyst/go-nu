package nu

import (
	"fmt"
	"golang.org/x/net/html"
)

type NovelTypeResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) NovelTypes() (novelTypes []NovelTypeResult, err error) {

	resp, err := s.client.Get("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("novel-types: %w", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("novel-types: %w", err)
	}

	aNovelTypeNodes, err := queryAll(doc, "a[genreid].typerank")
	if err != nil {
		return nil, fmt.Errorf("novel-types (a[genreid].typerank): %w", err)
	}

	for _, option := range aNovelTypeNodes {
		novelTypes = append(novelTypes, NovelTypeResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  option.LastChild.Data,
			Value: attr(option, "genreid"),
		})
	}

	return novelTypes, nil
}
