// Code generated by go-swagger; DO NOT EDIT.

package apikey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// PostV1APIKeyReader is a Reader for the PostV1APIKey structure.
type PostV1APIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1APIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1APIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/api_key] PostV1APIKey", response, response.Code())
	}
}

// NewPostV1APIKeyOK creates a PostV1APIKeyOK with default headers values
func NewPostV1APIKeyOK() *PostV1APIKeyOK {
	return &PostV1APIKeyOK{}
}

/*
PostV1APIKeyOK describes a response with status code 200, with default header values.

Success
*/
type PostV1APIKeyOK struct {
	Payload *models.RESTApikeyGeneratedData
}

// IsSuccess returns true when this post v1 Api key o k response has a 2xx status code
func (o *PostV1APIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 Api key o k response has a 3xx status code
func (o *PostV1APIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 Api key o k response has a 4xx status code
func (o *PostV1APIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 Api key o k response has a 5xx status code
func (o *PostV1APIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 Api key o k response a status code equal to that given
func (o *PostV1APIKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 Api key o k response
func (o *PostV1APIKeyOK) Code() int {
	return 200
}

func (o *PostV1APIKeyOK) Error() string {
	return fmt.Sprintf("[POST /v1/api_key][%d] postV1ApiKeyOK  %+v", 200, o.Payload)
}

func (o *PostV1APIKeyOK) String() string {
	return fmt.Sprintf("[POST /v1/api_key][%d] postV1ApiKeyOK  %+v", 200, o.Payload)
}

func (o *PostV1APIKeyOK) GetPayload() *models.RESTApikeyGeneratedData {
	return o.Payload
}

func (o *PostV1APIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTApikeyGeneratedData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}