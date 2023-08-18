// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ControllerIDStatsReader is a Reader for the GetV1ControllerIDStats structure.
type GetV1ControllerIDStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ControllerIDStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ControllerIDStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/controller/{id}/stats] GetV1ControllerIDStats", response, response.Code())
	}
}

// NewGetV1ControllerIDStatsOK creates a GetV1ControllerIDStatsOK with default headers values
func NewGetV1ControllerIDStatsOK() *GetV1ControllerIDStatsOK {
	return &GetV1ControllerIDStatsOK{}
}

/*
GetV1ControllerIDStatsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ControllerIDStatsOK struct {
	Payload *models.RESTWorkloadStatsData
}

// IsSuccess returns true when this get v1 controller Id stats o k response has a 2xx status code
func (o *GetV1ControllerIDStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 controller Id stats o k response has a 3xx status code
func (o *GetV1ControllerIDStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 controller Id stats o k response has a 4xx status code
func (o *GetV1ControllerIDStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 controller Id stats o k response has a 5xx status code
func (o *GetV1ControllerIDStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 controller Id stats o k response a status code equal to that given
func (o *GetV1ControllerIDStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 controller Id stats o k response
func (o *GetV1ControllerIDStatsOK) Code() int {
	return 200
}

func (o *GetV1ControllerIDStatsOK) Error() string {
	return fmt.Sprintf("[GET /v1/controller/{id}/stats][%d] getV1ControllerIdStatsOK  %+v", 200, o.Payload)
}

func (o *GetV1ControllerIDStatsOK) String() string {
	return fmt.Sprintf("[GET /v1/controller/{id}/stats][%d] getV1ControllerIdStatsOK  %+v", 200, o.Payload)
}

func (o *GetV1ControllerIDStatsOK) GetPayload() *models.RESTWorkloadStatsData {
	return o.Payload
}

func (o *GetV1ControllerIDStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWorkloadStatsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
