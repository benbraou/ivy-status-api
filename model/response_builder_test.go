// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, big.NewFloat(1.1), big.NewFloat(response.API.Version))

	// Check errors
	assert.Equal(t, []*Error{&Error{Code: "some code", Detail: "some detail"}}, response.Errors)

	// Check Data
	assert.Equal(t, &struct {
		short string
		full  string
	}{
		short: "Golang is cool",
		full:  "Golang is really cool!",
	}, response.Data)

}
