// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

type IvyStatusBuilder interface {
	RootFeatureGroup(root *FeatureGroup) IvyStatusBuilder
	LastUpdateDate(lastUpdateDate string) IvyStatusBuilder
	Build() *IvyStatus
}

type ivyStatusBuilder struct {
	Root       *FeatureGroup
	UpdateDate string
}

func (ib *ivyStatusBuilder) RootFeatureGroup(root *FeatureGroup) IvyStatusBuilder {
	ib.Root = root
	return ib
}

func (ib *ivyStatusBuilder) LastUpdateDate(lastUpdateDate string) IvyStatusBuilder {
	ib.UpdateDate = lastUpdateDate
	return ib
}

func (ib *ivyStatusBuilder) Build() *IvyStatus {
	return &IvyStatus{
		FeatureGroup:   ib.Root,
		LastUpdateDate: ib.UpdateDate,
	}
}

func NewIvyStatusBuilder() IvyStatusBuilder {
	return &ivyStatusBuilder{}
}
