// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1FileDlpConfigReader is a Reader for the PostV1FileDlpConfig structure.
type PostV1FileDlpConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1FileDlpConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1FileDlpConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/file/dlp/config] PostV1FileDlpConfig", response, response.Code())
	}
}

// NewPostV1FileDlpConfigOK creates a PostV1FileDlpConfigOK with default headers values
func NewPostV1FileDlpConfigOK() *PostV1FileDlpConfigOK {
	return &PostV1FileDlpConfigOK{}
}

/*
PostV1FileDlpConfigOK describes a response with status code 200, with default header values.

Success
*/
type PostV1FileDlpConfigOK struct {
}

// IsSuccess returns true when this post v1 file dlp config o k response has a 2xx status code
func (o *PostV1FileDlpConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 file dlp config o k response has a 3xx status code
func (o *PostV1FileDlpConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 file dlp config o k response has a 4xx status code
func (o *PostV1FileDlpConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 file dlp config o k response has a 5xx status code
func (o *PostV1FileDlpConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 file dlp config o k response a status code equal to that given
func (o *PostV1FileDlpConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 file dlp config o k response
func (o *PostV1FileDlpConfigOK) Code() int {
	return 200
}

func (o *PostV1FileDlpConfigOK) Error() string {
	return fmt.Sprintf("[POST /v1/file/dlp/config][%d] postV1FileDlpConfigOK ", 200)
}

func (o *PostV1FileDlpConfigOK) String() string {
	return fmt.Sprintf("[POST /v1/file/dlp/config][%d] postV1FileDlpConfigOK ", 200)
}

func (o *PostV1FileDlpConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
