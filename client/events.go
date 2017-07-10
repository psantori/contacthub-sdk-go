package client

import (
	"fmt"
	"net/http"

	"github.com/contactlab/contacthub-sdk-go/enums"
)

const (
	eventBasePath = "events"
)

type Event struct {
	EventID           string                 `json:"nodeId,required"`
	BringBackProperty BringBackProperty      `json:"bringBackProperties,required"`
	Type              string                 `json:"type,required"`
	Context           string                 `json:"context,required"`
	Properties        map[string]interface{} `json:"properties,required"`
	ContextInfo       map[string]interface{} `json:"contextInfo,omitempty"`
	Date              CustomDate             `json:"date,omitempty"`
}

type BringBackProperty struct {
	Type   enums.BringBackPropertyType `json:"type,required"`
	Value  string                      `json:"value,required"`
	NodeID string                      `json:"nodeId,required"`
}

// EventResponse represents a Event as returned by the ContactHub API
type EventResponse struct {
	Event
	ID string `json:"id,omitempty,required"`
}

// EventService provides access to the Events API
type EventService struct {
	client *Client
}

type eventListResponse struct {
	PageInfo PageInfo        `json:"page"`
	Events   []EventResponse `json:"elements"`
}

// Get returns an individual event by the ContactHub Event ID
func (s *EventService) Get(ID string) (*EventResponse, error) {
	path := fmt.Sprintf("%s/%s", eventBasePath, ID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	event := new(EventResponse)
	resp, err := s.client.Do(req, &event)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return event, nil
}

// Create creates a new Event on ContactHub
func (s *EventService) Create(event *Event) (*EventResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, eventBasePath, event)
	if err != nil {
		return nil, err
	}

	createdEvent := new(EventResponse)
	resp, err := s.client.Do(req, createdEvent)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return createdEvent, nil
}

// Update updates a Event on ContactHub
func (s *EventService) Update(ID string, event *Event) (*EventResponse, error) {
	path := fmt.Sprintf("%s/%s", eventBasePath, ID)
	req, err := s.client.NewRequest(http.MethodPost, path, event)
	if err != nil {
		return nil, err
	}

	createdEvent := new(EventResponse)
	resp, err := s.client.Do(req, createdEvent)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return createdEvent, nil
}

// List lists all events for a specified customer
func (s *EventService) List(customerID string, params *ListParams) ([]EventResponse, PageInfo, error) {
	params.preparePagination()
	params.QueryParams["customerId"] = customerID
	events, pageInfo, err := s.list(params, eventBasePath)
	if err != nil {
		return nil, PageInfo{}, err
	}

	return events, *pageInfo, err
}

func (s *EventService) list(params *ListParams, basePath string) ([]EventResponse, *PageInfo, error) {
	// build url
	/*if _, ok := params.QueryParams["nodeId"]; !ok {
		params.QueryParams["nodeId"] = s.client.Config.DefaultNodeID
	}*/
	path := addQuery(basePath, params.QueryParams)

	var List = &eventListResponse{}
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, &List)

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	return List.Events, &List.PageInfo, nil
}