// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"regexp"
	"strings"
)

// TableCells returns a list of table cells values
// e.g. | Annotation | `defineXXX()` | Run time | Spec | Compiler | Back Patch | will result in the
// slice [ Annotation, `defineXXX()`, Run time, Spec,  Compiler, Back Patch ]
func TableCells(line string) []string {
	substrings := regexp.MustCompile(`([^\\]\|\s*)|(^[\|])`).Split(line, -1)
	// cells uses the same backing array and capacity as substrings. We do this to have the storage
	// reused for the filtered slice.
	cells := substrings[:0]
	for _, substring := range substrings {
		if len(substring) > 0 {
			cells = append(cells, strings.TrimSpace(substring))
		}
	}
	return cells
}
