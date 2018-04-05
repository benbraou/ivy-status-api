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
	lines := markdown.Lines(content)

	// Add a fake line
	lines = append(lines, "")

	currentFeatureGroup := root
	currentFeatures := make([]*model.Feature, 0)
	currentCategories := make([]string, 0)

	for lineIndex, line := range lines {
		isHeader := markdown.IsLineHeader(line)
		hasFeatureStatusInfo := markdown.IsFeatureStatusLine(line)
		hasCategoriesInfo := markdown.IsCategoryHeaderLine(line)

		// At this point, we are parsing the fake empty line that we have added.
		// Append the features that we have built to the current feature group, and
		if lineIndex == len(lines)-1 {
			currentFeatureGroup.Data.Features = currentFeatures
			break
		}

		if isHeader {
			// Before switching to the new feature group, we need to attach all the features that we
			// have built to the previous feature group
			currentFeatureGroup.Data.Features = currentFeatures

			// Now reset features for the next feature group
			currentFeatures = make([]*model.Feature, 0)

			currentCategories = make([]string, 0)

			// Fresh start! Now, we have a new feature group
			currentFeatureGroup = attachToRootAndReturnFeatureGroup(root, line)
			continue
		}

		if !hasFeatureStatusInfo && !hasCategoriesInfo {
			continue
		}

		if hasCategoriesInfo {
			currentCategories = markdown.TableCells(line)
		} else {
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
			if markdown.IsSingleStatusLine(line) && len(currentCategories) == 0 {
				icon, featureName := markdown.GranularStatusIconAndDescription(line)
				currentFeatures = append(
					currentFeatures,
					model.
						NewFeatureBuilder().
						Name(featureName).
						Status(
							model.
								NewFeatureStatusBuilder().
								AddGranularStatus(
									model.NewGranularStatusBuilder().
										Code(markdown.IconToStatusCode(icon)).
										Build(),
								).
								Build(),
						).
						Build(),
				)
			} else {
				// We have several statuses. Each status is attached to one category
				featureBuilder := model.NewFeatureBuilder()
				featureStatusBuilder := model.NewFeatureStatusBuilder()
				statuses := markdown.TableCells(line)

				for i, status := range statuses {
					if i == 0 {
						featureBuilder.Name(status)
					} else {
						icon, des := markdown.GranularStatusIconAndDescription(status)
						cat := des
						if len(cat) == 0 {
							cat = currentCategories[i]
						}
						featureStatusBuilder.AddGranularStatus(
							model.NewGranularStatusBuilder().
								Code(markdown.IconToStatusCode(icon)).
								Category(cat).
								Build(),
						)
					}
				}
				featureBuilder.Status(featureStatusBuilder.Build())
				currentFeatures = append(currentFeatures, featureBuilder.Build())
			}
		}
	}
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

func attachToRootAndReturnFeatureGroup(root *model.FeatureGroup, line string) *model.FeatureGroup {
	level, _ := markdown.HeaderLevel(line)
	name := markdown.HeaderName(line)
	fg := &model.FeatureGroup{
		ChildrenStack: model.NewStack(),
		Data:          &model.FeatureGroupData{Name: name},
	}
	parentFeatureGroupNode(level, root).ChildrenStack.Push(fg)
	return fg
}

func parentFeatureGroupNode(level int, root *model.FeatureGroup) *model.FeatureGroup {
	parent := root
	for i := 1; i < level; i++ {
		top, _ := parent.ChildrenStack.Peek()
		parent = top.(*model.FeatureGroup)
	}
	return parent
}
