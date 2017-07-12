package client

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

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
			ID:           "758b6a0736350sdf972-d9cc0e815502",
			RegisteredAt: CustomDate{registeredAt},
			UpdatedAt:    CustomDate{customerAt},
			Customer: Customer{
				NodeID:             "fakenodeid",
				Enabled:            true,
				ExtendedProperties: nil,
				BaseProperties: &BaseProperties{
					FirstName: null.StringFrom("-"),
					LastName:  null.StringFrom("-"),
					Contacts: &Contacts{
						Email:         null.StringFrom("Erwdfgin@irizafgfgil.ie"),
						OtherContacts: []OtherContact{},
						MobileDevices: []MobileDevice{},
					},
					Educations:    []Education{},
					Likes:         []Like{},
					Jobs:          []Job{},
					Subscriptions: []Subscription{},
				},
			},
		},
		CustomerResponse{
			ID:           "daa982b7-e02fdsf17c01b8",
			RegisteredAt: CustomDate{registeredAt2},
			UpdatedAt:    CustomDate{customerAt2},
			Customer: Customer{
				NodeID:             "fakenodeid",
				Enabled:            true,
				ExtendedProperties: nil,
				BaseProperties: &BaseProperties{
					FirstName: null.StringFrom("Bobb23y"),
					LastName:  null.StringFrom(""),
					Contacts: &Contacts{
						Email:         null.StringFrom("sdf@asdf.ict"),
						OtherContacts: []OtherContact{},
						MobileDevices: []MobileDevice{},
					},
					Credential: &Credential{
						Username: null.StringFrom("my-ussername"),
					},
					Educations:    []Education{},
					Likes:         []Like{},
					Jobs:          []Job{},
					Subscriptions: []Subscription{},
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
		ID:           "my-customer-id",
		RegisteredAt: CustomDate{registeredAt},
		UpdatedAt:    CustomDate{customerAt},
		Customer: Customer{
			NodeID:             "fakenodeid",
			Enabled:            true,
			ExtendedProperties: nil,
			BaseProperties: &BaseProperties{
				FirstName: null.StringFrom("-"),
				LastName:  null.StringFrom("-"),
				Contacts: &Contacts{
					Email:         null.StringFrom("Erwdfgin@irizafgfgil.ie"),
					OtherContacts: []OtherContact{},
					MobileDevices: []MobileDevice{},
				},
				Educations:    []Education{},
				Likes:         []Like{},
				Jobs:          []Job{},
				Subscriptions: []Subscription{},
			},
		},
	}
	if diff := pretty.Compare(customer, expected); diff != "" {
		t.Errorf("Client.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}

}
