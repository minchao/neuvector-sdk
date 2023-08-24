// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewDeleteV1ServerNameParams creates a new DeleteV1ServerNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ServerNameParams() *DeleteV1ServerNameParams {
	return &DeleteV1ServerNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ServerNameParamsWithTimeout creates a new DeleteV1ServerNameParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ServerNameParamsWithTimeout(timeout time.Duration) *DeleteV1ServerNameParams {
	return &DeleteV1ServerNameParams{
		timeout: timeout,
	}
}

// NewDeleteV1ServerNameParamsWithContext creates a new DeleteV1ServerNameParams object
// with the ability to set a context for a request.
func NewDeleteV1ServerNameParamsWithContext(ctx context.Context) *DeleteV1ServerNameParams {
	return &DeleteV1ServerNameParams{
		Context: ctx,
	}
}

// NewDeleteV1ServerNameParamsWithHTTPClient creates a new DeleteV1ServerNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ServerNameParamsWithHTTPClient(client *http.Client) *DeleteV1ServerNameParams {
	return &DeleteV1ServerNameParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ServerNameParams contains all the parameters to send to the API endpoint

	for the delete v1 server name operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ServerNameParams struct {

	/* Name.

	   Name of the server
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 server name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ServerNameParams) WithDefaults() *DeleteV1ServerNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 server name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ServerNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 server name params
func (o *DeleteV1ServerNameParams) WithTimeout(timeout time.Duration) *DeleteV1ServerNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 server name params
func (o *DeleteV1ServerNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 server name params
func (o *DeleteV1ServerNameParams) WithContext(ctx context.Context) *DeleteV1ServerNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 server name params
func (o *DeleteV1ServerNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 server name params
func (o *DeleteV1ServerNameParams) WithHTTPClient(client *http.Client) *DeleteV1ServerNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 server name params
func (o *DeleteV1ServerNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete v1 server name params
func (o *DeleteV1ServerNameParams) WithName(name string) *DeleteV1ServerNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete v1 server name params
func (o *DeleteV1ServerNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ServerNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
