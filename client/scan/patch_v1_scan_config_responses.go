// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1ScanConfigReader is a Reader for the PatchV1ScanConfig structure.
type PatchV1ScanConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1ScanConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1ScanConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/scan/config] PatchV1ScanConfig", response, response.Code())
	}
}

// NewPatchV1ScanConfigOK creates a PatchV1ScanConfigOK with default headers values
func NewPatchV1ScanConfigOK() *PatchV1ScanConfigOK {
	return &PatchV1ScanConfigOK{}
}

/*
PatchV1ScanConfigOK describes a response with status code 200, with default header values.

Success
*/
type PatchV1ScanConfigOK struct {
}

// IsSuccess returns true when this patch v1 scan config o k response has a 2xx status code
func (o *PatchV1ScanConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 scan config o k response has a 3xx status code
func (o *PatchV1ScanConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 scan config o k response has a 4xx status code
func (o *PatchV1ScanConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 scan config o k response has a 5xx status code
func (o *PatchV1ScanConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 scan config o k response a status code equal to that given
func (o *PatchV1ScanConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 scan config o k response
func (o *PatchV1ScanConfigOK) Code() int {
	return 200
}

func (o *PatchV1ScanConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/scan/config][%d] patchV1ScanConfigOK ", 200)
}

func (o *PatchV1ScanConfigOK) String() string {
	return fmt.Sprintf("[PATCH /v1/scan/config][%d] patchV1ScanConfigOK ", 200)
}

func (o *PatchV1ScanConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
