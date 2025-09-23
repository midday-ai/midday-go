# Customers
(*Customers*)

## Overview

### Available Operations

* [List](#list) - List all customers
* [Create](#create) - Create customer
* [Get](#get) - Retrieve a customer
* [Delete](#delete) - Delete a customer
* [Update](#update) - Update a customer

## List

Retrieve a list of customers for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCustomers" method="get" path="/customers" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Customers.List(ctx, middaygo.Pointer("acme"), []string{
        "name",
        "asc",
    }, middaygo.Pointer("eyJpZCI6IjEyMyJ9"), middaygo.Pointer[float64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `q`                                                      | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | acme                                                     |
| `sort`                                                   | []*string*                                               | :heavy_minus_sign:                                       | N/A                                                      | [<br/>"name",<br/>"asc"<br/>]                            |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | eyJpZCI6IjEyMyJ9                                         |
| `pageSize`                                               | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      | 20                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListCustomersResponse](../../models/operations/listcustomersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new customer for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCustomer" method="post" path="/customers" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Customers.Create(ctx, &operations.CreateCustomerRequest{
        ID: middaygo.Pointer("b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4"),
        Name: "Acme Corporation",
        Email: "contact@acme.com",
        BillingEmail: middaygo.Pointer("finance@acme.com"),
        Country: middaygo.Pointer("United States"),
        AddressLine1: middaygo.Pointer("123 Main Street"),
        AddressLine2: middaygo.Pointer("Suite 400"),
        City: middaygo.Pointer("San Francisco"),
        State: middaygo.Pointer("California"),
        Zip: middaygo.Pointer("94105"),
        Phone: middaygo.Pointer("+1-555-123-4567"),
        Website: middaygo.Pointer("https://acme.com"),
        Note: middaygo.Pointer("Preferred contact method is email. Large enterprise client."),
        VatNumber: middaygo.Pointer("US123456789"),
        CountryCode: middaygo.Pointer("US"),
        Contact: middaygo.Pointer("John Smith"),
        Tags: []operations.CreateCustomerTagRequest{
            operations.CreateCustomerTagRequest{
                ID: "e7a9c1a2-4c2a-4e7a-9c1a-2b7c1e24c2a4",
                Name: "VIP",
            },
            operations.CreateCustomerTagRequest{
                ID: "f1b2c3d4-5678-4e7a-9c1a-2b7c1e24c2a4",
                Name: "Enterprise",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateCustomerRequest](../../models/operations/createcustomerrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateCustomerResponse](../../models/operations/createcustomerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a customer by ID for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerById" method="get" path="/customers/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Customers.Get(ctx, "b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCustomerByIDResponse](../../models/operations/getcustomerbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a customer by ID for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCustomer" method="delete" path="/customers/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Customers.Delete(ctx, "b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteCustomerResponse](../../models/operations/deletecustomerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a customer by ID for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCustomer" method="patch" path="/customers/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Customers.Update(ctx, "b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4", &operations.UpdateCustomerRequestBody{
        ID: middaygo.Pointer("b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4"),
        Name: "Acme Corporation",
        Email: "contact@acme.com",
        BillingEmail: middaygo.Pointer("finance@acme.com"),
        Country: middaygo.Pointer("United States"),
        AddressLine1: middaygo.Pointer("123 Main Street"),
        AddressLine2: middaygo.Pointer("Suite 400"),
        City: middaygo.Pointer("San Francisco"),
        State: middaygo.Pointer("California"),
        Zip: middaygo.Pointer("94105"),
        Phone: middaygo.Pointer("+1-555-123-4567"),
        Website: middaygo.Pointer("https://acme.com"),
        Note: middaygo.Pointer("Preferred contact method is email. Large enterprise client."),
        VatNumber: middaygo.Pointer("US123456789"),
        CountryCode: middaygo.Pointer("US"),
        Contact: middaygo.Pointer("John Smith"),
        Tags: []operations.UpdateCustomerTagRequest{
            operations.UpdateCustomerTagRequest{
                ID: "e7a9c1a2-4c2a-4e7a-9c1a-2b7c1e24c2a4",
                Name: "VIP",
            },
            operations.UpdateCustomerTagRequest{
                ID: "f1b2c3d4-5678-4e7a-9c1a-2b7c1e24c2a4",
                Name: "Enterprise",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           | b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4                                                          |
| `requestBody`                                                                                 | [*operations.UpdateCustomerRequestBody](../../models/operations/updatecustomerrequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.UpdateCustomerResponse](../../models/operations/updatecustomerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |