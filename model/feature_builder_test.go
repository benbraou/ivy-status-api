// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"reflect"
	"testing"
)

// GranularStatusBuilder tests
func TestGranularStatusBuilder(t *testing.T) {
	builder := NewGranularStatusBuilder()
	status := builder.
		Code("IMPLEMENTED").
		Category("Spec").
		Description("extra description").
		Build()

	if status.Code != "IMPLEMENTED" {
		t.Error("Expected code: IMPLEMENTED, got ", status.Code)
	}

	if status.Category != "Spec" {
		t.Error("Expected category: Spec, got ", status.Category)
	}

	if status.Description != "extra description" {
		t.Error("Expected description: extra description, got ", status.Description)
	}
}

// FeatureStatusBuilder tests
func TestFeatureStatusBuilderNotCompleteFeature(t *testing.T) {
	builder := NewFeatureStatusBuilder()
	featureStatus := builder.
		AddGranularStatus(
			NewGranularStatusBuilder().
				Code("NOT_IMPLEMENTED").
				Category("Spec").
				Description("extra description").
				Build(),
		).
		Build()

	if featureStatus.Completed == true {
		t.Error("Expected feature status to be not completed")
	}

	if len(featureStatus.GranularStatuses) != 1 {
		t.Error("Expected feature status to have 1 granular status, got ",
			len(featureStatus.GranularStatuses))
	}

	if !reflect.DeepEqual(featureStatus.GranularStatuses[0], NewGranularStatusBuilder().
		Code("NOT_IMPLEMENTED").
		Category("Spec").
		Description("extra description").
		Build()) {
		t.Error("Feature status does not contain correct granular status")
	}
}

func TestFeatureStatusBuilderCompleteFeature(t *testing.T) {
	builder := NewFeatureStatusBuilder()
	featureStatus := builder.
		AddGranularStatus(
			NewGranularStatusBuilder().
				Code("IMPLEMENTED").
				Category("Spec").
				Description("extra description").
				Build(),
		).
		AddGranularStatus(
			NewGranularStatusBuilder().
				Code("IMPLEMENTED").
				Category("Complier").
				Description("extra description").
				Build(),
		).
		Build()

	if featureStatus.Completed == false {
		t.Error("Expected feature status to be completed")
	}

	if len(featureStatus.GranularStatuses) != 2 {
		t.Error("Expected feature status to have 2 granular status, got ",
			len(featureStatus.GranularStatuses))
	}
}

// FeatureBuilder tests
func TestFeatureBuilder(t *testing.T) {
	builder := NewFeatureBuilder()
	feature := builder.
		Name("defineDirective()").
		Status(
			NewFeatureStatusBuilder().AddGranularStatus(
				NewGranularStatusBuilder().
					Code("IMPLEMENTED").
					Category("Spec").
					Description("extra description").
					Build(),
			).Build()).
		Build()

	if feature.Name != "defineDirective()" {
		t.Error("Expected feature name: defineDirective(), got ", feature.Name)
	}

	if !reflect.DeepEqual(feature.Status, NewFeatureStatusBuilder().AddGranularStatus(
		NewGranularStatusBuilder().
			Code("IMPLEMENTED").
			Category("Spec").
			Description("extra description").
			Build(),
	).Build()) {
		t.Error("Feature does not contain the correct status")
	}
}
