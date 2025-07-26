# TrackerEntries
(*TrackerEntries*)

## Overview

### Available Operations

* [List](#list) - List all tracker entries
* [Create](#create) - Create a tracker entry
* [CreateBulk](#createbulk) - Create multiple tracker entries
* [Delete](#delete) - Delete a tracker entry
* [Update](#update) - Update a tracker entry

## List

List all tracker entries for the authenticated team.

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

    res, err := s.TrackerEntries.List(ctx, "2024-04-01", "2024-04-30", middaygo.String("b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2"))
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
| `from`                                                   | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2024-04-01                                               |
| `to`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 2024-04-30                                               |
| `projectID`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      | b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListTrackerEntriesResponse](../../models/operations/listtrackerentriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a tracker entry for the authenticated team.

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

    res, err := s.TrackerEntries.Create(ctx, &operations.CreateTrackerEntryRequest{
        Start: types.MustTimeFromString("2024-04-15T09:00:00.000Z"),
        Stop: types.MustTimeFromString("2024-04-15T17:00:00.000Z"),
        Dates: []string{
            "2024-04-15",
            "2024-04-16",
        },
        AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
        ProjectID: "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2",
        Description: middaygo.String("Worked on implementing user authentication feature"),
        Duration: 28800,
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
| `request`                                                                                    | [operations.CreateTrackerEntryRequest](../../models/operations/createtrackerentryrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateTrackerEntryResponse](../../models/operations/createtrackerentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateBulk

Create multiple tracker entries in a single request for efficient data migration.

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

    res, err := s.TrackerEntries.CreateBulk(ctx, &operations.CreateTrackerEntriesBulkRequest{
        Entries: []operations.Entry{
            operations.Entry{
                Start: types.MustTimeFromString("2024-04-15T09:00:00.000Z"),
                Stop: types.MustTimeFromString("2024-04-15T17:00:00.000Z"),
                Dates: []string{
                    "2024-04-15",
                },
                AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
                ProjectID: "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2",
                Description: middaygo.String("Working on authentication feature"),
                Duration: 28800,
            },
            operations.Entry{
                Start: types.MustTimeFromString("2024-04-16T09:00:00.000Z"),
                Stop: types.MustTimeFromString("2024-04-16T17:00:00.000Z"),
                Dates: []string{
                    "2024-04-16",
                },
                AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
                ProjectID: "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2",
                Description: middaygo.String("Working on dashboard feature"),
                Duration: 28800,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.CreateTrackerEntriesBulkRequest](../../models/operations/createtrackerentriesbulkrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CreateTrackerEntriesBulkResponse](../../models/operations/createtrackerentriesbulkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a tracker entry for the authenticated team.

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

    res, err := s.TrackerEntries.Delete(ctx, "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteTrackerEntryResponse](../../models/operations/deletetrackerentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a tracker entry for the authenticated team.

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

    res, err := s.TrackerEntries.Update(ctx, "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2", &operations.UpdateTrackerEntryRequestBody{
        Start: types.MustTimeFromString("2024-04-15T09:00:00.000Z"),
        Stop: types.MustTimeFromString("2024-04-15T17:00:00.000Z"),
        Dates: []string{
            "2024-04-15",
            "2024-04-16",
        },
        AssignedID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
        ProjectID: "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2",
        Description: middaygo.String("Worked on implementing user authentication feature"),
        Duration: 28800,
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

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `id`                                                                                                  | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   | b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2                                                                  |
| `requestBody`                                                                                         | [*operations.UpdateTrackerEntryRequestBody](../../models/operations/updatetrackerentryrequestbody.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |                                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.UpdateTrackerEntryResponse](../../models/operations/updatetrackerentryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |