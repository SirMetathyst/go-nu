package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"testing"
)

func TestTags(t *testing.T) {
	client := nu.DefaultClient
	tags, err := client.Tags()

	fmt.Printf("List of tags:\n\n")
	for i, tag := range tags {
		fmt.Printf("Tag #%d: Slug: \"%s\", Name: \"%s\", Value: \"%s\"\n", i, tag.Slug, tag.Name, tag.Value)
	}

	if err != nil {
		panic(err)
	}
}
