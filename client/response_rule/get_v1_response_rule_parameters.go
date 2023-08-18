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

// NewGetV1ResponseRuleParams creates a new GetV1ResponseRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ResponseRuleParams() *GetV1ResponseRuleParams {
	return &GetV1ResponseRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ResponseRuleParamsWithTimeout creates a new GetV1ResponseRuleParams object
// with the ability to set a timeout on a request.
func NewGetV1ResponseRuleParamsWithTimeout(timeout time.Duration) *GetV1ResponseRuleParams {
	return &GetV1ResponseRuleParams{
		timeout: timeout,
	}
}

// NewGetV1ResponseRuleParamsWithContext creates a new GetV1ResponseRuleParams object
// with the ability to set a context for a request.
func NewGetV1ResponseRuleParamsWithContext(ctx context.Context) *GetV1ResponseRuleParams {
	return &GetV1ResponseRuleParams{
		Context: ctx,
	}
}

// NewGetV1ResponseRuleParamsWithHTTPClient creates a new GetV1ResponseRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ResponseRuleParamsWithHTTPClient(client *http.Client) *GetV1ResponseRuleParams {
	return &GetV1ResponseRuleParams{
		HTTPClient: client,
	}
}

/*
GetV1ResponseRuleParams contains all the parameters to send to the API endpoint

	for the get v1 response rule operation.

	Typically these are written to a http.Request.
*/
type GetV1ResponseRuleParams struct {

	// XAuthToken.
	XAuthToken string

	/* Scope.

	   When set to fed, returned fed response rules. When set to local, returned local response rules. If there is no query string 'scope', all response rules will be returned.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseRuleParams) WithDefaults() *GetV1ResponseRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 response rule params
func (o *GetV1ResponseRuleParams) WithTimeout(timeout time.Duration) *GetV1ResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 response rule params
func (o *GetV1ResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 response rule params
func (o *GetV1ResponseRuleParams) WithContext(ctx context.Context) *GetV1ResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 response rule params
func (o *GetV1ResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 response rule params
func (o *GetV1ResponseRuleParams) WithHTTPClient(client *http.Client) *GetV1ResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 response rule params
func (o *GetV1ResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get v1 response rule params
func (o *GetV1ResponseRuleParams) WithXAuthToken(xAuthToken string) *GetV1ResponseRuleParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get v1 response rule params
func (o *GetV1ResponseRuleParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithScope adds the scope to the get v1 response rule params
func (o *GetV1ResponseRuleParams) WithScope(scope *string) *GetV1ResponseRuleParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get v1 response rule params
func (o *GetV1ResponseRuleParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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