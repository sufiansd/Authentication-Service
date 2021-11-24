// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ValidateTokenOKCode is the HTTP code returned for type ValidateTokenOK
const ValidateTokenOKCode int = 200

/*ValidateTokenOK Valid Token

swagger:response validateTokenOK
*/
type ValidateTokenOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewValidateTokenOK creates ValidateTokenOK with default headers values
func NewValidateTokenOK() *ValidateTokenOK {

	return &ValidateTokenOK{}
}

// WithPayload adds the payload to the validate token o k response
func (o *ValidateTokenOK) WithPayload(payload string) *ValidateTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate token o k response
func (o *ValidateTokenOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ValidateTokenUnauthorizedCode is the HTTP code returned for type ValidateTokenUnauthorized
const ValidateTokenUnauthorizedCode int = 401

/*ValidateTokenUnauthorized Invalid Token

swagger:response validateTokenUnauthorized
*/
type ValidateTokenUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewValidateTokenUnauthorized creates ValidateTokenUnauthorized with default headers values
func NewValidateTokenUnauthorized() *ValidateTokenUnauthorized {

	return &ValidateTokenUnauthorized{}
}

// WithPayload adds the payload to the validate token unauthorized response
func (o *ValidateTokenUnauthorized) WithPayload(payload string) *ValidateTokenUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate token unauthorized response
func (o *ValidateTokenUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateTokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ValidateTokenInternalServerErrorCode is the HTTP code returned for type ValidateTokenInternalServerError
const ValidateTokenInternalServerErrorCode int = 500

/*ValidateTokenInternalServerError Internal Server Error

swagger:response validateTokenInternalServerError
*/
type ValidateTokenInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewValidateTokenInternalServerError creates ValidateTokenInternalServerError with default headers values
func NewValidateTokenInternalServerError() *ValidateTokenInternalServerError {

	return &ValidateTokenInternalServerError{}
}

// WithPayload adds the payload to the validate token internal server error response
func (o *ValidateTokenInternalServerError) WithPayload(payload string) *ValidateTokenInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate token internal server error response
func (o *ValidateTokenInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}