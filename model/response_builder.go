// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package model

// ResponseBuilder builds the response that will be sent to the client
type ResponseBuilder interface {
	Version(number float64) ResponseBuilder
	Data(data interface{}) ResponseBuilder
	Errors(errors []*Error) ResponseBuilder
	AddError(error *Error) ResponseBuilder
	Links(links map[string]string) ResponseBuilder
	Build() *Response
}

type responseBuilder struct {
	API            *API
	ResponseErrors []*Error
	ResponseData   interface{}
	ResponseLinks  map[string]string
}

func (rb *responseBuilder) Version(number float64) ResponseBuilder {
	rb.API = &API{Version: number}
	return rb
}

func (rb *responseBuilder) Data(data interface{}) ResponseBuilder {
	rb.ResponseData = data
	return rb
}

func (rb *responseBuilder) Errors(errors []*Error) ResponseBuilder {
	rb.ResponseErrors = errors
	return rb
}

func (rb *responseBuilder) AddError(error *Error) ResponseBuilder {
	rb.ResponseErrors = append(rb.ResponseErrors, error)
	return rb
}

func (rb *responseBuilder) Links(links map[string]string) ResponseBuilder {
	rb.ResponseLinks = links
	return rb
}

func (rb *responseBuilder) Build() *Response {
	return &Response{
		API:    rb.API,
		Data:   rb.ResponseData,
		Errors: rb.ResponseErrors,
		Links:  rb.ResponseLinks,
	}
}

// NewResponseBuilder returns a newly allocated Response Builder
func NewResponseBuilder() ResponseBuilder {
	return &responseBuilder{}
}
