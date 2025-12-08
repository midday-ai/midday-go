# Invoices
(*Invoices*)

## Overview

### Available Operations

* [List](#list) - List all invoices
* [Create](#create) - Create an invoice
* [GetInvoicesPaymentStatus](#getinvoicespaymentstatus) - Payment status
* [Summary](#summary) - Invoice summary
* [Get](#get) - Retrieve a invoice
* [Update](#update) - Update an invoice
* [Delete](#delete) - Delete a invoice

## List

Retrieve a list of invoices for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="listInvoices" method="get" path="/invoices" -->
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

    res, err := s.Invoices.List(ctx, operations.ListInvoicesRequest{
        Cursor: middaygo.String("25"),
        Sort: []string{
            "createdAt",
            "desc",
        },
        PageSize: middaygo.Float64(25),
        Q: middaygo.String("Acme"),
        Start: middaygo.String("2024-01-01"),
        End: middaygo.String("2024-01-31"),
        Statuses: []string{
            "paid",
            "unpaid",
        },
        Customers: []string{
            "customer-uuid-1",
            "customer-uuid-2",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListInvoicesRequest](../../models/operations/listinvoicesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListInvoicesResponse](../../models/operations/listinvoicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create an invoice for the authenticated team. The behavior depends on deliveryType: 'create' generates and finalizes the invoice immediately, 'create_and_send' also sends it to the customer, 'scheduled' schedules the invoice for automatic processing at the specified date.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInvoice" method="post" path="/invoices" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"github.com/midday-ai/midday-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Invoices.Create(ctx, &operations.CreateInvoiceRequest{
        Template: operations.Template{
            CustomerLabel: middaygo.String("Bill To"),
            Title: middaygo.String("Invoice"),
            FromLabel: middaygo.String("From"),
            InvoiceNoLabel: middaygo.String("Invoice #"),
            IssueDateLabel: middaygo.String("Issue Date"),
            DueDateLabel: middaygo.String("Due Date"),
            DescriptionLabel: middaygo.String("Description"),
            PriceLabel: middaygo.String("Rate"),
            QuantityLabel: middaygo.String("Qty"),
            TotalLabel: middaygo.String("Amount"),
            TotalSummaryLabel: middaygo.String("Total"),
            VatLabel: middaygo.String("VAT"),
            TaxLabel: middaygo.String("Sales Tax"),
            DiscountLabel: middaygo.String("Discount"),
            Timezone: middaygo.String("America/Los_Angeles"),
            PaymentLabel: middaygo.String("Payment Information"),
            NoteLabel: middaygo.String("Notes"),
            LogoURL: middaygo.String("https://example.com/logo.png"),
            Currency: middaygo.String("USD"),
            DateFormat: middaygo.String("MM/dd/yyyy"),
            IncludeVat: middaygo.Bool(false),
            IncludeTax: middaygo.Bool(true),
            IncludeDiscount: middaygo.Bool(false),
            IncludeDecimals: middaygo.Bool(true),
            IncludePdf: middaygo.Bool(true),
            SendCopy: middaygo.Bool(true),
            IncludeUnits: middaygo.Bool(true),
            IncludeQr: middaygo.Bool(false),
            TaxRate: middaygo.Float64(8.5),
            VatRate: middaygo.Float64(0),
            Size: operations.SizeLetter.ToPointer(),
            DeliveryType: operations.TemplateDeliveryTypeCreate.ToPointer(),
            Locale: middaygo.String("en-US"),
            PaymentDetails: &operations.TemplatePaymentDetails{},
            FromDetails: &operations.TemplateFromDetails{},
        },
        FromDetails: &operations.FromDetails{},
        CustomerID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
        PaymentDetails: &operations.PaymentDetails{},
        NoteDetails: &operations.NoteDetails{},
        DueDate: "2024-07-15T23:59:59.000Z",
        IssueDate: "2024-06-15T00:00:00.000Z",
        InvoiceNumber: middaygo.String("INV-2024-001"),
        LogoURL: middaygo.String("https://example.com/logo.png"),
        Tax: middaygo.Float64(85),
        TopBlock: &operations.TopBlock{},
        BottomBlock: &operations.BottomBlock{},
        Amount: middaygo.Float64(1085),
        LineItems: []operations.LineItem{
            operations.LineItem{
                Quantity: middaygo.Float64(40),
                Price: middaygo.Float64(75),
                Tax: middaygo.Float64(8.5),
                Name: &operations.Name{},
            },
            operations.LineItem{
                Quantity: middaygo.Float64(20),
                Price: middaygo.Float64(50),
                Tax: middaygo.Float64(8.5),
                Name: &operations.Name{},
            },
        },
        DeliveryType: operations.DeliveryTypeCreate,
        ScheduledAt: types.MustNewTimeFromString("2024-07-01T09:00:00.000Z"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateInvoiceRequest](../../models/operations/createinvoicerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateInvoiceResponse](../../models/operations/createinvoiceresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.CreateInvoiceBadRequestError     | 400                                        | application/json                           |
| apierrors.CreateInvoiceNotFoundError       | 404                                        | application/json                           |
| apierrors.ConflictError                    | 409                                        | application/json                           |
| apierrors.CreateInvoiceInternalServerError | 500                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## GetInvoicesPaymentStatus

Get payment status for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_/invoices/payment-status" method="get" path="/invoices/payment-status" -->
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

    res, err := s.Invoices.GetInvoicesPaymentStatus(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInvoicesPaymentStatusResponse](../../models/operations/getinvoicespaymentstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Summary

Get summary of invoices for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInvoiceSummary" method="get" path="/invoices/summary" -->
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

    res, err := s.Invoices.Summary(ctx, operations.GetInvoiceSummaryStatusPaid.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |                                                                                           |
| `status`                                                                                  | [*operations.GetInvoiceSummaryStatus](../../models/operations/getinvoicesummarystatus.md) | :heavy_minus_sign:                                                                        | Filter summary by invoice status                                                          | paid                                                                                      |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |                                                                                           |

### Response

**[*operations.GetInvoiceSummaryResponse](../../models/operations/getinvoicesummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a invoice by its unique identifier for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInvoiceById" method="get" path="/invoices/{id}" -->
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

    res, err := s.Invoices.Get(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInvoiceByIDResponse](../../models/operations/getinvoicebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an invoice by its unique identifier for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInvoice" method="put" path="/invoices/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/midday-ai/midday-go/models/components"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"github.com/midday-ai/midday-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity(components.Security{
            Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Invoices.Update(ctx, "<id>", &operations.UpdateInvoiceRequestBody{
        Status: operations.UpdateInvoiceStatusRequestPaid.ToPointer(),
        PaidAt: types.MustNewTimeFromString("2024-06-15T12:00:00.000Z"),
        InternalNote: middaygo.String("Payment received via bank transfer"),
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `requestBody`                                                                               | [*operations.UpdateInvoiceRequestBody](../../models/operations/updateinvoicerequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.UpdateInvoiceResponse](../../models/operations/updateinvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an invoice by its unique identifier for the authenticated team. Only invoices with status 'draft' or 'canceled' can be deleted directly. If the invoice is not in one of these statuses, update its status to 'canceled' before attempting deletion.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteInvoice" method="delete" path="/invoices/{id}" -->
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

    res, err := s.Invoices.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInvoiceResponse](../../models/operations/deleteinvoiceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |