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

// GetV1WorkloadIDComplianceReader is a Reader for the GetV1WorkloadIDCompliance structure.
type GetV1WorkloadIDComplianceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WorkloadIDComplianceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WorkloadIDComplianceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/workload/{id}/compliance] GetV1WorkloadIDCompliance", response, response.Code())
	}
}

// NewGetV1WorkloadIDComplianceOK creates a GetV1WorkloadIDComplianceOK with default headers values
func NewGetV1WorkloadIDComplianceOK() *GetV1WorkloadIDComplianceOK {
	return &GetV1WorkloadIDComplianceOK{}
}

/*
GetV1WorkloadIDComplianceOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WorkloadIDComplianceOK struct {
	Payload *models.RESTComplianceData
}

// IsSuccess returns true when this get v1 workload Id compliance o k response has a 2xx status code
func (o *GetV1WorkloadIDComplianceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 workload Id compliance o k response has a 3xx status code
func (o *GetV1WorkloadIDComplianceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 workload Id compliance o k response has a 4xx status code
func (o *GetV1WorkloadIDComplianceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 workload Id compliance o k response has a 5xx status code
func (o *GetV1WorkloadIDComplianceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 workload Id compliance o k response a status code equal to that given
func (o *GetV1WorkloadIDComplianceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 workload Id compliance o k response
func (o *GetV1WorkloadIDComplianceOK) Code() int {
	return 200
}

func (o *GetV1WorkloadIDComplianceOK) Error() string {
	return fmt.Sprintf("[GET /v1/workload/{id}/compliance][%d] getV1WorkloadIdComplianceOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDComplianceOK) String() string {
	return fmt.Sprintf("[GET /v1/workload/{id}/compliance][%d] getV1WorkloadIdComplianceOK  %+v", 200, o.Payload)
}

func (o *GetV1WorkloadIDComplianceOK) GetPayload() *models.RESTComplianceData {
	return o.Payload
}

func (o *GetV1WorkloadIDComplianceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTComplianceData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}