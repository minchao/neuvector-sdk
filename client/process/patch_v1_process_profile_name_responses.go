// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1ProcessProfileNameReader is a Reader for the PatchV1ProcessProfileName structure.
type PatchV1ProcessProfileNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1ProcessProfileNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1ProcessProfileNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/process_profile/{name}] PatchV1ProcessProfileName", response, response.Code())
	}
}

// NewPatchV1ProcessProfileNameOK creates a PatchV1ProcessProfileNameOK with default headers values
func NewPatchV1ProcessProfileNameOK() *PatchV1ProcessProfileNameOK {
	return &PatchV1ProcessProfileNameOK{}
}

/*
PatchV1ProcessProfileNameOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1ProcessProfileNameOK struct {
}

// IsSuccess returns true when this patch v1 process profile name o k response has a 2xx status code
func (o *PatchV1ProcessProfileNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 process profile name o k response has a 3xx status code
func (o *PatchV1ProcessProfileNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 process profile name o k response has a 4xx status code
func (o *PatchV1ProcessProfileNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 process profile name o k response has a 5xx status code
func (o *PatchV1ProcessProfileNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 process profile name o k response a status code equal to that given
func (o *PatchV1ProcessProfileNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 process profile name o k response
func (o *PatchV1ProcessProfileNameOK) Code() int {
	return 200
}

func (o *PatchV1ProcessProfileNameOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/process_profile/{name}][%d] patchV1ProcessProfileNameOK ", 200)
}

func (o *PatchV1ProcessProfileNameOK) String() string {
	return fmt.Sprintf("[PATCH /v1/process_profile/{name}][%d] patchV1ProcessProfileNameOK ", 200)
}

func (o *PatchV1ProcessProfileNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
