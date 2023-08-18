// Code generated by go-swagger; DO NOT EDIT.

package response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1ResponseRuleReader is a Reader for the PatchV1ResponseRule structure.
type PatchV1ResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1ResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1ResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/response/rule] PatchV1ResponseRule", response, response.Code())
	}
}

// NewPatchV1ResponseRuleOK creates a PatchV1ResponseRuleOK with default headers values
func NewPatchV1ResponseRuleOK() *PatchV1ResponseRuleOK {
	return &PatchV1ResponseRuleOK{}
}

/*
PatchV1ResponseRuleOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1ResponseRuleOK struct {
}

// IsSuccess returns true when this patch v1 response rule o k response has a 2xx status code
func (o *PatchV1ResponseRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 response rule o k response has a 3xx status code
func (o *PatchV1ResponseRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 response rule o k response has a 4xx status code
func (o *PatchV1ResponseRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 response rule o k response has a 5xx status code
func (o *PatchV1ResponseRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 response rule o k response a status code equal to that given
func (o *PatchV1ResponseRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 response rule o k response
func (o *PatchV1ResponseRuleOK) Code() int {
	return 200
}

func (o *PatchV1ResponseRuleOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/response/rule][%d] patchV1ResponseRuleOK ", 200)
}

func (o *PatchV1ResponseRuleOK) String() string {
	return fmt.Sprintf("[PATCH /v1/response/rule][%d] patchV1ResponseRuleOK ", 200)
}

func (o *PatchV1ResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
