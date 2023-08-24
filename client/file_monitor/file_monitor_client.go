// Code generated by go-swagger; DO NOT EDIT.

package file_monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new file monitor API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for file monitor API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1FileMonitor(params *GetV1FileMonitorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1FileMonitorOK, error)

	GetV1FileMonitorName(params *GetV1FileMonitorNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1FileMonitorNameOK, error)

	PatchV1FileMonitorName(params *PatchV1FileMonitorNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1FileMonitorNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1FileMonitor gets a list of file monitors
*/
func (a *Client) GetV1FileMonitor(params *GetV1FileMonitorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1FileMonitorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1FileMonitorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1FileMonitor",
		Method:             "GET",
		PathPattern:        "/v1/file_monitor",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1FileMonitorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1FileMonitorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1FileMonitor: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1FileMonitorName shows file monitor
*/
func (a *Client) GetV1FileMonitorName(params *GetV1FileMonitorNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1FileMonitorNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1FileMonitorNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1FileMonitorName",
		Method:             "GET",
		PathPattern:        "/v1/file_monitor/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1FileMonitorNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1FileMonitorNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1FileMonitorName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1FileMonitorName updates file monitor
*/
func (a *Client) PatchV1FileMonitorName(params *PatchV1FileMonitorNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1FileMonitorNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1FileMonitorNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1FileMonitorName",
		Method:             "PATCH",
		PathPattern:        "/v1/file_monitor/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1FileMonitorNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1FileMonitorNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1FileMonitorName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
