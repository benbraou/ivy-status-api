// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"github.com/benbraou/ivy-status-api/model"
	"github.com/benbraou/ivy-status-api/utils/markdown"
)

func populateFromHeader(
	line *string,
	root **model.FeatureGroup,
	featureGroup **model.FeatureGroup,
	features *[]*model.Feature,
	categories *[]string,
	featuresGroupedByLevels *map[int][]*model.Feature) {

	markdown.PrepareLine(line)
	// Before switching to the new feature group, we need to attach all the features that we
	// have built to the previous feature group
	(*featureGroup).Data.Features = *features

	// Fresh start! Now, we have a new feature group
	*features = make([]*model.Feature, 0)
	*featuresGroupedByLevels = make(map[int][]*model.Feature)
	*categories = make([]string, 0)
	*featureGroup = attachToRootAndReturnFeatureGroup(*root, *line)
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
