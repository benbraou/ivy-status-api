// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"fmt"
	"regexp"
)

// UnorderedListIndentationLevel returns the index of the last whitespace announcing a bullet list
// e.g:
// for '- li', we return 1
// for '  - li', we return 3
// If the provided line does not correspond t a list element, an error is returned
func UnorderedListIndentationLevel(line string) (int, error) {
	match := regexp.
		MustCompile(`^\s*(-|\*)`).
		FindStringIndex(line)
	if len(match) == 0 {
		return -1, fmt.Errorf("Encountered a line that is an unordered list element: %s", line)
	}
	return match[1], nil
}
