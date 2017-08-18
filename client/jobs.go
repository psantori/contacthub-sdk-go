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
	jobBasePath = customerBasePath + "/%s/jobs"
)

// Job contains info about the Customer job
type Job struct {
	ID              string       `json:"id,required"`
	CompanyIndustry *null.String `json:"companyIndustry,omitempty"`
	CompanyName     *null.String `json:"companyName,omitempty"`
	JobTitle        *null.String `json:"jobTitle,omitempty"`
	StartDate       *SimpleDate  `json:"startDate,omitempty"`
	EndDate         *SimpleDate  `json:"endDate,omitempty"`
	IsCurrent       *null.Bool   `json:"isCurrent,omitempty"`
}

type JobResponse struct {
	ID              string      `json:"id,required"`
	CompanyIndustry null.String `json:"companyIndustry,required"`
	CompanyName     null.String `json:"companyName,required"`
	JobTitle        null.String `json:"jobTitle,required"`
	StartDate       *SimpleDate `json:"startDate,required"`
	EndDate         *SimpleDate `json:"endDate,required"`
	IsCurrent       null.Bool   `json:"isCurrent,required"`
}

// JobService provides access to the Jobs API
type JobService struct {
	client *Client
}

// Get returns an individual job of a customer
func (s *JobService) Get(customerID, ID string) (*JobResponse, error) {
	path := fmt.Sprintf(jobBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	job := new(JobResponse)
	_, err = s.client.Do(req, job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Create creates a new Job for the Customer, returns the response
func (s *JobService) Create(customerID string, job *Job) (*JobResponse, error) {
	path := fmt.Sprintf(jobBasePath, customerID)
	req, err := s.client.NewRequest(http.MethodPost, path, job)
	if err != nil {
		return nil, err
	}

	createdJob := new(JobResponse)
	_, err = s.client.Do(req, createdJob)
	if err != nil {
		return nil, err
	}

	return createdJob, nil
}

// Update updates a Job via a put operation
func (s *JobService) Update(customerID, ID string, job *Job) (*JobResponse, error) {
	path := fmt.Sprintf(jobBasePath, customerID) + "/" + ID
	req, err := s.client.NewRequest(http.MethodPut, path, job)
	if err != nil {
		return nil, err
	}

	createdJob := new(JobResponse)
	_, err = s.client.Do(req, createdJob)
	if err != nil {
		return nil, err
	}

	return createdJob, nil
}

// Delete deletes a Job
func (s *JobService) Delete(customerID, ID string) error {
	path := fmt.Sprintf(jobBasePath, customerID) + "/" + ID
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
