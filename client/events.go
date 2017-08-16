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
	BringBackProperty *BringBackProperty      `json:"bringBackProperties,omitempty"`
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
	ID                string                  `json:"id,omitempty,required"`
	CustomerID        *null.String            `json:"customerId,required"`
	Type              enums.EventType         `json:"type,required"`
	Context           enums.EventContext      `json:"context,required"`
	Properties        map[string]interface{}  `json:"properties,required"`
	BringBackProperty *BringBackProperty      `json:"bringBackProperties,required"`
	ContextInfo       map[string]interface{}  `json:"contextInfo,required"`
	Date              CustomDate              `json:"date,required"`
	RegisteredAt      CustomDate              `json:"registeredAt,required"`
	UpdatedAt         CustomDate              `json:"updatedAt,required"`
	Tracking          *map[string]interface{} `json:"Tracking,omitempty"`
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
