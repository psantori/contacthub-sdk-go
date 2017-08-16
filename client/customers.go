/**
 * This file is part of contacthub-sdk-go.
 *
 * contacthub-sdk-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * contacthub-sdk-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with contacthub-sdk-go. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2017 Arduino AG
 *
 * @author Luca Osti
 *
 */

package client

import (
	"fmt"
	"net/http"

	"github.com/guregu/null"
)

const (
	customerBasePath = "customers"
)

// CustomerService provides access to the Customers API
type CustomerService struct {
	client *Client
}

// Customer contains all editable fields for ContactHub Customer objects
type Customer struct {
	NodeID             string                  `json:"nodeId,required"`
	ExternalID         *null.String            `json:"externalId,omitempty"`
	Enabled            *null.Bool              `json:"enabled,omitempty"`
	ExtendedProperties *map[string]interface{} `json:"extended,omitempty"`
	Extra              *null.String            `json:"extra,omitempty"`
	BaseProperties     *BaseProperties         `json:"base,omitempty"`
	Tags               *Tags                   `json:"tags,omitempty"`
}

func (c *Customer) toPatchRequest() *customerPatchRequest {
	return &customerPatchRequest{
		ExternalID:         c.ExternalID,
		Enabled:            c.Enabled,
		ExtendedProperties: c.ExtendedProperties,
		Extra:              c.Extra,
		BaseProperties:     c.BaseProperties,
		Tags:               c.Tags,
	}
}

type customerPatchRequest struct {
	ExternalID         *null.String            `json:"externalId,omitempty"`
	Enabled            *null.Bool              `json:"enabled,omitempty"`
	ExtendedProperties *map[string]interface{} `json:"extended,omitempty"`
	Extra              *null.String            `json:"extra,omitempty"`
	BaseProperties     *BaseProperties         `json:"base,omitempty"`
	Tags               *Tags                   `json:"tags,omitempty"`
}

// CustomerResponse represents a Customer as returned by the ContactHub API
type CustomerResponse struct {
	ID                 string                  `json:"id,omitempty,required"`
	NodeID             string                  `json:"nodeId,required"`
	ExternalID         *null.String            `json:"externalId,required"`
	Enabled            bool                    `json:"enabled,required"`
	ExtendedProperties *map[string]interface{} `json:"extended,required"`
	Extra              *null.String            `json:"extra,required"`
	BaseProperties     *BasePropertiesResponse `json:"base,required"`
	Tags               *Tags                   `json:"tags,required"`
	RegisteredAt       CustomDate              `json:"registeredAt,required"`
	UpdatedAt          CustomDate              `json:"updatedAt,required"`
}

type customerListResponse struct {
	PageInfo  PageInfo           `json:"page"`
	Customers []CustomerResponse `json:"elements"`
}

// Get returns an individual customer by the ContactHub Customer ID
func (s *CustomerService) Get(ID string) (*CustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", customerBasePath, ID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	customer := new(CustomerResponse)
	_, err = s.client.Do(req, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// Delete deletes a customer by the ContactHub Customer ID
// Note: the API returns an empty body
func (s *CustomerService) Delete(ID string) error {
	path := fmt.Sprintf("%s/%s", customerBasePath, ID)
	req, err := s.client.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Create creates a new Customer on ContactHub
func (s *CustomerService) Create(customer *Customer) (*CustomerResponse, error) {
	if len(customer.NodeID) == 0 {
		customer.NodeID = s.client.Config.DefaultNodeID
	}
	req, err := s.client.NewRequest(http.MethodPost, customerBasePath, customer)
	if err != nil {
		return nil, err
	}

	createdCustomer := new(CustomerResponse)
	_, err = s.client.Do(req, createdCustomer)
	if err != nil {
		return nil, err
	}

	return createdCustomer, nil
}

// Update updates a Customer on ContactHub, via a patch operation
func (s *CustomerService) Update(ID string, customer *Customer) (*CustomerResponse, error) {
	path := fmt.Sprintf("%s/%s", customerBasePath, ID)
	customerRequest := customer.toPatchRequest()
	req, err := s.client.NewRequest(http.MethodPatch, path, customerRequest)
	if err != nil {
		return nil, err
	}

	createdCustomer := new(CustomerResponse)
	_, err = s.client.Do(req, createdCustomer)
	if err != nil {
		return nil, err
	}

	return createdCustomer, nil
}

// List requests all customers from the default Node
// The Node ID can be overriden via the QueryParams
func (s *CustomerService) List(params *ListParams) ([]CustomerResponse, PageInfo, error) {
	params.preparePagination()
	customers, pageInfo, err := s.list(params, customerBasePath)
	if err != nil {
		return nil, PageInfo{}, err
	}

	return customers, *pageInfo, err
}

func (s *CustomerService) list(params *ListParams, basePath string) ([]CustomerResponse, *PageInfo, error) {
	// build url
	if _, ok := params.QueryParams["nodeId"]; !ok {
		params.QueryParams["nodeId"] = s.client.Config.DefaultNodeID
	}
	path := addQuery(basePath, params.QueryParams)

	List := &customerListResponse{}
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = s.client.Do(req, List)

	if err != nil {
		return nil, nil, err
	}

	return List.Customers, &List.PageInfo, nil
}
