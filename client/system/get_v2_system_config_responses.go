// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV2SystemConfigReader is a Reader for the GetV2SystemConfig structure.
type GetV2SystemConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2SystemConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2SystemConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v2/system/config] GetV2SystemConfig", response, response.Code())
	}
}

// NewGetV2SystemConfigOK creates a GetV2SystemConfigOK with default headers values
func NewGetV2SystemConfigOK() *GetV2SystemConfigOK {
	return &GetV2SystemConfigOK{}
}

/*
GetV2SystemConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetV2SystemConfigOK struct {
	Payload *models.RESTSystemConfigDataV2
}

// IsSuccess returns true when this get v2 system config o k response has a 2xx status code
func (o *GetV2SystemConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v2 system config o k response has a 3xx status code
func (o *GetV2SystemConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 system config o k response has a 4xx status code
func (o *GetV2SystemConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 system config o k response has a 5xx status code
func (o *GetV2SystemConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 system config o k response a status code equal to that given
func (o *GetV2SystemConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v2 system config o k response
func (o *GetV2SystemConfigOK) Code() int {
	return 200
}

func (o *GetV2SystemConfigOK) Error() string {
	return fmt.Sprintf("[GET /v2/system/config][%d] getV2SystemConfigOK  %+v", 200, o.Payload)
}

func (o *GetV2SystemConfigOK) String() string {
	return fmt.Sprintf("[GET /v2/system/config][%d] getV2SystemConfigOK  %+v", 200, o.Payload)
}

func (o *GetV2SystemConfigOK) GetPayload() *models.RESTSystemConfigDataV2 {
	return o.Payload
}

func (o *GetV2SystemConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTSystemConfigDataV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
