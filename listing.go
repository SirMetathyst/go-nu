package nu

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
)

type GenreExplanationResult struct {
	Slug        string
	Name        string
	Description string
}

func (s *Client) GenreExplanation() (results []GenreExplanationResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/genre-explanation/")
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	genreNodes, err := queryAll(doc, ".w-blog-content tr")
	if err != nil {
		return nil, fmt.Errorf("genre-explanation: %w", err)
	}

	for index, genreNode := range genreNodes {
		if index == 0 {
			continue
		}

		aData := queryFirstChildDataOrDefault(genreNode, ".genreexplain a", "")
		pData := queryFirstChildDataOrDefault(genreNode, "p", "")

		results = append(results, GenreExplanationResult{
			Slug:        normalisedSlug(aData),
			Name:        cases.Title(language.English).String(aData),
			Description: normalisedDescription(pData),
		})
	}

	return results, nil
}

type NovelTypeResult struct {
	Slug  string
	Name  string
	Count int
}

func (s *Client) ListNovelTypes() (results []NovelTypeResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/list-types/")
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	parentNodes, err := queryAll(doc, ".g-cols.wpb_row.offset_default")
	if err != nil {
		return nil, fmt.Errorf("list-languages: %w", err)
	}

	if len(parentNodes) != 2 {
		return nil, fmt.Errorf("list-novel-types: expected two `.g-cols.wpb_row.offset_default`")
	}

	liNodes, err := queryAll(parentNodes[1], "li")
	if err != nil {
		return nil, fmt.Errorf("list-novel-types: %w", err)
	}

	for _, liNode := range liNodes {

		aData := queryFirstChildDataOrDefault(liNode, "a", "")

		countRaw := strings.TrimSpace(bracketReplacer.Replace(liNode.LastChild.Data))
		count, err := strconv.Atoi(countRaw)
		if err != nil {
			count = -1
		}

		results = append(results, NovelTypeResult{
			Slug:  normalisedSlug(aData),
			Name:  cases.Title(language.English).String(aData),
			Count: count,
		})
	}

	return results, nil
}

type LanguageResult struct {
	Slug  string
	Name  string
	Count int
}

func (s *Client) ListLanguages() (results []LanguageResult, err error) {

	doc, err := s.request("https://www.novelupdates.com/list-languages/")
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	parentNodes, err := queryAll(doc, ".g-cols.wpb_row.offset_default")
	if err != nil {
		return nil, fmt.Errorf("list-languages: %w", err)
	}

	if len(parentNodes) != 2 {
		return nil, fmt.Errorf("list-languages: expected two `.g-cols.wpb_row.offset_default`")
	}

	liNodes, err := queryAll(parentNodes[1], "li")
	if err != nil {
		return nil, fmt.Errorf("list-languages: %w", err)
	}

	for _, liNode := range liNodes {

		aData := queryFirstChildDataOrDefault(liNode, "a", "")

		countRaw := strings.TrimSpace(bracketReplacer.Replace(liNode.LastChild.Data))
		count, err := strconv.Atoi(countRaw)
		if err != nil {
			count = -1
		}

		results = append(results, LanguageResult{
			Slug:  normalisedSlug(aData),
			Name:  cases.Title(language.English).String(aData),
			Count: count,
		})
	}

	return results, nil
}

type TagResult struct {
	Slug        string
	Name        string
	Description string
	Count       int
}

func (s *Client) ListTags(page int) (results []TagResult, err error) {

	if page < 1 {
		page = 1
	}

	doc, err := s.request(fmt.Sprintf("https://www.novelupdates.com/list-tags/?st=1&pg=%d", page))
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	liNodes, err := queryAll(doc, ".staglistall li")
	if err != nil {
		return nil, fmt.Errorf("list-tags: %w", err)
	}

	for _, liNode := range liNodes {

		aData := queryFirstChildDataOrDefault(liNode, "a", "")
		aTitle := queryAttrOrDefault(liNode, "a", "title", "")

		countRaw := strings.TrimSpace(bracketReplacer.Replace(liNode.LastChild.Data))
		count, err := strconv.Atoi(countRaw)
		if err != nil {
			count = -1
		}

		results = append(results, TagResult{
			Slug:        normalisedSlug(aData),
			Name:        cases.Title(language.English).String(aData),
			Description: normalisedDescription(aTitle),
			Count:       count,
		})
	}

	return results, nil
}
