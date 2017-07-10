package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mockServer *httptest.Server

	mux *http.ServeMux

	client *Client
)

func setup() {
	mux = http.NewServeMux()
	mockServer = httptest.NewServer(mux)

	client, _ = New(&Config{
		DefaultNodeID: "fakenodeid",
		WorkspaceID:   "fakeworkspaceid",
		APIkey:        "fakeapikey",
	})
	url, _ := url.Parse(mockServer.URL)
	client.BaseURL = url
}
func teardown() {
	mockServer.Close()
}

func testMethod(t *testing.T, r *http.Request, expectedMethod string) {
	if r.Method != expectedMethod {
		t.Errorf("Expected %v, got %v request method", expectedMethod, r.Method)
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

	req, _ := client.NewRequest(http.MethodGet, "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected error.")
	}
}
