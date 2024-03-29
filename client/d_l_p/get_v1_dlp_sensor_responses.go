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

// GetV1DlpSensorReader is a Reader for the GetV1DlpSensor structure.
type GetV1DlpSensorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DlpSensorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DlpSensorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/dlp/sensor] GetV1DlpSensor", response, response.Code())
	}
}

// NewGetV1DlpSensorOK creates a GetV1DlpSensorOK with default headers values
func NewGetV1DlpSensorOK() *GetV1DlpSensorOK {
	return &GetV1DlpSensorOK{}
}

/*
GetV1DlpSensorOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DlpSensorOK struct {
	Payload *models.RESTDlpSensorsData
}

// IsSuccess returns true when this get v1 dlp sensor o k response has a 2xx status code
func (o *GetV1DlpSensorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 dlp sensor o k response has a 3xx status code
func (o *GetV1DlpSensorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 dlp sensor o k response has a 4xx status code
func (o *GetV1DlpSensorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 dlp sensor o k response has a 5xx status code
func (o *GetV1DlpSensorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 dlp sensor o k response a status code equal to that given
func (o *GetV1DlpSensorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 dlp sensor o k response
func (o *GetV1DlpSensorOK) Code() int {
	return 200
}

func (o *GetV1DlpSensorOK) Error() string {
	return fmt.Sprintf("[GET /v1/dlp/sensor][%d] getV1DlpSensorOK  %+v", 200, o.Payload)
}

func (o *GetV1DlpSensorOK) String() string {
	return fmt.Sprintf("[GET /v1/dlp/sensor][%d] getV1DlpSensorOK  %+v", 200, o.Payload)
}

func (o *GetV1DlpSensorOK) GetPayload() *models.RESTDlpSensorsData {
	return o.Payload
}

func (o *GetV1DlpSensorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTDlpSensorsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
