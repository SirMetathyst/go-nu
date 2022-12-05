package nu

import (
	"fmt"
	"golang.org/x/net/html"
)

type LanguageResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) Languages() (results []LanguageResult, err error) {

	resp, err := s.client.Get("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("languages: %w", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("languages: %w", err)
	}

	aLanguageNodes, err := queryAll(doc, "a[genreid].langrank")
	if err != nil {
		return nil, fmt.Errorf("languages (a[genreid].langrank): %w", err)
	}

	for _, option := range aLanguageNodes {
		results = append(results, LanguageResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  englishTitleCaser.String(option.LastChild.Data),
			Value: attr(option, "genreid"),
		})
	}

	return results, nil
}
