// Code generated by go-swagger; DO NOT EDIT.

package compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1CustomCheckGroupReader is a Reader for the PatchV1CustomCheckGroup structure.
type PatchV1CustomCheckGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1CustomCheckGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1CustomCheckGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/custom_check/{group}] PatchV1CustomCheckGroup", response, response.Code())
	}
}

// NewPatchV1CustomCheckGroupOK creates a PatchV1CustomCheckGroupOK with default headers values
func NewPatchV1CustomCheckGroupOK() *PatchV1CustomCheckGroupOK {
	return &PatchV1CustomCheckGroupOK{}
}

/*
PatchV1CustomCheckGroupOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1CustomCheckGroupOK struct {
}

// IsSuccess returns true when this patch v1 custom check group o k response has a 2xx status code
func (o *PatchV1CustomCheckGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 custom check group o k response has a 3xx status code
func (o *PatchV1CustomCheckGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 custom check group o k response has a 4xx status code
func (o *PatchV1CustomCheckGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 custom check group o k response has a 5xx status code
func (o *PatchV1CustomCheckGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 custom check group o k response a status code equal to that given
func (o *PatchV1CustomCheckGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 custom check group o k response
func (o *PatchV1CustomCheckGroupOK) Code() int {
	return 200
}

func (o *PatchV1CustomCheckGroupOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/custom_check/{group}][%d] patchV1CustomCheckGroupOK ", 200)
}

func (o *PatchV1CustomCheckGroupOK) String() string {
	return fmt.Sprintf("[PATCH /v1/custom_check/{group}][%d] patchV1CustomCheckGroupOK ", 200)
}

func (o *PatchV1CustomCheckGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
