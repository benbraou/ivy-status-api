// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

import (
	"errors"
)

// Stack simplified implementation (at least for now) using linked list concept
// The following operations will be supported:
//
// | Operation | Description                                               | Time Complexity |
// |-----------|-----------------------------------------------------------|-----------------|
// |Pop        |Removes and returns the top element of the stack           | O(1)            |
// |Push       |Pushes an element to the top of stack                      | O(1)            |
// |Peek       |Returns the top element from the stack without removing it | O(1)            |
// |Empty      |Returns whether is empty                                   | O(1)            |
//
// For the time being, I don't need to implement search (O(n))
type Stack struct {
	top *StackNode
}

// Pop removes and returns the top element of the stack. Executes in constant time
func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("Empty Stack Exception")
	}
	item := s.top.data
	s.top = s.top.next
	return item, nil
}

// Push pushes an element to the top of stack. Executes in constant time
func (s *Stack) Push(data interface{}) {
	newTop := &StackNode{data: data, next: s.top}
	s.top = newTop
}

// Peek returns the top element from the stack without removing it. Executes in constant time
func (s *Stack) Peek() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("Empty Stack Exception")
	}
	return s.top.data, nil
}

// Empty returns whether is empty. Executes in constant time
func (s *Stack) Empty() bool {
	return s.top == nil
}

type StackNode struct {
	data interface{}
	next *StackNode
}
