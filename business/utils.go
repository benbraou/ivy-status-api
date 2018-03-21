// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"errors"
	"regexp"
	"strings"

	"github.com/benbraou/ivy-status-api/model"
)

// IconToStatusCode returns the feature status code (IMPLEMENTED, NOT_IMPLEMENTED, NOT_APPLICABLE)
// to be sent in the response.
// ✅ => IMPLEMENTED
// ❌ =>  NOT_IMPLEMENTED
// n/a => NOT_APPLICABLE
func IconToStatusCode(icon string) string {
	switch icon {
	case CHECK:
		return model.StatusToString(model.IMPLEMENTED)
	case CROSS:
		return model.StatusToString(model.NOT_IMPLEMENTED)
	case NA:
		return model.StatusToString(model.NOT_APPLICABLE)
	default:
		return model.StatusToString(model.NOT_IMPLEMENTED)
	}
}

// IsCategoryHeader checks whether a feature line corresponds to a list of categories
// (e.g. | Feature | Runtime | Spec | Compiler is a category header). Some feature lines do not
// represent a list of categories but rather the feature status e.g. Sanitization ✅
// We need to discard those.
func IsCategoryHeader(line string) bool {
	return !strings.Contains(line, CROSS) &&
		!strings.Contains(line, CHECK) &&
		!strings.Contains(line, NA)
}

// IsFeatureGroupHeader returns true when the provided line is a H2 markdown header
// e.g. ## Component Composition
func IsFeatureGroupHeader(line string) bool {
	return regexp.
		MustCompile(`^\s*#{2}`).
		MatchString(line)
}

// FeatureGroupName returns the title from a H2 markdown header (this is the convention used in
// https://github.com/angular/angular/blob/master/packages/core/src/render3/STATUS.md to declare a
// a feature group)
func FeatureGroupName(line string) string {
	return regexp.
		MustCompile(`^\s*#{2}\s*|\s*$`).
		ReplaceAllString(line, "")
}

// IsSeperator considers a line as a separator if it does not contain any word character
func IsSeperator(line string) bool {
	return !regexp.
		MustCompile(`\w`).
		MatchString(line)
}

// HasFeatureStatusInfo returns false when a feature group is found but no related features
// (This may happen when `STATUS.md` is in an incomplete state ? bad commit ?)
func HasFeatureStatusInfo(lines []string) bool {
	if len(lines) == 0 {
		return false
	}
	// Sometimes, Ivy STATUS.md provides no category header. It directly displays a line of features
	// e.g. for the section Missing Pieces, we have the line: Sanitization ✅
	// We need to make sure that in case one line is available per feature group, we have category
	// headers with no features
	if len(lines) == 1 && IsCategoryHeader(lines[0]) {
		return false
	}
	return true
}

// FeatureGroupIndices returns a slice of indices that describe the lines at which new feature group
// starts
func FeatureGroupIndices(lines []string) []int {
	indices := make([]int, 0)
	for i, line := range lines {
		if IsFeatureGroupHeader(line) {
			indices = append(indices, i)
		}
	}
	return indices
}

// GranularStatusIconAndDescription returns icon and description of a subfeature
// e.g. Sanitization ✅ => ( ✅  , Sanitization  )
func GranularStatusIconAndDescription(td string) (string, string) {
	for _, icon := range [3]string{CROSS, CHECK, NA} {
		description := strings.Replace(td, icon, "", 1)
		if description != td {
			return icon, strings.TrimSpace(description)
		}
	}
	return NA, td
}

// FeatureGroupLines return the list of lines containing information about the feature group and its
// related features
func FeatureGroupLines(lines []string, i int, featureGroupIndices []int) []string {
	var featureGroupLines []string
	if i == len(featureGroupIndices)-1 {
		featureGroupLines = lines[featureGroupIndices[i]:]

	} else {
		featureGroupLines = lines[featureGroupIndices[i]:featureGroupIndices[i+1]]
	}
	return featureGroupLines
}

// ValidateLines returns an error if the provided lines do no match validation criteria.
// For the time being, only check that the provided slice is not empty
func ValidateLines(lines []string) (err error) {
	if len(lines) == 0 {
		return errors.New("Expected an input of at least one line")
	}
	return
}

// FeatureGroupContent removes prefixes and trailing whitespaces in order to keep only the useful
// information
func FeatureGroupContent(line string) string {
	// e.g. "  -  Sanitization    " => "Sanitization"
	r := regexp.MustCompile(`^\s*-?\s*|\s*$`)
	return r.ReplaceAllString(line, "")
}

func categoryValue(value string) string {
	return regexp.
		MustCompile(`^\s*\|?\s*|\s*$`).
		ReplaceAllString(FeatureGroupContent(value), "")
}

// Categories returns the list of categories (rune time, compile, spec) that define whether a
// feature has been implemented
func Categories(line string) []string {
	rawValues := strings.Split(line, " | ")
	valuesToReturn := make([]string, 0)
	for i := range rawValues {
		rawValues[i] = categoryValue(rawValues[i])
		if rawValues[i] != "" {
			valuesToReturn = append(valuesToReturn, rawValues[i])
		}

	}
	return valuesToReturn
}
