// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"github.com/benbraou/ivy-status-api/model"
	"github.com/benbraou/ivy-status-api/utils/markdown"
)

func populateSingleStatusFeature(
	line *string,
	features *[]*model.Feature,
	featuresGroupedByLevels *map[int][]*model.Feature) {

	level, e := markdown.UnorderedListIndentationLevel(*line)
	markdown.PrepareLine(line)

	icon, featureName := markdown.GranularStatusIconAndDescription(*line)

	if e == nil {
		feature := model.
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
			Build()
		if level > 1 {
			// attach to the right parent
			levelToAttach := level - 2
			abc := *featuresGroupedByLevels
			abc[levelToAttach][len(abc[levelToAttach])-1].ChildFeatures = append(
				abc[levelToAttach][len(abc[levelToAttach])-1].ChildFeatures,
				feature,
			)
		} else {
			*features = append(*features, feature)
		}
		(*featuresGroupedByLevels)[level] = append((*featuresGroupedByLevels)[level], feature)
	}
}

func populateTableRowFeature(
	line *string,
	features *[]*model.Feature,
	categories *[]string) {

	markdown.PrepareLine(line)
	// We have several statuses. Each status is attached to one category
	featureBuilder := model.NewFeatureBuilder()
	featureStatusBuilder := model.NewFeatureStatusBuilder()
	statuses := markdown.TableCells(*line)

	for i, status := range statuses {
		if i == 0 {
			featureBuilder.Name(status)
		} else {
			icon, cat := markdown.GranularStatusIconAndDescription(status)
			if len(cat) == 0 {
				cat = (*categories)[i]
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
	*features = append(*features, featureBuilder.Build())
}
