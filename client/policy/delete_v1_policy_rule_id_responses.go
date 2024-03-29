// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1PolicyRuleIDReader is a Reader for the DeleteV1PolicyRuleID structure.
type DeleteV1PolicyRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PolicyRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1PolicyRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/policy/rule/{id}] DeleteV1PolicyRuleID", response, response.Code())
	}
}

// NewDeleteV1PolicyRuleIDOK creates a DeleteV1PolicyRuleIDOK with default headers values
func NewDeleteV1PolicyRuleIDOK() *DeleteV1PolicyRuleIDOK {
	return &DeleteV1PolicyRuleIDOK{}
}

/*
DeleteV1PolicyRuleIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteV1PolicyRuleIDOK struct {
}

// IsSuccess returns true when this delete v1 policy rule Id o k response has a 2xx status code
func (o *DeleteV1PolicyRuleIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 policy rule Id o k response has a 3xx status code
func (o *DeleteV1PolicyRuleIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 policy rule Id o k response has a 4xx status code
func (o *DeleteV1PolicyRuleIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 policy rule Id o k response has a 5xx status code
func (o *DeleteV1PolicyRuleIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 policy rule Id o k response a status code equal to that given
func (o *DeleteV1PolicyRuleIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 policy rule Id o k response
func (o *DeleteV1PolicyRuleIDOK) Code() int {
	return 200
}

func (o *DeleteV1PolicyRuleIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/policy/rule/{id}][%d] deleteV1PolicyRuleIdOK ", 200)
}

func (o *DeleteV1PolicyRuleIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/policy/rule/{id}][%d] deleteV1PolicyRuleIdOK ", 200)
}

func (o *DeleteV1PolicyRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
