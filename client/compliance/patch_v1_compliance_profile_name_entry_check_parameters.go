// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

// NewPatchV1ComplianceProfileNameEntryCheckParams creates a new PatchV1ComplianceProfileNameEntryCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ComplianceProfileNameEntryCheckParams() *PatchV1ComplianceProfileNameEntryCheckParams {
	return &PatchV1ComplianceProfileNameEntryCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ComplianceProfileNameEntryCheckParamsWithTimeout creates a new PatchV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a timeout on a request.
func NewPatchV1ComplianceProfileNameEntryCheckParamsWithTimeout(timeout time.Duration) *PatchV1ComplianceProfileNameEntryCheckParams {
	return &PatchV1ComplianceProfileNameEntryCheckParams{
		timeout: timeout,
	}
}

// NewPatchV1ComplianceProfileNameEntryCheckParamsWithContext creates a new PatchV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a context for a request.
func NewPatchV1ComplianceProfileNameEntryCheckParamsWithContext(ctx context.Context) *PatchV1ComplianceProfileNameEntryCheckParams {
	return &PatchV1ComplianceProfileNameEntryCheckParams{
		Context: ctx,
	}
}

// NewPatchV1ComplianceProfileNameEntryCheckParamsWithHTTPClient creates a new PatchV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ComplianceProfileNameEntryCheckParamsWithHTTPClient(client *http.Client) *PatchV1ComplianceProfileNameEntryCheckParams {
	return &PatchV1ComplianceProfileNameEntryCheckParams{
		HTTPClient: client,
	}
}

/*
PatchV1ComplianceProfileNameEntryCheckParams contains all the parameters to send to the API endpoint

	for the patch v1 compliance profile name entry check operation.

	Typically these are written to a http.Request.
*/
type PatchV1ComplianceProfileNameEntryCheckParams struct {

	/* Body.

	   Compliance profile entry data
	*/
	Body *models.RESTComplianceProfileEntryConfigData

	/* Check.

	   Compliance profile entry check name
	*/
	Check string

	/* Name.

	   Compliance profile name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 compliance profile name entry check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithDefaults() *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 compliance profile name entry check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithTimeout(timeout time.Duration) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithContext(ctx context.Context) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithHTTPClient(client *http.Client) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithBody(body *models.RESTComplianceProfileEntryConfigData) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetBody(body *models.RESTComplianceProfileEntryConfigData) {
	o.Body = body
}

// WithCheck adds the check to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithCheck(check string) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetCheck(check)
	return o
}

// SetCheck adds the check to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetCheck(check string) {
	o.Check = check
}

// WithName adds the name to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WithName(name string) *PatchV1ComplianceProfileNameEntryCheckParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 compliance profile name entry check params
func (o *PatchV1ComplianceProfileNameEntryCheckParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ComplianceProfileNameEntryCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param check
	if err := r.SetPathParam("check", o.Check); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
