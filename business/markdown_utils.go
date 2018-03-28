// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"errors"
	"regexp"
	"strings"
)

// MaxHeaderLevel corresponds to greatest valid markdown header level.
// Markdown headers can be: H1, H2, H3, H4, H5, H6
const MaxHeaderLevel = 6

// HeaderName returns the markdown header name (e.g. # Go is cool => Go is cool)
func HeaderName(line string) string {
	return regexp.
		MustCompile(`^\s*#*\s*|\s*$`).
		ReplaceAllString(line, "")
}

// HeaderLevel returns the markdown header level (this corresponds to H1, H2, H3, H4, H5, H6)
func HeaderLevel(header string) (int, error) {
	count := 0
	for _, char := range header {
		if string(char) == "#" {
			count++
		} else {
			break
		}
	}
	if count == 0 || count > MaxHeaderLevel {
		return count, errors.New("Not a header")
	}
	return count, nil
}

// IsLineHeader returns whether the given line corresponds to a markdown header
func IsLineHeader(line string) bool {
	return regexp.
		MustCompile(`^\s*#`).
		MatchString(line)
}

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
			lines = append(lines, UsefulContent(line))
		}
	}
	return lines
}

// UsefulContent removes prefixes and trailing whitespaces and only keep  the useful information
func UsefulContent(line string) string {
	// e.g. "  -  Sanitization    " => "Sanitization"
	r := regexp.MustCompile(`^\s*-?\s*|\s*$`)
	return r.ReplaceAllString(line, "")
}
