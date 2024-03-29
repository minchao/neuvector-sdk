// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1WorkloadReader is a Reader for the GetV1Workload structure.
type GetV1WorkloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WorkloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WorkloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/workload] GetV1Workload", response, response.Code())
	}
}

// NewGetV1WorkloadOK creates a GetV1WorkloadOK with default headers values
func NewGetV1WorkloadOK() *GetV1WorkloadOK {
	return &GetV1WorkloadOK{}
}

/*
GetV1WorkloadOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WorkloadOK struct {
	Payload *models.RESTWorkloadsData
}

// IsSuccess returns true when this get v1 workload o k response has a 2xx status code
func (o *GetV1WorkloadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 workload o k response has a 3xx status code
func (o *GetV1WorkloadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 workload o k response has a 4xx status code
func (o *GetV1WorkloadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 workload o k response has a 5xx status code
func (o *GetV1WorkloadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 workload o k response a status code equal to that given
func (o *GetV1WorkloadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 workload o k response
func (o *GetV1WorkloadOK) Code() int {
	return 200
}

func (o *GetV1WorkloadOK) Error() string {
	return fmt.Sprintf("[GET /v1/workload][%d] getV1WorkloadOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadOK) String() string {
	return fmt.Sprintf("[GET /v1/workload][%d] getV1WorkloadOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadOK) GetPayload() *models.RESTWorkloadsData {
	return o.Payload
}

func (o *GetV1WorkloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWorkloadsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
