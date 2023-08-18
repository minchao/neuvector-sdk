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

// GetV1ScanSigstoreRootOfTrustRootNameReader is a Reader for the GetV1ScanSigstoreRootOfTrustRootName structure.
type GetV1ScanSigstoreRootOfTrustRootNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ScanSigstoreRootOfTrustRootNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ScanSigstoreRootOfTrustRootNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/scan/sigstore/root_of_trust/{root_name}] GetV1ScanSigstoreRootOfTrustRootName", response, response.Code())
	}
}

// NewGetV1ScanSigstoreRootOfTrustRootNameOK creates a GetV1ScanSigstoreRootOfTrustRootNameOK with default headers values
func NewGetV1ScanSigstoreRootOfTrustRootNameOK() *GetV1ScanSigstoreRootOfTrustRootNameOK {
	return &GetV1ScanSigstoreRootOfTrustRootNameOK{}
}

/*
GetV1ScanSigstoreRootOfTrustRootNameOK describes a response with status code 200, with default header values.

Success
*/
type GetV1ScanSigstoreRootOfTrustRootNameOK struct {
	Payload *models.RESTSigstoreRootOfTrustGet
}

// IsSuccess returns true when this get v1 scan sigstore root of trust root name o k response has a 2xx status code
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 scan sigstore root of trust root name o k response has a 3xx status code
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 scan sigstore root of trust root name o k response has a 4xx status code
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 scan sigstore root of trust root name o k response has a 5xx status code
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 scan sigstore root of trust root name o k response a status code equal to that given
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 scan sigstore root of trust root name o k response
func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) Code() int {
	return 200
}

func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/scan/sigstore/root_of_trust/{root_name}][%d] getV1ScanSigstoreRootOfTrustRootNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) String() string {
	return fmt.Sprintf("[GET /v1/scan/sigstore/root_of_trust/{root_name}][%d] getV1ScanSigstoreRootOfTrustRootNameOK  %+v", 200, o.Payload)
}

func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) GetPayload() *models.RESTSigstoreRootOfTrustGet {
	return o.Payload
}

func (o *GetV1ScanSigstoreRootOfTrustRootNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RESTSigstoreRootOfTrustGet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
