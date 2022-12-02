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

func (s *Client) Languages() (languages []LanguageResult, err error) {

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
		languages = append(languages, LanguageResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  option.LastChild.Data,
			Value: attr(option, "genreid"),
		})
	}

	return languages, nil
}
