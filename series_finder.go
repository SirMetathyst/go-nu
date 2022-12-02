package nu

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strconv"
	"strings"
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
}

func (s *Client) SeriesFinder(r SeriesFinderSearchRequest) (seriesResults []SeriesFinderResult, err error) {

	// URL Params:

	// ?sf=1 		-> series-finder: 1

	// nt=2443,2833 -> novel-type

	// org=495,496  -> novel-language

	// This is the number of releases (chapters).
	// rl=10
	// (default) mrl=min

	// The release frequency of a novel. Higher frequency means the novel is updated more often.
	// rf=20
	// (default) mrf=min

	// The number of reviews of a novel
	// rvc=30
	// mrvc=min

	// Novel rating on a scale of 1 to 5.
	// rt=2
	// mrt=min

	// The amount of ratings for a novel.
	// rtc=2
	// mrtc=min

	// The number of readers a novel has.
	// rct=10
	// mrct=min

	// First release date. MM/DD/YYYY
	// dtf=12/01/2022
	// mdtf=min

	// Last release date
	// dt=12/14/2022
	// mdt=min

	// The genres of the novel.
	// (include) gi=8
	// (exclude) ge=280
	// mgi=and

	// The tags of the novel.
	// (include) tgi=1248
	// (exclude) tge=11325
	// (default) mtgi=or

	// Status of novel (translation)
	// all (default): ss=1
	// completed: ss=2
	// ongoing: ss=3
	// hiatus: ss=4

	// Group
	// grp=31278
	// include: grpi=1
	// exclude: grpi=2

	// Original Publisher
	// op=45114
	// (default) include: opi=1
	// exclude: opi=2

	// English Publisher
	// enp=29159
	// (default) include: enpi=1
	// exclude: enpi=2

	// Series which contains these words will be returned as results.
	// This will also search the associated names field.
	// sh=blah

	// sort=sdate
	// srel: 	chapters
	// sfrel:	frequency
	// srank: 	rank
	// srate:	rating
	// sread: 	readers
	// sreview: reviews
	// abc: 	title
	// sdate: 	last updated

	// order=desc
	// desc: 	descending
	// asc:  	ascending

	v := url.Values{}
	SetPresent(v, "sf", "1")
	SetPresent(v, "nt", Join(r.NovelType, ","))
	SetPresent(v, "org", Join(r.Language, ","))
	SetPresent(v, "mrl", r.ChaptersRange.EncodeWithDefault(RangeMin), r.Chapters > 0)
	SetPresent(v, "rl", r.Chapters, r.Chapters > 0)
	SetPresent(v, "mrf", r.FrequencyRange.EncodeWithDefault(RangeMax), r.Frequency > 0)
	SetPresent(v, "rf", r.Frequency, r.Frequency > 0)
	SetPresent(v, "mrvc", r.ReviewsRange.EncodeWithDefault(RangeMin), r.Reviews != 0 || r.ReviewsRange != RangeMin)
	SetPresent(v, "rvc", r.Reviews, r.Reviews != 0 || r.ReviewsRange != RangeMin)
	SetPresent(v, "mrt", r.RatingRange.EncodeWithDefault(RangeMin), r.Rating != 0 || r.RatingRange != RangeMin)
	SetPresent(v, "rt", r.Rating.EncodeWithDefault(Star0), r.Rating != 0 || r.RatingRange != RangeMin)
	SetPresent(v, "mrct", r.ReadersRange.EncodeWithDefault(RangeMin), r.Readers != 0 || r.ReadersRange != RangeMin)
	SetPresent(v, "rct", r.Readers, r.Readers != 0 || r.ReadersRange != RangeMin)
	SetPresent(v, "mdtf", r.FirstReleaseDateRange.EncodeWithDefault(RangeMin), !r.FirstReleaseDate.IsZero() || r.FirstReleaseDateRange != RangeMin)
	SetPresent(v, "dtf", r.FirstReleaseDate.Format("01/02/2006"), !r.FirstReleaseDate.IsZero() || r.FirstReleaseDateRange != RangeMin)
	SetPresent(v, "mdtf", r.LastReleaseDateRange.EncodeWithDefault(RangeMin), !r.LastReleaseDate.IsZero() || r.LastReleaseDateRange != RangeMin)
	SetPresent(v, "dtf", r.LastReleaseDate.Format("01/02/2006"), !r.LastReleaseDate.IsZero() || r.LastReleaseDateRange != RangeMin)
	SetPresent(v, "mgi", r.GenreOperator.EncodeWithDefault(OpAND), len(r.GenreInclude) > 0)
	SetPresent(v, "gi", Join(r.GenreInclude, ","))
	SetPresent(v, "ge", Join(r.GenreExclude, ","))
	SetPresent(v, "mtgi", r.TagOperator.EncodeWithDefault(OpOR), len(r.TagInclude) > 0)
	SetPresent(v, "tgi", Join(r.TagInclude, ","))
	SetPresent(v, "tge", Join(r.TagExclude, ","))
	SetPresent(v, "ss", r.Status.Encode(), r.Status != StatusAll)
	SetPresent(v, "grpi", r.GroupFilter.EncodeWithDefault(FilterInclude), len(r.Groups) > 0)
	SetPresent(v, "grp", Join(r.Groups, ","))
	SetPresent(v, "opi", r.OriginalPublisherFilter.EncodeWithDefault(FilterInclude), len(r.OriginalPublishers) > 0)
	SetPresent(v, "op", Join(r.OriginalPublishers, ","))
	SetPresent(v, "enpi", r.EnglishPublisherFilter.EncodeWithDefault(FilterInclude), len(r.EnglishPublishers) > 0)
	SetPresent(v, "enp", Join(r.EnglishPublishers, ","))
	SetPresent(v, "sh", r.SeriesContains)
	SetPresent(v, "sort", r.SortBy.EncodeWithDefault(SortLastUpdated))
	SetPresent(v, "order", r.OrderBy.EncodeWithDefault(OrderDesc))

	response, err := s.client.Get(fmt.Sprintf("https://www.novelupdates.com/series-finder/?%s", v.Encode()))
	if err != nil {
		return nil, fmt.Errorf("series-finder: %w", err)
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("series-finder: %w", err)
	}

	//tagOptionNodes, err := queryAll(doc, "#tags_include option")
	//if err != nil {
	//	return nil, fmt.Errorf("tags (#tags_include option): %w", err)
	//}
	//
	//for _, option := range tagOptionNodes {
	//	seriesResults = append(seriesResults, SeriesFinderResult{})
	//}

	return seriesResults, nil
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
	// all on remote server
	return "1"
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

func SetPresent(target url.Values, key string, value interface{}, condition ...bool) {
	dcondition := true
	if len(condition) > 0 {
		dcondition = condition[0]
	}
	if dcondition == false {
		return
	}
	switch svalue := value.(type) {
	case string:
		if len(svalue) != 0 {
			target.Set(key, svalue)
		}
	case int:
		target.Set(key, strconv.Itoa(svalue))
	case float32:
		target.Set(key, fmt.Sprintf("%.1f", svalue))
	default:
		panic("This type is not supported")
	}
}

func Join[T ~string](elems []T, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return string(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(string(s))
	}
	return b.String()
}
