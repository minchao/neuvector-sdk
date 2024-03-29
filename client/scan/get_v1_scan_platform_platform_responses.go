// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/neuvector-sdk/models"
)

// GetV1ScanPlatformPlatformReader is a Reader for the GetV1ScanPlatformPlatform structure.
type GetV1ScanPlatformPlatformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ScanPlatformPlatformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ScanPlatformPlatformOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/scan/platform/platform] GetV1ScanPlatformPlatform", response, response.Code())
	}
}

// NewGetV1ScanPlatformPlatformOK creates a GetV1ScanPlatformPlatformOK with default headers values
func NewGetV1ScanPlatformPlatformOK() *GetV1ScanPlatformPlatformOK {
	return &GetV1ScanPlatformPlatformOK{}
}

/*
GetV1ScanPlatformPlatformOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ScanPlatformPlatformOK struct {
	Payload *models.RESTScanReportData
}

// IsSuccess returns true when this get v1 scan platform platform o k response has a 2xx status code
func (o *GetV1ScanPlatformPlatformOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 scan platform platform o k response has a 3xx status code
func (o *GetV1ScanPlatformPlatformOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 scan platform platform o k response has a 4xx status code
func (o *GetV1ScanPlatformPlatformOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 scan platform platform o k response has a 5xx status code
func (o *GetV1ScanPlatformPlatformOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 scan platform platform o k response a status code equal to that given
func (o *GetV1ScanPlatformPlatformOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 scan platform platform o k response
func (o *GetV1ScanPlatformPlatformOK) Code() int {
	return 200
}

func (o *GetV1ScanPlatformPlatformOK) Error() string {
	return fmt.Sprintf("[GET /v1/scan/platform/platform][%d] getV1ScanPlatformPlatformOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanPlatformPlatformOK) String() string {
	return fmt.Sprintf("[GET /v1/scan/platform/platform][%d] getV1ScanPlatformPlatformOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanPlatformPlatformOK) GetPayload() *models.RESTScanReportData {
	return o.Payload
}

func (o *GetV1ScanPlatformPlatformOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTScanReportData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
