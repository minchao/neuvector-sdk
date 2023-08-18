// Code generated by go-swagger; DO NOT EDIT.

package response_rule

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

// NewDeleteV1ResponseRuleParams creates a new DeleteV1ResponseRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ResponseRuleParams() *DeleteV1ResponseRuleParams {
	return &DeleteV1ResponseRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ResponseRuleParamsWithTimeout creates a new DeleteV1ResponseRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ResponseRuleParamsWithTimeout(timeout time.Duration) *DeleteV1ResponseRuleParams {
	return &DeleteV1ResponseRuleParams{
		timeout: timeout,
	}
}

// NewDeleteV1ResponseRuleParamsWithContext creates a new DeleteV1ResponseRuleParams object
// with the ability to set a context for a request.
func NewDeleteV1ResponseRuleParamsWithContext(ctx context.Context) *DeleteV1ResponseRuleParams {
	return &DeleteV1ResponseRuleParams{
		Context: ctx,
	}
}

// NewDeleteV1ResponseRuleParamsWithHTTPClient creates a new DeleteV1ResponseRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ResponseRuleParamsWithHTTPClient(client *http.Client) *DeleteV1ResponseRuleParams {
	return &DeleteV1ResponseRuleParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ResponseRuleParams contains all the parameters to send to the API endpoint

	for the delete v1 response rule operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ResponseRuleParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set to fed, all fed response rules get removed. When set to local or no query string, local response rules will be removed.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ResponseRuleParams) WithDefaults() *DeleteV1ResponseRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ResponseRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) WithTimeout(timeout time.Duration) *DeleteV1ResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) WithContext(ctx context.Context) *DeleteV1ResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) WithHTTPClient(client *http.Client) *DeleteV1ResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) WithXAuthToken(xAuthToken string) *DeleteV1ResponseRuleParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) WithScope(scope *string) *DeleteV1ResponseRuleParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete v1 response rule params
func (o *DeleteV1ResponseRuleParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
