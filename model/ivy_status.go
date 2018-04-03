// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

type IvyStatus struct {
	FeatureGroup   *FeatureGroup `json:"featureGroup"`
	LastUpdateDate string        `json:"lastUpdateDate"`
}
