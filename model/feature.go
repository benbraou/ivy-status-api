// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

// FeatureGroup groups related features together (e.g. Angular Annotations consists of @Component,
// @Directive, @Pipe, @Injectable...)
type FeatureGroup struct {
	Name     string     `json:"name"`
	Features []*Feature `json:"features"`
}

// Feature describes a granular Angular feature (e.g. creation reordering based on injection)
type Feature struct {
	Name   string         `json:"name"`
	Status *FeatureStatus `json:"status"`
}

// FeatureStatus describes the status of a feature in regarding to the criteria: runtime, spec,
// compiler
type FeatureStatus struct {
	Completed        bool              `json:"completed"`
	GranularStatuses []*GranularStatus `json:"granularStatuses"`
}

// GranularStatus describes the progress on part of a feature.
// e.g. {Category: 'Run time', Code: 'IMPLEMENTED', Description ''}
type GranularStatus struct {
	Category    string `json:"category"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// StatusCode describes whether something is implemented, not yet implemented or not applicable
type StatusCode int

const (
	IMPLEMENTED StatusCode = iota + 1
	NOT_IMPLEMENTED
	NOT_APPLICABLE
)

func StatusToString(code StatusCode) string {
	switch code {
	case IMPLEMENTED:
		return "IMPLEMENTED"
	case NOT_IMPLEMENTED:
		return "NOT_IMPLEMENTED"
	case NOT_APPLICABLE:
		return "NOT_APPLICABLE"
	default:
		return "NOT_IMPLEMENTED"
	}
}
