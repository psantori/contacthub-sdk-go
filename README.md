# Basic usage
Import this library
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
You can access the API methods via apiClient.Customers, apiClient.Events, etc.

# Create Customer
```go
newCustomer := client.Customer{
  ExternalID: "my-external-id",
  Enabled:    true,
  BaseProperties: &client.BaseProperties{
    FirstName: null.StringFrom("John"),
    LastName:  null.StringFrom("Von Neumann"),
    Credential: &client.Credential{
      Username: null.StringFrom("john-von-neumann")},
    Contacts: &client.Contacts{
      Email: null.StringFrom("email@email.it")}}}
createdCustomer, err := apiClient.Customers.Create(&newCustomer)
```
# Retrieve Customers
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



