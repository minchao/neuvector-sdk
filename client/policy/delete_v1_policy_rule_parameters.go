// Code generated by go-swagger; DO NOT EDIT.

package policy

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

// NewDeleteV1PolicyRuleParams creates a new DeleteV1PolicyRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1PolicyRuleParams() *DeleteV1PolicyRuleParams {
	return &DeleteV1PolicyRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1PolicyRuleParamsWithTimeout creates a new DeleteV1PolicyRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteV1PolicyRuleParamsWithTimeout(timeout time.Duration) *DeleteV1PolicyRuleParams {
	return &DeleteV1PolicyRuleParams{
		timeout: timeout,
	}
}

// NewDeleteV1PolicyRuleParamsWithContext creates a new DeleteV1PolicyRuleParams object
// with the ability to set a context for a request.
func NewDeleteV1PolicyRuleParamsWithContext(ctx context.Context) *DeleteV1PolicyRuleParams {
	return &DeleteV1PolicyRuleParams{
		Context: ctx,
	}
}

// NewDeleteV1PolicyRuleParamsWithHTTPClient creates a new DeleteV1PolicyRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1PolicyRuleParamsWithHTTPClient(client *http.Client) *DeleteV1PolicyRuleParams {
	return &DeleteV1PolicyRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteV1PolicyRuleParams contains all the parameters to send to the API endpoint

	for the delete v1 policy rule operation.

	Typically these are written to a http.Request.
*/
type DeleteV1PolicyRuleParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set to fed, fed policy rules get removed. When set to local or no query string, local policy rules get removed.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 policy rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PolicyRuleParams) WithDefaults() *DeleteV1PolicyRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 policy rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1PolicyRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) WithTimeout(timeout time.Duration) *DeleteV1PolicyRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) WithContext(ctx context.Context) *DeleteV1PolicyRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) WithHTTPClient(client *http.Client) *DeleteV1PolicyRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) WithXAuthToken(xAuthToken string) *DeleteV1PolicyRuleParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) WithScope(scope *string) *DeleteV1PolicyRuleParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete v1 policy rule params
func (o *DeleteV1PolicyRuleParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1PolicyRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Auth-Token
	if err := r.SetHeaderParam("X-Auth-Token", o.XAuthToken); err != nil {
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
