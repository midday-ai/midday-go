# TrackerEntries
(*TrackerEntries*)

## Overview

### Available Operations

* [List](#list) - List all tracker entries

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