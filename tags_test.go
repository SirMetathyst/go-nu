package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

func TestClient_SeriesFinderTags(t *testing.T) {

	client := nu.DefaultClient
	tags, err := client.SeriesFinderTags()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, tag := range tags {
			Lowercase(t, tag.Slug)
			NotContainsAny(t, tag.Slug, "/' \t\n\r")
			Title(t, tag.Name)
			NotContainsAny(t, tag.Name, "\t\n\r")
			Number(t, tag.Value)
		}
	})
}

func TestClient_ListTags(t *testing.T) {

	client := nu.DefaultClient
	tags, err := client.ListTags(1)

	if err != nil {
		t.Errorf("got: %v", err)
	}

	if len(tags) == 0 {
		t.Fatalf("list tags should return data on first page")
	}
}
