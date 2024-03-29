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

// GetV1LogIncidentReader is a Reader for the GetV1LogIncident structure.
type GetV1LogIncidentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1LogIncidentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1LogIncidentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/log/incident] GetV1LogIncident", response, response.Code())
	}
}

// NewGetV1LogIncidentOK creates a GetV1LogIncidentOK with default headers values
func NewGetV1LogIncidentOK() *GetV1LogIncidentOK {
	return &GetV1LogIncidentOK{}
}

/*
GetV1LogIncidentOK describes a response with status code 200, with default header values.

Success
*/
type GetV1LogIncidentOK struct {
	Payload *models.RESTIncidentsData
}

// IsSuccess returns true when this get v1 log incident o k response has a 2xx status code
func (o *GetV1LogIncidentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 log incident o k response has a 3xx status code
func (o *GetV1LogIncidentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 log incident o k response has a 4xx status code
func (o *GetV1LogIncidentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 log incident o k response has a 5xx status code
func (o *GetV1LogIncidentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 log incident o k response a status code equal to that given
func (o *GetV1LogIncidentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 log incident o k response
func (o *GetV1LogIncidentOK) Code() int {
	return 200
}

func (o *GetV1LogIncidentOK) Error() string {
	return fmt.Sprintf("[GET /v1/log/incident][%d] getV1LogIncidentOK  %+v", 200, o.Payload)
}

func (o *GetV1LogIncidentOK) String() string {
	return fmt.Sprintf("[GET /v1/log/incident][%d] getV1LogIncidentOK  %+v", 200, o.Payload)
}

func (o *GetV1LogIncidentOK) GetPayload() *models.RESTIncidentsData {
	return o.Payload
}

func (o *GetV1LogIncidentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTIncidentsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
