package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"testing"
)

func TestSeriesFinder(t *testing.T) {
	client := nu.DefaultClient
	client.SeriesFinder(nu.SeriesFinderSearchRequest{
		//NovelType: []nu.NovelType{
		//	nu.NovelTypeLightNovel,
		//	nu.NovelTypePublishedNovel,
		//},
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
}