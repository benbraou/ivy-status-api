// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"math/big"
	"reflect"
	"testing"
)

func TestResponseBuilder(t *testing.T) {
	builder := NewResponseBuilder()
	content := struct {
		short string
		full  string
	}{
		short: "Golang is cool",
		full:  "Golang is really cool!",
	}
	response := builder.
		Version(1.1).
		AddError(&Error{Code: "some code", Detail: "some detail"}).
		Data(&content).
		Build()

	// Check api version
	apiVersion := big.NewFloat(response.API.Version)
	if apiVersion.Cmp(big.NewFloat(1.1)) != 0 {
		t.Error("Expected version 1.1, got ", apiVersion)
	}

	// Check errors
	if !reflect.DeepEqual(response.Errors,
		[]*Error{&Error{Code: "some code", Detail: "some detail"}}) {
		t.Error("Expected  errors to be correctly build")
	}

	// Check Data
	if !reflect.DeepEqual(response.Data, &struct {
		short string
		full  string
	}{
		short: "Golang is cool",
		full:  "Golang is really cool!",
	}) {
		t.Error("Expected  data to be correctly set")
	}
}
