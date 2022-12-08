package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

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
