// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1LogThreatIDReader is a Reader for the GetV1LogThreatID structure.
type GetV1LogThreatIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1LogThreatIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1LogThreatIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/log/threat/{id}] GetV1LogThreatID", response, response.Code())
	}
}

// NewGetV1LogThreatIDOK creates a GetV1LogThreatIDOK with default headers values
func NewGetV1LogThreatIDOK() *GetV1LogThreatIDOK {
	return &GetV1LogThreatIDOK{}
}

/*
GetV1LogThreatIDOK describes a response with status code 200, with default header values.

Success
*/
type GetV1LogThreatIDOK struct {
	Payload *models.RESTThreatData
}

// IsSuccess returns true when this get v1 log threat Id o k response has a 2xx status code
func (o *GetV1LogThreatIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 log threat Id o k response has a 3xx status code
func (o *GetV1LogThreatIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 log threat Id o k response has a 4xx status code
func (o *GetV1LogThreatIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 log threat Id o k response has a 5xx status code
func (o *GetV1LogThreatIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 log threat Id o k response a status code equal to that given
func (o *GetV1LogThreatIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 log threat Id o k response
func (o *GetV1LogThreatIDOK) Code() int {
	return 200
}

func (o *GetV1LogThreatIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/log/threat/{id}][%d] getV1LogThreatIdOK  %+v", 200, o.Payload)
}

func (o *GetV1LogThreatIDOK) String() string {
	return fmt.Sprintf("[GET /v1/log/threat/{id}][%d] getV1LogThreatIdOK  %+v", 200, o.Payload)
}

func (o *GetV1LogThreatIDOK) GetPayload() *models.RESTThreatData {
	return o.Payload
}

func (o *GetV1LogThreatIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTThreatData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}