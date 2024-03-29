// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1ScanSigstoreRootOfTrustRootNameVerifierReader is a Reader for the PostV1ScanSigstoreRootOfTrustRootNameVerifier structure.
type PostV1ScanSigstoreRootOfTrustRootNameVerifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1ScanSigstoreRootOfTrustRootNameVerifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/scan/sigstore/root_of_trust/{root_name}/verifier] PostV1ScanSigstoreRootOfTrustRootNameVerifier", response, response.Code())
	}
}

// NewPostV1ScanSigstoreRootOfTrustRootNameVerifierOK creates a PostV1ScanSigstoreRootOfTrustRootNameVerifierOK with default headers values
func NewPostV1ScanSigstoreRootOfTrustRootNameVerifierOK() *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK {
	return &PostV1ScanSigstoreRootOfTrustRootNameVerifierOK{}
}

/*
PostV1ScanSigstoreRootOfTrustRootNameVerifierOK describes a response with status code 200, with default header values.

Success
*/
type PostV1ScanSigstoreRootOfTrustRootNameVerifierOK struct {
}

// IsSuccess returns true when this post v1 scan sigstore root of trust root name verifier o k response has a 2xx status code
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 scan sigstore root of trust root name verifier o k response has a 3xx status code
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 scan sigstore root of trust root name verifier o k response has a 4xx status code
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 scan sigstore root of trust root name verifier o k response has a 5xx status code
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 scan sigstore root of trust root name verifier o k response a status code equal to that given
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 scan sigstore root of trust root name verifier o k response
func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) Code() int {
	return 200
}

func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) Error() string {
	return fmt.Sprintf("[POST /v1/scan/sigstore/root_of_trust/{root_name}/verifier][%d] postV1ScanSigstoreRootOfTrustRootNameVerifierOK ", 200)
}

func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) String() string {
	return fmt.Sprintf("[POST /v1/scan/sigstore/root_of_trust/{root_name}/verifier][%d] postV1ScanSigstoreRootOfTrustRootNameVerifierOK ", 200)
}

func (o *PostV1ScanSigstoreRootOfTrustRootNameVerifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
