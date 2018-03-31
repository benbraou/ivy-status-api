// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCells(t *testing.T) {
	tr := "| Annotation     | `defineXXX()`  | Run time | Spec     | Compiler | Back Patch |"
	assert.Equal(
		t,
		[]string{
			"Annotation",
			"`defineXXX()`",
			"Run time",
			"Spec",
			"Compiler",
			"Back Patch",
		},
		TableCells(tr),
	)

	tr = "| `{{ exp \\| pipe: arg }}`                |  ✅     |  ✅      |  ✅      |"
	assert.Equal(
		t,
		[]string{
			"`{{ exp \\| pipe: arg }}`",
			"✅",
			"✅",
			"✅",
		},
		TableCells(tr),
	)

}
