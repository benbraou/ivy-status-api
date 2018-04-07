// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// GranularStatusBuilder tests
func TestGranularStatusBuilder(t *testing.T) {
	builder := NewGranularStatusBuilder()
	status := builder.
		Code("IMPLEMENTED").
		Category("Spec").
		Description("extra description").
		Build()

	assert.Equal(t, "IMPLEMENTED", status.Code)
	assert.Equal(t, "Spec", status.Category)
	assert.Equal(t, "extra description", status.Description)
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

	assert.False(t, featureStatus.Completed)
	assert.Equal(t, 1, len(featureStatus.GranularStatuses))
	assert.Equal(t, NewGranularStatusBuilder().
		Code("NOT_IMPLEMENTED").
		Category("Spec").
		Description("extra description").
		Build(), featureStatus.GranularStatuses[0])
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

	assert.True(t, featureStatus.Completed)
	assert.Equal(t, 2, len(featureStatus.GranularStatuses))
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

	assert.Equal(t, "defineDirective()", feature.Name)

	assert.Equal(t, NewFeatureStatusBuilder().AddGranularStatus(
		NewGranularStatusBuilder().
			Code("IMPLEMENTED").
			Category("Spec").
			Description("extra description").
			Build(),
	).Build(), feature.Status)
}
