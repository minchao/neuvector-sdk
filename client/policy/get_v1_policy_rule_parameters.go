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

// NewGetV1PolicyRuleParams creates a new GetV1PolicyRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1PolicyRuleParams() *GetV1PolicyRuleParams {
	return &GetV1PolicyRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PolicyRuleParamsWithTimeout creates a new GetV1PolicyRuleParams object
// with the ability to set a timeout on a request.
func NewGetV1PolicyRuleParamsWithTimeout(timeout time.Duration) *GetV1PolicyRuleParams {
	return &GetV1PolicyRuleParams{
		timeout: timeout,
	}
}

// NewGetV1PolicyRuleParamsWithContext creates a new GetV1PolicyRuleParams object
// with the ability to set a context for a request.
func NewGetV1PolicyRuleParamsWithContext(ctx context.Context) *GetV1PolicyRuleParams {
	return &GetV1PolicyRuleParams{
		Context: ctx,
	}
}

// NewGetV1PolicyRuleParamsWithHTTPClient creates a new GetV1PolicyRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1PolicyRuleParamsWithHTTPClient(client *http.Client) *GetV1PolicyRuleParams {
	return &GetV1PolicyRuleParams{
		HTTPClient: client,
	}
}

/*
GetV1PolicyRuleParams contains all the parameters to send to the API endpoint

	for the get v1 policy rule operation.

	Typically these are written to a http.Request.
*/
type GetV1PolicyRuleParams struct {

	/* Scope.

	   When set to fed, returned fed policy rules. When set to local, returned local policy rules. If there is no query string 'scope', all policy rules will be returned.
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 policy rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PolicyRuleParams) WithDefaults() *GetV1PolicyRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 policy rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PolicyRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) WithTimeout(timeout time.Duration) *GetV1PolicyRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) WithContext(ctx context.Context) *GetV1PolicyRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) WithHTTPClient(client *http.Client) *GetV1PolicyRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScope adds the scope to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) WithScope(scope *string) *GetV1PolicyRuleParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the get v1 policy rule params
func (o *GetV1PolicyRuleParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PolicyRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
