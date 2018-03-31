// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderName(t *testing.T) {
	line := "## `@angular/compiler` changes"
	assert.Equal(t, "`@angular/compiler` changes", HeaderName(line))
	line = "## Decorators"
	assert.Equal(t, "Decorators", HeaderName(line))
}

func TestHeaderLevel(t *testing.T) {
	line := "## `@angular/compiler` changes"
	l, e := HeaderLevel(line)
	assert.Nil(t, e)
	assert.Equal(t, 2, l)

	line = "Fake header"
	_, e = HeaderLevel(line)
	assert.NotNil(t, e)
}

func TestIsLineHeader(t *testing.T) {
	assert.True(t, IsLineHeader("## `@angular/compiler` changes"))
	assert.False(t, IsLineHeader("`@angular/compiler` changes"))
}
