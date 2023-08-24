// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParams creates a new GetV1ScanSigstoreRootOfTrustRootNameVerifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParams() *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	return &GetV1ScanSigstoreRootOfTrustRootNameVerifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithTimeout creates a new GetV1ScanSigstoreRootOfTrustRootNameVerifierParams object
// with the ability to set a timeout on a request.
func NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithTimeout(timeout time.Duration) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	return &GetV1ScanSigstoreRootOfTrustRootNameVerifierParams{
		timeout: timeout,
	}
}

// NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithContext creates a new GetV1ScanSigstoreRootOfTrustRootNameVerifierParams object
// with the ability to set a context for a request.
func NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithContext(ctx context.Context) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	return &GetV1ScanSigstoreRootOfTrustRootNameVerifierParams{
		Context: ctx,
	}
}

// NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithHTTPClient creates a new GetV1ScanSigstoreRootOfTrustRootNameVerifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScanSigstoreRootOfTrustRootNameVerifierParamsWithHTTPClient(client *http.Client) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	return &GetV1ScanSigstoreRootOfTrustRootNameVerifierParams{
		HTTPClient: client,
	}
}

/*
GetV1ScanSigstoreRootOfTrustRootNameVerifierParams contains all the parameters to send to the API endpoint

	for the get v1 scan sigstore root of trust root name verifier operation.

	Typically these are written to a http.Request.
*/
type GetV1ScanSigstoreRootOfTrustRootNameVerifierParams struct {

	/* RootName.

	   Root Of Trust Name
	*/
	RootName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scan sigstore root of trust root name verifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WithDefaults() *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scan sigstore root of trust root name verifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WithTimeout(timeout time.Duration) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WithContext(ctx context.Context) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WithHTTPClient(client *http.Client) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRootName adds the rootName to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WithRootName(rootName string) *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams {
	o.SetRootName(rootName)
	return o
}

// SetRootName adds the rootName to the get v1 scan sigstore root of trust root name verifier params
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) SetRootName(rootName string) {
	o.RootName = rootName
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScanSigstoreRootOfTrustRootNameVerifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param root_name
	if err := r.SetPathParam("root_name", o.RootName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
