// Code generated by go-swagger; DO NOT EDIT.

package admission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1AdmissionStatsReader is a Reader for the GetV1AdmissionStats structure.
type GetV1AdmissionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AdmissionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AdmissionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/admission/stats] GetV1AdmissionStats", response, response.Code())
	}
}

// NewGetV1AdmissionStatsOK creates a GetV1AdmissionStatsOK with default headers values
func NewGetV1AdmissionStatsOK() *GetV1AdmissionStatsOK {
	return &GetV1AdmissionStatsOK{}
}

/*
GetV1AdmissionStatsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1AdmissionStatsOK struct {
	Payload *models.RESTAdmissionStatsData
}

// IsSuccess returns true when this get v1 admission stats o k response has a 2xx status code
func (o *GetV1AdmissionStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 admission stats o k response has a 3xx status code
func (o *GetV1AdmissionStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 admission stats o k response has a 4xx status code
func (o *GetV1AdmissionStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 admission stats o k response has a 5xx status code
func (o *GetV1AdmissionStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 admission stats o k response a status code equal to that given
func (o *GetV1AdmissionStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 admission stats o k response
func (o *GetV1AdmissionStatsOK) Code() int {
	return 200
}

func (o *GetV1AdmissionStatsOK) Error() string {
	return fmt.Sprintf("[GET /v1/admission/stats][%d] getV1AdmissionStatsOK  %+v", 200, o.Payload)
}

func (o *GetV1AdmissionStatsOK) String() string {
	return fmt.Sprintf("[GET /v1/admission/stats][%d] getV1AdmissionStatsOK  %+v", 200, o.Payload)
}

func (o *GetV1AdmissionStatsOK) GetPayload() *models.RESTAdmissionStatsData {
	return o.Payload
}

func (o *GetV1AdmissionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTAdmissionStatsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
