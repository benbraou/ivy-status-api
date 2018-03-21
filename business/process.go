// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"strings"

	"github.com/benbraou/ivy-status-api/model"
)

// ProduceFeatureGroups builds the ivy renderer implementation status
func ProduceFeatureGroups(content string) []*model.FeatureGroup {
	featureGroups := make([]*model.FeatureGroup, 0)
	populateFeatureGroups(lines(content), &featureGroups)
	return featureGroups
}

// Builds feature groups from the lines that make the file `STATUS.md`
// https://raw.githubusercontent.com/angular/angular/master/packages/core/src/render3/STATUS.md
// e.g.
// ## Bootstrap API
// | Feature                             | Runtime |
// | ----------------------------------- | ------- |
// | `renderComponent()`                 |  ✅     |
// | `getHostElement()`                  |  ✅     |
// | `createInjector()`                  |  ❌     |
//
// ## I18N
// | Feature                             | Runtime | Spec     | Compiler |
// | ----------------------------------- | ------- | -------- | -------- |
// | translate text literals             |  ❌     |  ❌      |  ❌      |
// | rearrange text nodes                |  ❌     |  ❌      |  ❌      |
// | ICU                                 |  ❌     |  ❌      |  ❌      |
func populateFeatureGroups(lines []string, featureGroups *[]*model.FeatureGroup) {
	featureGroupIndices := FeatureGroupIndices(lines)
	if len(featureGroupIndices) == 0 {
		return
	}

	for i := range featureGroupIndices {
		featureGroup := buildFeatureGroup(FeatureGroupLines(lines, i, featureGroupIndices))
		if featureGroup != nil {
			*featureGroups = append(*featureGroups, featureGroup)
		}
	}

}

// Given a slice that holds information about the implementation status of a group of features, we
// build the corresponding FeatureGroup
func buildFeatureGroup(lines []string) *model.FeatureGroup {
	err := ValidateLines(lines)
	if err != nil {
		return nil
	}

	fgb := model.
		NewFeatureGroupBuilder().
		// The first line always contains the feature group name e.g.
		// ## Component Composition
		Name(FeatureGroupName(lines[0]))

	// The rest of the lines contain information about the related features. e.g.
	// | Feature                                  | Runtime | Spec     | Compiler |
	// | ---------------------------------------- | ------- | -------- | -------- |
	// | creation reordering based on injection   |   ❌    |    ❌    |    ✅    |
	// | `class CompA extends CompB {}`           |   ❌    |    ❌    |    ❌    |
	// | `class CompA extends CompB { @Input }`   |   ❌    |    ❌    |    ❌    |
	// | `class CompA extends CompB { @Output }`  |   ❌    |    ❌    |    ❌    |
	lines = lines[1:]

	// Handle the case when only a feature group is found but no related features (This may happen
	// when `STATUS.md` is in an incomplete state ?)
	if !HasFeatureStatusInfo(lines) {
		return fgb.Build()
	}

	// At this point, we know that we have a list of features. But, do we have a category header?
	// In case, there is none, we create a default empty list of categories
	// e.g.
	// ## Missing Pieces
	// - Sanitization ✅
	// - Back patching in tree shakable way. ❌
	// - attribute namespace ❌

	hasCategoryHeader := IsCategoryHeader(lines[0])

	categories := Categories(lines[0])

	if !hasCategoryHeader {
		categories = make([]string, len(categories))
	}

	if hasCategoryHeader {
		lines = lines[1:]
	}

	populateFeatures(fgb, categories, lines)

	return fgb.Build()
}

// Populates a feature group with the implentation status of the included features. This will be
// done based on the information provided in the corresponding markdown liness
func populateFeatures(fgb model.FeatureGroupBuilder, categories []string, lines []string) {
	err := ValidateLines(lines)
	if err != nil {
		return
	}

	for _, line := range lines {
		statuses := Categories(line)
		if len(statuses) == 1 {
			// special case: the feature name and status are provided one shot (e.g. Sanitization ✅)
			icon, featureName := GranularStatusIconAndDescription(line)
			fgb.AddFeature(
				model.
					NewFeatureBuilder().
					Name(featureName).
					Status(model.NewFeatureStatusBuilder().
						AddGranularStatus(model.NewGranularStatusBuilder().
							Code(
								IconToStatusCode(icon),
							).
							Build(),
						).
						Build(),
					).
					Build(),
			)
		} else {
			featureBuilder := model.NewFeatureBuilder()
			featureStatusBuilder := model.NewFeatureStatusBuilder()
			for i, status := range statuses {
				if i == 0 {
					featureBuilder.Name(status)
				} else {
					icon, _ := GranularStatusIconAndDescription(status)
					featureStatusBuilder.AddGranularStatus(
						model.NewGranularStatusBuilder().
							Code(IconToStatusCode(icon)).
							Category(categories[i]).
							Build(),
					)
				}
			}
			featureBuilder.Status(featureStatusBuilder.Build())
			fgb.AddFeature(featureBuilder.Build())
		}
	}
}

func lines(markdownContent string) []string {
	lines := make([]string, 0)
	for _, line := range strings.Split(markdownContent, "\n") {
		if !IsSeperator(line) {
			lines = append(lines, FeatureGroupContent(line))
		}
	}
	return lines
}
