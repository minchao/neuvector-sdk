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
)

// NewDeleteV1ComplianceProfileNameEntryCheckParams creates a new DeleteV1ComplianceProfileNameEntryCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ComplianceProfileNameEntryCheckParams() *DeleteV1ComplianceProfileNameEntryCheckParams {
	return &DeleteV1ComplianceProfileNameEntryCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ComplianceProfileNameEntryCheckParamsWithTimeout creates a new DeleteV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ComplianceProfileNameEntryCheckParamsWithTimeout(timeout time.Duration) *DeleteV1ComplianceProfileNameEntryCheckParams {
	return &DeleteV1ComplianceProfileNameEntryCheckParams{
		timeout: timeout,
	}
}

// NewDeleteV1ComplianceProfileNameEntryCheckParamsWithContext creates a new DeleteV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a context for a request.
func NewDeleteV1ComplianceProfileNameEntryCheckParamsWithContext(ctx context.Context) *DeleteV1ComplianceProfileNameEntryCheckParams {
	return &DeleteV1ComplianceProfileNameEntryCheckParams{
		Context: ctx,
	}
}

// NewDeleteV1ComplianceProfileNameEntryCheckParamsWithHTTPClient creates a new DeleteV1ComplianceProfileNameEntryCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ComplianceProfileNameEntryCheckParamsWithHTTPClient(client *http.Client) *DeleteV1ComplianceProfileNameEntryCheckParams {
	return &DeleteV1ComplianceProfileNameEntryCheckParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ComplianceProfileNameEntryCheckParams contains all the parameters to send to the API endpoint

	for the delete v1 compliance profile name entry check operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ComplianceProfileNameEntryCheckParams struct {

	// XAuthToken.
	XAuthToken string

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

// WithDefaults hydrates default values in the delete v1 compliance profile name entry check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithDefaults() *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 compliance profile name entry check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithTimeout(timeout time.Duration) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithContext(ctx context.Context) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithHTTPClient(client *http.Client) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithXAuthToken(xAuthToken string) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithCheck adds the check to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithCheck(check string) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetCheck(check)
	return o
}

// SetCheck adds the check to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetCheck(check string) {
	o.Check = check
}

// WithName adds the name to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WithName(name string) *DeleteV1ComplianceProfileNameEntryCheckParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete v1 compliance profile name entry check params
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ComplianceProfileNameEntryCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
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
