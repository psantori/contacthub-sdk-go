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
	likeBasePath = customerBasePath + "/%s/likes"
)

// Like represents a thing the Customer liked
type Like struct {
	ID          string       `json:"id,required"`
	Category    *null.String `json:"category,required"`
	Name        *null.String `json:"name,required"`
	CreatedTime *CustomDate  `json:"createdTime,omitempty"`
}

type LikeResponse struct {
	ID          string      `json:"id,required"`
	Category    null.String `json:"category,required"`
	Name        null.String `json:"name,required"`
	CreatedTime *CustomDate `json:"createdTime,required"`
}

// LikeService provides access to the Likes API
type LikeService struct {
	client *Client
}

// Get returns an individual like of a customer
func (s *LikeService) Get(customerID, ID string) (*LikeResponse, error) {
	path := fmt.Sprintf(likeBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	like := new(LikeResponse)
	_, err = s.client.Do(req, like)
	if err != nil {
		return nil, err
	}

	return like, nil
}

// Create creates a new Like for the Customer, returns the response
func (s *LikeService) Create(customerID string, like *Like) (*LikeResponse, error) {
	path := fmt.Sprintf(likeBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodPost, path, like)
	if err != nil {
		return nil, err
	}

	createdLike := new(LikeResponse)
	_, err = s.client.Do(req, createdLike)
	if err != nil {
		return nil, err
	}

	return createdLike, nil
}

// Update updates a Like via a put operation
func (s *LikeService) Update(customerID, ID string, like *Like) (*LikeResponse, error) {
	path := fmt.Sprintf(likeBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodPut, path, like)
	if err != nil {
		return nil, err
	}

	createdLike := new(LikeResponse)
	_, err = s.client.Do(req, createdLike)
	if err != nil {
		return nil, err
	}

	return createdLike, nil
}

// Delete deletes a Like
func (s *LikeService) Delete(customerID, ID string) error {
	path := fmt.Sprintf(likeBasePath, customerID) + "/" + ID
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
