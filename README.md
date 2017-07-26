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

#Customers
## Create Customer
```go
newCustomer := client.Customer{
  ExternalID: "my-external-id",
  Enabled:    true,
  BaseProperties: &client.BaseProperties{
    FirstName: nullable.StringFrom("John"),
    LastName:  nullable.StringFrom("Von Neumann"),
    Credential: &client.Credential{
      Username: nullable.StringFrom("john-von-neumann")},
    Contacts: &client.Contacts{
      Email: nullable.StringFrom("email@email.it")}}}
createdCustomer, err := apiClient.Customers.Create(&newCustomer)
```
## Retrieve Customer by ID
```go
customerResponse, err := apiClient.Customers.Get("customerID")
```

## Update a Customer
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


