package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"

	"testing"
)

func TestClient_SeriesFinderGenres_Generated(t *testing.T) {

	client := nu.DefaultClient
	genres, err := client.SeriesFinderGenres()

	be.NilErr(t, err)

	t.Run("generated genres are valid", func(t *testing.T) {
		be.Equal(t, len(genres), len(nu.SlugToGenre))
		for _, genre := range genres {

			generatedGenre, _ := nu.ValueToGenre[genre.Value]
			be.Equal(t, generatedGenre, nu.Genre(genre.Value))

			generatedSlug, _ := nu.GenreToSlug[generatedGenre]
			be.Equal(t, generatedSlug, genre.Slug)

			generatedTitle, _ := nu.GenreToTitle[generatedGenre]
			be.Equal(t, generatedTitle, genre.Name)
		}
	})
}

func TestClient_SeriesFinderLanguages_Generated(t *testing.T) {

	client := nu.DefaultClient
	languages, err := client.SeriesFinderLanguages()

	be.NilErr(t, err)

	t.Run("generated languages are valid", func(t *testing.T) {
		be.Equal(t, len(languages), len(nu.SlugToLanguage))
		for _, language := range languages {

			generatedLanguage, _ := nu.ValueToLanguage[language.Value]
			be.Equal(t, generatedLanguage, nu.Language(language.Value))

			generatedSlug, _ := nu.LanguageToSlug[generatedLanguage]
			be.Equal(t, generatedSlug, language.Slug)

			generatedTitle, _ := nu.LanguageToTitle[generatedLanguage]
			be.Equal(t, generatedTitle, language.Name)
		}
	})
}

func TestClient_SeriesFinderNovelTypes_Generated(t *testing.T) {

	client := nu.DefaultClient
	novelTypes, err := client.SeriesFinderNovelTypes()

	be.NilErr(t, err)

	t.Run("generated novel types are valid", func(t *testing.T) {
		be.Equal(t, len(novelTypes), len(nu.SlugToNovelType))
		for _, novelType := range novelTypes {

			generatedNovelType, _ := nu.ValueToNovelType[novelType.Value]
			be.Equal(t, generatedNovelType, nu.NovelType(novelType.Value))

			generatedSlug, _ := nu.NovelTypeToSlug[generatedNovelType]
			be.Equal(t, generatedSlug, novelType.Slug)

			generatedTitle, _ := nu.NovelTypeToTitle[generatedNovelType]
			be.Equal(t, generatedTitle, novelType.Name)
		}
	})
}

func TestClient_SeriesFinderTags_Generated(t *testing.T) {

	client := nu.DefaultClient
	tags, err := client.SeriesFinderTags()

	be.NilErr(t, err)

	t.Run("generated tags are valid", func(t *testing.T) {
		be.Equal(t, len(tags), len(nu.SlugToTag))
		for _, tag := range tags {

			generatedTag, _ := nu.ValueToTag[tag.Value]
			be.Equal(t, generatedTag, nu.Tag(tag.Value))

			generatedSlug, _ := nu.TagToSlug[generatedTag]
			be.Equal(t, generatedSlug, tag.Slug)

			generatedTitle, _ := nu.TagToTitle[generatedTag]
			be.Equal(t, generatedTitle, tag.Name)
		}
	})
}
