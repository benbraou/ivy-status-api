// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{}

	if !s.Empty() {
		t.Error("Expected the stack to empty but it was not")
	}

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected empty Stack to return an error when Pop is called, but it did not")
	}

	s.Push(10)

	if s.Empty() {
		t.Error("Expected the stack not to be empty but it was")
	}

	s.Push(13)

	top, _ := s.Pop()

	if top != 13 {
		t.Error("Expected the top stack element to be 13, but it was ", top)
	}

}
