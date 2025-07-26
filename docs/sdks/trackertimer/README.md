# TrackerTimer
(*TrackerTimer*)

## Overview

### Available Operations

* [StartTimer](#starttimer) - Start a timer
* [StopTimer](#stoptimer) - Stop a timer
* [GetCurrentTimer](#getcurrenttimer) - Get current timer
* [GetTimerStatus](#gettimerstatus) - Get timer status

## StartTimer

Start a new timer or continue from a paused entry.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/types"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.TrackerTimer.StartTimer(ctx, &operations.StartTimerRequest{
        ProjectID: "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2",
        AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
        Description: middaygo.String("Working on implementing timer feature"),
        Start: types.MustNewTimeFromString("2024-04-15T09:00:00.000Z"),
        ContinueFromEntry: middaygo.String("c4d5e6f7-2a3b-4c5d-8e9f-3a4b5c6d7e8f"),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.StartTimerRequest](../../models/operations/starttimerrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.StartTimerResponse](../../models/operations/starttimerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## StopTimer

Stop the current running timer or a specific timer entry.

### Example Usage

```go
package main

import(
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/types"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := middaygo.New(
        middaygo.WithSecurity("MIDDAY_API_KEY"),
    )

    res, err := s.TrackerTimer.StopTimer(ctx, &operations.StopTimerRequest{
        EntryID: middaygo.String("b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2"),
        AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
        Stop: types.MustNewTimeFromString("2024-04-15T17:00:00.000Z"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.StopTimerRequest](../../models/operations/stoptimerrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.StopTimerResponse](../../models/operations/stoptimerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCurrentTimer

Get the currently running timer for the authenticated user.

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

    res, err := s.TrackerTimer.GetCurrentTimer(ctx, middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"))
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
| `assignedID`                                             | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | a1b2c3d4-e5f6-7890-abcd-ef1234567890                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetCurrentTimerResponse](../../models/operations/getcurrenttimerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetTimerStatus

Get timer status including elapsed time for the authenticated user.

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

    res, err := s.TrackerTimer.GetTimerStatus(ctx, middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"))
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
| `assignedID`                                             | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | a1b2c3d4-e5f6-7890-abcd-ef1234567890                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTimerStatusResponse](../../models/operations/gettimerstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |