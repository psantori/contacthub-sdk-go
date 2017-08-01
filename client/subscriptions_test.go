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
