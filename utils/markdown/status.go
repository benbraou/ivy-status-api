// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/benbraou/ivy-status-api/constants"
	"github.com/benbraou/ivy-status-api/model"
)

// IsFeatureStatusLine considers that the given line holds information on the feature status if ✅ ,
// ❌ or n/a is present
func IsFeatureStatusLine(line string) bool {
	return regexp.
		MustCompile(fmt.Sprintf("%s|%s|%s", constants.Cross, constants.Check, constants.Na)).
		MatchString(strings.ToLower(line))
}

// IsSingleStatusLine checks whether a line contains one piece of information regarding the status
// of a feature
func IsSingleStatusLine(line string) bool {
	res := 0
	if strings.Contains(line, constants.Cross) {
		res++
	}
	if strings.Contains(line, constants.Check) {
		res++
	}
	if strings.Contains(line, constants.Na) {
		res++
	}
	return res == 1
}

// GranularStatusIconAndDescription returns icon and description of a subfeature
// e.g. Sanitization ✅ => ( ✅  , Sanitization  )
func GranularStatusIconAndDescription(td string) (string, string) {
	for _, icon := range [3]string{constants.Cross, constants.Check, constants.Na} {
		description := strings.Replace(td, icon, "", 1)
		if description != td {
			return icon, strings.TrimSpace(description)
		}
	}
	return constants.Na, td
}

// IconToStatusCode returns the feature status code (IMPLEMENTED, NOT_IMPLEMENTED, NOT_APPLICABLE)
// to be sent in the response.
// ✅ => IMPLEMENTED
// ❌ =>  NOT_IMPLEMENTED
// n/a => NOT_APPLICABLE
func IconToStatusCode(icon string) string {
	icon = strings.ToLower(icon)
	switch icon {
	case constants.Check:
		return model.StatusToString(model.IMPLEMENTED)
	case constants.Cross:
		return model.StatusToString(model.NOT_IMPLEMENTED)
	case constants.Na:
		return model.StatusToString(model.NOT_APPLICABLE)
	default:
		return model.StatusToString(model.NOT_IMPLEMENTED)
	}
}
