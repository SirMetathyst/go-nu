package nu_test

import (
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

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
