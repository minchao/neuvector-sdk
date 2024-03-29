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

// GetV1WorkloadIDReader is a Reader for the GetV1WorkloadID structure.
type GetV1WorkloadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WorkloadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WorkloadIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/workload/{id}] GetV1WorkloadID", response, response.Code())
	}
}

// NewGetV1WorkloadIDOK creates a GetV1WorkloadIDOK with default headers values
func NewGetV1WorkloadIDOK() *GetV1WorkloadIDOK {
	return &GetV1WorkloadIDOK{}
}

/*
GetV1WorkloadIDOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WorkloadIDOK struct {
	Payload *models.RESTWorkloadDetailData
}

// IsSuccess returns true when this get v1 workload Id o k response has a 2xx status code
func (o *GetV1WorkloadIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 workload Id o k response has a 3xx status code
func (o *GetV1WorkloadIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 workload Id o k response has a 4xx status code
func (o *GetV1WorkloadIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 workload Id o k response has a 5xx status code
func (o *GetV1WorkloadIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 workload Id o k response a status code equal to that given
func (o *GetV1WorkloadIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 workload Id o k response
func (o *GetV1WorkloadIDOK) Code() int {
	return 200
}

func (o *GetV1WorkloadIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/workload/{id}][%d] getV1WorkloadIdOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDOK) String() string {
	return fmt.Sprintf("[GET /v1/workload/{id}][%d] getV1WorkloadIdOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDOK) GetPayload() *models.RESTWorkloadDetailData {
	return o.Payload
}

func (o *GetV1WorkloadIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWorkloadDetailData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
