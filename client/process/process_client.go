// Code generated by go-swagger; DO NOT EDIT.

package process

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new process API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for process API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1ProcessProfile(params *GetV1ProcessProfileParams, opts ...ClientOption) (*GetV1ProcessProfileOK, error)

	GetV1ProcessProfileName(params *GetV1ProcessProfileNameParams, opts ...ClientOption) (*GetV1ProcessProfileNameOK, error)

	GetV1ProcessRulesUUID(params *GetV1ProcessRulesUUIDParams, opts ...ClientOption) (*GetV1ProcessRulesUUIDOK, error)

	PatchV1ProcessProfileName(params *PatchV1ProcessProfileNameParams, opts ...ClientOption) (*PatchV1ProcessProfileNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1ProcessProfile gets a list of process profiles
*/
func (a *Client) GetV1ProcessProfile(params *GetV1ProcessProfileParams, opts ...ClientOption) (*GetV1ProcessProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ProcessProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ProcessProfile",
		Method:             "GET",
		PathPattern:        "/v1/process_profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ProcessProfileReader{formats: a.formats},
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
	success, ok := result.(*GetV1ProcessProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ProcessProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ProcessProfileName gets a process profile
*/
func (a *Client) GetV1ProcessProfileName(params *GetV1ProcessProfileNameParams, opts ...ClientOption) (*GetV1ProcessProfileNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ProcessProfileNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ProcessProfileName",
		Method:             "GET",
		PathPattern:        "/v1/process_profile/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ProcessProfileNameReader{formats: a.formats},
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
	success, ok := result.(*GetV1ProcessProfileNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ProcessProfileName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ProcessRulesUUID gets a process rule
*/
func (a *Client) GetV1ProcessRulesUUID(params *GetV1ProcessRulesUUIDParams, opts ...ClientOption) (*GetV1ProcessRulesUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ProcessRulesUUIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1ProcessRulesUUID",
		Method:             "GET",
		PathPattern:        "/v1/process_rules/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ProcessRulesUUIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1ProcessRulesUUIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1ProcessRulesUUID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ProcessProfileName updates a process profile
*/
func (a *Client) PatchV1ProcessProfileName(params *PatchV1ProcessProfileNameParams, opts ...ClientOption) (*PatchV1ProcessProfileNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ProcessProfileNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchV1ProcessProfileName",
		Method:             "PATCH",
		PathPattern:        "/v1/process_profile/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ProcessProfileNameReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ProcessProfileNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchV1ProcessProfileName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
