// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1ScanRegistryNameScanReader is a Reader for the DeleteV1ScanRegistryNameScan structure.
type DeleteV1ScanRegistryNameScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ScanRegistryNameScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1ScanRegistryNameScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/scan/registry/{name}/scan] DeleteV1ScanRegistryNameScan", response, response.Code())
	}
}

// NewDeleteV1ScanRegistryNameScanOK creates a DeleteV1ScanRegistryNameScanOK with default headers values
func NewDeleteV1ScanRegistryNameScanOK() *DeleteV1ScanRegistryNameScanOK {
	return &DeleteV1ScanRegistryNameScanOK{}
}

/*
DeleteV1ScanRegistryNameScanOK describes a response with status code 200, with default header values.

Success
*/
type DeleteV1ScanRegistryNameScanOK struct {
}

// IsSuccess returns true when this delete v1 scan registry name scan o k response has a 2xx status code
func (o *DeleteV1ScanRegistryNameScanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 scan registry name scan o k response has a 3xx status code
func (o *DeleteV1ScanRegistryNameScanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 scan registry name scan o k response has a 4xx status code
func (o *DeleteV1ScanRegistryNameScanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 scan registry name scan o k response has a 5xx status code
func (o *DeleteV1ScanRegistryNameScanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 scan registry name scan o k response a status code equal to that given
func (o *DeleteV1ScanRegistryNameScanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 scan registry name scan o k response
func (o *DeleteV1ScanRegistryNameScanOK) Code() int {
	return 200
}

func (o *DeleteV1ScanRegistryNameScanOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/scan/registry/{name}/scan][%d] deleteV1ScanRegistryNameScanOK ", 200)
}

func (o *DeleteV1ScanRegistryNameScanOK) String() string {
	return fmt.Sprintf("[DELETE /v1/scan/registry/{name}/scan][%d] deleteV1ScanRegistryNameScanOK ", 200)
}

func (o *DeleteV1ScanRegistryNameScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}