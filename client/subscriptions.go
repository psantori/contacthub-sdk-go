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
	subscriptionBasePath = customerBasePath + "/%s/subscriptions"
)

// Subscription contains info about the Customer subscriptions
type Subscription struct {
	ID           string                    `json:"id,required"`
	Name         *null.String              `json:"name,omitempty"`
	Type         *null.String              `json:"type,omitempty"`
	Kind         *enums.SubscriptionKind   `json:"kind,omitempty"`
	Subscribed   *null.Bool                `json:"subscribed,omitempty"`
	StartDate    *CustomDate               `json:"startDate,omitempty"`
	EndDate      *CustomDate               `json:"endDate,omitempty"`
	SubscriberID *null.String              `json:"subscriberId,omitempty"`
	RegisteredAt *CustomDate               `json:"registeredAt,omitempty"`
	UpdatedAt    *CustomDate               `json:"updatedAt,omitempty"`
	Preferences  *[]map[string]interface{} `json:"preferences,omitempty"`
}

// SubscriptionResponse is actually identical to the Subscription payload
type SubscriptionResponse Subscription

// SubscriptionService provides access to the Subscriptions API
type SubscriptionService struct {
	client *Client
}

// Get returns an individual subscription of a customer
func (s *SubscriptionService) Get(customerID, ID string) (*SubscriptionResponse, error) {
	path := fmt.Sprintf(subscriptionBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	subscription := new(SubscriptionResponse)
	_, err = s.client.Do(req, subscription)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

// Create creates a new Subscription for the Customer, returns the response
func (s *SubscriptionService) Create(customerID string, subscription *Subscription) (*SubscriptionResponse, error) {
	path := fmt.Sprintf(subscriptionBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodPost, path, subscription)
	if err != nil {
		return nil, err
	}

	createdSubscription := new(SubscriptionResponse)
	_, err = s.client.Do(req, createdSubscription)
	if err != nil {
		return nil, err
	}

	return createdSubscription, nil
}

// Update updates a Subscription via a patch operation
func (s *SubscriptionService) Update(customerID, ID string, subscription *Subscription) (*SubscriptionResponse, error) {
	path := fmt.Sprintf(subscriptionBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodPatch, path, subscription)
	if err != nil {
		return nil, err
	}

	createdSubscription := new(SubscriptionResponse)
	_, err = s.client.Do(req, createdSubscription)
	if err != nil {
		return nil, err
	}

	return createdSubscription, nil
}

// Delete deletes a Subscription
func (s *SubscriptionService) Delete(customerID, ID string) error {
	path := fmt.Sprintf(subscriptionBasePath, customerID) + "/" + ID
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
