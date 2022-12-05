package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

func TestClient_Languages(t *testing.T) {

	client := nu.DefaultClient
	languages, err := client.Languages()

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
