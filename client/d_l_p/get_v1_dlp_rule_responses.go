// Code generated by go-swagger; DO NOT EDIT.

package d_l_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1DlpRuleReader is a Reader for the GetV1DlpRule structure.
type GetV1DlpRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DlpRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DlpRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/dlp/rule] GetV1DlpRule", response, response.Code())
	}
}

// NewGetV1DlpRuleOK creates a GetV1DlpRuleOK with default headers values
func NewGetV1DlpRuleOK() *GetV1DlpRuleOK {
	return &GetV1DlpRuleOK{}
}

/*
GetV1DlpRuleOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DlpRuleOK struct {
	Payload *models.RESTDlpRulesData
}

// IsSuccess returns true when this get v1 dlp rule o k response has a 2xx status code
func (o *GetV1DlpRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 dlp rule o k response has a 3xx status code
func (o *GetV1DlpRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 dlp rule o k response has a 4xx status code
func (o *GetV1DlpRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 dlp rule o k response has a 5xx status code
func (o *GetV1DlpRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 dlp rule o k response a status code equal to that given
func (o *GetV1DlpRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 dlp rule o k response
func (o *GetV1DlpRuleOK) Code() int {
	return 200
}

func (o *GetV1DlpRuleOK) Error() string {
	return fmt.Sprintf("[GET /v1/dlp/rule][%d] getV1DlpRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1DlpRuleOK) String() string {
	return fmt.Sprintf("[GET /v1/dlp/rule][%d] getV1DlpRuleOK  %+v", 200, o.Payload)
}

func (o *GetV1DlpRuleOK) GetPayload() *models.RESTDlpRulesData {
	return o.Payload
}

func (o *GetV1DlpRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTDlpRulesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
