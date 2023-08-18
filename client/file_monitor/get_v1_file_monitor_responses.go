// Code generated by go-swagger; DO NOT EDIT.

package file_monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1FileMonitorReader is a Reader for the GetV1FileMonitor structure.
type GetV1FileMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1FileMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1FileMonitorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/file_monitor] GetV1FileMonitor", response, response.Code())
	}
}

// NewGetV1FileMonitorOK creates a GetV1FileMonitorOK with default headers values
func NewGetV1FileMonitorOK() *GetV1FileMonitorOK {
	return &GetV1FileMonitorOK{}
}

/*
GetV1FileMonitorOK describes a response with status code 200, with default header values.

Success
*/
type GetV1FileMonitorOK struct {
	Payload *models.RESTFileMonitorFileData
}

// IsSuccess returns true when this get v1 file monitor o k response has a 2xx status code
func (o *GetV1FileMonitorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 file monitor o k response has a 3xx status code
func (o *GetV1FileMonitorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 file monitor o k response has a 4xx status code
func (o *GetV1FileMonitorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 file monitor o k response has a 5xx status code
func (o *GetV1FileMonitorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 file monitor o k response a status code equal to that given
func (o *GetV1FileMonitorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 file monitor o k response
func (o *GetV1FileMonitorOK) Code() int {
	return 200
}

func (o *GetV1FileMonitorOK) Error() string {
	return fmt.Sprintf("[GET /v1/file_monitor][%d] getV1FileMonitorOK  %+v", 200, o.Payload)
}

func (o *GetV1FileMonitorOK) String() string {
	return fmt.Sprintf("[GET /v1/file_monitor][%d] getV1FileMonitorOK  %+v", 200, o.Payload)
}

func (o *GetV1FileMonitorOK) GetPayload() *models.RESTFileMonitorFileData {
	return o.Payload
}

func (o *GetV1FileMonitorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTFileMonitorFileData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}