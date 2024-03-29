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

// GetV1WafSersorNameReader is a Reader for the GetV1WafSersorName structure.
type GetV1WafSersorNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WafSersorNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WafSersorNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/waf/sersor/{name}] GetV1WafSersorName", response, response.Code())
	}
}

// NewGetV1WafSersorNameOK creates a GetV1WafSersorNameOK with default headers values
func NewGetV1WafSersorNameOK() *GetV1WafSersorNameOK {
	return &GetV1WafSersorNameOK{}
}

/*
GetV1WafSersorNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WafSersorNameOK struct {
	Payload *models.RESTWafSensorData
}

// IsSuccess returns true when this get v1 waf sersor name o k response has a 2xx status code
func (o *GetV1WafSersorNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 waf sersor name o k response has a 3xx status code
func (o *GetV1WafSersorNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 waf sersor name o k response has a 4xx status code
func (o *GetV1WafSersorNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 waf sersor name o k response has a 5xx status code
func (o *GetV1WafSersorNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 waf sersor name o k response a status code equal to that given
func (o *GetV1WafSersorNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 waf sersor name o k response
func (o *GetV1WafSersorNameOK) Code() int {
	return 200
}

func (o *GetV1WafSersorNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/waf/sersor/{name}][%d] getV1WafSersorNameOK  %+v", 200, o.Payload)
}

func (o *GetV1WafSersorNameOK) String() string {
	return fmt.Sprintf("[GET /v1/waf/sersor/{name}][%d] getV1WafSersorNameOK  %+v", 200, o.Payload)
}

func (o *GetV1WafSersorNameOK) GetPayload() *models.RESTWafSensorData {
	return o.Payload
}

func (o *GetV1WafSersorNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWafSensorData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
