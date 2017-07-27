package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestEventList(t *testing.T) {
	setup()
	defer teardown()

	response := `{"page":{"size":2,"totalElements":10,"totalUnfilteredElements":0,"totalPages":5,"number":0},"elements":[{"id":"my-event-id1","customerId":"my-customer-id","bringBackProperties":null,"type":"viewedProduct","context":"ECOMMERCE","date":"2017-07-13T16:14:46.019+0000","properties":{"id":"999","sku":"test-something","name":"kit bundle (Test)","price":0,"imageUrl":"https://something.jpg","linkUrl":"https://store.example","shortDescription":"Test bundle","category":[]},"contextInfo":{"client":{"ip":"111.111.111.111","userAgent":"Somebrowser/1.1"}},"registeredAt":"2017-07-13T08:19:11.019+0000","updatedAt":"2017-07-13T08:19:11.019+0000","tracking":null},{"id":"my-event-id2","customerId":"my-customer-id","bringBackProperties":null,"type":"reviewedProduct","context":"ECOMMERCE","date":"2017-07-13T16:14:48.019+0000","properties":{"id":"10000","sku":"test-something2","name":"kit bundle2 (Test)","price":0,"imageUrl":"https://something2.jpg","linkUrl":"https://store.example","shortDescription":"Test bundle 2","category":[]},"contextInfo":{"client":{"ip":"111.111.111.111","userAgent":"Somebrowser/1.1"}},"registeredAt":"2017-07-13T08:19:11.01+0000","updatedAt":"2017-07-13T08:19:11.019+0000","tracking":null}]}`
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
