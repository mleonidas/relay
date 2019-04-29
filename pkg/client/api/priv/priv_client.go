// Code generated by go-swagger; DO NOT EDIT.

package priv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new priv API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for priv API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
HealthCheck checks service health
*/
func (a *Client) HealthCheck(params *HealthCheckParams) (*HealthCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "healthCheck",
		Method:             "GET",
		PathPattern:        "/_priv/healthz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*HealthCheckOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}