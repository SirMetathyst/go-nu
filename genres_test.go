package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"

	"testing"
)

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
