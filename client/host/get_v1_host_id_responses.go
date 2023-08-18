// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1HostIDReader is a Reader for the GetV1HostID structure.
type GetV1HostIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1HostIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1HostIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/host/{id}] GetV1HostID", response, response.Code())
	}
}

// NewGetV1HostIDOK creates a GetV1HostIDOK with default headers values
func NewGetV1HostIDOK() *GetV1HostIDOK {
	return &GetV1HostIDOK{}
}

/*
GetV1HostIDOK describes a response with status code 200, with default header values.

Success
*/
type GetV1HostIDOK struct {
	Payload *models.RESTHostData
}

// IsSuccess returns true when this get v1 host Id o k response has a 2xx status code
func (o *GetV1HostIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 host Id o k response has a 3xx status code
func (o *GetV1HostIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 host Id o k response has a 4xx status code
func (o *GetV1HostIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 host Id o k response has a 5xx status code
func (o *GetV1HostIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 host Id o k response a status code equal to that given
func (o *GetV1HostIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 host Id o k response
func (o *GetV1HostIDOK) Code() int {
	return 200
}

func (o *GetV1HostIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/host/{id}][%d] getV1HostIdOK  %+v", 200, o.Payload)
}

func (o *GetV1HostIDOK) String() string {
	return fmt.Sprintf("[GET /v1/host/{id}][%d] getV1HostIdOK  %+v", 200, o.Payload)
}

func (o *GetV1HostIDOK) GetPayload() *models.RESTHostData {
	return o.Payload
}

func (o *GetV1HostIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTHostData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
