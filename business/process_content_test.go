// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package business

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/benbraou/ivy-status-api/model"
	"github.com/stretchr/testify/assert"
)

func testProcessHeadersOnly(t *testing.T) {
	content := `# Overview
	# Implementation Status
	## @angular/compiler-cli changes
	### ngtsc TSC compiler transformer
	### ngcc Angular node_module compatibility compiler
	## @angular/compiler changes
	## @angular/core changes
	`
	root := model.NewRootFeatureGroup()
	process(content, root)

	implementationStatus := CheckAndReturnTopFeatureGroup(t, root, "Implementation Status")
	overview := CheckAndReturnTopFeatureGroup(t, root, "Overview")

	assert.True(t, true, overview.ChildrenStack.Empty())

	core := CheckAndReturnTopFeatureGroup(t, implementationStatus, "@angular/core changes")
	compiler := CheckAndReturnTopFeatureGroup(t, implementationStatus, "@angular/compiler changes")
	compilerCli := CheckAndReturnTopFeatureGroup(t, implementationStatus, "@angular/compiler-cli changes")

	assert.True(t, true, core.ChildrenStack.Empty())
	assert.True(t, true, compiler.ChildrenStack.Empty())

	CheckAndReturnTopFeatureGroup(t, compilerCli, "ngcc Angular node_module compatibility compiler")
	CheckAndReturnTopFeatureGroup(t, compilerCli, "ngtsc TSC compiler transformer")
}

func CheckAndReturnTopFeatureGroup(
	t *testing.T,
	root *model.FeatureGroup,
	name string,
) *model.FeatureGroup {
	var topFg *model.FeatureGroup
	assert.NotNil(t, root.ChildrenStack)
	top, err := root.ChildrenStack.Pop()
	assert.Nil(t, err)
	topFg = top.(*model.FeatureGroup)
	assert.Equal(t, name, topFg.Data.Name)
	return topFg
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestProduceFeatureGroups(t *testing.T) {
	testProduceFeatureGroupsUsingMockDate(t, "04_01_2018")
}

func testProduceFeatureGroupsUsingMockDate(t *testing.T, mockDate string) {
	pwd, _ := os.Getwd()

	markdown, err := ioutil.ReadFile(fmt.Sprintf("%s/../mocks/status_%s.md", pwd, mockDate))
	checkError(err)

	var expectedIvyStatus *model.FeatureGroup
	var expectedBytes []byte

	expectedBytes, err = ioutil.ReadFile(
		fmt.Sprintf("%s/../mocks/feature_groups_%s.json", pwd, mockDate),
	)
	checkError(err)
	json.Unmarshal(expectedBytes, &expectedIvyStatus)
	ivyStatus := ProduceIvyStatus(string(markdown))

	compareFeatureGroups(t, ivyStatus, expectedIvyStatus)
}

func compareFeatureGroups(
	t *testing.T,
	ivyStatus *model.FeatureGroup,
	expectedIvyStatus *model.FeatureGroup) {

	if len(ivyStatus.FeatureGroups) != len(expectedIvyStatus.FeatureGroups) {
		t.Error(
			"Expected to have ",
			len(expectedIvyStatus.FeatureGroups),
			"groups, but got ", len(ivyStatus.FeatureGroups),
		)
	}

	assert.Equal(t, expectedIvyStatus.Data.Name, ivyStatus.Data.Name)
	compareFeatures(t, expectedIvyStatus.Data.Features, ivyStatus.Data.Features)

	for i, fg := range ivyStatus.FeatureGroups {
		compareFeatureGroups(t, expectedIvyStatus.FeatureGroups[i], fg)
	}
}

func compareFeatures(
	t *testing.T,
	features []*model.Feature,
	expectedFeatures []*model.Feature) {
	for i, feature := range features {
		assert.Equal(t, feature.Name, expectedFeatures[i].Name)
		compareStatuses(t, feature.Status, expectedFeatures[i].Status)
	}
}

func compareStatuses(
	t *testing.T,
	status *model.FeatureStatus,
	expectedStatus *model.FeatureStatus) {
	assert.Equal(t, status.Completed, expectedStatus.Completed)
	assert.Equal(t, status.GranularStatuses, expectedStatus.GranularStatuses)
}
