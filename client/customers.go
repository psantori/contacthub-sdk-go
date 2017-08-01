package client

import (
	"fmt"
	"net/http"

	"github.com/guregu/null"
)

const (
	customerBasePath = "customers"
)

// Customer contains all editable fields for ContactHub Customer objects
type Customer struct {
	NodeID             string                  `json:"nodeId,required"`
	ExternalID         *null.String            `json:"externalId,omitempty"`
	Enabled            bool                    `json:"enabled,required"`
	ExtendedProperties *map[string]interface{} `json:"extended,omitempty"`
	Extra              *null.String            `json:"extra,omitempty"`
	BaseProperties     *BaseProperties         `json:"base,omitempty"`
	Tags               *Tags                   `json:"tags,omitempty"`
}

type customerPutRequest struct {
	*Customer
	ID string `json:"id,omitempty"`
}

// CustomerResponse represents a Customer as returned by the ContactHub API
type CustomerResponse struct {
	*Customer
	ID           string     `json:"id,omitempty,required"`
	RegisteredAt CustomDate `json:"registeredAt,omitempty"`
	UpdatedAt    CustomDate `json:"updatedAt,omitempty"`
}

// CustomerService provides access to the Customers API
type CustomerService struct {
	client *Client
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
	customerRequest := &customerPutRequest{customer, ID}
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
