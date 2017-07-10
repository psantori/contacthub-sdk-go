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
