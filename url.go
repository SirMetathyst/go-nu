package nu

import (
	"net/url"
	"sort"
	"strings"
)

type Values map[string][]string

func Get[T ~string](v Values, key T) string {
	if v == nil {
		return ""
	}
	vs := v[string(key)]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func Set[T ~string](v Values, key T, value T) {
	if value != "" {
		v[string(key)] = []string{string(value)}
	}
}

func Add[T ~string](s Values, key T, value ...T) {
	for _, vv := range value {
		if vv != "" {
			s[string(key)] = append(s[string(key)], string(vv))
		}
	}
}

func Del[T ~string](v Values, key T) {
	delete(v, string(key))
}

func Has[T ~string](v Values, key T) bool {
	_, ok := v[string(key)]
	return ok
}

func join[T ~string](elems []T, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return url.QueryEscape(string(elems[0]))
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}
	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(elems[0]))
	for _, s := range elems[1:] {
		if s != "" {
			b.WriteString(sep)
			b.WriteString(url.QueryEscape(string(s)))
		}
	}
	return b.String()
}

func Encode(v Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.QueryEscape(k)
		if len(vs) == 1 && vs[0] != "" {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(vs[0]))
		} else if len(vs) > 1 {
			jv := join(vs, ",")
			if jv != "" {
				if buf.Len() > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(keyEscaped)
				buf.WriteByte('=')
				buf.WriteString(jv)
			}
		}
	}
	return buf.String()
}
