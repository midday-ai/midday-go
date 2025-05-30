# Metrics
(*Metrics*)

## Overview

### Available Operations

* [Revenue](#revenue) - Revenue metrics
* [Profit](#profit) - Profit metrics
* [BurnRate](#burnrate) - Burn rate metrics
* [Runway](#runway) - Runway metrics
* [Expenses](#expenses) - Expense metrics
* [Spending](#spending) - Spending metrics

## Revenue

Revenue metrics for the authenticated team.

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

    res, err := s.Metrics.Revenue(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetRevenueMetricsResponse](../../models/operations/getrevenuemetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Profit

Profit metrics for the authenticated team.

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

    res, err := s.Metrics.Profit(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetProfitMetricsResponse](../../models/operations/getprofitmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## BurnRate

Burn rate metrics for the authenticated team.

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

    res, err := s.Metrics.BurnRate(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetBurnRateMetricsResponse](../../models/operations/getburnratemetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Runway

Runway metrics for the authenticated team.

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

    res, err := s.Metrics.Runway(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetRunwayMetricsResponse](../../models/operations/getrunwaymetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Expenses

Expense metrics for the authenticated team.

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

    res, err := s.Metrics.Expenses(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetExpensesMetricsResponse](../../models/operations/getexpensesmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Spending

Spending metrics for the authenticated team.

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

    res, err := s.Metrics.Spending(ctx, "2023-01-01", "2023-12-31", middaygo.String("USD"))
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

**[*operations.GetSpendingMetricsResponse](../../models/operations/getspendingmetricsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |