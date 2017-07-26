package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout time.Duration = 5000
	version                      = "0.0.1"
	defaultBaseURL               = "https://api.contactlab.it/"
	contentType                  = "application/json"
	userAgent                    = "golang+contacthub/" + version
)

// Config contains the basic Client configuration
type Config struct {
	APIkey        string
	APIVersion    string
	DefaultNodeID string
	WorkspaceID   string
	Timeout       time.Duration
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
func New(config *Config) (*Client, error) {

	if config.APIVersion == "" {
		config.APIVersion = "v1"
	}
	if config.APIkey == "" {
		return nil, errors.New("APIkey is a required field")
	}
	if config.WorkspaceID == "" {
		return nil, errors.New("WorkspaceID is a required field")
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	baseURL.Path = "hub/" + config.APIVersion + "/" + "workspaces/" + config.WorkspaceID + "/"

	if config.Timeout < 1 {
		config.Timeout = DefaultTimeout
	}
	httpClient := &http.Client{
		Timeout: config.Timeout * time.Millisecond,
	}
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Config: config}

	c.Customers = &CustomerService{client: c}
	c.Events = &EventService{client: c}
	return c, nil
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

	req.Header.Add("Accept", contentType)
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", userAgent)
	return req, nil
}

// Do actually perform the request
func (c *Client) Do(req *http.Request, into interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	defer closeResponse(resp)

	if err != nil {
		return nil, err
	}

	// Handle API errors
	err = handleErrors(resp)
	if err != nil {
		return resp, err
	}

	// No decoding needed
	if into == nil {
		return resp, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
		if err == io.EOF { // Empty body is not necessarily an error
			err = nil
		}
		return resp, err
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

// closeResponse makes sure the Body was completely read and closed, in order to be able to reuse the connection
// See https://github.com/google/go-github/pull/317
func closeResponse(r *http.Response) {
	if r != nil {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}
}
