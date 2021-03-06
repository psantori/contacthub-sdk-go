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
	"strconv"
	"testing"
	"time"

	"strings"

	"github.com/contactlab/contacthub-sdk-go/nullable"
	"github.com/guregu/null"
	"github.com/kylelemons/godebug/pretty"
)

func TestCustomerList(t *testing.T) {
	setup()
	defer teardown()

	response := `{"page":{"size":2,"totalElements":2,"totalUnfilteredElements":0,"totalPages":1,"number":0},"elements":[{"id":"758b6a0736350sdf972-d9cc0e815502","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2017-06-29T20:23:09.215+0000","updatedAt":"2017-06-29T20:23:09.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"-","lastName":"-","middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":{"email":"Erwdfgin@irizafgfgil.ie","fax":null,"mobilePhone":null,"phone":null,"otherContacts":[],"mobileDevices":[]},"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null},{"id":"daa982b7-e02fdsf17c01b8","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2017-07-07T07:19:54.475+0000","updatedAt":"2017-07-07T07:19:54.475+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"Bobb23y","lastName":"","middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":{"email":"sdf@asdf.ict","fax":null,"mobilePhone":null,"phone":null,"otherContacts":[],"mobileDevices":[]},"address":null,"credential":{"username":"my-ussername","password":null},"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null}]}`
	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testQueryStringValue(t, r, "nodeId", "fakenodeid")
		fmt.Fprint(w, response)
	})

	params := ListParams{}
	customers, pageInfo, err := testClient.Customers.List(&params)

	if err != nil {
		t.Errorf("Unexpected error. Customers.List: %v", err)
	}

	if len(customers) != 2 {
		t.Errorf("Wrong list size. Expected 2, got %v", err)
	}

	registeredAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-06-29T20:23:09.215+0000")
	registeredAt2, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-07-07T07:19:54.475+0000")
	customerAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-06-29T20:23:09.215+0000")
	customerAt2, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-07-07T07:19:54.475+0000")
	expectedCustomers := []CustomerResponse{
		CustomerResponse{
			ID:                 "758b6a0736350sdf972-d9cc0e815502",
			RegisteredAt:       CustomDate{registeredAt},
			UpdatedAt:          CustomDate{customerAt},
			NodeID:             "fakenodeid",
			Enabled:            true,
			ExtendedProperties: nil,
			BaseProperties: &BasePropertiesResponse{
				FirstName: null.StringFrom("-"),
				LastName:  null.StringFrom("-"),
				Contacts: &ContactsResponse{
					Email: null.StringFrom("Erwdfgin@irizafgfgil.ie"),
				},
			},
		},
		CustomerResponse{
			ID:                 "daa982b7-e02fdsf17c01b8",
			RegisteredAt:       CustomDate{registeredAt2},
			UpdatedAt:          CustomDate{customerAt2},
			NodeID:             "fakenodeid",
			Enabled:            true,
			ExtendedProperties: nil,
			BaseProperties: &BasePropertiesResponse{
				FirstName: null.StringFrom("Bobb23y"),
				LastName:  null.StringFrom(""),
				Contacts: &ContactsResponse{
					Email: null.StringFrom("sdf@asdf.ict"),
				},
				Credential: &CredentialResponse{
					Username: null.StringFrom("my-ussername"),
				},
			},
		},
	}
	if diff := pretty.Compare(customers, expectedCustomers); diff != "" {
		t.Errorf("Client.List: invalid value for struct: (-got +expected)\n%s", diff)
	}

	expectedPageInfo := PageInfo{
		Size:                    2,
		TotalElements:           2,
		TotalUnfilteredElements: 0,
		TotalPages:              1,
		Page:                    0,
	}

	if diff := pretty.Compare(pageInfo, expectedPageInfo); diff != "" {
		t.Errorf("Client.List: invalid value for struct: (-got +expected)\n%s", diff)
	}

}
func TestCustomerListByPage(t *testing.T) {
	setup()
	defer teardown()

	responses := make([]string, 2)
	responses[0] = `{"page":{"size":1,"totalElements":3,"totalUnfilteredElements":3,"totalPages":3,"number":0},"elements":[{"id":"758b6a0736350sdf972-d9cc0e815502","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2017-06-29T20:23:09.215+0000","updatedAt":"2017-06-29T20:23:09.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"-","lastName":"-","middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":{"email":"Erwdfgin@irizafgfgil.ie","fax":null,"mobilePhone":null,"phone":null,"otherContacts":[],"mobileDevices":[]},"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null}]}`
	responses[1] = `{"page":{"size":1,"totalElements":3,"totalUnfilteredElements":3,"totalPages":3,"number":1},"elements":[{"id":"758b6a0736350sdf972-d9cc0e815502","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2017-06-29T20:23:09.215+0000","updatedAt":"2017-06-29T20:23:09.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"-","lastName":"-","middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":{"email":"Erwdfgin@irizafgfgil.ie","fax":null,"mobilePhone":null,"phone":null,"otherContacts":[],"mobileDevices":[]},"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null}]}`
	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testQueryStringValue(t, r, "nodeId", "fakenodeid")
		testQueryStringPositiveInt(t, r, "page")
		page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
		fmt.Fprint(w, responses[page])
	})

	params := ListParams{Page: 1}
	_, pageInfo, err := testClient.Customers.List(&params)

	if err != nil {
		t.Errorf("Unexpected error. Customers.List: %v", err)
	}

	if pageInfo.Page != 1 {
		t.Errorf("Wrong page. Expected 1, got %v", pageInfo.Page)
	}
}

func TestCustomerGet(t *testing.T) {
	setup()
	defer teardown()
	response := `{"id":"my-customer-id","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2017-06-29T20:23:09.215+0000","updatedAt":"2017-06-29T20:23:09.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"-","lastName":"-","middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":{"email":"Erwdfgin@irizafgfgil.ie","fax":null,"mobilePhone":null,"phone":null,"otherContacts":[],"mobileDevices":[]},"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null}`
	mux.HandleFunc("/customers/my-customer-id", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	customer, err := testClient.Customers.Get("my-customer-id")
	if err != nil {
		t.Errorf("Unexpected error. Customers.Get: %v", err)
	}

	registeredAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-06-29T20:23:09.215+0000")
	customerAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2017-06-29T20:23:09.215+0000")
	expected := &CustomerResponse{
		ID:                 "my-customer-id",
		RegisteredAt:       CustomDate{registeredAt},
		UpdatedAt:          CustomDate{customerAt},
		NodeID:             "fakenodeid",
		Enabled:            true,
		ExtendedProperties: nil,
		BaseProperties: &BasePropertiesResponse{
			FirstName: null.StringFrom("-"),
			LastName:  null.StringFrom("-"),
			Contacts: &ContactsResponse{
				Email: null.StringFrom("Erwdfgin@irizafgfgil.ie"),
			},
		},
	}

	if diff := pretty.Compare(customer, expected); diff != "" {
		t.Errorf("Customers.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestCustomerDelete(t *testing.T) {
	setup()
	defer teardown()

	response := ``
	mux.HandleFunc("/customers/my-customer-id", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Customers.Delete("my-customer-id")
	if err != nil {
		t.Errorf("Unexpected error. Customers.Delete: %v", err)
	}
}

// Needs more test cases, with different field combinations
func TestCustomerCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"nodeId":"fakenodeid","enabled":true,"extended":{"test":"value"}}`

	response := `{"id":"my-new-customer-id","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2022-02-22T20:22:22.215+0000","updatedAt":"2022-02-22T20:22:22.215+0000","enabled":true,"extended":{"test":"value"},"base":null,"tags":null}`
	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Client.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	registeredAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	expectedCustomerResponse := CustomerResponse{
		ID:           "my-new-customer-id",
		RegisteredAt: CustomDate{registeredAt},
		UpdatedAt:    CustomDate{registeredAt},
		Enabled:      true,
		NodeID:       "fakenodeid",
		ExtendedProperties: &map[string]interface{}{
			"test": "value",
		},
	}

	customer := Customer{
		NodeID:  testClient.Config.DefaultNodeID,
		Enabled: nullable.BoolFrom(true),
		ExtendedProperties: &map[string]interface{}{
			"test": "value",
		},
	}
	customerResponse, err := testClient.Customers.Create(&customer)

	if err != nil {
		t.Errorf("Unexpected error. Customers.Create: %v", err)
	}

	if diff := pretty.Compare(customerResponse, expectedCustomerResponse); diff != "" {
		t.Errorf("Client.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestCustomerCreateWithBaseProperties(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"nodeId":"fakenodeid","enabled":true,"base":{"firstName":"John","lastName":null}}`

	response := `{"id":"my-new-customer-id","nodeId":"fakenodeid","externalId":null,"extra":null,"registeredAt":"2022-02-22T20:22:22.215+0000","updatedAt":"2022-02-22T20:22:22.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"John","lastName":null,"middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":null,"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"tags":null}`
	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Client.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	registeredAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	expectedCustomerResponse := CustomerResponse{
		ID:           "my-new-customer-id",
		RegisteredAt: CustomDate{registeredAt},
		UpdatedAt:    CustomDate{registeredAt},
		Enabled:      true,
		NodeID:       "fakenodeid",
		BaseProperties: &BasePropertiesResponse{
			FirstName: null.StringFrom("John"),
		},
	}

	customer := Customer{
		NodeID:  testClient.Config.DefaultNodeID,
		Enabled: nullable.BoolFrom(true),
		BaseProperties: &BaseProperties{
			FirstName: nullable.StringFrom("John"),
			LastName:  nullable.NullString(),
		},
	}
	customerResponse, err := testClient.Customers.Create(&customer)

	if err != nil {
		t.Errorf("Unexpected error. Customers.Create: %v", err)
	}

	if diff := pretty.Compare(customerResponse, expectedCustomerResponse); diff != "" {
		t.Errorf("Client.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestCustomerUpdate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"externalId":"my-external-id","enabled":true,"base":{"pictureUrl":null,"firstName":"John"}}`
	response := `{"id":"my-new-customer-id","nodeId":"fakenodeid","externalId":"my-external-id","extra":null,"registeredAt":"2022-02-22T20:22:22.215+0000","updatedAt":"2022-02-22T23:23:22.215+0000","enabled":true,"base":{"pictureUrl":null,"title":null,"prefix":null,"firstName":"John","lastName":null,"middleName":null,"gender":null,"dob":null,"locale":null,"timezone":null,"contacts":null,"address":null,"credential":null,"educations":[],"likes":[],"socialProfile":null,"jobs":[],"subscriptions":[]},"extended":null,"tags":null}`

	mux.HandleFunc("/customers/my-new-customer-id", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Client.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}
		fmt.Fprint(w, response)
	})
	customer := Customer{
		ExternalID: nullable.StringFrom("my-external-id"),
		Enabled:    nullable.BoolFrom(true),
		BaseProperties: &BaseProperties{
			FirstName:  nullable.StringFrom("John"),
			PictureURL: nullable.NullString(),
		},
	}
	customerResponse, err := testClient.Customers.Update("my-new-customer-id", &customer)

	if err != nil {
		t.Errorf("Unexpected error. Customers.Update: %v", err)
	}

	registeredAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
	updatedAt, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T23:23:22.215+0000")
	expectedCustomerResponse := CustomerResponse{
		ID:           "my-new-customer-id",
		RegisteredAt: CustomDate{registeredAt},
		UpdatedAt:    CustomDate{updatedAt},
		Enabled:      true,
		NodeID:       "fakenodeid",
		ExternalID:   nullable.StringFrom("my-external-id"),
		BaseProperties: &BasePropertiesResponse{
			FirstName: null.StringFrom("John"),
		},
	}
	if diff := pretty.Compare(customerResponse, expectedCustomerResponse); diff != "" {
		t.Errorf("Client.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}
