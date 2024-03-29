// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new namespace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Domain(params *GetV1DomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1DomainOK, error)

	PatchV1Domain(params *PatchV1DomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1DomainOK, error)

	PatchV1DomainName(params *PatchV1DomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1DomainNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1Domain gets namespace list
*/
func (a *Client) GetV1Domain(params *GetV1DomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1DomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Domain",
		Method:             "GET",
		PathPattern:        "/v1/domain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1DomainReader{formats: a.formats},
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
	success, ok := result.(*GetV1DomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Domain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1Domain configures namespace setting
*/
func (a *Client) PatchV1Domain(params *PatchV1DomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1DomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1DomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1Domain",
		Method:             "PATCH",
		PathPattern:        "/v1/domain",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1DomainReader{formats: a.formats},
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
	success, ok := result.(*PatchV1DomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1Domain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1DomainName updates namespace
*/
func (a *Client) PatchV1DomainName(params *PatchV1DomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1DomainNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1DomainNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1DomainName",
		Method:             "PATCH",
		PathPattern:        "/v1/domain/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1DomainNameReader{formats: a.formats},
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
	success, ok := result.(*PatchV1DomainNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1DomainName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
