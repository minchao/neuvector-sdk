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

// GetV2WorkloadIDReader is a Reader for the GetV2WorkloadID structure.
type GetV2WorkloadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2WorkloadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2WorkloadIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v2/workload/{id}] GetV2WorkloadID", response, response.Code())
	}
}

// NewGetV2WorkloadIDOK creates a GetV2WorkloadIDOK with default headers values
func NewGetV2WorkloadIDOK() *GetV2WorkloadIDOK {
	return &GetV2WorkloadIDOK{}
}

/*
GetV2WorkloadIDOK describes a response with status code 200, with default header values.

Success
*/
type GetV2WorkloadIDOK struct {
	Payload *models.RESTWorkloadDetailDataV2
}

// IsSuccess returns true when this get v2 workload Id o k response has a 2xx status code
func (o *GetV2WorkloadIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v2 workload Id o k response has a 3xx status code
func (o *GetV2WorkloadIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 workload Id o k response has a 4xx status code
func (o *GetV2WorkloadIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 workload Id o k response has a 5xx status code
func (o *GetV2WorkloadIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 workload Id o k response a status code equal to that given
func (o *GetV2WorkloadIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v2 workload Id o k response
func (o *GetV2WorkloadIDOK) Code() int {
	return 200
}

func (o *GetV2WorkloadIDOK) Error() string {
	return fmt.Sprintf("[GET /v2/workload/{id}][%d] getV2WorkloadIdOK  %+v", 200, o.Payload)
}

func (o *GetV2WorkloadIDOK) String() string {
	return fmt.Sprintf("[GET /v2/workload/{id}][%d] getV2WorkloadIdOK  %+v", 200, o.Payload)
}

func (o *GetV2WorkloadIDOK) GetPayload() *models.RESTWorkloadDetailDataV2 {
	return o.Payload
}

func (o *GetV2WorkloadIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWorkloadDetailDataV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
