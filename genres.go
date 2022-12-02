package nu

import (
	"fmt"
	"golang.org/x/net/html"
)

type GenreResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) Genres() (genres []GenreResult, err error) {

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
		genres = append(genres, GenreResult{
			Slug:  normalisedSlug(option.LastChild.Data),
			Name:  option.LastChild.Data,
			Value: attr(option, "genreid"),
		})
	}

	return genres, nil
}
