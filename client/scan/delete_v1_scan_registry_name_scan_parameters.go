// Code generated by go-swagger; DO NOT EDIT.

package scan

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

// NewDeleteV1ScanRegistryNameScanParams creates a new DeleteV1ScanRegistryNameScanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ScanRegistryNameScanParams() *DeleteV1ScanRegistryNameScanParams {
	return &DeleteV1ScanRegistryNameScanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ScanRegistryNameScanParamsWithTimeout creates a new DeleteV1ScanRegistryNameScanParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ScanRegistryNameScanParamsWithTimeout(timeout time.Duration) *DeleteV1ScanRegistryNameScanParams {
	return &DeleteV1ScanRegistryNameScanParams{
		timeout: timeout,
	}
}

// NewDeleteV1ScanRegistryNameScanParamsWithContext creates a new DeleteV1ScanRegistryNameScanParams object
// with the ability to set a context for a request.
func NewDeleteV1ScanRegistryNameScanParamsWithContext(ctx context.Context) *DeleteV1ScanRegistryNameScanParams {
	return &DeleteV1ScanRegistryNameScanParams{
		Context: ctx,
	}
}

// NewDeleteV1ScanRegistryNameScanParamsWithHTTPClient creates a new DeleteV1ScanRegistryNameScanParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ScanRegistryNameScanParamsWithHTTPClient(client *http.Client) *DeleteV1ScanRegistryNameScanParams {
	return &DeleteV1ScanRegistryNameScanParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ScanRegistryNameScanParams contains all the parameters to send to the API endpoint

	for the delete v1 scan registry name scan operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ScanRegistryNameScanParams struct {

	// XAuthToken.
	XAuthToken string

	/* Name.

	   Name of the registry
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 scan registry name scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ScanRegistryNameScanParams) WithDefaults() *DeleteV1ScanRegistryNameScanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 scan registry name scan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ScanRegistryNameScanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) WithTimeout(timeout time.Duration) *DeleteV1ScanRegistryNameScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) WithContext(ctx context.Context) *DeleteV1ScanRegistryNameScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) WithHTTPClient(client *http.Client) *DeleteV1ScanRegistryNameScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) WithXAuthToken(xAuthToken string) *DeleteV1ScanRegistryNameScanParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) SetXAuthToken(xAuthToken string) {
	o.XAuthToken = xAuthToken
}

// WithName adds the name to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) WithName(name string) *DeleteV1ScanRegistryNameScanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete v1 scan registry name scan params
func (o *DeleteV1ScanRegistryNameScanParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ScanRegistryNameScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}