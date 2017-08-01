package client

import (
	"fmt"
	"net/http"

	"github.com/contactlab/contacthub-sdk-go/enums"
	"github.com/guregu/null"
)

const (
	eventBasePath = "events"
)

// Event represents a Contacthub Event
type Event struct {
	CustomerID        *null.String            `json:"customerId,omitempty"`
	Type              enums.EventType         `json:"type,required"`
	Context           enums.EventContext      `json:"context,required"`
	Properties        map[string]interface{}  `json:"properties,required"`
	BringBackProperty *BringBackProperty      `json:"bringBackProperties"`
	ContextInfo       *map[string]interface{} `json:"contextInfo,omitempty"`
	Date              *CustomDate             `json:"date,omitempty"`
}

// BringBackProperty represents a ContactHub event BringBackProperty, used to match the event with existing users
type BringBackProperty struct {
	Type   enums.BringBackPropertyType `json:"type,required"`
	Value  string                      `json:"value,required"`
	NodeID string                      `json:"nodeId,required"`
}

// EventResponse represents a Event as returned by the ContactHub API
type EventResponse struct {
	*Event
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
	_, err = s.client.Do(req, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

// Create creates a new Event on ContactHub
func (s *EventService) Create(event *Event) (*EventResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, eventBasePath, event)
	if err != nil {
		return nil, err
	}

	createdEvent := new(EventResponse)
	_, err = s.client.Do(req, createdEvent)
	if err != nil {
		return nil, err
	}

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
	path := addQuery(basePath, params.QueryParams)

	List := &eventListResponse{}
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	_, err = s.client.Do(req, List)

	if err != nil {
		return nil, nil, err
	}

	return List.Events, &List.PageInfo, nil
}

// Delete remove an Event from ContactHub
func (s *EventService) Delete(ID string) error {
	path := fmt.Sprintf("%s/%s", eventBasePath, ID)
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
