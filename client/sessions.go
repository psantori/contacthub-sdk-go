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
)

const (
	sessionBasePath = customerBasePath + "/%s/sessions"
)

// Session contains info about the Customer sessions
type Session struct {
	Value string `json:"value,required"`
}

type SessionResponse struct {
	ID    string `json:"id,required"`
	Value string `json:"value,required"`
}

// SessionService provides access to the Sessions API
type SessionService struct {
	client *Client
}

// Get returns an individual session of a customer
func (s *SessionService) Get(customerID, ID string) (*SessionResponse, error) {
	path := fmt.Sprintf(sessionBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	session := new(SessionResponse)
	_, err = s.client.Do(req, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// Create creates a new Session for the Customer, returns the response
func (s *SessionService) Create(customerID string, session *Session) (*SessionResponse, error) {
	path := fmt.Sprintf(sessionBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodPost, path, session)
	if err != nil {
		return nil, err
	}

	createdSession := new(SessionResponse)
	_, err = s.client.Do(req, createdSession)
	if err != nil {
		return nil, err
	}

	return createdSession, nil
}

// List gets a list of sessions assigned to the customer
func (s *SessionService) List(customerID string) ([]SessionResponse, error) {
	path := fmt.Sprintf(sessionBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	sessions := new([]SessionResponse)
	_, err = s.client.Do(req, sessions)
	if err != nil {
		return nil, err
	}

	return *sessions, nil
}

// Delete deletes a Session
func (s *SessionService) Delete(customerID, ID string) error {
	path := fmt.Sprintf(sessionBasePath, customerID) + "/" + ID
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
