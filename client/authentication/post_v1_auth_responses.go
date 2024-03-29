// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// PostV1AuthReader is a Reader for the PostV1Auth structure.
type PostV1AuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1AuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1AuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostV1AuthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/auth] PostV1Auth", response, response.Code())
	}
}

// NewPostV1AuthOK creates a PostV1AuthOK with default headers values
func NewPostV1AuthOK() *PostV1AuthOK {
	return &PostV1AuthOK{}
}

/*
PostV1AuthOK describes a response with status code 200, with default header values.

Success
*/
type PostV1AuthOK struct {
	Payload *models.RESTTokenData
}

// IsSuccess returns true when this post v1 auth o k response has a 2xx status code
func (o *PostV1AuthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 auth o k response has a 3xx status code
func (o *PostV1AuthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 auth o k response has a 4xx status code
func (o *PostV1AuthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 auth o k response has a 5xx status code
func (o *PostV1AuthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 auth o k response a status code equal to that given
func (o *PostV1AuthOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 auth o k response
func (o *PostV1AuthOK) Code() int {
	return 200
}

func (o *PostV1AuthOK) Error() string {
	return fmt.Sprintf("[POST /v1/auth][%d] postV1AuthOK  %+v", 200, o.Payload)
}

func (o *PostV1AuthOK) String() string {
	return fmt.Sprintf("[POST /v1/auth][%d] postV1AuthOK  %+v", 200, o.Payload)
}

func (o *PostV1AuthOK) GetPayload() *models.RESTTokenData {
	return o.Payload
}

func (o *PostV1AuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTTokenData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1AuthUnauthorized creates a PostV1AuthUnauthorized with default headers values
func NewPostV1AuthUnauthorized() *PostV1AuthUnauthorized {
	return &PostV1AuthUnauthorized{}
}

/*
PostV1AuthUnauthorized describes a response with status code 401, with default header values.

Authentication failed
*/
type PostV1AuthUnauthorized struct {
	Payload *models.RESTError
}

// IsSuccess returns true when this post v1 auth unauthorized response has a 2xx status code
func (o *PostV1AuthUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 auth unauthorized response has a 3xx status code
func (o *PostV1AuthUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 auth unauthorized response has a 4xx status code
func (o *PostV1AuthUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 auth unauthorized response has a 5xx status code
func (o *PostV1AuthUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 auth unauthorized response a status code equal to that given
func (o *PostV1AuthUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post v1 auth unauthorized response
func (o *PostV1AuthUnauthorized) Code() int {
	return 401
}

func (o *PostV1AuthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/auth][%d] postV1AuthUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1AuthUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/auth][%d] postV1AuthUnauthorized  %+v", 401, o.Payload)
}

func (o *PostV1AuthUnauthorized) GetPayload() *models.RESTError {
	return o.Payload
}

func (o *PostV1AuthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
