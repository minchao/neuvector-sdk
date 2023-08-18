// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1PolicyRuleReader is a Reader for the DeleteV1PolicyRule structure.
type DeleteV1PolicyRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PolicyRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1PolicyRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/policy/rule] DeleteV1PolicyRule", response, response.Code())
	}
}

// NewDeleteV1PolicyRuleOK creates a DeleteV1PolicyRuleOK with default headers values
func NewDeleteV1PolicyRuleOK() *DeleteV1PolicyRuleOK {
	return &DeleteV1PolicyRuleOK{}
}

/*
DeleteV1PolicyRuleOK describes a response with status code 200, with default header values.

Success
*/
type DeleteV1PolicyRuleOK struct {
}

// IsSuccess returns true when this delete v1 policy rule o k response has a 2xx status code
func (o *DeleteV1PolicyRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 policy rule o k response has a 3xx status code
func (o *DeleteV1PolicyRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 policy rule o k response has a 4xx status code
func (o *DeleteV1PolicyRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 policy rule o k response has a 5xx status code
func (o *DeleteV1PolicyRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 policy rule o k response a status code equal to that given
func (o *DeleteV1PolicyRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 policy rule o k response
func (o *DeleteV1PolicyRuleOK) Code() int {
	return 200
}

func (o *DeleteV1PolicyRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/policy/rule][%d] deleteV1PolicyRuleOK ", 200)
}

func (o *DeleteV1PolicyRuleOK) String() string {
	return fmt.Sprintf("[DELETE /v1/policy/rule][%d] deleteV1PolicyRuleOK ", 200)
}

func (o *DeleteV1PolicyRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
