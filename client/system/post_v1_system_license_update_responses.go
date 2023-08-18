// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1SystemLicenseUpdateReader is a Reader for the PostV1SystemLicenseUpdate structure.
type PostV1SystemLicenseUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SystemLicenseUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1SystemLicenseUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/system/license/update] PostV1SystemLicenseUpdate", response, response.Code())
	}
}

// NewPostV1SystemLicenseUpdateOK creates a PostV1SystemLicenseUpdateOK with default headers values
func NewPostV1SystemLicenseUpdateOK() *PostV1SystemLicenseUpdateOK {
	return &PostV1SystemLicenseUpdateOK{}
}

/*
PostV1SystemLicenseUpdateOK describes a response with status code 200, with default header values.

Success
*/
type PostV1SystemLicenseUpdateOK struct {
}

// IsSuccess returns true when this post v1 system license update o k response has a 2xx status code
func (o *PostV1SystemLicenseUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 system license update o k response has a 3xx status code
func (o *PostV1SystemLicenseUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 system license update o k response has a 4xx status code
func (o *PostV1SystemLicenseUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 system license update o k response has a 5xx status code
func (o *PostV1SystemLicenseUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 system license update o k response a status code equal to that given
func (o *PostV1SystemLicenseUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 system license update o k response
func (o *PostV1SystemLicenseUpdateOK) Code() int {
	return 200
}

func (o *PostV1SystemLicenseUpdateOK) Error() string {
	return fmt.Sprintf("[POST /v1/system/license/update][%d] postV1SystemLicenseUpdateOK ", 200)
}

func (o *PostV1SystemLicenseUpdateOK) String() string {
	return fmt.Sprintf("[POST /v1/system/license/update][%d] postV1SystemLicenseUpdateOK ", 200)
}

func (o *PostV1SystemLicenseUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}