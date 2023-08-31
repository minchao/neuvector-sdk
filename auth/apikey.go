package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ApiKeyWriter implements runtime.ClientAuthInfoWriter interface.
type ApiKeyWriter struct {
	ApiKey string
}

func NewApiKeyWriter(apiKey string) *ApiKeyWriter {
	return &ApiKeyWriter{ApiKey: apiKey}
}

func (a ApiKeyWriter) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	return req.SetHeaderParam("X-Auth-Apikey", a.ApiKey)
}
