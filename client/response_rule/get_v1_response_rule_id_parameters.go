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

// NewGetV1ResponseRuleIDParams creates a new GetV1ResponseRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ResponseRuleIDParams() *GetV1ResponseRuleIDParams {
	return &GetV1ResponseRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ResponseRuleIDParamsWithTimeout creates a new GetV1ResponseRuleIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ResponseRuleIDParamsWithTimeout(timeout time.Duration) *GetV1ResponseRuleIDParams {
	return &GetV1ResponseRuleIDParams{
		timeout: timeout,
	}
}

// NewGetV1ResponseRuleIDParamsWithContext creates a new GetV1ResponseRuleIDParams object
// with the ability to set a context for a request.
func NewGetV1ResponseRuleIDParamsWithContext(ctx context.Context) *GetV1ResponseRuleIDParams {
	return &GetV1ResponseRuleIDParams{
		Context: ctx,
	}
}

// NewGetV1ResponseRuleIDParamsWithHTTPClient creates a new GetV1ResponseRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ResponseRuleIDParamsWithHTTPClient(client *http.Client) *GetV1ResponseRuleIDParams {
	return &GetV1ResponseRuleIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ResponseRuleIDParams contains all the parameters to send to the API endpoint

	for the get v1 response rule ID operation.

	Typically these are written to a http.Request.
*/
type GetV1ResponseRuleIDParams struct {

	/* ID.

	   Rule ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 response rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseRuleIDParams) WithDefaults() *GetV1ResponseRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 response rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ResponseRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) WithTimeout(timeout time.Duration) *GetV1ResponseRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) WithContext(ctx context.Context) *GetV1ResponseRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) WithHTTPClient(client *http.Client) *GetV1ResponseRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) WithID(id string) *GetV1ResponseRuleIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 response rule ID params
func (o *GetV1ResponseRuleIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ResponseRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
