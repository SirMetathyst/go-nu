package nu

import (
	"fmt"
)

type Language string

type SeriesFinderLanguageResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) SeriesFinderLanguages() (results []SeriesFinderSearchPropertyResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("series-finder-languages: %w", err)
	}

	aLanguageNodes, err := queryAll(doc, "a[genreid].langrank")
	if err != nil {
		return nil, fmt.Errorf("series-finder-languages: %w", err)
	}

	for _, option := range aLanguageNodes {
		results = append(results, SeriesFinderSearchPropertyResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  englishTitleCaser.String(option.LastChild.Data),
			Value: attr(option, "genreid"),
		})
	}

	return results, nil
}
