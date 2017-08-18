# Setup

Import the client library
```go
import "github.com/contactlab/contacthub-sdk-go/client"
```
Create a new ContactHub client:
```go
config := &client.Config{
  DefaultNodeID: "*** (Optional) NODE ID ***",
  WorkspaceID:   "*** WORKSPACE ID ***",
  APIkey:        "*** API KEY ****",
}

apiClient, err := client.New(config)
```

# General usage
The library provides access to the ContactHub API via apiClient.Customers, apiClient.Events, etc., using a set
of structs modelled after the API schema.

All optional fields are pointers to guregu/null/ types, in order to differentiate null from empty values and to support patch operations. The `github.com/contactlab/contacthub-sdk-go/nullable` package provides helper methods to instantiate those types.

# Customers API
## Create a Customer
```go
newCustomer := client.Customer{
  ExternalID: "my-external-id",
  Enabled:    nullable.BoolFrom(true),
  BaseProperties: &client.BaseProperties{
    FirstName: nullable.StringFrom("John"),
    LastName:  nullable.StringFrom("Von Neumann"),
    Credential: &client.Credential{
      Username: nullable.StringFrom("john-von-neumann"),
    },
    Contacts: &client.Contacts{
      Email: nullable.StringFrom("email@email.it"),
    }
  },
}
createdCustomer, err := apiClient.Customers.Create(&newCustomer)
```
## Retrieve Customer by ID
```go
customerResponse, err := apiClient.Customers.Get("customerID")
```

## Delete a Customer
```go
err := apiClient.Customers.Delete("customerID")
```

## Update a Customer (patch)
Note that almost all fields are optional and nullable. See docs for [Customer Patch](http://developer.contactlab.com/documentation/contacthub/api/index#operation/patchCustomerWorkspaces) and [BaseProperties schema](http://developer.contactlab.com/documentation/contacthub/schemas/customer.baseProperties.html)
```go
customerPatch := Customer{
  NodeID:     testClient.Config.DefaultNodeID,
  Enabled:    true,
  BaseProperties: &BaseProperties{
    FirstName:  nullable.StringFrom("John"), // Change the firstName
    PictureURL: nullable.NullString(), // This unsets the previous PictureURL
  },
}
customerResponse, err := apiClient.Customers.Update("customerID", customerPatch)
```

## Retrieve a list of Customers
```go
params := api.ListParams{PageSize: 50, Page: 0}
// Pagination loop to get all Customers
for {
  customers, pageInfo, err := apiClient.Customers.List(&params)

  if err != nil {
    // Handle any errors
    panic(err)
  }

  for _, customer := range customers {
    // Do something with the customer
    fmt.Println(customer.ID)
  }
  if !pageInfo.HasNextPage() {
    break
  }
  params.Page++
}
```

## Retrieve a list of Customers matching an ExternalId
```go
params := api.ListParams{PageSize: 50, QueryParams{
  "externalId": "my-external-id"
}}
customers, pageInfo, err := apiClient.Customers.List(&params)
```

# Events API

## Create an Event
```go
event := Event{
  CustomerID: nullable.StringFrom("customer-id"),
  Type:       enums.AbandonedCart,
  Properties: map[string]interface{}{},
  Context:    enums.Ecommerce,
  Date:       &CustomDate{createdAt},
}
eventResponse, err := testClient.Events.Create(&event)
```

## Delete an Event

```go
err := testClient.Events.Delete("eventId")
```
## Retrieve a list of Events for a Customer
```go
params := api.ListParams{PageSize: 50}
// Pagination loop to get all Events
for {
  events, pageInfo, err := apiClient.Events.List("customerID", &params)

  if err != nil {
    // Handle any errors
    panic(err)
  }

  for _, event := range events {
    // Do something with the event
    fmt.Println(event.ID)
  }
  if !pageInfo.HasNextPage() {
    break
  }
  params.Page++
}
```

# Sessions API

## Create session for a Customer
This can be used to reconcile anonymous events with an existing Customer. The Session Value field should coincide with the SessionId set as a BringBackProperty for the events.

```go
sessionResponse, err := apiClient.Sessions.Create("my-customer-id", &Session{"my-session"})
```

## List sessions for a Customer
Note there is no pagination
```go
sessionResponses, err := apiClient.Sessions.List("my-customer-id")
```

## Delete session
```go
err := apiClient.Sessions.Delete("my-customer-id", "my-session")
```

# Other APIs
This library implements all the other sub-endpoints for the Customers:
- Subscriptions
- Jobs
- Educations
- Likes

Note that all operations on the Customer Subscriptions, Jobs, Educations and Likes can be also performed via Customers.Update, even with partial updates.


## Subscriptions API
### Add a subscription for a Customer

```go
startDate, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
subscriptionKind := enums.DigitalMessage
subscription := Subscription{
  ID:         "my-subscription",
  Subscribed: nullable.BoolFrom(true), // optional
  Kind:       &subscriptionKind, // optional
  StartDate:  &CustomDate{startDate}, // optional
}

subscriptionResponse, err := apiClient.Subscriptions.Create("my-customer-id", &subscription)
```

### Get a Subscription
```go
subscriptionResponse, err := apiClient.Subscriptions.Get("my-customer-id", "my-subscription")
```

### Update a subscription (put)
```go
startDate, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
endDate, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-07-22T20:22:22.215+0000")
subscriptionKind := enums.DigitalMessage
subscription := Subscription{
  ID:         "my-subscription",
  Subscribed: nullable.BoolFrom(false),
  StartDate:  &CustomDate{startDate},
  Kind:       &subscriptionKind,
  EndDate:    &CustomDate{endDate},
}
	err := apiClient.Subscriptions.Update("my-customer-id", "my-subscription", &subscription)
```

### Delete a subscription
```go
err := apiClient.Subscriptions.Delete("my-customer-id", "my-subscription")
```

## Jobs API

### Add a Job
```go
startDate, _ := time.Parse("2006-01-02", "2012-02-22")
job := Job{
  ID:              "job",
  IsCurrent:       nullable.BoolFrom(true),
  CompanyIndustry: nullable.StringFrom("ICT"),
  CompanyName:     nullable.StringFrom("Google"),
  StartDate:       &SimpleDate{startDate},
}

jobResponse, err := apiClient.Jobs.Create("my-customer-id", &job)
```

### Get a Job
```go
jobResponse, err := apiClient.Jobs.Get("my-customer-id", "you-really-should")
```

### Update a Job (put)
```go
startDate, _ := time.Parse("2006-01-02", "2012-02-22")
endDate, _ := time.Parse("2006-01-02", "2022-02-22")
job := Job{
  ID:              "job",
  IsCurrent:       nullable.BoolFrom(false),
  CompanyIndustry: nullable.StringFrom("ICT"),
  CompanyName:     nullable.StringFrom("Google"),
  StartDate:       &SimpleDate{startDate},
  EndDate:         &SimpleDate{endDate},
  EndDate:         nil,
}

jobResponse, err := apiClient.Jobs.Update("my-customer-id", "job", &job)
```

### Delete a Job
```go
err := apiClient.Jobs.Delete("my-customer-id", "job")
```

## Educations API

### Add an Education
```go
schoolType := enums.PrimarySchool
education := Education{
  ID:                  "education",
  SchoolType:          &schoolType,
  SchoolName:          nullable.StringFrom("School"),
  SchoolConcentration: nullable.StringFrom("stuff"),
  StartYear:           nullable.IntFrom(1996),
  EndYear:             nullable.IntFrom(2001),
  IsCurrent:           nullable.BoolFrom(false),
}

educationResponse, err := testClient.Educations.Create("my-customer-id", &education)
```

### Get an Education
```go
educationResponse, err := apiClient.Educations.Get("my-customer-id", "education")
```

### Update an Education (put)
```go
schoolType := enums.PrimarySchool
education := Education{
  ID:                  "education",
  SchoolType:          &schoolType,
  SchoolName:          nullable.StringFrom("School name"),
  SchoolConcentration: nullable.StringFrom("something"),
  StartYear:           nullable.IntFrom(1996),
  EndYear:             nullable.IntFrom(2001),
  IsCurrent:           nullable.BoolFrom(false),
}

educationResponse, err := testClient.Educations.Update("my-customer-id", "education", &education)
```

### Delete an Education
```go
err := apiClient.Educations.Delete("my-customer-id", "education")
```

## Likes API

### Add a Like
```go
createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
like := Like{
  ID:          "like",
  Category:    nullable.StringFrom("category"),
  Name:        nullable.StringFrom("name"),
  CreatedTime: &CustomDate{createdTime},
}

likeResponse, err := testClient.Likes.Create("my-customer-id", &like)
```
### Get a Like
```go
likeResponse, err := apiClient.Likes.Get("my-customer-id", "like")
```

### Update a Like (put)
```go
createdTime, _ := time.Parse("2006-01-02T15:04:05.999-0700", "2022-02-22T20:22:22.215+0000")
like := Like{
  ID:          "like",
  Name:        nullable.StringFrom("other name"),
  CreatedTime: &CustomDate{createdTime},
}

likeResponse, err := testClient.Likes.Update("my-customer-id", "like", &like)
```

### Delete a Like
```go
err := apiClient.Likes.Delete("my-customer-id", "like")
```
