// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE
package business

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/benbraou/ivy-status-api/model"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Assertions will look better with testify (https://github.com/benbraou/ivy-status-api/issues/1)
func TestProduceFeatureGroups(t *testing.T) {
	pwd, _ := os.Getwd()

	markdown, err := ioutil.ReadFile(pwd + "/../mocks/status_03_18_2018.md")
	checkError(err)

	var expectedFeatureGroups []*model.FeatureGroup
	var expectedBytes []byte
	expectedBytes, err = ioutil.ReadFile(pwd + "/../mocks/feature_groups_03_18_2018.json")
	checkError(err)

	json.Unmarshal(expectedBytes, &expectedFeatureGroups)

	producedFeatureGroups := ProduceFeatureGroups(string(markdown))

	if len(producedFeatureGroups) != len(expectedFeatureGroups) {
		t.Error(
			"Expected to have ",
			len(expectedFeatureGroups),
			"groups, but got ", len(producedFeatureGroups),
		)
	}

	for i, producedFeatureGroup := range producedFeatureGroups {
		assertFeatureGroup(t, producedFeatureGroup, expectedFeatureGroups[i])
	}
}

func assertFeatureGroup(
	t *testing.T,
	producedFeatureGroup *model.FeatureGroup,
	expectedFeatureGroup *model.FeatureGroup,
) {
	if producedFeatureGroup.Name != expectedFeatureGroup.Name {
		t.Error(
			"Expected group to have the name: ",
			expectedFeatureGroup.Name,
			"but got: ",
			producedFeatureGroup.Name,
		)
	}

	for i, feature := range producedFeatureGroup.Features {
		assertFeature(t, feature, expectedFeatureGroup.Features[i])
	}
}

func assertFeature(
	t *testing.T,
	producedFeature *model.Feature,
	expectedFeature *model.Feature,
) {
	if producedFeature.Name != expectedFeature.Name {
		t.Error(
			"Expected feature to have the name: ",
			expectedFeature.Name,
			"but got: ",
			producedFeature.Name,
		)
	}

	if producedFeature.Status.Completed != expectedFeature.Status.Completed {
		t.Error(
			"Expected feature completion to be: ",
			expectedFeature.Status.Completed,
			"but got: ",
			producedFeature.Status.Completed,
		)
	}

	for i, granularStatus := range producedFeature.Status.GranularStatuses {
		assertGranularStatus(t, granularStatus, expectedFeature.Status.GranularStatuses[i])
	}
}

func assertGranularStatus(
	t *testing.T,
	granularStatus *model.GranularStatus,
	expectedGranularStatus *model.GranularStatus,
) {
	if granularStatus.Code != expectedGranularStatus.Code {
		t.Error(
			"Expected granular status to be: ",
			expectedGranularStatus.Code,
			"but got: ",
			granularStatus.Code,
		)
	}
}
