// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

// FeatureBuilder builds a group of features e.g. Annotations, Component Composition...
type FeatureBuilder interface {
	Name(name string) FeatureBuilder
	Status(status *FeatureStatus) FeatureBuilder
	Build() *Feature
}

type featureBuilder struct {
	FeatureName   string
	FeatureStatus *FeatureStatus
}

func (fb *featureBuilder) Name(name string) FeatureBuilder {
	fb.FeatureName = name
	return fb
}

func (fb *featureBuilder) Status(status *FeatureStatus) FeatureBuilder {
	fb.FeatureStatus = status
	return fb
}

func (fb *featureBuilder) Build() *Feature {
	return &Feature{
		Name:   fb.FeatureName,
		Status: fb.FeatureStatus,
	}
}

func NewFeatureBuilder() FeatureBuilder {
	return &featureBuilder{}
}

type FeatureStatusBuilder interface {
	GranularStatuses(granularStatuses []*GranularStatus) FeatureStatusBuilder
	Categories(Categories []string) FeatureStatusBuilder
	AddGranularStatus(granularStatus *GranularStatus) FeatureStatusBuilder
	Build() *FeatureStatus
}

type featureStatusBuilder struct {
	FeatureGranularStatuses []*GranularStatus
	FeatureStatusCategories []string
}

func (fb *featureStatusBuilder) GranularStatuses(granularStatuses []*GranularStatus) FeatureStatusBuilder {
	fb.FeatureGranularStatuses = granularStatuses
	return fb
}

func (fb *featureStatusBuilder) Categories(categories []string) FeatureStatusBuilder {
	fb.FeatureStatusCategories = categories
	return fb
}

func (fb *featureStatusBuilder) AddGranularStatus(granularStatus *GranularStatus) FeatureStatusBuilder {
	fb.FeatureGranularStatuses = append(fb.FeatureGranularStatuses, granularStatus)
	return fb
}

func (fb *featureStatusBuilder) Build() *FeatureStatus {
	featureCompleted := true
	for _, status := range fb.FeatureGranularStatuses {
		if status.Code == StatusToString(NOT_IMPLEMENTED) {
			featureCompleted = false
			break
		}
	}
	return &FeatureStatus{
		Completed:        featureCompleted,
		Categories:       fb.FeatureStatusCategories,
		GranularStatuses: fb.FeatureGranularStatuses,
	}
}

func NewFeatureStatusBuilder() FeatureStatusBuilder {
	return &featureStatusBuilder{}
}

// GranularStatusBuilder allows to build fro example {Category: "Compiler", Code: "IMPLEMENTED"}
type GranularStatusBuilder interface {
	Category(category string) GranularStatusBuilder
	Code(code string) GranularStatusBuilder
	Description(description string) GranularStatusBuilder
	Build() *GranularStatus
}

type granularStatusBuilder struct {
	GranularCategory    string
	GranularCode        string
	GranularDescription string
}

func (gb *granularStatusBuilder) Category(category string) GranularStatusBuilder {
	gb.GranularCategory = category
	return gb
}

func (gb *granularStatusBuilder) Code(code string) GranularStatusBuilder {
	gb.GranularCode = code
	return gb
}

func (gb *granularStatusBuilder) Description(description string) GranularStatusBuilder {
	gb.GranularDescription = description
	return gb
}

func (gb *granularStatusBuilder) Build() *GranularStatus {
	return &GranularStatus{
		Category:    gb.GranularCategory,
		Code:        gb.GranularCode,
		Description: gb.GranularDescription,
	}
}

func NewGranularStatusBuilder() GranularStatusBuilder {
	return &granularStatusBuilder{}
}
