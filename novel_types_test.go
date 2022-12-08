package nu_test

import (
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
