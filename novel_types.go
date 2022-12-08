package nu

import (
	"fmt"
)

type NovelType string

type SeriesFinderNovelTypeResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) SeriesFinderNovelTypes() (results []SeriesFinderSearchPropertyResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("series-finder-novel-types: %w", err)
	}

	aNovelTypeNodes, err := queryAll(doc, "a[genreid].typerank")
	if err != nil {
		return nil, fmt.Errorf("series-finder-novel-types: %w", err)
	}

	for _, option := range aNovelTypeNodes {
		results = append(results, SeriesFinderSearchPropertyResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  englishTitleCaser.String(option.LastChild.Data),
			Value: attr(option, "genreid"),
		})
	}

	return results, nil
}
