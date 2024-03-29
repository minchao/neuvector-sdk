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

// GetV1ScanRegistryNameReader is a Reader for the GetV1ScanRegistryName structure.
type GetV1ScanRegistryNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ScanRegistryNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ScanRegistryNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/scan/registry/{name}] GetV1ScanRegistryName", response, response.Code())
	}
}

// NewGetV1ScanRegistryNameOK creates a GetV1ScanRegistryNameOK with default headers values
func NewGetV1ScanRegistryNameOK() *GetV1ScanRegistryNameOK {
	return &GetV1ScanRegistryNameOK{}
}

/*
GetV1ScanRegistryNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ScanRegistryNameOK struct {
	Payload *models.RESTRegistrySummaryData
}

// IsSuccess returns true when this get v1 scan registry name o k response has a 2xx status code
func (o *GetV1ScanRegistryNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 scan registry name o k response has a 3xx status code
func (o *GetV1ScanRegistryNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 scan registry name o k response has a 4xx status code
func (o *GetV1ScanRegistryNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 scan registry name o k response has a 5xx status code
func (o *GetV1ScanRegistryNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 scan registry name o k response a status code equal to that given
func (o *GetV1ScanRegistryNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 scan registry name o k response
func (o *GetV1ScanRegistryNameOK) Code() int {
	return 200
}

func (o *GetV1ScanRegistryNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/scan/registry/{name}][%d] getV1ScanRegistryNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanRegistryNameOK) String() string {
	return fmt.Sprintf("[GET /v1/scan/registry/{name}][%d] getV1ScanRegistryNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanRegistryNameOK) GetPayload() *models.RESTRegistrySummaryData {
	return o.Payload
}

func (o *GetV1ScanRegistryNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTRegistrySummaryData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
