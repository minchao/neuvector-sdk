// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1LogActivity(params *GetV1LogActivityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogActivityOK, error)

	GetV1LogAudit(params *GetV1LogAuditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogAuditOK, error)

	GetV1LogEvent(params *GetV1LogEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogEventOK, error)

	GetV1LogIncident(params *GetV1LogIncidentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogIncidentOK, error)

	GetV1LogSecurity(params *GetV1LogSecurityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogSecurityOK, error)

	GetV1LogThreat(params *GetV1LogThreatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogThreatOK, error)

	GetV1LogThreatID(params *GetV1LogThreatIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogThreatIDOK, error)

	GetV1LogViolation(params *GetV1LogViolationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogViolationOK, error)

	GetV1LogViolationWorkload(params *GetV1LogViolationWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogViolationWorkloadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1LogActivity gets activity list
*/
func (a *Client) GetV1LogActivity(params *GetV1LogActivityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogActivityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogActivityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogActivity",
		Method:             "GET",
		PathPattern:        "/v1/log/activity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogActivityReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogActivityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogActivity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogAudit gets a list of audits
*/
func (a *Client) GetV1LogAudit(params *GetV1LogAuditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogAuditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogAuditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogAudit",
		Method:             "GET",
		PathPattern:        "/v1/log/audit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogAuditReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogAuditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogAudit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogEvent gets a list of events
*/
func (a *Client) GetV1LogEvent(params *GetV1LogEventParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogEventParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogEvent",
		Method:             "GET",
		PathPattern:        "/v1/log/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogEventReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogIncident gets a list of incidents
*/
func (a *Client) GetV1LogIncident(params *GetV1LogIncidentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogIncidentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogIncidentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogIncident",
		Method:             "GET",
		PathPattern:        "/v1/log/incident",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogIncidentReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogIncidentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogIncident: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogSecurity gets security event list
*/
func (a *Client) GetV1LogSecurity(params *GetV1LogSecurityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogSecurityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogSecurityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogSecurity",
		Method:             "GET",
		PathPattern:        "/v1/log/security",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogSecurityReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogSecurityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogSecurity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogThreat gets a list of threats
*/
func (a *Client) GetV1LogThreat(params *GetV1LogThreatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogThreatOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogThreatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogThreat",
		Method:             "GET",
		PathPattern:        "/v1/log/threat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogThreatReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogThreatOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogThreat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogThreatID shows threat
*/
func (a *Client) GetV1LogThreatID(params *GetV1LogThreatIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogThreatIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogThreatIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogThreatID",
		Method:             "GET",
		PathPattern:        "/v1/log/threat/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogThreatIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogThreatIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogThreatID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogViolation gets a list of violations
*/
func (a *Client) GetV1LogViolation(params *GetV1LogViolationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogViolationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogViolationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogViolation",
		Method:             "GET",
		PathPattern:        "/v1/log/violation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogViolationReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogViolationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogViolation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1LogViolationWorkload gets violation workloads
*/
func (a *Client) GetV1LogViolationWorkload(params *GetV1LogViolationWorkloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1LogViolationWorkloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1LogViolationWorkloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1LogViolationWorkload",
		Method:             "GET",
		PathPattern:        "/v1/log/violation/workload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1LogViolationWorkloadReader{formats: a.formats},
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
	success, ok := result.(*GetV1LogViolationWorkloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1LogViolationWorkload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
