package create

import (
	"fmt"
	"regexp"
	"strings"
)

// SanitiseFileName removes any spaces or special characters so the format
// is valid to use as a file name.
func SanitiseFileName(name string) string {
	baseName := strings.Trim(name, " ")

	hasExtention := false
	if strings.HasSuffix(baseName, ".md") {
		hasExtention = true
	}

	separators := regexp.MustCompile(`[ &=+:*/]`)
	baseName = separators.ReplaceAllString(baseName, "-")

	doubleSeparators := regexp.MustCompile(`--`)
	baseName = doubleSeparators.ReplaceAllString(baseName, "-")

	baseName = strings.Trim(baseName, "-")

	if !hasExtention {
		return fmt.Sprintf("%s.md", baseName)
	}

	return baseName
}
