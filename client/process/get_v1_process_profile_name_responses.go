// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ProcessProfileNameReader is a Reader for the GetV1ProcessProfileName structure.
type GetV1ProcessProfileNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ProcessProfileNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ProcessProfileNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/process_profile/{name}] GetV1ProcessProfileName", response, response.Code())
	}
}

// NewGetV1ProcessProfileNameOK creates a GetV1ProcessProfileNameOK with default headers values
func NewGetV1ProcessProfileNameOK() *GetV1ProcessProfileNameOK {
	return &GetV1ProcessProfileNameOK{}
}

/*
GetV1ProcessProfileNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ProcessProfileNameOK struct {
	Payload *models.RESTProcessProfileData
}

// IsSuccess returns true when this get v1 process profile name o k response has a 2xx status code
func (o *GetV1ProcessProfileNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 process profile name o k response has a 3xx status code
func (o *GetV1ProcessProfileNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 process profile name o k response has a 4xx status code
func (o *GetV1ProcessProfileNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 process profile name o k response has a 5xx status code
func (o *GetV1ProcessProfileNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 process profile name o k response a status code equal to that given
func (o *GetV1ProcessProfileNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 process profile name o k response
func (o *GetV1ProcessProfileNameOK) Code() int {
	return 200
}

func (o *GetV1ProcessProfileNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/process_profile/{name}][%d] getV1ProcessProfileNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ProcessProfileNameOK) String() string {
	return fmt.Sprintf("[GET /v1/process_profile/{name}][%d] getV1ProcessProfileNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ProcessProfileNameOK) GetPayload() *models.RESTProcessProfileData {
	return o.Payload
}

func (o *GetV1ProcessProfileNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTProcessProfileData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
