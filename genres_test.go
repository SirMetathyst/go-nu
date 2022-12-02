package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"testing"
)

func TestGenres(t *testing.T) {
	client := nu.DefaultClient
	genres, err := client.Genres()

	fmt.Printf("List of genres:\n\n")
	for i, genre := range genres {
		fmt.Printf("Genre #%d: Slug: \"%s\", Name: \"%s\", Value: \"%s\"\n", i, genre.Slug, genre.Name, genre.Value)
	}

	if err != nil {
		panic(err)
	}
}
