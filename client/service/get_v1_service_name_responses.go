// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ServiceNameReader is a Reader for the GetV1ServiceName structure.
type GetV1ServiceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServiceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ServiceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/service/{name}] GetV1ServiceName", response, response.Code())
	}
}

// NewGetV1ServiceNameOK creates a GetV1ServiceNameOK with default headers values
func NewGetV1ServiceNameOK() *GetV1ServiceNameOK {
	return &GetV1ServiceNameOK{}
}

/*
GetV1ServiceNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ServiceNameOK struct {
	Payload *models.RESTServiceData
}

// IsSuccess returns true when this get v1 service name o k response has a 2xx status code
func (o *GetV1ServiceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 service name o k response has a 3xx status code
func (o *GetV1ServiceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 service name o k response has a 4xx status code
func (o *GetV1ServiceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 service name o k response has a 5xx status code
func (o *GetV1ServiceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 service name o k response a status code equal to that given
func (o *GetV1ServiceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 service name o k response
func (o *GetV1ServiceNameOK) Code() int {
	return 200
}

func (o *GetV1ServiceNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/service/{name}][%d] getV1ServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ServiceNameOK) String() string {
	return fmt.Sprintf("[GET /v1/service/{name}][%d] getV1ServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ServiceNameOK) GetPayload() *models.RESTServiceData {
	return o.Payload
}

func (o *GetV1ServiceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTServiceData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}