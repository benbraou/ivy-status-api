// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"github.com/benbraou/ivy-status-api/model"
	"github.com/benbraou/ivy-status-api/utils/markdown"
)

// ProduceIvyStatus builds the ivy renderer implementation status
func ProduceIvyStatus(content string) *model.FeatureGroup {
	root := model.NewRootFeatureGroup()
	process(content, root)
	postprocess(root)
	return root
}

func process(content string, root *model.FeatureGroup) {
	rawLines := markdown.RawLines(content)
	currentFeatureGroup := root
	currentFeatures := make([]*model.Feature, 0)
	// currentCategories represents the criteria according to which the status of a feature is
	// tracked. e.g. [ Annotation, `defineXXX()`, Run time, Spec,  Compiler, Back Patch ]
	currentCategories := make([]string, 0)
	// Group the single line features together by their level of indentation
	// e.g. the following features will be grouped into three separate levels:
	// - ❌ Rewrite existing code by interpreting the associated STORING_METADATA_IN_D.TS
	//   - ❌ `ComponentCompiler`: `@Component` => `defineComponent`
	//     - ❌ `TemplateCompiler`
	featuresGroupedByLevels := make(map[int][]*model.Feature)

	for _, line := range rawLines {
		processRawLine(
			&line,
			&root,
			&currentFeatureGroup,
			&currentFeatures,
			&currentCategories,
			&featuresGroupedByLevels,
		)
	}
	currentFeatureGroup.Data.Features = currentFeatures
}

func processRawLine(
	line *string,
	root **model.FeatureGroup,
	featureGroup **model.FeatureGroup,
	features *[]*model.Feature,
	categories *[]string,
	featuresGroupedByLevels *map[int][]*model.Feature,
) {

	hasFeatureStatusInfo := markdown.IsFeatureStatusLine(*line)
	hasCategoriesInfo := markdown.IsCategoryHeaderLine(*line)

	if markdown.IsLineHeader(*line) {
		populateFromHeader(line, root, featureGroup, features, categories, featuresGroupedByLevels)
		return
	}

	if !hasFeatureStatusInfo && !hasCategoriesInfo {
		return
	}

	if hasCategoriesInfo {
		markdown.PrepareLine(line)
		*categories = markdown.TableCells(*line)
		return
	}

	// We have a feature status info line. Check whether:
	// 1- it contains one single status,
	// or it has several statuses: in that case, it is attached to the categories built
	// previously

	// Single status line does not mean:
	// | Feature                             | Runtime |
	// | ----------------------------------- | ------- |
	// | `markDirty()`                       |  ✅     |
	// It means instead that it is not preceded by a category line: e.g.
	// ### `ngcc` Angular `node_module` compatibility compiler
	// A tool which "upgrades" `node_module` compiled with non-ivy `ngc` into ivy compliant format.
	// - ❌ Basic setup of stand alone executable
	if markdown.IsSingleStatusLine(*line) && len(*categories) == 0 {
		populateSingleStatusFeature(line, features, featuresGroupedByLevels)
		return
	}

	populateTableRowFeature(line, features, categories)

}

func postprocess(root *model.FeatureGroup) {

	for !root.ChildrenStack.Empty() {
		top, _ := root.ChildrenStack.Pop()
		topFg := top.(*model.FeatureGroup)

		if !topFg.ChildrenStack.Empty() || len(topFg.Data.Features) > 0 {
			root.FeatureGroups = append(root.FeatureGroups, topFg)
			postprocess(topFg)
		}
	}
	for i := len(root.FeatureGroups)/2 - 1; i >= 0; i-- {
		opp := len(root.FeatureGroups) - 1 - i
		root.FeatureGroups[i], root.FeatureGroups[opp] = root.FeatureGroups[opp], root.FeatureGroups[i]
	}
}
