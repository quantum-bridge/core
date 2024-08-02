package convert

import (
	"regexp"
	"strings"
)

var (
	// matchFirstCap matches the first capital letter of a string.
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	// matchAllCap matches all capital letters of a string.
	matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
