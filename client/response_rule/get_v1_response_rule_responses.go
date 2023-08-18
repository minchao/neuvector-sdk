// Code generated by go-swagger; DO NOT EDIT.

package response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ResponseRuleReader is a Reader for the GetV1ResponseRule structure.
type GetV1ResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/response/rule] GetV1ResponseRule", response, response.Code())
	}
}

// NewGetV1ResponseRuleOK creates a GetV1ResponseRuleOK with default headers values
func NewGetV1ResponseRuleOK() *GetV1ResponseRuleOK {
	return &GetV1ResponseRuleOK{}
}

/*
GetV1ResponseRuleOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ResponseRuleOK struct {
	Payload *models.RESTResponseRulesData
}

// IsSuccess returns true when this get v1 response rule o k response has a 2xx status code
func (o *GetV1ResponseRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 response rule o k response has a 3xx status code
func (o *GetV1ResponseRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 response rule o k response has a 4xx status code
func (o *GetV1ResponseRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 response rule o k response has a 5xx status code
func (o *GetV1ResponseRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 response rule o k response a status code equal to that given
func (o *GetV1ResponseRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 response rule o k response
func (o *GetV1ResponseRuleOK) Code() int {
	return 200
}

func (o *GetV1ResponseRuleOK) Error() string {
	return fmt.Sprintf("[GET /v1/response/rule][%d] getV1ResponseRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1ResponseRuleOK) String() string {
	return fmt.Sprintf("[GET /v1/response/rule][%d] getV1ResponseRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1ResponseRuleOK) GetPayload() *models.RESTResponseRulesData {
	return o.Payload
}

func (o *GetV1ResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTResponseRulesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}