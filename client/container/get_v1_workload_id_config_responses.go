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

// GetV1WorkloadIDConfigReader is a Reader for the GetV1WorkloadIDConfig structure.
type GetV1WorkloadIDConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WorkloadIDConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WorkloadIDConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/workload/{id}/config] GetV1WorkloadIDConfig", response, response.Code())
	}
}

// NewGetV1WorkloadIDConfigOK creates a GetV1WorkloadIDConfigOK with default headers values
func NewGetV1WorkloadIDConfigOK() *GetV1WorkloadIDConfigOK {
	return &GetV1WorkloadIDConfigOK{}
}

/*
GetV1WorkloadIDConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WorkloadIDConfigOK struct {
	Payload *models.RESTWorkloadConfigData
}

// IsSuccess returns true when this get v1 workload Id config o k response has a 2xx status code
func (o *GetV1WorkloadIDConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 workload Id config o k response has a 3xx status code
func (o *GetV1WorkloadIDConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 workload Id config o k response has a 4xx status code
func (o *GetV1WorkloadIDConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 workload Id config o k response has a 5xx status code
func (o *GetV1WorkloadIDConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 workload Id config o k response a status code equal to that given
func (o *GetV1WorkloadIDConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 workload Id config o k response
func (o *GetV1WorkloadIDConfigOK) Code() int {
	return 200
}

func (o *GetV1WorkloadIDConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/workload/{id}/config][%d] getV1WorkloadIdConfigOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDConfigOK) String() string {
	return fmt.Sprintf("[GET /v1/workload/{id}/config][%d] getV1WorkloadIdConfigOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDConfigOK) GetPayload() *models.RESTWorkloadConfigData {
	return o.Payload
}

func (o *GetV1WorkloadIDConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWorkloadConfigData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
