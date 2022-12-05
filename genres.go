package nu

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var englishTitleCaser = cases.Title(language.English)

type GenreResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) Genres() (results []GenreResult, err error) {

	resp, err := s.client.Get("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("genres: %w", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("genres: %w", err)
	}

	aGenreNodes, err := queryAll(doc, "a[genreid].genreme")
	if err != nil {
		return nil, fmt.Errorf("genres (a[genreid]): %w", err)
	}

	for _, option := range aGenreNodes {
		results = append(results, GenreResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  englishTitleCaser.String(option.LastChild.Data),
			Value: attr(option, "genreid"),
		})
	}

	return results, nil
}
