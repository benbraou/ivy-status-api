// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{}

	assertEmpty(true, s, t)

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected empty Stack to return an error when Pop is called, but it did not")
	}

	s.Push(10)

	assertEmpty(false, s, t)

	s.Push(13)

	top, _ := s.Pop()
	assertTopValue(top, 13, t)
	top, _ = s.Pop()
	assertTopValue(top, 10, t)

}

func assertTopValue(top interface{}, value int, t *testing.T) {
	if top != value {
		t.Error(fmt.Sprintf("Expected the top stack element to be %d, but it was %d", value, top))
	}
}

func assertEmpty(empty bool, s *Stack, t *testing.T) {
	msg := "Expected the stack to empty but it was not"
	if !empty {
		msg = "Expected the stack not to be empty but it was"
	}
	if s.Empty() != empty {
		t.Error(msg)
	}
}
