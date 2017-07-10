package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/guregu/null"
)

type ErrorResponse struct {
	*http.Response
	Message string      `json:"message"`
	Logref  string      `json:"logref"`
	Data    null.String `json:"data"`
	Errors  []ApiError  `json:"errors"`
}

type ApiError struct {
	Message string      `json:"message,omitempty"`
	Path    string      `json:"path,omitempty"`
	Data    interface{} `json:"data"`
}

func (r *ErrorResponse) Error() string {
	messages := make([]string, 1+len(r.Errors))
	messages[0] = r.Message

	for i, apiError := range r.Errors {
		messages[i+1] = fmt.Sprintf("%v (%v)", apiError.Message, apiError.Path)
	}
	return fmt.Sprintf("(%d) %v %v: %v",
		r.Response.StatusCode, r.Response.Request.Method, r.Response.Request.URL, strings.Join(messages, ", "))
}

func handleErrors(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}