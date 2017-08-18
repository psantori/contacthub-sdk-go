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
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/contactlab/contacthub-sdk-go/nullable"
	"github.com/guregu/null"
	"github.com/kylelemons/godebug/pretty"
)

func TestLikeCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"like","category":"category","name":"name","createdTime":"2022-02-22T20:22:22.215+0000"}`
	response := `{"id":"like","category":"category","name":"name","createdTime":"2022-02-22T20:22:22.215+0000"}`
	mux.HandleFunc("/customers/my-customer-id/likes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Likes.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")

	expectedLikeResponse := LikeResponse{
		ID:          "like",
		Category:    null.StringFrom("category"),
		Name:        null.StringFrom("name"),
		CreatedTime: &CustomDate{createdTime},
	}

	like := Like{
		ID:          "like",
		Category:    nullable.StringFrom("category"),
		Name:        nullable.StringFrom("name"),
		CreatedTime: &CustomDate{createdTime},
	}

	likeResponse, err := testClient.Likes.Create("my-customer-id", &like)

	if err != nil {
		t.Errorf("Unexpected error. Likes.Create: %v", err)
	}

	if diff := pretty.Compare(likeResponse, expectedLikeResponse); diff != "" {
		t.Errorf("Likes.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestLikeGet(t *testing.T) {
	setup()
	defer teardown()

	response := `{"id":"like","category":"category","name":"name","createdTime":"2022-02-22T20:22:22.215+0000"}`
	mux.HandleFunc("/customers/my-customer-id/likes/like", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")

	expectedLikeResponse := LikeResponse{
		ID:          "like",
		Category:    null.StringFrom("category"),
		Name:        null.StringFrom("name"),
		CreatedTime: &CustomDate{createdTime},
	}

	likeResponse, err := testClient.Likes.Get("my-customer-id", "like")

	if err != nil {
		t.Errorf("Unexpected error. Likes.Get: %v", err)
	}

	if diff := pretty.Compare(likeResponse, expectedLikeResponse); diff != "" {
		t.Errorf("Client.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestLikeUpdate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"like","category":"category2","name":"name","createdTime":"2022-02-22T20:22:22.215+0000"}`
	response := `{"id":"like","category":"category2","name":"name","createdTime":"2022-02-22T20:22:22.215+0000"}`
	mux.HandleFunc("/customers/my-customer-id/likes/like", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Likes.Update: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")

	expectedLikeResponse := LikeResponse{
		ID:          "like",
		Category:    null.StringFrom("category2"),
		Name:        null.StringFrom("name"),
		CreatedTime: &CustomDate{createdTime},
	}

	like := Like{
		ID:          "like",
		Category:    nullable.StringFrom("category2"),
		Name:        nullable.StringFrom("name"),
		CreatedTime: &CustomDate{createdTime},
	}

	likeResponse, err := testClient.Likes.Update("my-customer-id", "like", &like)

	if err != nil {
		t.Errorf("Unexpected error. Likes.Update: %v", err)
	}

	if diff := pretty.Compare(likeResponse, expectedLikeResponse); diff != "" {
		t.Errorf("Likes.Update: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestLikeDelete(t *testing.T) {
	setup()
	defer teardown()

	response := ``
	mux.HandleFunc("/customers/my-customer-id/likes/like", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Likes.Delete("my-customer-id", "like")

	if err != nil {
		t.Errorf("Unexpected error. Likes.Delete: %v", err)
	}
}
