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

	"github.com/minchao/neuvector-sdk/models"
)

// NewPostV1SystemRequestParams creates a new PostV1SystemRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1SystemRequestParams() *PostV1SystemRequestParams {
	return &PostV1SystemRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1SystemRequestParamsWithTimeout creates a new PostV1SystemRequestParams object
// with the ability to set a timeout on a request.
func NewPostV1SystemRequestParamsWithTimeout(timeout time.Duration) *PostV1SystemRequestParams {
	return &PostV1SystemRequestParams{
		timeout: timeout,
	}
}

// NewPostV1SystemRequestParamsWithContext creates a new PostV1SystemRequestParams object
// with the ability to set a context for a request.
func NewPostV1SystemRequestParamsWithContext(ctx context.Context) *PostV1SystemRequestParams {
	return &PostV1SystemRequestParams{
		Context: ctx,
	}
}

// NewPostV1SystemRequestParamsWithHTTPClient creates a new PostV1SystemRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1SystemRequestParamsWithHTTPClient(client *http.Client) *PostV1SystemRequestParams {
	return &PostV1SystemRequestParams{
		HTTPClient: client,
	}
}

/*
PostV1SystemRequestParams contains all the parameters to send to the API endpoint

	for the post v1 system request operation.

	Typically these are written to a http.Request.
*/
type PostV1SystemRequestParams struct {

	/* Body.

	   System request data
	*/
	Body *models.RESTSystemRequestData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 system request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SystemRequestParams) WithDefaults() *PostV1SystemRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 system request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SystemRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 system request params
func (o *PostV1SystemRequestParams) WithTimeout(timeout time.Duration) *PostV1SystemRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 system request params
func (o *PostV1SystemRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 system request params
func (o *PostV1SystemRequestParams) WithContext(ctx context.Context) *PostV1SystemRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 system request params
func (o *PostV1SystemRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 system request params
func (o *PostV1SystemRequestParams) WithHTTPClient(client *http.Client) *PostV1SystemRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 system request params
func (o *PostV1SystemRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post v1 system request params
func (o *PostV1SystemRequestParams) WithBody(body *models.RESTSystemRequestData) *PostV1SystemRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 system request params
func (o *PostV1SystemRequestParams) SetBody(body *models.RESTSystemRequestData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1SystemRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
