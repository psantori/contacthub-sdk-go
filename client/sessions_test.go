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

	"github.com/kylelemons/godebug/pretty"
)

func TestSessionCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"value":"session"}`
	response := `{"id":"5ae33d78-49e5-413a-9e88-317ece2525c9","value":"session"}`
	mux.HandleFunc("/customers/my-customer-id/sessions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Sessions.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	expectedSessionResponse := SessionResponse{
		ID:    "5ae33d78-49e5-413a-9e88-317ece2525c9",
		Value: "session",
	}

	session := Session{"session"}

	sessionResponse, err := testClient.Sessions.Create("my-customer-id", &session)

	if err != nil {
		t.Errorf("Unexpected error. Sessions.Create: %v", err)
	}

	if diff := pretty.Compare(sessionResponse, expectedSessionResponse); diff != "" {
		t.Errorf("Sessions.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestSessionGet(t *testing.T) {
	setup()
	defer teardown()

	response := `{"id":"5ae33d78-49e5-413a-9e88-317ece2525c9","value":"session"}`
	mux.HandleFunc("/customers/my-customer-id/sessions/session", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	expectedSessionResponse := SessionResponse{
		ID:    "5ae33d78-49e5-413a-9e88-317ece2525c9",
		Value: "session",
	}

	sessionResponse, err := testClient.Sessions.Get("my-customer-id", "session")

	if err != nil {
		t.Errorf("Unexpected error. Sessions.Get: %v", err)
	}

	if diff := pretty.Compare(sessionResponse, expectedSessionResponse); diff != "" {
		t.Errorf("Client.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestSessionList(t *testing.T) {
	setup()
	defer teardown()

	response := `[{"id":"5ae33d78-49e5-413a-9e88-317ece2525c9","value":"session"}]`
	mux.HandleFunc("/customers/my-customer-id/sessions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	expectedResponse := []SessionResponse{
		{
			ID:    "5ae33d78-49e5-413a-9e88-317ece2525c9",
			Value: "session",
		},
	}

	sessions, err := testClient.Sessions.List("my-customer-id")

	if err != nil {
		t.Errorf("Unexpected error. Sessions.Update: %v", err)
	}

	if diff := pretty.Compare(sessions, expectedResponse); diff != "" {
		t.Errorf("Sessions.List: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestSessionDelete(t *testing.T) {
	setup()
	defer teardown()

	response := ``
	mux.HandleFunc("/customers/my-customer-id/sessions/session", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Sessions.Delete("my-customer-id", "session")

	if err != nil {
		t.Errorf("Unexpected error. Sessions.Delete: %v", err)
	}
}
