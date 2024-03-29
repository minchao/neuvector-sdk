// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1SystemConfigReader is a Reader for the PatchV1SystemConfig structure.
type PatchV1SystemConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1SystemConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1SystemConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/system/config] PatchV1SystemConfig", response, response.Code())
	}
}

// NewPatchV1SystemConfigOK creates a PatchV1SystemConfigOK with default headers values
func NewPatchV1SystemConfigOK() *PatchV1SystemConfigOK {
	return &PatchV1SystemConfigOK{}
}

/*
PatchV1SystemConfigOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1SystemConfigOK struct {
}

// IsSuccess returns true when this patch v1 system config o k response has a 2xx status code
func (o *PatchV1SystemConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 system config o k response has a 3xx status code
func (o *PatchV1SystemConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 system config o k response has a 4xx status code
func (o *PatchV1SystemConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 system config o k response has a 5xx status code
func (o *PatchV1SystemConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 system config o k response a status code equal to that given
func (o *PatchV1SystemConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 system config o k response
func (o *PatchV1SystemConfigOK) Code() int {
	return 200
}

func (o *PatchV1SystemConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/system/config][%d] patchV1SystemConfigOK ", 200)
}

func (o *PatchV1SystemConfigOK) String() string {
	return fmt.Sprintf("[PATCH /v1/system/config][%d] patchV1SystemConfigOK ", 200)
}

func (o *PatchV1SystemConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
