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

// NewGetV1BenchHostIDKubernetesParams creates a new GetV1BenchHostIDKubernetesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1BenchHostIDKubernetesParams() *GetV1BenchHostIDKubernetesParams {
	return &GetV1BenchHostIDKubernetesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1BenchHostIDKubernetesParamsWithTimeout creates a new GetV1BenchHostIDKubernetesParams object
// with the ability to set a timeout on a request.
func NewGetV1BenchHostIDKubernetesParamsWithTimeout(timeout time.Duration) *GetV1BenchHostIDKubernetesParams {
	return &GetV1BenchHostIDKubernetesParams{
		timeout: timeout,
	}
}

// NewGetV1BenchHostIDKubernetesParamsWithContext creates a new GetV1BenchHostIDKubernetesParams object
// with the ability to set a context for a request.
func NewGetV1BenchHostIDKubernetesParamsWithContext(ctx context.Context) *GetV1BenchHostIDKubernetesParams {
	return &GetV1BenchHostIDKubernetesParams{
		Context: ctx,
	}
}

// NewGetV1BenchHostIDKubernetesParamsWithHTTPClient creates a new GetV1BenchHostIDKubernetesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1BenchHostIDKubernetesParamsWithHTTPClient(client *http.Client) *GetV1BenchHostIDKubernetesParams {
	return &GetV1BenchHostIDKubernetesParams{
		HTTPClient: client,
	}
}

/*
GetV1BenchHostIDKubernetesParams contains all the parameters to send to the API endpoint

	for the get v1 bench host ID kubernetes operation.

	Typically these are written to a http.Request.
*/
type GetV1BenchHostIDKubernetesParams struct {

	/* ID.

	   Host ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 bench host ID kubernetes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1BenchHostIDKubernetesParams) WithDefaults() *GetV1BenchHostIDKubernetesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 bench host ID kubernetes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1BenchHostIDKubernetesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) WithTimeout(timeout time.Duration) *GetV1BenchHostIDKubernetesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) WithContext(ctx context.Context) *GetV1BenchHostIDKubernetesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) WithHTTPClient(client *http.Client) *GetV1BenchHostIDKubernetesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) WithID(id string) *GetV1BenchHostIDKubernetesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 bench host ID kubernetes params
func (o *GetV1BenchHostIDKubernetesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1BenchHostIDKubernetesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
