package iotutil

import (
	"bytes"
	"regexp"
	"unicode/utf8"
)

var variablesRegexp = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)

func Length(val string) int {
	return utf8.RuneCountInString(val)
}

func InsertColon(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune(':')
		}
	}
	return buffer.String()
}

func VerifyVariableName(name string) bool {
	return variablesRegexp.MatchString(name)
}
