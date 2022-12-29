package main

import (
	"bytes"
	"fmt"
	"github.com/SirMetathyst/go-nu"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	client := nu.DefaultClient

	//if err := WriteNovelTypesToFile(client); err != nil {
	//	log.Fatalln(err)
	//}
	//if err := WriteLanguagesToFile(client); err != nil {
	//	log.Fatalln(err)
	//}
	//if err := WriteTagsToFile(client); err != nil {
	//	log.Fatalln(err)
	//}
	//if err := WriteGenresToFile(client); err != nil {
	//	log.Fatalln(err)
	//}

	if err := WriteAllToFile(client); err != nil {
		log.Fatalln(err)
	}
}

func normalisedName(n string) string {

	n = strcase.ToCamel(n)
	n = strings.Replace(n, " ", "", -1)
	n = strings.Replace(n, "-", "", -1)
	n = strings.Replace(n, "/", "", -1)
	n = strings.Replace(n, "'", "", -1)

	return n
}

func WriteAllToFile(client *nu.Client) error {

	buf := &bytes.Buffer{}
	WriteHeader(buf, "nu")
	if err := WriteNovelTypes(client, buf); err != nil {
		return err
	}
	if err := WriteLanguages(client, buf); err != nil {
		return err
	}
	if err := WriteGenres(client, buf); err != nil {
		return err
	}
	if err := WriteTags(client, buf); err != nil {
		return err
	}

	return WriteFormattedFile("generated.go", buf.Bytes())
}

func WriteLanguagesToFile(client *nu.Client) error {

	buf := &bytes.Buffer{}
	WriteHeader(buf, "nu")
	if err := WriteLanguages(client, buf); err != nil {
		return err
	}

	return WriteFormattedFile("languages_generated.go", buf.Bytes())
}

func WriteNovelTypesToFile(client *nu.Client) error {

	buf := &bytes.Buffer{}
	WriteHeader(buf, "nu")
	if err := WriteNovelTypes(client, buf); err != nil {
		return err
	}

	return WriteFormattedFile("novel_types_generated.go", buf.Bytes())
}

func WriteTagsToFile(client *nu.Client) error {

	buf := &bytes.Buffer{}
	WriteHeader(buf, "nu")
	if err := WriteTags(client, buf); err != nil {
		return err
	}

	b := &bytes.Buffer{}
	WriteHeader(b, "nu")
	WriteTags(client, b)

	_ = os.WriteFile("tags_generated.go", b.Bytes(), 0666)
	cmd := exec.Command("go", "fmt", "./tags_generated.go")
	return cmd.Run()
}

func WriteGenresToFile(client *nu.Client) error {

	buf := &bytes.Buffer{}
	WriteHeader(buf, "nu")
	if err := WriteGenres(client, buf); err != nil {
		return err
	}

	return WriteFormattedFile("genres_generated.go", buf.Bytes())
}

func WriteFormattedFile(name string, data []byte) error {

	_ = os.WriteFile(name, data, 0666)
	cmd := exec.Command("go", "fmt", name)

	return cmd.Run()
}

func WriteHeader(b *bytes.Buffer, packageName string) {

	//////// Header
	b.WriteString("// Code generated by nug (NovelUpdates Generator). DO NOT EDIT.\n")
	b.WriteString(fmt.Sprintf("package %s\n\n", packageName))
}

func WriteNovelTypes(client *nu.Client, b *bytes.Buffer) error {

	s := "NovelType"

	results, err := client.SeriesFinderNovelTypes()
	if err != nil {
		return err
	}

	//////// Main
	b.WriteString(fmt.Sprintf("// %s: Total(%d)\n", s, len(results)))
	b.WriteString("const (\n")

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", s, normalisedName(result.Name), s, result.Value))
	}
	b.WriteString(")\n\n")

	/////////// ValueTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tValueTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Value, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	//////// ToDisplayString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToTitle = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Name))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// DisplayStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tTitleTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Name, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// SlugStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tSlugTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Slug, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// ToSlugString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToSlug = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Slug))
	}

	b.WriteString("}\n")
	b.WriteString(")\n")

	return err
}

func WriteLanguages(client *nu.Client, b *bytes.Buffer) error {

	s := "Language"

	results, err := client.SeriesFinderLanguages()
	if err != nil {
		return err
	}

	//////// Main
	b.WriteString(fmt.Sprintf("// %s: Total(%d)\n", s, len(results)))
	b.WriteString("const (\n")

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", s, normalisedName(result.Name), s, result.Value))
	}
	b.WriteString(")\n\n")

	/////////// ValueTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tValueTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Value, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	//////// ToDisplayString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToTitle = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Name))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// DisplayStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tTitleTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Name, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// SlugStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tSlugTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Slug, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// ToSlugString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToSlug = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Slug))
	}

	b.WriteString("}\n")
	b.WriteString(")\n")

	return nil
}

func WriteGenres(client *nu.Client, b *bytes.Buffer) error {

	s := "Genre"

	type genreDescResult struct {
		result nu.GenreExplanationResult
		url    string
	}

	genreResultMap := make(map[string]genreDescResult)
	genreResults, err := client.GenreExplanation()
	if err != nil {
		return err
	}
	for _, genreResult := range genreResults {
		genreResultMap[genreResult.Name] = genreDescResult{result: genreResult, url: "https://www.novelupdates.com/genre-explanation/"}
	}

	results, err := client.SeriesFinderGenres()
	if err != nil {
		return err
	}

	//////// Main
	b.WriteString(fmt.Sprintf("// %s: Total(%d)\n", s, len(results)))
	b.WriteString("const (\n")

	//for _, result := range results {
	//	b.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", s, normalisedName(result.Name), s, result.Value))
	//}

	for _, result := range results {
		b.WriteString(fmt.Sprintf("// %s... \n// Description generated from: %s\n",
			normalisedTagDescription(genreResultMap[result.Name].result.Description, normalisedName(result.Name)), genreResultMap[result.Name].url))
		b.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", s, normalisedName(result.Name), s, result.Value))
	}

	b.WriteString(")\n\n")

	/////////// ValueTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tValueTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Value, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	//////// ToDisplayString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToTitle = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Name))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// DisplayStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tTitleTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Name, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// SlugStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tSlugTo%s = map[string]%s{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Slug, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// ToSlugString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToSlug = map[%s]string{\n", s, s))

	for _, result := range results {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Slug))
	}

	b.WriteString("}\n")
	b.WriteString(")\n")

	return err
}

func normalisedTagDescription(description string, tagName string) string {

	s := strings.ReplaceAll(description, "Tag is", fmt.Sprintf("Tag%s is", tagName))
	s = strings.ReplaceAll(s, "Tag should", fmt.Sprintf("Tag%s should", tagName))
	s = strings.ReplaceAll(s, "Tag that", fmt.Sprintf("Tag%s that", tagName))
	s = strings.ReplaceAll(s, "Tag refers", fmt.Sprintf("Tag%s refers", tagName))
	s = strings.ReplaceAll(s, "Tag to", fmt.Sprintf("Tag%s to", tagName))

	return s
}

func WriteTags(client *nu.Client, b *bytes.Buffer) error {

	s := "Tag"

	seriesFinderResults, err := client.SeriesFinderTags()
	if err != nil {
		return err
	}

	type tagDescResult struct {
		result nu.TagResult
		url    string
	}

	tagResultMap := make(map[string]tagDescResult)
	page := 1
	for tagResults, err := client.ListTags(page); tagResults != nil || err != nil; tagResults, err = client.ListTags(page) {
		if err != nil {
			return err
		}
		for _, tagResult := range tagResults {
			tagResultMap[tagResult.Name] = tagDescResult{result: tagResult, url: fmt.Sprintf("https://www.novelupdates.com/list-tags/?st=1&pg=%d", page)}
		}
		page++
	}

	//////// Main
	b.WriteString(fmt.Sprintf("// %s: Total(%d)\n", s, len(seriesFinderResults)))
	b.WriteString("const (\n")

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("// %s... \n// Description generated from: %s\n",
			normalisedTagDescription(tagResultMap[result.Name].result.Description, normalisedName(result.Name)), tagResultMap[result.Name].url))
		b.WriteString(fmt.Sprintf("\t%s%s %s = \"%s\"\n", s, normalisedName(result.Name), s, result.Value))
	}
	b.WriteString(")\n\n")

	/////////// ValueTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tValueTo%s = map[string]%s{\n", s, s))

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Value, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	//////// ToDisplayString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToTitle = map[%s]string{\n", s, s))

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Name))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// DisplayStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tTitleTo%s = map[string]%s{\n", s, s))

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Name, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// SlugStringTo
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\tSlugTo%s = map[string]%s{\n", s, s))

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("\t\"%s\":%s%s,\n", result.Slug, s, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	/////////// ToSlugString
	b.WriteString("var (\n")
	b.WriteString(fmt.Sprintf("\t%sToSlug = map[%s]string{\n", s, s))

	for _, result := range seriesFinderResults {
		b.WriteString(fmt.Sprintf("\t%s%s:\"%s\",\n", s, normalisedName(result.Name), result.Slug))
	}

	b.WriteString("}\n")
	b.WriteString(")\n")

	return err
}
