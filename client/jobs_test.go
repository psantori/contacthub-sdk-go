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

func TestJobCreate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"job","companyIndustry":"ict","companyName":"google","startDate":"2012-02-22","isCurrent":true}`
	response := `{"id":"job","companyIndustry":"ict","companyName":"google","startDate":"2012-02-22","endDate":null,"isCurrent":true}`
	mux.HandleFunc("/customers/my-customer-id/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Jobs.Create: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	startDate, _ := time.Parse("2006-01-02", "2012-02-22")
	expectedJobResponse := JobResponse{
		ID:              "job",
		IsCurrent:       null.BoolFrom(true),
		CompanyIndustry: null.StringFrom("ict"),
		CompanyName:     null.StringFrom("google"),
		StartDate:       &SimpleDate{startDate},
		EndDate:         nil,
	}

	job := Job{
		ID:              "job",
		IsCurrent:       nullable.BoolFrom(true),
		CompanyIndustry: nullable.StringFrom("ict"),
		CompanyName:     nullable.StringFrom("google"),
		StartDate:       &SimpleDate{startDate},
	}

	jobResponse, err := testClient.Jobs.Create("my-customer-id", &job)

	if err != nil {
		t.Errorf("Unexpected error. Jobs.Create: %v", err)
	}

	if diff := pretty.Compare(jobResponse, expectedJobResponse); diff != "" {
		t.Errorf("Jobs.Create: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestJobGet(t *testing.T) {
	setup()
	defer teardown()

	response := `{"id":"job","companyIndustry":"ict","companyName":"google","startDate":"2012-02-22","endDate":null,"isCurrent":true}`
	mux.HandleFunc("/customers/my-customer-id/jobs/job", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, response)
	})

	startDate, _ := time.Parse("2006-01-02", "2012-02-22")
	expectedJobResponse := JobResponse{
		ID:              "job",
		IsCurrent:       null.BoolFrom(true),
		CompanyIndustry: null.StringFrom("ict"),
		CompanyName:     null.StringFrom("google"),
		StartDate:       &SimpleDate{startDate},
		EndDate:         nil,
	}

	jobResponse, err := testClient.Jobs.Get("my-customer-id", "job")

	if err != nil {
		t.Errorf("Unexpected error. Jobs.Get: %v", err)
	}

	if diff := pretty.Compare(jobResponse, expectedJobResponse); diff != "" {
		t.Errorf("Client.Get: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestJobUpdate(t *testing.T) {
	setup()
	defer teardown()

	expectedRequestBody := `{"id":"job","companyIndustry":"ICT","companyName":"Google","startDate":"2012-02-22","isCurrent":true}`
	response := `{"id":"job","companyIndustry":"ICT","companyName":"Google","startDate":"2012-02-22","endDate":null,"isCurrent":true}`
	mux.HandleFunc("/customers/my-customer-id/jobs/job", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)

		body, _ := ioutil.ReadAll(r.Body)
		if trimmedBody := strings.TrimSpace(string(body)); trimmedBody != expectedRequestBody {
			t.Errorf("Jobs.Update: invalid body. \nGot: %v\nExpected: %v", trimmedBody, expectedRequestBody)
		}

		fmt.Fprint(w, response)
	})

	startDate, _ := time.Parse("2006-01-02", "2012-02-22")
	expectedJobResponse := JobResponse{
		ID:              "job",
		IsCurrent:       null.BoolFrom(true),
		CompanyIndustry: null.StringFrom("ICT"),
		CompanyName:     null.StringFrom("Google"),
		StartDate:       &SimpleDate{startDate},
		EndDate:         nil,
	}

	job := Job{
		ID:              "job",
		IsCurrent:       nullable.BoolFrom(true),
		CompanyIndustry: nullable.StringFrom("ICT"),
		CompanyName:     nullable.StringFrom("Google"),
		StartDate:       &SimpleDate{startDate},
		EndDate:         nil,
	}

	jobResponse, err := testClient.Jobs.Update("my-customer-id", "job", &job)

	if err != nil {
		t.Errorf("Unexpected error. Jobs.Update: %v", err)
	}

	if diff := pretty.Compare(jobResponse, expectedJobResponse); diff != "" {
		t.Errorf("Jobs.Update: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestJobDelete(t *testing.T) {
	setup()
	defer teardown()

	response := ``
	mux.HandleFunc("/customers/my-customer-id/jobs/job", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, response)
	})

	err := testClient.Jobs.Delete("my-customer-id", "job")

	if err != nil {
		t.Errorf("Unexpected error. Jobs.Delete: %v", err)
	}
}
