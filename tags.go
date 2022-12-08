package nu

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Tag string

type SeriesFinderTagResult struct {
	Slug  string
	Name  string
	Value string
}

func (s *Client) SeriesFinderTags() (results []SeriesFinderSearchPropertyResult, err error) {

	response, err := s.client.Get("https://www.novelupdates.com/series-finder/")
	if err != nil {
		return nil, fmt.Errorf("series-finder-tags: %w", err)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("series-finder-tags: %w", err)
	}

	tagOptionNodes, err := queryAll(doc, "#tags_include option")
	if err != nil {
		return nil, err
	}

	for _, option := range tagOptionNodes {
		results = append(results, SeriesFinderSearchPropertyResult{
			Slug:  normalisedSlug(option.FirstChild.Data),
			Name:  cases.Title(language.English).String(option.FirstChild.Data),
			Value: attr(option, "value"),
		})
	}

	return results, nil
}

type TagResult struct {
	Slug        string
	Name        string
	Description string
}

func (s *Client) ListTags(page int) (results []TagResult, err error) {

	if page < 1 {
		page = 1
	}

	doc, err := s.request(fmt.Sprintf("https://www.novelupdates.com/list-tags/?st=1&pg=%d", page))
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	tagNodes, err := queryAll(doc, ".staglistall a")
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	for _, tagNode := range tagNodes {
		results = append(results, TagResult{
			Slug:        normalisedSlug(tagNode.FirstChild.Data),
			Name:        cases.Title(language.English).String(tagNode.FirstChild.Data),
			Description: normalisedDescription(attr(tagNode, "title")),
		})
	}

	return results, nil
}
