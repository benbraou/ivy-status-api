// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"regexp"
	"strings"
)

// IsLineSeperator considers a line as a separator if it does not contain any word character
func IsLineSeperator(line string) bool {
	return !regexp.
		MustCompile(`\w`).
		MatchString(line)
}

// Lines returns a slice of meaningful lines present in the markdown file content
func Lines(markdownContent string) []string {
	lines := make([]string, 0)
	for _, line := range strings.Split(markdownContent, "\n") {
		if !IsLineSeperator(line) {
			prepareLine(&line)
			lines = append(lines, line)
		}
	}
	return lines
}

// prepareLine removes prefixes and trailing whitespaces and only keep the useful information
func prepareLine(line *string) {
	// e.g. "  -  Sanitization    " => "Sanitization"
	r := regexp.MustCompile(`^\s*-?\s*|\s*$`)
	*line = r.ReplaceAllString(*line, "")
}
