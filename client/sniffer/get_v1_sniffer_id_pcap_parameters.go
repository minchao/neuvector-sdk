// Code generated by go-swagger; DO NOT EDIT.

package sniffer

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

// NewGetV1SnifferIDPcapParams creates a new GetV1SnifferIDPcapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SnifferIDPcapParams() *GetV1SnifferIDPcapParams {
	return &GetV1SnifferIDPcapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SnifferIDPcapParamsWithTimeout creates a new GetV1SnifferIDPcapParams object
// with the ability to set a timeout on a request.
func NewGetV1SnifferIDPcapParamsWithTimeout(timeout time.Duration) *GetV1SnifferIDPcapParams {
	return &GetV1SnifferIDPcapParams{
		timeout: timeout,
	}
}

// NewGetV1SnifferIDPcapParamsWithContext creates a new GetV1SnifferIDPcapParams object
// with the ability to set a context for a request.
func NewGetV1SnifferIDPcapParamsWithContext(ctx context.Context) *GetV1SnifferIDPcapParams {
	return &GetV1SnifferIDPcapParams{
		Context: ctx,
	}
}

// NewGetV1SnifferIDPcapParamsWithHTTPClient creates a new GetV1SnifferIDPcapParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SnifferIDPcapParamsWithHTTPClient(client *http.Client) *GetV1SnifferIDPcapParams {
	return &GetV1SnifferIDPcapParams{
		HTTPClient: client,
	}
}

/*
GetV1SnifferIDPcapParams contains all the parameters to send to the API endpoint

	for the get v1 sniffer ID pcap operation.

	Typically these are written to a http.Request.
*/
type GetV1SnifferIDPcapParams struct {

	/* ID.

	   Sniffer ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 sniffer ID pcap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SnifferIDPcapParams) WithDefaults() *GetV1SnifferIDPcapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 sniffer ID pcap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SnifferIDPcapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) WithTimeout(timeout time.Duration) *GetV1SnifferIDPcapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) WithContext(ctx context.Context) *GetV1SnifferIDPcapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) WithHTTPClient(client *http.Client) *GetV1SnifferIDPcapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) WithID(id string) *GetV1SnifferIDPcapParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 sniffer ID pcap params
func (o *GetV1SnifferIDPcapParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SnifferIDPcapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
