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
	educationBasePath = customerBasePath + "/%s/educations"
)

// Education contains the Education info of a Customer
type Education struct {
	ID                  string            `json:"id,required"`
	SchoolType          *enums.SchoolType `json:"schoolType,omitempty"`
	SchoolName          *null.String      `json:"schoolName,omitempty"`
	SchoolConcentration *null.String      `json:"schoolConcentration,omitempty"`
	StartYear           *null.Int         `json:"startYear,omitempty"`
	EndYear             *null.Int         `json:"endYear,omitempty"`
	IsCurrent           *null.Bool        `json:"isCurrent,omitempty"`
}

type EducationResponse struct {
	ID                  string            `json:"id,required"`
	SchoolType          *enums.SchoolType `json:"schoolType,required"`
	SchoolName          null.String       `json:"schoolName,required"`
	SchoolConcentration null.String       `json:"schoolConcentration,required"`
	StartYear           null.Int          `json:"startYear,required"`
	EndYear             null.Int          `json:"endYear,required"`
	IsCurrent           null.Bool         `json:"isCurrent,required"`
}

// EducationService provides access to the Educations API
type EducationService struct {
	client *Client
}

// Get returns an individual education of a customer
func (s *EducationService) Get(customerID, ID string) (*EducationResponse, error) {
	path := fmt.Sprintf(educationBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	education := new(EducationResponse)
	_, err = s.client.Do(req, education)
	if err != nil {
		return nil, err
	}

	return education, nil
}

// Create creates a new Education for the Customer, returns the response
func (s *EducationService) Create(customerID string, education *Education) (*EducationResponse, error) {
	path := fmt.Sprintf(educationBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodPost, path, education)
	if err != nil {
		return nil, err
	}

	createdEducation := new(EducationResponse)
	_, err = s.client.Do(req, createdEducation)
	if err != nil {
		return nil, err
	}

	return createdEducation, nil
}

// Update updates a Education via a put operation
func (s *EducationService) Update(customerID, ID string, education *Education) (*EducationResponse, error) {
	path := fmt.Sprintf(educationBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodPut, path, education)
	if err != nil {
		return nil, err
	}

	createdEducation := new(EducationResponse)
	_, err = s.client.Do(req, createdEducation)
	if err != nil {
		return nil, err
	}

	return createdEducation, nil
}

// Delete deletes a Education
func (s *EducationService) Delete(customerID, ID string) error {
	path := fmt.Sprintf(educationBasePath, customerID) + "/" + ID
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
