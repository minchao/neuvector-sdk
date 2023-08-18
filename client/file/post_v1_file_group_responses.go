// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1FileGroupReader is a Reader for the PostV1FileGroup structure.
type PostV1FileGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1FileGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1FileGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/file/group] PostV1FileGroup", response, response.Code())
	}
}

// NewPostV1FileGroupOK creates a PostV1FileGroupOK with default headers values
func NewPostV1FileGroupOK() *PostV1FileGroupOK {
	return &PostV1FileGroupOK{}
}

/*
PostV1FileGroupOK describes a response with status code 200, with default header values.

Success. Get a yaml file.
*/
type PostV1FileGroupOK struct {
}

// IsSuccess returns true when this post v1 file group o k response has a 2xx status code
func (o *PostV1FileGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 file group o k response has a 3xx status code
func (o *PostV1FileGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 file group o k response has a 4xx status code
func (o *PostV1FileGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 file group o k response has a 5xx status code
func (o *PostV1FileGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 file group o k response a status code equal to that given
func (o *PostV1FileGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 file group o k response
func (o *PostV1FileGroupOK) Code() int {
	return 200
}

func (o *PostV1FileGroupOK) Error() string {
	return fmt.Sprintf("[POST /v1/file/group][%d] postV1FileGroupOK ", 200)
}

func (o *PostV1FileGroupOK) String() string {
	return fmt.Sprintf("[POST /v1/file/group][%d] postV1FileGroupOK ", 200)
}

func (o *PostV1FileGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}