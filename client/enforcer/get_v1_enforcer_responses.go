// Code generated by go-swagger; DO NOT EDIT.

package enforcer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1EnforcerReader is a Reader for the GetV1Enforcer structure.
type GetV1EnforcerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1EnforcerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1EnforcerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/enforcer] GetV1Enforcer", response, response.Code())
	}
}

// NewGetV1EnforcerOK creates a GetV1EnforcerOK with default headers values
func NewGetV1EnforcerOK() *GetV1EnforcerOK {
	return &GetV1EnforcerOK{}
}

/*
GetV1EnforcerOK describes a response with status code 200, with default header values.

Success
*/
type GetV1EnforcerOK struct {
	Payload *models.RESTAgentsData
}

// IsSuccess returns true when this get v1 enforcer o k response has a 2xx status code
func (o *GetV1EnforcerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 enforcer o k response has a 3xx status code
func (o *GetV1EnforcerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 enforcer o k response has a 4xx status code
func (o *GetV1EnforcerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 enforcer o k response has a 5xx status code
func (o *GetV1EnforcerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 enforcer o k response a status code equal to that given
func (o *GetV1EnforcerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 enforcer o k response
func (o *GetV1EnforcerOK) Code() int {
	return 200
}

func (o *GetV1EnforcerOK) Error() string {
	return fmt.Sprintf("[GET /v1/enforcer][%d] getV1EnforcerOK  %+v", 200, o.Payload)
}

func (o *GetV1EnforcerOK) String() string {
	return fmt.Sprintf("[GET /v1/enforcer][%d] getV1EnforcerOK  %+v", 200, o.Payload)
}

func (o *GetV1EnforcerOK) GetPayload() *models.RESTAgentsData {
	return o.Payload
}

func (o *GetV1EnforcerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTAgentsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
