// Code generated by go-swagger; DO NOT EDIT.

package compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1CustomCheckReader is a Reader for the GetV1CustomCheck structure.
type GetV1CustomCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CustomCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CustomCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/custom_check] GetV1CustomCheck", response, response.Code())
	}
}

// NewGetV1CustomCheckOK creates a GetV1CustomCheckOK with default headers values
func NewGetV1CustomCheckOK() *GetV1CustomCheckOK {
	return &GetV1CustomCheckOK{}
}

/*
GetV1CustomCheckOK describes a response with status code 200, with default header values.

Success
*/
type GetV1CustomCheckOK struct {
	Payload *models.RESTCustomCheckListData
}

// IsSuccess returns true when this get v1 custom check o k response has a 2xx status code
func (o *GetV1CustomCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 custom check o k response has a 3xx status code
func (o *GetV1CustomCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 custom check o k response has a 4xx status code
func (o *GetV1CustomCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 custom check o k response has a 5xx status code
func (o *GetV1CustomCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 custom check o k response a status code equal to that given
func (o *GetV1CustomCheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 custom check o k response
func (o *GetV1CustomCheckOK) Code() int {
	return 200
}

func (o *GetV1CustomCheckOK) Error() string {
	return fmt.Sprintf("[GET /v1/custom_check][%d] getV1CustomCheckOK  %+v", 200, o.Payload)
}

func (o *GetV1CustomCheckOK) String() string {
	return fmt.Sprintf("[GET /v1/custom_check][%d] getV1CustomCheckOK  %+v", 200, o.Payload)
}

func (o *GetV1CustomCheckOK) GetPayload() *models.RESTCustomCheckListData {
	return o.Payload
}

func (o *GetV1CustomCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTCustomCheckListData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
