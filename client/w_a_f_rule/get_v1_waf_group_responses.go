// Code generated by go-swagger; DO NOT EDIT.

package w_a_f_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1WafGroupReader is a Reader for the GetV1WafGroup structure.
type GetV1WafGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WafGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WafGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/waf/group] GetV1WafGroup", response, response.Code())
	}
}

// NewGetV1WafGroupOK creates a GetV1WafGroupOK with default headers values
func NewGetV1WafGroupOK() *GetV1WafGroupOK {
	return &GetV1WafGroupOK{}
}

/*
GetV1WafGroupOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WafGroupOK struct {
	Payload *models.RESTWafGroupsData
}

// IsSuccess returns true when this get v1 waf group o k response has a 2xx status code
func (o *GetV1WafGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 waf group o k response has a 3xx status code
func (o *GetV1WafGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 waf group o k response has a 4xx status code
func (o *GetV1WafGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 waf group o k response has a 5xx status code
func (o *GetV1WafGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 waf group o k response a status code equal to that given
func (o *GetV1WafGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 waf group o k response
func (o *GetV1WafGroupOK) Code() int {
	return 200
}

func (o *GetV1WafGroupOK) Error() string {
	return fmt.Sprintf("[GET /v1/waf/group][%d] getV1WafGroupOK  %+v", 200, o.Payload)
}

func (o *GetV1WafGroupOK) String() string {
	return fmt.Sprintf("[GET /v1/waf/group][%d] getV1WafGroupOK  %+v", 200, o.Payload)
}

func (o *GetV1WafGroupOK) GetPayload() *models.RESTWafGroupsData {
	return o.Payload
}

func (o *GetV1WafGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWafGroupsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
