// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// PostV1GroupReader is a Reader for the PostV1Group structure.
type PostV1GroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1GroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1GroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1GroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/group] PostV1Group", response, response.Code())
	}
}

// NewPostV1GroupOK creates a PostV1GroupOK with default headers values
func NewPostV1GroupOK() *PostV1GroupOK {
	return &PostV1GroupOK{}
}

/*
PostV1GroupOK describes a response with status code 200, with default header values.

Success
*/
type PostV1GroupOK struct {
}

// IsSuccess returns true when this post v1 group o k response has a 2xx status code
func (o *PostV1GroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 group o k response has a 3xx status code
func (o *PostV1GroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 group o k response has a 4xx status code
func (o *PostV1GroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 group o k response has a 5xx status code
func (o *PostV1GroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 group o k response a status code equal to that given
func (o *PostV1GroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 group o k response
func (o *PostV1GroupOK) Code() int {
	return 200
}

func (o *PostV1GroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/group][%d] postV1GroupOK ", 200)
}

func (o *PostV1GroupOK) String() string {
	return fmt.Sprintf("[POST /v1/group][%d] postV1GroupOK ", 200)
}

func (o *PostV1GroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostV1GroupBadRequest creates a PostV1GroupBadRequest with default headers values
func NewPostV1GroupBadRequest() *PostV1GroupBadRequest {
	return &PostV1GroupBadRequest{}
}

/*
PostV1GroupBadRequest describes a response with status code 400, with default header values.

Error
*/
type PostV1GroupBadRequest struct {
	Payload *models.RESTError
}

// IsSuccess returns true when this post v1 group bad request response has a 2xx status code
func (o *PostV1GroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 group bad request response has a 3xx status code
func (o *PostV1GroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 group bad request response has a 4xx status code
func (o *PostV1GroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 group bad request response has a 5xx status code
func (o *PostV1GroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 group bad request response a status code equal to that given
func (o *PostV1GroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 group bad request response
func (o *PostV1GroupBadRequest) Code() int {
	return 400
}

func (o *PostV1GroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/group][%d] postV1GroupBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1GroupBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/group][%d] postV1GroupBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1GroupBadRequest) GetPayload() *models.RESTError {
	return o.Payload
}

func (o *PostV1GroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}