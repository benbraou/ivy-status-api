// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLineSeperator(t *testing.T) {
	assert.True(t, IsLineSeperator("------------"))
	assert.False(t, IsLineSeperator("-------test-----"))
}

func TestRawLines(t *testing.T) {
	assert.Equal(t, []string{"abc     ", "def", "ghi"}, RawLines("abc     \ndef\nghi"))
}

func TestPrepareLine(t *testing.T) {
	line := "   I am a line     "
	PrepareLine(&line)
	assert.Equal(t, "I am a line", line)

	line = "  - I am a line     "
	PrepareLine(&line)
	assert.Equal(t, "I am a line", line)

	line = "  * I am a line     "
	PrepareLine(&line)
	assert.Equal(t, "I am a line", line)
}

func TestIsCategoryHeaderLine(t *testing.T) {
	assert.True(t, IsCategoryHeaderLine("abc | def | ghi"))
	assert.False(t, IsCategoryHeaderLine(" - ❌ `@Directive`"))
}
