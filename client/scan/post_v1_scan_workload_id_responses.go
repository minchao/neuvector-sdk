// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1ScanWorkloadIDReader is a Reader for the PostV1ScanWorkloadID structure.
type PostV1ScanWorkloadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ScanWorkloadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1ScanWorkloadIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/scan/workload/{id}] PostV1ScanWorkloadID", response, response.Code())
	}
}

// NewPostV1ScanWorkloadIDOK creates a PostV1ScanWorkloadIDOK with default headers values
func NewPostV1ScanWorkloadIDOK() *PostV1ScanWorkloadIDOK {
	return &PostV1ScanWorkloadIDOK{}
}

/*
PostV1ScanWorkloadIDOK describes a response with status code 200, with default header values.

Success
*/
type PostV1ScanWorkloadIDOK struct {
}

// IsSuccess returns true when this post v1 scan workload Id o k response has a 2xx status code
func (o *PostV1ScanWorkloadIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 scan workload Id o k response has a 3xx status code
func (o *PostV1ScanWorkloadIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 scan workload Id o k response has a 4xx status code
func (o *PostV1ScanWorkloadIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 scan workload Id o k response has a 5xx status code
func (o *PostV1ScanWorkloadIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 scan workload Id o k response a status code equal to that given
func (o *PostV1ScanWorkloadIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 scan workload Id o k response
func (o *PostV1ScanWorkloadIDOK) Code() int {
	return 200
}

func (o *PostV1ScanWorkloadIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/scan/workload/{id}][%d] postV1ScanWorkloadIdOK ", 200)
}

func (o *PostV1ScanWorkloadIDOK) String() string {
	return fmt.Sprintf("[POST /v1/scan/workload/{id}][%d] postV1ScanWorkloadIdOK ", 200)
}

func (o *PostV1ScanWorkloadIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
