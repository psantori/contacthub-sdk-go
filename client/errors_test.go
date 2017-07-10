package client

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/guregu/null/zero"
)

func TestErrorResponse(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{
    "message": "source customer is not valid",
    "logref": "xxxxx-a4bc-xxxx-a6df-xxxxxxxx",
    "data": null,
    "errors": [
        {
            "message": "unique customer property is required",
            "path": "/base/credential/username",
            "data": null,
            "code": null
        }
    ]
}`)),
	}
	err := handleErrors(res).(*ErrorResponse)

	if err == nil {
		t.Fatalf("Expected error response.")
	}

	expected := &ErrorResponse{
		Response: res,
		Message:  "source customer is not valid",
		Logref:   "xxxxx-a4bc-xxxx-a6df-xxxxxxxx",
		Data:     zero.StringFrom(""),
		Errors:   make([]ApiError, 1)}

	expected.Errors[0] = ApiError{
		Message: "unique customer property is required",
		Path:    "/base/credential/username",
		Data:    nil,
		Code:    zero.StringFrom("")}

	if !reflect.DeepEqual(err, expected) {
		t.Errorf("Got: %#v", err)
		t.Errorf("Expected: %#v", expected)
	}
}
