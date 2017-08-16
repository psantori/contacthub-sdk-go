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

	"github.com/contactlab/contacthub-sdk-go/enums"
	"github.com/contactlab/contacthub-sdk-go/nullable"
	"github.com/kylelemons/godebug/pretty"
)

func TestSubscriptionCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"subscription","kind":"DIGITAL_MESSAGE","subscribed":true,"startDate":"2022-02-22T20:22:22.215+0000"}`
	response := `{"id":"subscription","name":null,"type":null,"kind":"DIGITAL_MESSAGE","subscribed":true,"startDate":"2022-02-22T20:22:22.215+0000","endDate":null,"subscriberId":null,"registeredAt":null,"updatedAt":null,"preferences":[]}`
	mux.HandleFunc("/customers/my-customer-id/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Client.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	startDate, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	subscriptionKind := enums.DigitalMessage
	expectedSubscriptionResponse := SubscriptionResponse{
		ID:          "subscription",
		Subscribed:  nullable.BoolFrom(true),
		Kind:        &subscriptionKind,
		StartDate:   &CustomDate{startDate},
		Preferences: &[]map[string]interface{}{},
	}

	subscription := Subscription{
		ID:         "subscription",
		Subscribed: nullable.BoolFrom(true),
		Kind:       &subscriptionKind,
		StartDate:  &CustomDate{startDate},
	}

	subscriptionResponse, err := testClient.Subscriptions.Create("my-customer-id", &subscription)

	if err != nil {
		t.Errorf("Unexpected error. Subscriptions.Create: %v", err)
	}

	if diff := pretty.Compare(subscriptionResponse, expectedSubscriptionResponse); diff != "" {
		t.Errorf("Client.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}
