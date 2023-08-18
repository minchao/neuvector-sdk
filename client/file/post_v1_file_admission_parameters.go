// Code generated by go-swagger; DO NOT EDIT.

package file

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

// NewPostV1FileAdmissionParams creates a new PostV1FileAdmissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1FileAdmissionParams() *PostV1FileAdmissionParams {
	return &PostV1FileAdmissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1FileAdmissionParamsWithTimeout creates a new PostV1FileAdmissionParams object
// with the ability to set a timeout on a request.
func NewPostV1FileAdmissionParamsWithTimeout(timeout time.Duration) *PostV1FileAdmissionParams {
	return &PostV1FileAdmissionParams{
		timeout: timeout,
	}
}

// NewPostV1FileAdmissionParamsWithContext creates a new PostV1FileAdmissionParams object
// with the ability to set a context for a request.
func NewPostV1FileAdmissionParamsWithContext(ctx context.Context) *PostV1FileAdmissionParams {
	return &PostV1FileAdmissionParams{
		Context: ctx,
	}
}

// NewPostV1FileAdmissionParamsWithHTTPClient creates a new PostV1FileAdmissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1FileAdmissionParamsWithHTTPClient(client *http.Client) *PostV1FileAdmissionParams {
	return &PostV1FileAdmissionParams{
		HTTPClient: client,
	}
}

/*
PostV1FileAdmissionParams contains all the parameters to send to the API endpoint

	for the post v1 file admission operation.

	Typically these are written to a http.Request.
*/
type PostV1FileAdmissionParams struct {

	// XAuthToken.
	XAuthToken string

	/* Body.

	   Configuration data
	*/
	Body *models.RESTAdmCtrlRulesExport

	/* Scope.

	   It exports the admission control rules & state configurations when the scope is local.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 file admission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileAdmissionParams) WithDefaults() *PostV1FileAdmissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 file admission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FileAdmissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithTimeout(timeout time.Duration) *PostV1FileAdmissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithContext(ctx context.Context) *PostV1FileAdmissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithHTTPClient(client *http.Client) *PostV1FileAdmissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithXAuthToken(xAuthToken string) *PostV1FileAdmissionParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithBody adds the body to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithBody(body *models.RESTAdmCtrlRulesExport) *PostV1FileAdmissionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetBody(body *models.RESTAdmCtrlRulesExport) {
	o.Body = body
}

// WithScope adds the scope to the post v1 file admission params
func (o *PostV1FileAdmissionParams) WithScope(scope *string) *PostV1FileAdmissionParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the post v1 file admission params
func (o *PostV1FileAdmissionParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1FileAdmissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
