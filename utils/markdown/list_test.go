// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnorderedListIndentationLevel(t *testing.T) {

	l, e := UnorderedListIndentationLevel("abc")
	assert.NotNil(t, e)
	assert.Equal(t, -1, l)

	l, e = UnorderedListIndentationLevel("- abc")
	assert.Nil(t, e)
	assert.Equal(t, 1, l)

	l, e = UnorderedListIndentationLevel("  - abc")
	assert.Nil(t, e)
	assert.Equal(t, 3, l)

	l, e = UnorderedListIndentationLevel("* abc")
	assert.Nil(t, e)
	assert.Equal(t, 1, l)

	l, e = UnorderedListIndentationLevel("  * abc")
	assert.Nil(t, e)
	assert.Equal(t, 3, l)
}
