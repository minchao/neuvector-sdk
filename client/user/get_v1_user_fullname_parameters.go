// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetV1UserFullnameParams creates a new GetV1UserFullnameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1UserFullnameParams() *GetV1UserFullnameParams {
	return &GetV1UserFullnameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1UserFullnameParamsWithTimeout creates a new GetV1UserFullnameParams object
// with the ability to set a timeout on a request.
func NewGetV1UserFullnameParamsWithTimeout(timeout time.Duration) *GetV1UserFullnameParams {
	return &GetV1UserFullnameParams{
		timeout: timeout,
	}
}

// NewGetV1UserFullnameParamsWithContext creates a new GetV1UserFullnameParams object
// with the ability to set a context for a request.
func NewGetV1UserFullnameParamsWithContext(ctx context.Context) *GetV1UserFullnameParams {
	return &GetV1UserFullnameParams{
		Context: ctx,
	}
}

// NewGetV1UserFullnameParamsWithHTTPClient creates a new GetV1UserFullnameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1UserFullnameParamsWithHTTPClient(client *http.Client) *GetV1UserFullnameParams {
	return &GetV1UserFullnameParams{
		HTTPClient: client,
	}
}

/*
GetV1UserFullnameParams contains all the parameters to send to the API endpoint

	for the get v1 user fullname operation.

	Typically these are written to a http.Request.
*/
type GetV1UserFullnameParams struct {

	/* Fullname.

	   User name
	*/
	Fullname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 user fullname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserFullnameParams) WithDefaults() *GetV1UserFullnameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 user fullname params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UserFullnameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 user fullname params
func (o *GetV1UserFullnameParams) WithTimeout(timeout time.Duration) *GetV1UserFullnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 user fullname params
func (o *GetV1UserFullnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 user fullname params
func (o *GetV1UserFullnameParams) WithContext(ctx context.Context) *GetV1UserFullnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 user fullname params
func (o *GetV1UserFullnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 user fullname params
func (o *GetV1UserFullnameParams) WithHTTPClient(client *http.Client) *GetV1UserFullnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 user fullname params
func (o *GetV1UserFullnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFullname adds the fullname to the get v1 user fullname params
func (o *GetV1UserFullnameParams) WithFullname(fullname string) *GetV1UserFullnameParams {
	o.SetFullname(fullname)
	return o
}

// SetFullname adds the fullname to the get v1 user fullname params
func (o *GetV1UserFullnameParams) SetFullname(fullname string) {
	o.Fullname = fullname
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1UserFullnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fullname
	if err := r.SetPathParam("fullname", o.Fullname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
