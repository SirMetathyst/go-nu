package nu_test

import (
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"github.com/carlmjohnson/be"
	"testing"
)

func TestClient_GenreExplanation(t *testing.T) {

	client := nu.DefaultClient
	genres, err := client.GenreExplanation()

	be.NilErr(t, err)

	t.Run("data scraped successfully", func(t *testing.T) {
		for _, genre := range genres {
			if len(genre.Slug) == 0 {
				t.Errorf("genre slug is empty")
			}
			Lowercase(t, genre.Slug)
			NotContainsAny(t, genre.Slug, "/' \t\n\r")
			if len(genre.Name) == 0 {
				t.Errorf("genre name is empty")
			}
			Title(t, genre.Name)
			NotContainsAny(t, genre.Name, "\t\n\r")
			if len(genre.Description) == 0 {
				t.Errorf("genre description is empty")
			}
			NotContainsAny(t, genre.Description, "\t\n\r")
			fmt.Println(genre)
		}
	})
}

func TestClient_ListTags(t *testing.T) {

	client := nu.DefaultClient
	tags, err := client.ListTags(1)

	fmt.Println(tags)

	if err != nil {
		t.Errorf("got: %v", err)
	}

	if len(tags) == 0 {
		t.Fatalf("list tags should return data on first page")
	}
}

func TestClient_ListNovelTypes(t *testing.T) {

	client := nu.DefaultClient
	novelTypes, err := client.ListNovelTypes()

	fmt.Println(novelTypes)

	if err != nil {
		t.Errorf("got: %v", err)
	}
}

func TestClient_ListLanguages(t *testing.T) {

	client := nu.DefaultClient
	languages, err := client.ListLanguages()

	fmt.Println(languages)

	if err != nil {
		t.Errorf("got: %v", err)
	}
}
