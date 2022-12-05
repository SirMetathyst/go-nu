package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"

	"testing"
)

func TestClient_Genres(t *testing.T) {

	client := nu.DefaultClient
	genres, err := client.Genres()

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
