// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ServerNameReader is a Reader for the GetV1ServerName structure.
type GetV1ServerNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServerNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ServerNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/server/{name}] GetV1ServerName", response, response.Code())
	}
}

// NewGetV1ServerNameOK creates a GetV1ServerNameOK with default headers values
func NewGetV1ServerNameOK() *GetV1ServerNameOK {
	return &GetV1ServerNameOK{}
}

/*
GetV1ServerNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ServerNameOK struct {
	Payload *models.RESTServerData
}

// IsSuccess returns true when this get v1 server name o k response has a 2xx status code
func (o *GetV1ServerNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 server name o k response has a 3xx status code
func (o *GetV1ServerNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 server name o k response has a 4xx status code
func (o *GetV1ServerNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 server name o k response has a 5xx status code
func (o *GetV1ServerNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 server name o k response a status code equal to that given
func (o *GetV1ServerNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 server name o k response
func (o *GetV1ServerNameOK) Code() int {
	return 200
}

func (o *GetV1ServerNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/server/{name}][%d] getV1ServerNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ServerNameOK) String() string {
	return fmt.Sprintf("[GET /v1/server/{name}][%d] getV1ServerNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ServerNameOK) GetPayload() *models.RESTServerData {
	return o.Payload
}

func (o *GetV1ServerNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTServerData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
