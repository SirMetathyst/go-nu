package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

func TestClient_SeriesFinderNovelTypes(t *testing.T) {

	client := nu.DefaultClient
	novelTypes, err := client.SeriesFinderNovelTypes()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, novelType := range novelTypes {
			Lowercase(t, novelType.Slug)
			NotContainsAny(t, novelType.Slug, "/' \t\n\r")
			Title(t, novelType.Name)
			NotContainsAny(t, novelType.Name, "\t\n\r")
			Number(t, novelType.Value)
		}
	})
}

func TestClient_Languages(t *testing.T) {

	client := nu.DefaultClient
	languages, err := client.SeriesFinderLanguages()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, language := range languages {
			Lowercase(t, language.Slug)
			NotContainsAny(t, language.Slug, "/' \t\n\r")
			Title(t, language.Name)
			NotContainsAny(t, language.Name, "\t\n\r")
			Number(t, language.Value)
		}
	})
}

func TestClient_SeriesFinderGenres(t *testing.T) {

	client := nu.DefaultClient
	genres, err := client.SeriesFinderGenres()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, genre := range genres {
			Lowercase(t, genre.Slug)
			NotContainsAny(t, genre.Slug, "/' \t\n\r")
			Title(t, genre.Name)
			NotContainsAny(t, genre.Name, "\t\n\r")
			Number(t, genre.Value)
		}
	})
}

func TestClient_SeriesFinderTags(t *testing.T) {

	client := nu.DefaultClient
	tags, err := client.SeriesFinderTags()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, tag := range tags {
			Lowercase(t, tag.Slug)
			NotContainsAny(t, tag.Slug, "/' \t\n\r")
			Title(t, tag.Name)
			NotContainsAny(t, tag.Name, "\t\n\r")
			Number(t, tag.Value)
		}
	})
}

func TestClient_SeriesFinderSearch(t *testing.T) {

	client := nu.DefaultClient
	results, err := client.SeriesFinderSearch(nu.SeriesFinderSearchRequest{
		NovelType: []nu.NovelType{
			nu.NovelTypeWebNovel,
			nu.NovelTypeLightNovel,
		},
		//Language: []nu.Language{
		//	nu.LanguageJapanese,
		//	nu.LanguageChinese,
		//},
		//ChaptersRange: nu.RangeMax,
		//Chapters:      10,
		//FrequencyRange: nu.RangeMax,
		//Frequency:      1,
		//ReviewsRange: nu.RangeMin,
		//Reviews:      0,
		//RatingRange: nu.RangeMax,
		//Rating:      nu.Star0,
		//ReadersRange: nu.RangeMin,
		//Readers:      0,
		//FirstReleaseDateRange: nu.RangeMax,
		//FirstReleaseDate:      time.Now(),
		//LastReleaseDateRange: nu.RangeMax,
		//LastReleaseDate:      time.Now(),
		//GenreOperator: nu.OpAND,
		//GenreInclude:  []nu.Genre{},
		//GenreExclude:  []nu.Genre{nu.GenreAction},
		//TagOperator: nu.OpAND,
		//TagInclude:  []nu.Tag{},
		//TagExclude:  []nu.Tag{nu.TagAcademy},
		//Status: nu.StatusOngoing,
		//Groups:      []string{"1022", "931"},
		//GroupFilter: nu.FilterExclude,
		//OriginalPublisherFilter: nu.FilterInclude,
		//OriginalPublishers:      []string{"2381"},
		//EnglishPublisherFilter: nu.FilterInclude,
		//EnglishPublishers:      []string{"24"},
		//SeriesContains: "test",
		//SortBy:  nu.SortTitle,
		//OrderBy: nu.OrderAsc,
	})

	be.NilErr(t, err)

	for _, result := range results {
		fmt.Println(result)
	}
}
