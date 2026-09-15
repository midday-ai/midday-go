# InvoicePayments
(*InvoicePayments*)

## Overview

Invoice payment processing

### Available Operations

* [GetStripeConnectURL](#getstripeconnecturl) - Get Stripe Connect URL
* [StripeConnectCallback](#stripeconnectcallback) - Stripe Connect OAuth callback
* [DisconnectStripe](#disconnectstripe) - Disconnect Stripe account
* [CreateInvoicePaymentIntent](#createinvoicepaymentintent) - Create payment intent for invoice
* [GetStripeConnectStatus](#getstripeconnectstatus) - Get Stripe Connect status

## GetStripeConnectURL

Generates OAuth URL for Stripe Connect Standard integration. Allows teams to connect their Stripe account for accepting invoice payments.

### Example Usage

<!-- UsageSnippet language="go" operationID="getStripeConnectUrl" method="get" path="/invoice-payments/connect-stripe" -->
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

    res, err := s.InvoicePayments.GetStripeConnectURL(ctx)
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

**[*operations.GetStripeConnectURLResponse](../../models/operations/getstripeconnecturlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## StripeConnectCallback

Handles OAuth callback from Stripe Connect after user authorization. Exchanges authorization code for connected account ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="stripeConnectCallback" method="get" path="/invoice-payments/connect-stripe/callback" -->
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

    res, err := s.InvoicePayments.StripeConnectCallback(ctx, "California", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth state parameter for CSRF protection                |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from Stripe                     |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `errorDescription`                                       | **string*                                                | :heavy_minus_sign:                                       | OAuth error description                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.StripeConnectCallbackResponse](../../models/operations/stripeconnectcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DisconnectStripe

Disconnects the team's Stripe Connect account.

### Example Usage

<!-- UsageSnippet language="go" operationID="disconnectStripe" method="post" path="/invoice-payments/disconnect-stripe" -->
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

    res, err := s.InvoicePayments.DisconnectStripe(ctx)
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

**[*operations.DisconnectStripeResponse](../../models/operations/disconnectstriperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateInvoicePaymentIntent

Creates a Stripe PaymentIntent for paying an invoice. This is a public endpoint that uses the invoice token for authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInvoicePaymentIntent" method="post" path="/invoice-payments/payment-intent" -->
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

    res, err := s.InvoicePayments.CreateInvoicePaymentIntent(ctx, operations.CreateInvoicePaymentIntentRequest{
        Token: "<value>",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.CreateInvoicePaymentIntentRequest](../../models/operations/createinvoicepaymentintentrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.CreateInvoicePaymentIntentResponse](../../models/operations/createinvoicepaymentintentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetStripeConnectStatus

Gets the current Stripe Connect status for the team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getStripeConnectStatus" method="get" path="/invoice-payments/stripe-status" -->
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

    res, err := s.InvoicePayments.GetStripeConnectStatus(ctx)
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

**[*operations.GetStripeConnectStatusResponse](../../models/operations/getstripeconnectstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |