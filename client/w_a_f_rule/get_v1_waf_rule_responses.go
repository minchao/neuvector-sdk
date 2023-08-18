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

// GetV1WafRuleReader is a Reader for the GetV1WafRule structure.
type GetV1WafRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WafRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WafRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/waf/rule] GetV1WafRule", response, response.Code())
	}
}

// NewGetV1WafRuleOK creates a GetV1WafRuleOK with default headers values
func NewGetV1WafRuleOK() *GetV1WafRuleOK {
	return &GetV1WafRuleOK{}
}

/*
GetV1WafRuleOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WafRuleOK struct {
	Payload *models.RESTWafRulesData
}

// IsSuccess returns true when this get v1 waf rule o k response has a 2xx status code
func (o *GetV1WafRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 waf rule o k response has a 3xx status code
func (o *GetV1WafRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 waf rule o k response has a 4xx status code
func (o *GetV1WafRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 waf rule o k response has a 5xx status code
func (o *GetV1WafRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 waf rule o k response a status code equal to that given
func (o *GetV1WafRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 waf rule o k response
func (o *GetV1WafRuleOK) Code() int {
	return 200
}

func (o *GetV1WafRuleOK) Error() string {
	return fmt.Sprintf("[GET /v1/waf/rule][%d] getV1WafRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1WafRuleOK) String() string {
	return fmt.Sprintf("[GET /v1/waf/rule][%d] getV1WafRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1WafRuleOK) GetPayload() *models.RESTWafRulesData {
	return o.Payload
}

func (o *GetV1WafRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTWafRulesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}