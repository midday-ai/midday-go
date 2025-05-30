# BankAccounts
(*BankAccounts*)

## Overview

### Available Operations

* [List](#list) - List all bank accounts
* [Create](#create) - Create a bank account
* [Get](#get) - Retrieve a bank account
* [Delete](#delete) - Delete a bank account
* [Update](#update) - Update a bank account

## List

Retrieve a list of bank accounts for the authenticated team.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.BankAccounts.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `enabled`                                                | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `manual`                                                 | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListBankAccountsResponse](../../models/operations/listbankaccountsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new bank account for the authenticated team.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.BankAccounts.Create(ctx, &operations.CreateBankAccountRequest{
        Name: "Checking Account",
        Currency: middaygo.String("USD"),
        Manual: middaygo.Bool(false),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateBankAccountRequest](../../models/operations/createbankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateBankAccountResponse](../../models/operations/createbankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a bank account by ID for the authenticated team.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.BankAccounts.Get(ctx, "b7e6c2a0-1f2d-4c3b-9a8e-123456789abc")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b7e6c2a0-1f2d-4c3b-9a8e-123456789abc                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBankAccountByIDResponse](../../models/operations/getbankaccountbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a bank account by ID for the authenticated team.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.BankAccounts.Delete(ctx, "b7e6c2a0-1f2d-4c3b-9a8e-123456789abc")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b7e6c2a0-1f2d-4c3b-9a8e-123456789abc                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteBankAccountResponse](../../models/operations/deletebankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a bank account by ID for the authenticated team.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.BankAccounts.Update(ctx, "b7e6c2a0-1f2d-4c3b-9a8e-123456789abc", &operations.UpdateBankAccountRequestBody{
        ID: middaygo.String("b7e6c2a0-1f2d-4c3b-9a8e-123456789abc"),
        Name: middaygo.String("Checking Account"),
        Enabled: middaygo.Bool(true),
        Balance: middaygo.Float64(1500.75),
        Type: operations.UpdateBankAccountTypeDepository.ToPointer(),
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

| Parameter                                                                                                                               | Type                                                                                                                                    | Required                                                                                                                                | Description                                                                                                                             | Example                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                   | :heavy_check_mark:                                                                                                                      | The context to use for the request.                                                                                                     |                                                                                                                                         |
| `id`                                                                                                                                    | *string*                                                                                                                                | :heavy_check_mark:                                                                                                                      | N/A                                                                                                                                     | b7e6c2a0-1f2d-4c3b-9a8e-123456789abc                                                                                                    |
| `requestBody`                                                                                                                           | [*operations.UpdateBankAccountRequestBody](../../models/operations/updatebankaccountrequestbody.md)                                     | :heavy_minus_sign:                                                                                                                      | N/A                                                                                                                                     | {<br/>"id": "b7e6c2a0-1f2d-4c3b-9a8e-123456789abc",<br/>"name": "Checking Account",<br/>"enabled": true,<br/>"balance": 1500.75,<br/>"type": "depository"<br/>} |
| `opts`                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                | :heavy_minus_sign:                                                                                                                      | The options for this request.                                                                                                           |                                                                                                                                         |

### Response

**[*operations.UpdateBankAccountResponse](../../models/operations/updatebankaccountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |