# Reports
(*Reports*)

## Overview

### Available Operations

* [Revenue](#revenue) - Revenue reports
* [Profit](#profit) - Profit reports
* [BurnRate](#burnrate) - Burn rate reports
* [Runway](#runway) - Runway reports
* [Expenses](#expenses) - Expense reports
* [Spending](#spending) - Spending reports

## Revenue

Revenue reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRevenueReports" method="get" path="/reports/revenue" -->
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

    res, err := s.Reports.Revenue(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRevenueResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetRevenueReportsResponse](../../models/operations/getrevenuereportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Profit

Profit reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getProfitReports" method="get" path="/reports/profit" -->
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

    res, err := s.Reports.Profit(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetProfitResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetProfitReportsResponse](../../models/operations/getprofitreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## BurnRate

Burn rate reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getBurnRateReports" method="get" path="/reports/burn-rate" -->
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

    res, err := s.Reports.BurnRate(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBurnRateResponseSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBurnRateReportsResponse](../../models/operations/getburnratereportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Runway

Runway reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRunwayReports" method="get" path="/reports/runway" -->
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

    res, err := s.Reports.Runway(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRunwayResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetRunwayReportsResponse](../../models/operations/getrunwayreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Expenses

Expense reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getExpensesReports" method="get" path="/reports/expenses" -->
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

    res, err := s.Reports.Expenses(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetExpensesResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetExpensesReportsResponse](../../models/operations/getexpensesreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Spending

Spending reports for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSpendingReports" method="get" path="/reports/spending" -->
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

    res, err := s.Reports.Spending(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SpendingResultArray != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-01-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2023-12-31                                               |
| `currency`                                               | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | USD                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetSpendingReportsResponse](../../models/operations/getspendingreportsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |