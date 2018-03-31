// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFeatureStatusLine(t *testing.T) {
	assert.True(t, IsFeatureStatusLine("  - ❌ `@Directive`"))
	assert.True(t, IsFeatureStatusLine("✅ `defineComponent()`  "))
	assert.True(t, IsFeatureStatusLine(" n/a `defineComponent()`  "))
	assert.True(t, IsFeatureStatusLine(" N/a `defineComponent()`  "))
}

func TestIsSingleStatusLine(t *testing.T) {
	assert.True(t, IsSingleStatusLine("  - ❌ `@Directive`"))
	assert.False(t, IsSingleStatusLine("  - ❌ `@Directive` ✅ `defineComponent()`"))
}

func TestGranularStatusIconAndDescription(t *testing.T) {
	assert.True(t, IsSingleStatusLine("  - ❌ `@Directive`"))
	assert.False(t, IsSingleStatusLine("  - ❌ `@Directive` ✅ `defineComponent()`"))
}

func TestIconToStatusCode(t *testing.T) {
	assert.Equal(t, "NOT_IMPLEMENTED", IconToStatusCode("❌"))
	assert.Equal(t, "IMPLEMENTED", IconToStatusCode("✅"))
	assert.Equal(t, "NOT_APPLICABLE", IconToStatusCode("n/a"))
}
