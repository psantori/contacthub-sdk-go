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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/guregu/null/zero"

	"strings"
)

// ErrorResponse represents an error response from the ContactHub API, which may contain multiple errors
type ErrorResponse struct {
	*http.Response
	Message string      `json:"message"`
	Logref  string      `json:"logref"`
	Data    zero.String `json:"data"`
	Errors  []APIError  `json:"errors"`
}

// APIError contains info about a ContactHub error
type APIError struct {
	Message string      `json:"message,omitempty"`
	Path    string      `json:"path,omitempty"`
	Data    interface{} `json:"data"`
	Code    zero.String `json:"code"`
}

func (r *ErrorResponse) Error() string {
	messages := make([]string, 1+len(r.Errors))
	messages[0] = r.Message

	for i, apiError := range r.Errors {
		messages[i+1] = fmt.Sprintf("%v (%v)", apiError.Message, apiError.Path)
	}
	return fmt.Sprintf("(%d) %v %v: %v",
		r.Response.StatusCode, r.Response.Request.Method, r.Response.Request.URL, strings.Join(messages, ", "))
}

func handleErrors(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}
