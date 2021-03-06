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
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

var (
	mockServer *httptest.Server

	mux *http.ServeMux

	testClient *Client
)

func setup() {
	mux = http.NewServeMux()
	mockServer = httptest.NewServer(mux)

	testClient, _ = New(&Config{
		DefaultNodeID: "fakenodeid",
		WorkspaceID:   "fakeworkspaceid",
		APIkey:        "fakeapikey",
		Timeout:       50,
		Debug:         false,
	})
	url, _ := url.Parse(mockServer.URL)
	testClient.BaseURL = url
}
func teardown() {
	mockServer.Close()
}

func testMethod(t *testing.T, r *http.Request, expectedMethod string) {
	if r.Method != expectedMethod {
		t.Errorf("Expected %v, got %v request method", expectedMethod, r.Method)
	}
}

func testQueryString(t *testing.T, r *http.Request, expectedQuery string) {
	queryParams := r.URL.Query()

	if _, ok := queryParams[expectedQuery]; !ok {
		t.Errorf("Expected '%v' querystring param", expectedQuery)
	}
}

func testQueryStringValue(t *testing.T, r *http.Request, expectedQuery string, expectedValue string) {
	testQueryString(t, r, expectedQuery)
	queryParams := r.URL.Query()

	if queryParams.Get(expectedQuery) != expectedValue {
		t.Errorf("Expected '%v' for querystring param '%v'. Got '%v'", expectedValue, expectedQuery, queryParams.Get(expectedQuery))
	}
}

func testQueryStringPositiveInt(t *testing.T, r *http.Request, expectedQuery string) {
	testQueryString(t, r, expectedQuery)
	queryValue := r.URL.Query().Get(expectedQuery)
	if value, err := strconv.ParseInt(queryValue, 10, 64); err != nil {
		t.Errorf("Expected positive int for querystring param '%v'. Error: %v", expectedQuery, err)
	} else if value < 0 {
		t.Errorf("Expected positive int for querystring param '%v'. Got %v", expectedQuery, value)
	}
}

func TestNew(t *testing.T) {
	_, err := New(&Config{APIkey: "key", WorkspaceID: "ID"})

	if err != nil {
		t.Fatalf("Client New(): %v", err)
	}
}
func TestHttpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Nope", 400)
	})

	req, _ := testClient.NewRequest(http.MethodGet, "/", nil)
	_, err := testClient.Do(req, nil)

	if err == nil {
		t.Error("Expected error.")
	}
}

func TestEmptyBody(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(202)
		fmt.Fprint(w, "")
	})

	req, err := testClient.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Error("TestEmptyBody: Unexpected error.")
	}

	object := make(map[string]interface{})
	_, err = testClient.Do(req, &object)
	if err != nil {
		t.Error("TestEmptyBody: Unexpected error.")
	}
}

func TestClientTimeout(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, "{}")
	})

	req, _ := testClient.NewRequest(http.MethodGet, "/", nil)
	_, err := testClient.Do(req, nil)

	if err == nil {
		t.Error("Expected timeout.")
	}
}
