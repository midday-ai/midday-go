# Transactions
(*Transactions*)

## Overview

Manage transactions

### Available Operations

* [List](#list) - List all transactions
* [Create](#create) - Create a transaction
* [Get](#get) - Retrieve a transaction
* [Delete](#delete) - Delete a transaction
* [Update](#update) - Update a transaction
* [GetAttachmentPreSignedURL](#getattachmentpresignedurl) - Generate pre-signed URL for transaction attachment
* [CreateMany](#createmany) - Bulk create transactions
* [DeleteMany](#deletemany) - Bulk delete transactions
* [UpdateMany](#updatemany) - Bulk update transactions

## List

Retrieve a list of transactions for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTransactions" method="get" path="/transactions" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
        Cursor: middaygo.String("eyJpZCI6IjEyMyJ9"),
        Sort: []string{
            "date",
            "desc",
        },
        PageSize: middaygo.Float64(50),
        Q: middaygo.String("office supplies"),
        Categories: []string{
            "office-supplies",
            "travel",
        },
        Tags: []string{
            "tag-1",
            "tag-2",
        },
        Start: middaygo.String("2024-04-01T00:00:00.000Z"),
        End: middaygo.String("2024-04-30T23:59:59.999Z"),
        Accounts: []string{
            "account-1",
            "account-2",
        },
        Assignees: []string{
            "user-1",
            "user-2",
        },
        Statuses: []operations.ListTransactionsStatus{
            operations.ListTransactionsStatusInReview,
            operations.ListTransactionsStatusExportError,
        },
        Recurring: []string{
            "monthly",
            "annually",
        },
        Attachments: operations.AttachmentsInclude.ToPointer(),
        AmountRange: []*float64{
            middaygo.Float64(100),
            middaygo.Float64(1000),
        },
        Amount: []string{
            "150.75",
            "299.99",
        },
        Type: operations.ListTransactionsTypeExpense.ToPointer(),
        Manual: operations.ManualInclude.ToPointer(),
        Exported: middaygo.Bool(false),
        Fulfilled: middaygo.Bool(true),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListTransactionsRequest](../../models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListTransactionsResponse](../../models/operations/listtransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a transaction

### Example Usage

<!-- UsageSnippet language="go" operationID="createTransaction" method="post" path="/transactions" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.Create(ctx, operations.CreateTransactionRequest{
        Name: "<value>",
        Amount: 5744.12,
        Currency: "Forint",
        Date: "2024-01-12",
        BankAccountID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateTransactionRequest](../../models/operations/createtransactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateTransactionResponse](../../models/operations/createtransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a transaction by its ID for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTransactionById" method="get" path="/transactions/{id}" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.Get(ctx, "391723c9-de99-4039-b8e2-4fa5bbdf9480")
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Transaction ID (UUID).                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTransactionByIDResponse](../../models/operations/gettransactionbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a transaction for the authenticated team. Only manually created transactions can be deleted via this endpoint or the form. Transactions inserted by bank connections cannot be deleted, but can be excluded by updating the status.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTransaction" method="delete" path="/transactions/{id}" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.Delete(ctx, "92766ee2-a2bc-44aa-97af-6891695fc321")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Transaction ID (UUID).                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTransactionResponse](../../models/operations/deletetransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a transaction for the authenticated team. If there's no change, returns it as it is.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTransaction" method="patch" path="/transactions/{id}" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.Update(ctx, "f0c1d0ef-5679-4c1b-9698-2c64e97e8c1d", operations.UpdateTransactionRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | *string*                                                                                           | :heavy_check_mark:                                                                                 | Transaction ID (UUID).                                                                             |
| `requestBody`                                                                                      | [operations.UpdateTransactionRequestBody](../../models/operations/updatetransactionrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateTransactionResponse](../../models/operations/updatetransactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAttachmentPreSignedURL

Generate a pre-signed URL for accessing a transaction attachment. The URL is valid for 60 seconds and allows secure temporary access to the attachment file.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTransactionAttachmentPreSignedUrl" method="post" path="/transactions/{transactionId}/attachments/{attachmentId}/presigned-url" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.GetAttachmentPreSignedURL(ctx, "b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4", "a43dc3a5-6925-4d91-ac9c-4c1a34bdb388", middaygo.Bool(true))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           | Example                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |                                                                                                                                                       |
| `transactionID`                                                                                                                                       | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | Unique identifier of the transaction                                                                                                                  | b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4                                                                                                                  |
| `attachmentID`                                                                                                                                        | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | Unique identifier of the attachment to generate a pre-signed URL for                                                                                  | a43dc3a5-6925-4d91-ac9c-4c1a34bdb388                                                                                                                  |
| `download`                                                                                                                                            | **bool*                                                                                                                                               | :heavy_minus_sign:                                                                                                                                    | Whether to force download the file. If true, the file will be downloaded. If false or omitted, the file will be displayed in the browser if possible. | true                                                                                                                                                  |
| `opts`                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                              | :heavy_minus_sign:                                                                                                                                    | The options for this request.                                                                                                                         |                                                                                                                                                       |

### Response

**[*operations.GetTransactionAttachmentPreSignedURLResponse](../../models/operations/gettransactionattachmentpresignedurlresponse.md), error**

### Errors

| Error Type                                                        | Status Code                                                       | Content Type                                                      |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| apierrors.GetTransactionAttachmentPreSignedURLBadRequestError     | 400                                                               | application/json                                                  |
| apierrors.GetTransactionAttachmentPreSignedURLNotFoundError       | 404                                                               | application/json                                                  |
| apierrors.GetTransactionAttachmentPreSignedURLInternalServerError | 500                                                               | application/json                                                  |
| apierrors.APIError                                                | 4XX, 5XX                                                          | \*/\*                                                             |

## CreateMany

Bulk create transactions for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTransactions" method="post" path="/transactions/bulk" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.CreateMany(ctx, []operations.RequestBody{
        operations.RequestBody{
            Name: "<value>",
            Amount: 5142.41,
            Currency: "Gourde",
            Date: "2024-03-22",
            BankAccountID: "<id>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]operations.RequestBody](../../.md)                    | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateTransactionsResponse](../../models/operations/createtransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteMany

Bulk delete transactions for the authenticated team. Only manually created transactions can be deleted via this endpoint or the form. Transactions inserted by bank connections cannot be deleted, but can be excluded by updating the status.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTransactions" method="delete" path="/transactions/bulk" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.DeleteMany(ctx, []string{
        "c0db9ee1-75c5-4621-84a2-0c38d2dc3106",
        "e5581754-1917-44fa-a324-166437019d98",
        "a8bccdfb-07ed-4359-8712-46b82316d8f9",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]string](../../.md)                                    | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteTransactionsResponse](../../models/operations/deletetransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMany

Bulk update transactions for the authenticated team. If there's no change, returns it as it is.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTransactions" method="patch" path="/transactions/bulk" -->
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
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Transactions.UpdateMany(ctx, operations.UpdateTransactionsRequest{
        Ids: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateTransactionsRequest](../../models/operations/updatetransactionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateTransactionsResponse](../../models/operations/updatetransactionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |