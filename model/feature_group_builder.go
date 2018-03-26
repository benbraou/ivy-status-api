// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

// FeatureGroupBuilder builds a group of features e.g. Annotations, Component Composition...
type FeatureGroupBuilder interface {
	Name(name string) FeatureGroupBuilder
	Features(features []*Feature) FeatureGroupBuilder
	AddFeature(feature *Feature) FeatureGroupBuilder
	Build() *FeatureGroup
}

type featureGroupBuilder struct {
	GroupName     string
	GroupFeatures []*Feature
}

func (fb *featureGroupBuilder) Name(name string) FeatureGroupBuilder {
	fb.GroupName = name
	return fb
}

func (fb *featureGroupBuilder) Features(features []*Feature) FeatureGroupBuilder {
	fb.GroupFeatures = features
	return fb
}

func (fb *featureGroupBuilder) AddFeature(feature *Feature) FeatureGroupBuilder {
	fb.GroupFeatures = append(fb.GroupFeatures, feature)
	return fb
}

func (fb *featureGroupBuilder) Build() *FeatureGroup {
	return &FeatureGroup{
		Data: &FeatureGroupData{
			Name:     fb.GroupName,
			Features: fb.GroupFeatures,
		},
	}
}

// newFeatureGroupBuilder returns a newly allocated Response Builder
func NewFeatureGroupBuilder() FeatureGroupBuilder {
	return &featureGroupBuilder{}
}
