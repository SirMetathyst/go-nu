package nu

import (
	"fmt"
	"golang.org/x/net/html"
	"strconv"
	"time"
)

type SeriesFinderSearchRequest struct {
	NovelType               []NovelType
	Language                []Language
	ChaptersRange           Range
	Chapters                int
	FrequencyRange          Range
	Frequency               float32
	ReviewsRange            Range
	Reviews                 int
	RatingRange             Range
	Rating                  Rating
	ReadersRange            Range
	Readers                 int
	FirstReleaseDateRange   Range
	FirstReleaseDate        time.Time
	LastReleaseDateRange    Range
	LastReleaseDate         time.Time
	GenreOperator           Operator
	GenreInclude            []Genre
	GenreExclude            []Genre
	TagOperator             Operator
	TagInclude              []Tag
	TagExclude              []Tag
	Status                  Status
	GroupFilter             Filter
	Groups                  []string
	OriginalPublisherFilter Filter
	OriginalPublishers      []string
	EnglishPublisherFilter  Filter
	EnglishPublishers       []string
	SeriesContains          string
	SortBy                  Sort
	OrderBy                 Order
}

type SeriesFinderResult struct {
	Title string
}

func (s *Client) SeriesFinder(req SeriesFinderSearchRequest) (results []SeriesFinderResult, err error) {

	v := encodeSeriesFinderSearchRequest(req)

	fmt.Println(Encode(v))

	response, err := s.client.Get(fmt.Sprintf("https://www.novelupdates.com/series-finder/?%s", Encode(v)))
	if err != nil {
		return nil, fmt.Errorf("series-finder: %w", err)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("series-finder: %w", err)
	}

	searchResultNodes, err := queryAll(doc, ".search_main_box_nu")
	if err != nil {
		return nil, fmt.Errorf("series-finder (.search_main_box_nu): %w", err)
	}

	for _, searchResultNode := range searchResultNodes {

		titleNode, err := query(searchResultNode, ".search_title a")
		if err != nil {
			return nil, fmt.Errorf("series-finder (.search_title): %w", err)
		}

		results = append(results, SeriesFinderResult{
			Title: titleNode.FirstChild.Data,
		})
	}

	return results, nil
}

func encodeSeriesFinderSearchRequest(req SeriesFinderSearchRequest) Values {

	v := Values{}
	Set(v, "sf", "1")
	Add(v, "nt", req.NovelType...)
	Add(v, "nt", req.Language...)
	//Set(v, "mrl", req.ChaptersRange.EncodeWithDefault(RangeMin), req.Chapters > 0)
	//Set(v, "rl", req.Chapters, req.Chapters > 0)
	Set(v, "mrl", req.ChaptersRange.EncodeWithDefault(RangeMin))
	Set(v, "rl", strconv.Itoa(req.Chapters))
	//Set(v, "mrf", req.FrequencyRange.EncodeWithDefault(RangeMax), req.Frequency > 0)
	//Set(v, "rf", req.Frequency, req.Frequency > 0)
	Set(v, "mrf", req.FrequencyRange.EncodeWithDefault(RangeMax))
	Set(v, "rf", strconv.FormatFloat(float64(req.Frequency), 'f', -1, 32))
	//Set(v, "mrvc", req.ReviewsRange.EncodeWithDefault(RangeMin), req.Reviews != 0 || req.ReviewsRange != RangeMin)
	//Set(v, "rvc", req.Reviews, req.Reviews != 0 || req.ReviewsRange != RangeMin)
	Set(v, "mrvc", req.ReviewsRange.EncodeWithDefault(RangeMin))
	Set(v, "rvc", strconv.Itoa(req.Reviews))
	//Set(v, "mrt", req.RatingRange.EncodeWithDefault(RangeMin), req.Rating != 0 || req.RatingRange != RangeMin)
	//Set(v, "rt", req.Rating.EncodeWithDefault(Star0), req.Rating != 0 || req.RatingRange != RangeMin)
	Set(v, "mrt", req.RatingRange.EncodeWithDefault(RangeMin))
	Set(v, "rt", req.Rating.EncodeWithDefault(Star0))
	//Set(v, "mrct", req.ReadersRange.EncodeWithDefault(RangeMin), req.Readers != 0 || req.ReadersRange != RangeMin)
	//Set(v, "rct", req.Readers, req.Readers != 0 || req.ReadersRange != RangeMin)
	Set(v, "mrct", req.ReadersRange.EncodeWithDefault(RangeMin))
	Set(v, "rct", strconv.Itoa(req.Readers))
	//Set(v, "mdtf", req.FirstReleaseDateRange.EncodeWithDefault(RangeMin), !req.FirstReleaseDate.IsZero() || req.FirstReleaseDateRange != RangeMin)
	//Set(v, "dtf", req.FirstReleaseDate.Format("01/02/2006"), !req.FirstReleaseDate.IsZero() || req.FirstReleaseDateRange != RangeMin)
	Set(v, "mdtf", req.FirstReleaseDateRange.EncodeWithDefault(RangeMin))
	Set(v, "dtf", req.FirstReleaseDate.Format("01/02/2006"))
	//SetPresent(v, "mdtf", req.LastReleaseDateRange.EncodeWithDefault(RangeMin), !req.LastReleaseDate.IsZero() || req.LastReleaseDateRange != RangeMin)
	//SetPresent(v, "dtf", req.LastReleaseDate.Format("01/02/2006"), !req.LastReleaseDate.IsZero() || req.LastReleaseDateRange != RangeMin)
	Set(v, "mdtf", req.LastReleaseDateRange.EncodeWithDefault(RangeMin))
	Set(v, "dtf", req.LastReleaseDate.Format("01/02/2006"))
	//SetPresent(v, "mgi", req.GenreOperator.EncodeWithDefault(OpAND), len(req.GenreInclude) > 0)
	//SetPresent(v, "gi", Join(req.GenreInclude, ","))
	//SetPresent(v, "ge", Join(req.GenreExclude, ","))
	Set(v, "mgi", req.GenreOperator.EncodeWithDefault(OpAND))
	Add(v, "gi", req.GenreInclude...)
	Add(v, "ge", req.GenreExclude...)
	//Set(v, "mtgi", req.TagOperator.EncodeWithDefault(OpOR), len(req.TagInclude) > 0)
	//Add(v, "tgi", Join(req.TagInclude, ","))
	//Add(v, "tge", Join(req.TagExclude, ","))
	Set(v, "mtgi", req.TagOperator.EncodeWithDefault(OpOR))
	Add(v, "tgi", req.TagInclude...)
	Add(v, "tge", req.TagExclude...)
	//Set(v, "ss", req.Status.EncodeWithDefault(StatusAll), req.Status != StatusAll)
	Set(v, "ss", req.Status.EncodeWithDefault(StatusAll))
	//Set(v, "grpi", req.GroupFilter.EncodeWithDefault(FilterInclude), len(req.Groups) > 0)
	//Add(v, "grp", Join(req.Groups, ","))
	Set(v, "grpi", req.GroupFilter.EncodeWithDefault(FilterInclude))
	Add(v, "grp", req.Groups...)
	//Set(v, "opi", req.OriginalPublisherFilter.EncodeWithDefault(FilterInclude), len(req.OriginalPublishers) > 0)
	//Add(v, "op", Join(req.OriginalPublishers, ","))
	Set(v, "opi", req.OriginalPublisherFilter.EncodeWithDefault(FilterInclude))
	Add(v, "op", req.OriginalPublishers...)
	//Set(v, "enpi", req.EnglishPublisherFilter.EncodeWithDefault(FilterInclude), len(req.EnglishPublishers) > 0)
	//Add(v, "enp", Join(req.EnglishPublishers, ","))
	Set(v, "enpi", req.EnglishPublisherFilter.EncodeWithDefault(FilterInclude))
	Add(v, "enp", req.EnglishPublishers...)
	Set(v, "sh", req.SeriesContains)
	Set(v, "sort", req.SortBy.EncodeWithDefault(SortLastUpdated))
	Set(v, "order", req.OrderBy.EncodeWithDefault(OrderDesc))

	return v
}

type Range int

const (
	RangeMin Range = iota
	RangeMax
)

func (s Range) Encode() string {
	if s == RangeMin {
		return "min"
	} else if s == RangeMax {
		return "max"
	}
	return ""
}

func (s Range) EncodeWithDefault(def Range) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Rating int

const (
	Star0 Rating = iota
	Star1
	Star2
	Star3
	Star4
	Star5
)

func (s Rating) Encode() string {
	if s >= 0 && s <= 5 {
		return strconv.Itoa(int(s))
	}
	return "0"
}

func (s Rating) EncodeWithDefault(def Rating) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Operator int

const (
	OpAND Operator = iota
	OpOR
)

func (s Operator) Encode() string {
	if s == OpAND {
		return "and"
	} else if s == OpOR {
		return "or"
	}
	return ""
}

func (s Operator) EncodeWithDefault(def Operator) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Status int

const (
	StatusAll Status = iota
	StatusCompleted
	StatusOngoing
	StatusHiatus
)

func (s Status) Encode() string {
	if s >= 0 && s <= 3 {
		return strconv.Itoa(int(s) + 1)
	}
	return ""
}

func (s Status) EncodeWithDefault(def Status) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Filter string

const (
	FilterInclude Filter = "include"
	FilterExclude Filter = "exclude"
)

func (s Filter) Encode() string {
	if s == FilterInclude {
		return "1"
	} else if s == FilterExclude {
		return "2"
	}
	return ""
}

func (s Filter) EncodeWithDefault(def Filter) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Sort string

const (
	SortChapters    Sort = "chapters"
	SortFrequency   Sort = "frequency"
	SortRank        Sort = "rank"
	SortRating      Sort = "rating"
	SortReaders     Sort = "readers"
	SortReviews     Sort = "reviews"
	SortTitle       Sort = "title"
	SortLastUpdated Sort = "date"
)

func (s Sort) Encode() string {
	switch s {
	case SortChapters:
		return "srel"
	case SortFrequency:
		return "sfrel"
	case SortRank:
		return "srank"
	case SortRating:
		return "srate"
	case SortReaders:
		return "sread"
	case SortReviews:
		return "sreview"
	case SortTitle:
		return "abc"
	case SortLastUpdated:
		return "sdate"
	}
	return ""
}

func (s Sort) EncodeWithDefault(def Sort) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}

type Order string

const (
	OrderDesc Order = "desc"
	OrderAsc  Order = "asc"
)

func (s Order) Encode() string {
	if s == OrderDesc || s == OrderAsc {
		return string(s)
	}
	return ""
}

func (s Order) EncodeWithDefault(def Order) string {
	v := s.Encode()
	if v == "" {
		return def.Encode()
	}
	return v
}
