// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

// Response struct is used for all HTTP responses
type Response struct {
	API    *API     `json:"API"`
	Errors []*Error `json:"errors"`
	Data   `json:"data"`
	Links  map[string]string `json:"links"`
}

// Data corresponds to the successful content
type Data interface{}

// Error element
type Error struct {
	Code        string `json:"code"`
	ErrorSource `json:"source"`
	Title       string `json:"title"`
	Detail      string `json:"deail"`
}

// ErrorSource describes the source of the error
type ErrorSource struct {
	Pointer string `json:"pointer"`
}

// API provides information on the api version
type API struct {
	Version float64 `json:"version"`
}
