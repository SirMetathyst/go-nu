package nu

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var englishTitleCaser = cases.Title(language.English)

type Genre string

type SeriesFinderGenreResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) SeriesFinderGenres() (results []SeriesFinderSearchPropertyResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("series-finder-genres: %w", err)
	}

	aGenreNodes, err := queryAll(doc, "a[genreid].genreme")
	if err != nil {
		return nil, fmt.Errorf("series-finder-genres: %w", err)
	}

	for _, option := range aGenreNodes {
		results = append(results, SeriesFinderSearchPropertyResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  englishTitleCaser.String(option.LastChild.Data),
			Value: attr(option, "genreid"),
		})
	}

	return results, nil
}
