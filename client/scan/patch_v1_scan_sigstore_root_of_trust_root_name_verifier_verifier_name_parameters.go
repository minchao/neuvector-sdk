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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams creates a new PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams() *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	return &PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithTimeout creates a new PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams object
// with the ability to set a timeout on a request.
func NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithTimeout(timeout time.Duration) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	return &PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams{
		timeout: timeout,
	}
}

// NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithContext creates a new PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams object
// with the ability to set a context for a request.
func NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithContext(ctx context.Context) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	return &PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams{
		Context: ctx,
	}
}

// NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithHTTPClient creates a new PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParamsWithHTTPClient(client *http.Client) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	return &PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams{
		HTTPClient: client,
	}
}

/*
PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams contains all the parameters to send to the API endpoint

	for the patch v1 scan sigstore root of trust root name verifier verifier name operation.

	Typically these are written to a http.Request.
*/
type PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Verifier patch data
	*/
	Body *models.RESTSigstoreVerifierPatch

	/* RootName.

	   Root Of Trust Name
	*/
	RootName string

	/* VerifierName.

	   Verifier Name
	*/
	VerifierName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 scan sigstore root of trust root name verifier verifier name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithDefaults() *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 scan sigstore root of trust root name verifier verifier name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithTimeout(timeout time.Duration) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithContext(ctx context.Context) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithHTTPClient(client *http.Client) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithXAuthToken(xAuthToken string) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithBody(body *models.RESTSigstoreVerifierPatch) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetBody(body *models.RESTSigstoreVerifierPatch) {
	o.Body = body
}

// WithRootName adds the rootName to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithRootName(rootName string) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetRootName(rootName)
	return o
}

// SetRootName adds the rootName to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetRootName(rootName string) {
	o.RootName = rootName
}

// WithVerifierName adds the verifierName to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WithVerifierName(verifierName string) *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams {
	o.SetVerifierName(verifierName)
	return o
}

// SetVerifierName adds the verifierName to the patch v1 scan sigstore root of trust root name verifier verifier name params
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) SetVerifierName(verifierName string) {
	o.VerifierName = verifierName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ScanSigstoreRootOfTrustRootNameVerifierVerifierNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param root_name
	if err := r.SetPathParam("root_name", o.RootName); err != nil {
		return err
	}

	// path param verifier_name
	if err := r.SetPathParam("verifier_name", o.VerifierName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}