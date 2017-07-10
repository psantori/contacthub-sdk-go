package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	version        = "0.0.1"
	defaultBaseURL = "https://api.contactlab.it/"
	mediaType      = "application/json"
	format         = "json"
	userAgent      = "golang+contacthub/" + version
)

// Config contains the basic Client configuration
type Config struct {
	APIkey        string
	APIVersion    string
	WorkspaceID   string
	DefaultNodeID string
}

// QueryParams is simply a map of query paramss
type QueryParams map[string]string

// Client is the actual ContactHub API client
type Client struct {

	// Base URL for API requests.
	BaseURL *url.URL

	// HTTP client
	client *http.Client

	// User agent
	UserAgent string

	// Basic config for contacthub API
	Config *Config

	Customers *CustomerService

	Events *EventService
}

// New creates a new API client
func New(config *Config) *Client {

	if config.APIVersion == "" {
		config.APIVersion = "v1"
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	baseURL.Path = "hub/" + config.APIVersion + "/" + "workspaces/" + config.WorkspaceID + "/"

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Config: config}

	c.Customers = &CustomerService{client: c}
	c.Events = &EventService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	requestedURL := c.BaseURL.ResolveReference(rel)

	encBody, err := prepareBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, requestedURL.String(), encBody)
	if err != nil {
		return nil, err
	}

	bearerToken := "Bearer " + c.Config.APIkey
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)
	return req, nil
}

// Do actually perform the request
func (c *Client) Do(req *http.Request, into interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	err = handleErrors(resp)
	if err != nil {
		return resp, err
	}
	if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
		return nil, err
	}

	return resp, nil
}

// addQuery sets the query string parameters
func addQuery(basePath string, queryParams QueryParams) string {
	// Specify URL query string parameters
	params := url.Values{}
	for k, v := range queryParams {
		params.Add(k, v)
	}

	path := basePath + "?" + params.Encode()
	return path
}

func prepareBody(body interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return buf, err
		}
	}
	return buf, nil
}
