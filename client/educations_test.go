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

	"github.com/contactlab/contacthub-sdk-go/enums"
	"github.com/contactlab/contacthub-sdk-go/nullable"
	"github.com/guregu/null"
	"github.com/kylelemons/godebug/pretty"
)

func TestEducationCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"education","schoolType":"PRIMARY_SCHOOL","schoolName":"school","schoolConcentration":"something","startYear":1996,"endYear":2001,"isCurrent":false}`
	response := `{"id":"education","schoolType":"PRIMARY_SCHOOL","schoolName":"school","schoolConcentration":"something","startYear":1996,"endYear":2001,"isCurrent":false}`
	mux.HandleFunc("/customers/my-customer-id/educations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Educations.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	schoolType := enums.PrimarySchool
	expectedEducationResponse := EducationResponse{
		ID:                  "education",
		SchoolType:          &schoolType,
		SchoolName:          null.StringFrom("school"),
		SchoolConcentration: null.StringFrom("something"),
		StartYear:           null.IntFrom(1996),
		EndYear:             null.IntFrom(2001),
		IsCurrent:           null.BoolFrom(false),
	}

	education := Education{
		ID:                  "education",
		SchoolType:          &schoolType,
		SchoolName:          nullable.StringFrom("school"),
		SchoolConcentration: nullable.StringFrom("something"),
		StartYear:           nullable.IntFrom(1996),
		EndYear:             nullable.IntFrom(2001),
		IsCurrent:           nullable.BoolFrom(false),
	}

	educationResponse, err := testClient.Educations.Create("my-customer-id", &education)

	if err != nil {
		t.Errorf("Unexpected error. Educations.Create: %v", err)
	}

	if diff := pretty.Compare(educationResponse, expectedEducationResponse); diff != "" {
		t.Errorf("Educations.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestEducationGet(t *testing.T) {
	setup()
	defer teardown()

	response := `{"id":"education","schoolType":"PRIMARY_SCHOOL","schoolName":"school","schoolConcentration":"something","startYear":1996,"endYear":2001,"isCurrent":false}`
	mux.HandleFunc("/customers/my-customer-id/educations/education", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	schoolType := enums.PrimarySchool
	expectedEducationResponse := EducationResponse{
		ID:                  "education",
		SchoolType:          &schoolType,
		SchoolName:          null.StringFrom("school"),
		SchoolConcentration: null.StringFrom("something"),
		StartYear:           null.IntFrom(1996),
		EndYear:             null.IntFrom(2001),
		IsCurrent:           null.BoolFrom(false),
	}

	educationResponse, err := testClient.Educations.Get("my-customer-id", "education")

	if err != nil {
		t.Errorf("Unexpected error. Educations.Get: %v", err)
	}

	if diff := pretty.Compare(educationResponse, expectedEducationResponse); diff != "" {
		t.Errorf("Educations.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestEducationUpdate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"education","schoolType":"PRIMARY_SCHOOL","schoolName":"School","schoolConcentration":"stuff","startYear":1996,"endYear":2001,"isCurrent":true}`
	response := `{"id":"education","schoolType":"PRIMARY_SCHOOL","schoolName":"School","schoolConcentration":"stuff","startYear":1996,"endYear":2001,"isCurrent":true}`
	mux.HandleFunc("/customers/my-customer-id/educations/education", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Educations.Update: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	schoolType := enums.PrimarySchool
	expectedEducationResponse := EducationResponse{
		ID:                  "education",
		SchoolType:          &schoolType,
		SchoolName:          null.StringFrom("School"),
		SchoolConcentration: null.StringFrom("stuff"),
		StartYear:           null.IntFrom(1996),
		EndYear:             null.IntFrom(2001),
		IsCurrent:           null.BoolFrom(true),
	}

	education := Education{
		ID:                  "education",
		SchoolType:          &schoolType,
		SchoolName:          nullable.StringFrom("School"),
		SchoolConcentration: nullable.StringFrom("stuff"),
		StartYear:           nullable.IntFrom(1996),
		EndYear:             nullable.IntFrom(2001),
		IsCurrent:           nullable.BoolFrom(true),
	}

	educationResponse, err := testClient.Educations.Update("my-customer-id", "education", &education)

	if err != nil {
		t.Errorf("Unexpected error. Educations.Update: %v", err)
	}

	if diff := pretty.Compare(educationResponse, expectedEducationResponse); diff != "" {
		t.Errorf("Educations.Update: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestEducationDelete(t *testing.T) {
	setup()
	defer teardown()

	response := ``
	mux.HandleFunc("/customers/my-customer-id/educations/education", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Educations.Delete("my-customer-id", "education")

	if err != nil {
		t.Errorf("Unexpected error. Educations.Delete: %v", err)
	}
}
