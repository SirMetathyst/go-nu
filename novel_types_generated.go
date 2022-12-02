// Code generated by nug (NovelUpdates Generator). DO NOT EDIT.
package nu

type NovelType string

// NovelType: Total(3)
const (
	NovelTypeLightNovel     NovelType = "2443"
	NovelTypePublishedNovel NovelType = "26874"
	NovelTypeWebNovel       NovelType = "2444"
)

var (
	NovelTypeToTitle = map[NovelType]string{
		NovelTypeLightNovel:     "Light Novel",
		NovelTypePublishedNovel: "Published Novel",
		NovelTypeWebNovel:       "Web Novel",
	}
)

var (
	TitleToNovelType = map[string]NovelType{
		"Light Novel":     NovelTypeLightNovel,
		"Published Novel": NovelTypePublishedNovel,
		"Web Novel":       NovelTypeWebNovel,
	}
)

var (
	SlugToNovelType = map[string]NovelType{
		"light-novel":     NovelTypeLightNovel,
		"published-novel": NovelTypePublishedNovel,
		"web-novel":       NovelTypeWebNovel,
	}
)

var (
	NovelTypeToSlug = map[NovelType]string{
		NovelTypeLightNovel:     "light-novel",
		NovelTypePublishedNovel: "published-novel",
		NovelTypeWebNovel:       "web-novel",
	}
)