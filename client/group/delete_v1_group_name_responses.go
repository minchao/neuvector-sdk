// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1GroupNameReader is a Reader for the DeleteV1GroupName structure.
type DeleteV1GroupNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1GroupNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1GroupNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/group/{name}] DeleteV1GroupName", response, response.Code())
	}
}

// NewDeleteV1GroupNameOK creates a DeleteV1GroupNameOK with default headers values
func NewDeleteV1GroupNameOK() *DeleteV1GroupNameOK {
	return &DeleteV1GroupNameOK{}
}

/*
DeleteV1GroupNameOK describes a response with status code 200, with default header values.

Success
*/
type DeleteV1GroupNameOK struct {
}

// IsSuccess returns true when this delete v1 group name o k response has a 2xx status code
func (o *DeleteV1GroupNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 group name o k response has a 3xx status code
func (o *DeleteV1GroupNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 group name o k response has a 4xx status code
func (o *DeleteV1GroupNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 group name o k response has a 5xx status code
func (o *DeleteV1GroupNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 group name o k response a status code equal to that given
func (o *DeleteV1GroupNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 group name o k response
func (o *DeleteV1GroupNameOK) Code() int {
	return 200
}

func (o *DeleteV1GroupNameOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/group/{name}][%d] deleteV1GroupNameOK ", 200)
}

func (o *DeleteV1GroupNameOK) String() string {
	return fmt.Sprintf("[DELETE /v1/group/{name}][%d] deleteV1GroupNameOK ", 200)
}

func (o *DeleteV1GroupNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
