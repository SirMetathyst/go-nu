package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

func TestClient_NovelTypes(t *testing.T) {

	client := nu.DefaultClient
	novelTypes, err := client.NovelTypes()

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
