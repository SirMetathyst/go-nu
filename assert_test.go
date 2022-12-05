package nu_test

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
	"testing"
)

func Uppercase(t testing.TB, s string) {
	t.Helper()
	if strings.ToUpper(s) != s {
		t.Fatalf("uppercase: want: `%s`, got: `%s`", strings.ToUpper(s), s)
	}
}

func Lowercase(t testing.TB, s string) {
	t.Helper()
	if strings.ToLower(s) != s {
		t.Fatalf("lowercase: want: `%s`, got: `%s`", strings.ToLower(s), s)
	}
}

func Title(t testing.TB, s string) {
	t.Helper()
	title := cases.Title(language.English).String(s)
	if title != s {
		t.Fatalf("title: want: %s`, got: `%s`", title, s)
	}
}

func Number(t testing.TB, s string) {
	t.Helper()
	if _, err := strconv.Atoi(s); err != nil {
		t.Fatalf("number: want: number, got: `%s`", s)
	}
}

func ContainsAny(t testing.TB, s string, chars string) {
	t.Helper()
	if !strings.ContainsAny(s, chars) {
		t.Fatalf("contains-any: want any: `%s`, got: `%s`", chars, s)
	}
}

func NotContainsAny(t testing.TB, s string, chars string) {
	t.Helper()
	if strings.ContainsAny(s, chars) {
		t.Fatalf("not-contains-any: not want any: `%s`, got: `%s`", chars, s)
	}
}
