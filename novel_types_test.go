package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"testing"
)

func TestNovelTypes(t *testing.T) {
	client := nu.DefaultClient
	tags, err := client.NovelTypes()

	fmt.Printf("List of novel types:\n\n")
	for i, tag := range tags {
		fmt.Printf("Novel Type #%d: Slug: \"%s\", Name: \"%s\", Value: \"%s\"\n", i, tag.Slug, tag.Name, tag.Value)
	}

	if err != nil {
		panic(err)
	}
}
