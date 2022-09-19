// Package date contains logic for date based requirements/formatting not
// handled by the standard lib.
package date

import (
	"regexp"
	"strconv"
)

// DaySuffix returns the string for the current date when you need a more
// human friendly date suffix.
func DaySuffix(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}

// IncludesSuffixFormat checks if the supplied date format string includes a
// day suffix identifier so you can check if it should be replaced with the
// correct value or not.
func IncludesSuffixFormat(input string) bool {
	re := regexp.MustCompile(`.*\d+(?:th|st|nd|rd){1}`)
	return re.MatchString(input)
}
