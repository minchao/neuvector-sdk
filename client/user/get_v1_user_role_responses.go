// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1UserRoleReader is a Reader for the GetV1UserRole structure.
type GetV1UserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1UserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1UserRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/user_role] GetV1UserRole", response, response.Code())
	}
}

// NewGetV1UserRoleOK creates a GetV1UserRoleOK with default headers values
func NewGetV1UserRoleOK() *GetV1UserRoleOK {
	return &GetV1UserRoleOK{}
}

/*
GetV1UserRoleOK describes a response with status code 200, with default header values.

Success
*/
type GetV1UserRoleOK struct {
	Payload *models.RESTUserRolesData
}

// IsSuccess returns true when this get v1 user role o k response has a 2xx status code
func (o *GetV1UserRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 user role o k response has a 3xx status code
func (o *GetV1UserRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 user role o k response has a 4xx status code
func (o *GetV1UserRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 user role o k response has a 5xx status code
func (o *GetV1UserRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 user role o k response a status code equal to that given
func (o *GetV1UserRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 user role o k response
func (o *GetV1UserRoleOK) Code() int {
	return 200
}

func (o *GetV1UserRoleOK) Error() string {
	return fmt.Sprintf("[GET /v1/user_role][%d] getV1UserRoleOK  %+v", 200, o.Payload)
}

func (o *GetV1UserRoleOK) String() string {
	return fmt.Sprintf("[GET /v1/user_role][%d] getV1UserRoleOK  %+v", 200, o.Payload)
}

func (o *GetV1UserRoleOK) GetPayload() *models.RESTUserRolesData {
	return o.Payload
}

func (o *GetV1UserRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTUserRolesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}