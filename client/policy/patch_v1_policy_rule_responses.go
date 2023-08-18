// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1PolicyRuleReader is a Reader for the PatchV1PolicyRule structure.
type PatchV1PolicyRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1PolicyRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1PolicyRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/policy/rule] PatchV1PolicyRule", response, response.Code())
	}
}

// NewPatchV1PolicyRuleOK creates a PatchV1PolicyRuleOK with default headers values
func NewPatchV1PolicyRuleOK() *PatchV1PolicyRuleOK {
	return &PatchV1PolicyRuleOK{}
}

/*
PatchV1PolicyRuleOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1PolicyRuleOK struct {
}

// IsSuccess returns true when this patch v1 policy rule o k response has a 2xx status code
func (o *PatchV1PolicyRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 policy rule o k response has a 3xx status code
func (o *PatchV1PolicyRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 policy rule o k response has a 4xx status code
func (o *PatchV1PolicyRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 policy rule o k response has a 5xx status code
func (o *PatchV1PolicyRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 policy rule o k response a status code equal to that given
func (o *PatchV1PolicyRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 policy rule o k response
func (o *PatchV1PolicyRuleOK) Code() int {
	return 200
}

func (o *PatchV1PolicyRuleOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/policy/rule][%d] patchV1PolicyRuleOK ", 200)
}

func (o *PatchV1PolicyRuleOK) String() string {
	return fmt.Sprintf("[PATCH /v1/policy/rule][%d] patchV1PolicyRuleOK ", 200)
}

func (o *PatchV1PolicyRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
