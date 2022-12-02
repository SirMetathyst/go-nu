package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"testing"
)

func TestLanguages(t *testing.T) {
	client := nu.DefaultClient
	languages, err := client.Languages()

	fmt.Printf("List of languages:\n\n")
	for i, language := range languages {
		fmt.Printf("Language #%d: Slug: \"%s\", Name: \"%s\", Value: \"%s\"\n", i, language.Slug, language.Name, language.Value)
	}

	if err != nil {
		panic(err)
	}
}
