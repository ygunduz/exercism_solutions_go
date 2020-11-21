package main

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '_' || c == ',' || c == '-'
	}
	fields := strings.FieldsFunc(s, f)
	res := make([]byte, len(fields))
	for i, field := range fields {
		res[i] = []byte(strings.ToUpper(field))[0]
	}
	return string(res)
}
