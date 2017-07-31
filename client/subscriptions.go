package client

import (
	"fmt"
	"net/http"
)

const (
	subscriptionBasePath = customerBasePath + "/%s/subscriptions"
)

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
