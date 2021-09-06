// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plugins API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugins API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PluginsIPReservationsRelatedIPV6Read(params *PluginsIPReservationsRelatedIPV6ReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsIPReservationsRelatedIPV6ReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PluginsIPReservationsRelatedIPV6Read plugins ip reservations related ipv6 read API
*/
func (a *Client) PluginsIPReservationsRelatedIPV6Read(params *PluginsIPReservationsRelatedIPV6ReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PluginsIPReservationsRelatedIPV6ReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPluginsIPReservationsRelatedIPV6ReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "plugins_ip-reservations_related-ipv6_read",
		Method:             "GET",
		PathPattern:        "/plugins/ip-reservations/related-ipv6/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PluginsIPReservationsRelatedIPV6ReadReader{formats: a.formats},
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
	success, ok := result.(*PluginsIPReservationsRelatedIPV6ReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for plugins_ip-reservations_related-ipv6_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
