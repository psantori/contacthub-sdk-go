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
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/contactlab/contacthub-sdk-go/nullable"

	"github.com/contactlab/contacthub-sdk-go/enums"

	"github.com/kylelemons/godebug/pretty"
)

func TestEventList(t *testing.T) {
	setup()
	defer teardown()

	response := `{"page":{"size":2,"totalElements":10,"totalUnfilteredElements":0,"totalPages":5,"number":0},"elements":[{"id":"my-event-id1","customerId":"my-customer-id","bringBackProperties":null,"type":"viewedProduct","context":"ECOMMERCE","date":"2017-08-01T20:23:09.215+0000","properties":{"id":"999","sku":"test-something","name":"kit bundle (Test)","price":0,"imageUrl":"https://something.jpg","linkUrl":"https://store.example","shortDescription":"Test bundle","category":[]},"contextInfo":{"client":{"ip":"111.111.111.111","userAgent":"Somebrowser/1.1"}},"registeredAt":"0001-01-01T00:00:00.00+0000","updatedAt":"0001-01-01T00:00:00.00+0000","tracking":null},{"id":"my-event-id2","customerId":"my-customer-id","bringBackProperties":null,"type":"reviewedProduct","context":"ECOMMERCE","date":"2017-08-01T20:23:09.215+0000","properties":{"id":"10000","sku":"test-something2","name":"kit bundle2 (Test)","price":0,"imageUrl":"https://something2.jpg","linkUrl":"https://store.example","shortDescription":"Test bundle 2","category":[]},"contextInfo":{"client":{"ip":"111.111.111.111","userAgent":"Somebrowser/1.1"}},"registeredAt":"0001-01-01T00:00:00.00+0000","updatedAt":"0001-01-01T00:00:00.00+0000","tracking":null}]}`
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testQueryStringValue(t, r, "customerId", "my-customer-id")
		fmt.Fprint(w, response)
	})

	params := ListParams{}
	events, pageInfo, err := testClient.Events.List("my-customer-id", &params)

	if err != nil {
		t.Errorf("Unexpected error. Events.List: %v", err)
	}

	if len(events) != 2 {
		t.Errorf("Wrong list size. Expected 2, got %v", err)
	}

	expectedPageInfo := PageInfo{
		Size:                    2,
		TotalElements:           10,
		TotalUnfilteredElements: 0,
		TotalPages:              5,
		Page:                    0,
	}

	if diff := pretty.Compare(pageInfo, expectedPageInfo); diff != "" {
		t.Errorf("Events.List: invalid value for struct: (-got +expected)\n%s", diff)
	}
	date1, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-08-01T20:23:09.215+0000")
	date2, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-08-01T20:23:09.215+0000")
	dateZero, _ := time.Parse("2006-01-02T15:04:05.999-0700", "0001-01-01T00:00:00.00+0000")
	expectedList := []EventResponse{
		EventResponse{
			ID:         "my-event-id1",
			CustomerID: nullable.StringFrom("my-customer-id"),
			Type:       enums.ViewedProduct,
			Context:    enums.Ecommerce,
			Properties: map[string]interface{}{
				"id":               "999",
				"sku":              "test-something",
				"name":             "kit bundle (Test)",
				"price":            0,
				"imageUrl":         "https://something.jpg",
				"linkUrl":          "https://store.example",
				"shortDescription": "Test bundle",
				"category":         []string{},
			},
			ContextInfo: map[string]interface{}{
				"client": map[string]interface{}{
					"ip":        "111.111.111.111",
					"userAgent": "Somebrowser/1.1",
				},
			},
			Date:         CustomDate{date1},
			RegisteredAt: CustomDate{dateZero},
			UpdatedAt:    CustomDate{dateZero},
		},
		EventResponse{
			ID:         "my-event-id2",
			CustomerID: nullable.StringFrom("my-customer-id"),
			Type:       enums.ReviewedProduct,
			Context:    enums.Ecommerce,
			Properties: map[string]interface{}{
				"id":               "10000",
				"sku":              "test-something2",
				"name":             "kit bundle2 (Test)",
				"price":            0,
				"imageUrl":         "https://something2.jpg",
				"linkUrl":          "https://store.example",
				"shortDescription": "Test bundle 2",
				"category":         []string{},
			},
			ContextInfo: map[string]interface{}{
				"client": map[string]interface{}{
					"ip":        "111.111.111.111",
					"userAgent": "Somebrowser/1.1",
				},
			},
			Date:         CustomDate{date2},
			RegisteredAt: CustomDate{dateZero},
			UpdatedAt:    CustomDate{dateZero},
		},
	}

	if diff := pretty.Compare(events, expectedList); diff != "" {
		t.Errorf("Events.List: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

// This method actually returns an empty body on the actual ContactHub API, but the documentation states otherwise.
func TestEventCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"customerId":"aaa","type":"abandonedCart","context":"ECOMMERCE","properties":{},"date":"2022-02-22T20:22:22.215+0000"}`
	response := `{"id": "my-new-event-id", "customerId":"aaa","type":"abandonedCart","context":"ECOMMERCE","properties":{},"bringBackProperties":null, "date":"2022-02-22T20:22:22.215+0000"}`
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Client.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	createdAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	expectedEventResponse := EventResponse{
		ID:         "my-new-event-id",
		CustomerID: nullable.StringFrom("aaa"),
		Type:       enums.AbandonedCart,
		Properties: map[string]interface{}{},
		Context:    enums.Ecommerce,
		Date:       CustomDate{createdAt},
	}

	event := Event{
		CustomerID: nullable.StringFrom("aaa"),
		Type:       enums.AbandonedCart,
		Properties: map[string]interface{}{},
		Context:    enums.Ecommerce,
		Date:       &CustomDate{createdAt},
	}
	eventResponse, err := testClient.Events.Create(&event)

	if err != nil {
		t.Errorf("Unexpected error. Events.Create: %v", err)
	}

	if diff := pretty.Compare(eventResponse, expectedEventResponse); diff != "" {
		t.Errorf("Client.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestEventDelete(t *testing.T) {
	setup()
	defer teardown()
	response := ``
	mux.HandleFunc("/events/my-event-id", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Events.Delete("my-event-id")
	if err != nil {
		t.Errorf("Unexpected error. Events.Delete: %v", err)
	}
}

func TestEventGet(t *testing.T) {
	setup()
	defer teardown()
	response := `{"id": "my-event-id", "customerId":"aaa","type":"abandonedCart","context":"ECOMMERCE","properties":{},"bringBackProperties":null, "date":"2022-02-22T20:22:22.215+0000"}`
	mux.HandleFunc("/events/my-event-id", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	event, err := testClient.Events.Get("my-event-id")
	if err != nil {
		t.Errorf("Unexpected error. Events.Get: %v", err)
	}

	date, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	expected := EventResponse{
		ID:         "my-event-id",
		CustomerID: nullable.StringFrom("aaa"),
		Type:       enums.AbandonedCart,
		Properties: map[string]interface{}{},
		Context:    enums.Ecommerce,
		Date:       CustomDate{date},
	}
	if diff := pretty.Compare(event, expected); diff != "" {
		t.Errorf("Events.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}
