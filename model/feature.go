// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import "github.com/benbraou/ivy-status-api/constants"

// FeatureGroupData holds information about:
// 1. the feature group name e.g. Angular Annotations
// 2. the related features e.g. (@Component, @Directive, @Pipe, @Injectable...)
type FeatureGroupData struct {
	Name     string     `json:"name"`
	Features []*Feature `json:"features"`
}

// FeatureGroup can hold child feature groups. It is in reality a tree. e.g.
//                 [Implementation Status]
//                          / \
//                         /   \
//                        /     \
//                       /       \
//     [`@angular/compiler-cli`  [`@angular/compiler` changes]
//                changes]
//                  /
//                 /
//   [`ngtsc` TSC compiler transformer]
// Internally, nested feature groups will be stored in a Stack data structure. Before sending the
// response to the client, a slice of feature groups will be created based on this Stack
// Why child feature groups is handled by a stack ? Let's look at this example:
//
// H1 AAA
// H2 BBB
// H2 CCC
// H3 DDD
// The children of H1 (BBB and CCC) are stored in a Stack
// When we encounter H3 DDD, we need to nest it to the parent H2 CCC and not H2 BBB. Easy, we just
// peek to the top of H1 children Stack. It's H2 CCC: the last H2 added
type FeatureGroup struct {
	Data          *FeatureGroupData `json:"data"`
	ChildrenStack *Stack            `json:"-"`
	FeatureGroups []*FeatureGroup   `json:"featureGroups"`
}

func NewRootFeatureGroup() *FeatureGroup {
	return &FeatureGroup{
		ChildrenStack: NewStack(),

		Data: &FeatureGroupData{Name: constants.RootFeatureGroupName},
	}
}

// Feature describes a granular Angular feature (e.g. creation reordering based on injection)
type Feature struct {
	Name          string         `json:"name"`
	Status        *FeatureStatus `json:"status"`
	ChildFeatures []*Feature     `json:"childFeatures"`
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
