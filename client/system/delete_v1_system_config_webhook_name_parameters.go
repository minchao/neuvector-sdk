// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewDeleteV1SystemConfigWebhookNameParams creates a new DeleteV1SystemConfigWebhookNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1SystemConfigWebhookNameParams() *DeleteV1SystemConfigWebhookNameParams {
	return &DeleteV1SystemConfigWebhookNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1SystemConfigWebhookNameParamsWithTimeout creates a new DeleteV1SystemConfigWebhookNameParams object
// with the ability to set a timeout on a request.
func NewDeleteV1SystemConfigWebhookNameParamsWithTimeout(timeout time.Duration) *DeleteV1SystemConfigWebhookNameParams {
	return &DeleteV1SystemConfigWebhookNameParams{
		timeout: timeout,
	}
}

// NewDeleteV1SystemConfigWebhookNameParamsWithContext creates a new DeleteV1SystemConfigWebhookNameParams object
// with the ability to set a context for a request.
func NewDeleteV1SystemConfigWebhookNameParamsWithContext(ctx context.Context) *DeleteV1SystemConfigWebhookNameParams {
	return &DeleteV1SystemConfigWebhookNameParams{
		Context: ctx,
	}
}

// NewDeleteV1SystemConfigWebhookNameParamsWithHTTPClient creates a new DeleteV1SystemConfigWebhookNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1SystemConfigWebhookNameParamsWithHTTPClient(client *http.Client) *DeleteV1SystemConfigWebhookNameParams {
	return &DeleteV1SystemConfigWebhookNameParams{
		HTTPClient: client,
	}
}

/*
DeleteV1SystemConfigWebhookNameParams contains all the parameters to send to the API endpoint

	for the delete v1 system config webhook name operation.

	Typically these are written to a http.Request.
*/
type DeleteV1SystemConfigWebhookNameParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   System webhook name
	*/
	Name string

	/* Scope.

	   When set the scope to be fed, it will delete the fed level webhook. When set the scope to be local, it will delete the local webhook. If there is no query string 'scope', it will use 'local' as the default value.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 system config webhook name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SystemConfigWebhookNameParams) WithDefaults() *DeleteV1SystemConfigWebhookNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 system config webhook name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SystemConfigWebhookNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithTimeout(timeout time.Duration) *DeleteV1SystemConfigWebhookNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithContext(ctx context.Context) *DeleteV1SystemConfigWebhookNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithHTTPClient(client *http.Client) *DeleteV1SystemConfigWebhookNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithXAuthToken(xAuthToken string) *DeleteV1SystemConfigWebhookNameParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithName(name string) *DeleteV1SystemConfigWebhookNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) WithScope(scope *string) *DeleteV1SystemConfigWebhookNameParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete v1 system config webhook name params
func (o *DeleteV1SystemConfigWebhookNameParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1SystemConfigWebhookNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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