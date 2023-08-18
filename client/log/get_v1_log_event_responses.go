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

// GetV1LogEventReader is a Reader for the GetV1LogEvent structure.
type GetV1LogEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1LogEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1LogEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/log/event] GetV1LogEvent", response, response.Code())
	}
}

// NewGetV1LogEventOK creates a GetV1LogEventOK with default headers values
func NewGetV1LogEventOK() *GetV1LogEventOK {
	return &GetV1LogEventOK{}
}

/*
GetV1LogEventOK describes a response with status code 200, with default header values.

Success
*/
type GetV1LogEventOK struct {
	Payload *models.RESTEventsData
}

// IsSuccess returns true when this get v1 log event o k response has a 2xx status code
func (o *GetV1LogEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 log event o k response has a 3xx status code
func (o *GetV1LogEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 log event o k response has a 4xx status code
func (o *GetV1LogEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 log event o k response has a 5xx status code
func (o *GetV1LogEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 log event o k response a status code equal to that given
func (o *GetV1LogEventOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 log event o k response
func (o *GetV1LogEventOK) Code() int {
	return 200
}

func (o *GetV1LogEventOK) Error() string {
	return fmt.Sprintf("[GET /v1/log/event][%d] getV1LogEventOK  %+v", 200, o.Payload)
}

func (o *GetV1LogEventOK) String() string {
	return fmt.Sprintf("[GET /v1/log/event][%d] getV1LogEventOK  %+v", 200, o.Payload)
}

func (o *GetV1LogEventOK) GetPayload() *models.RESTEventsData {
	return o.Payload
}

func (o *GetV1LogEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTEventsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
