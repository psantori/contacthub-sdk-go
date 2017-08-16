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
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/guregu/null/zero"
	"github.com/kylelemons/godebug/pretty"
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
		Errors:   make([]APIError, 1)}

	expected.Errors[0] = APIError{
		Message: "unique customer property is required",
		Path:    "/base/credential/username",
		Data:    nil,
		Code:    zero.StringFrom("")}

	if diff := pretty.Compare(err, expected); diff != "" {
		t.Errorf("TestErrorResponse: invalid value for struct: (-got +expected)\n%s", diff)
	}

}
