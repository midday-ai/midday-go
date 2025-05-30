# TrackerProjects
(*TrackerProjects*)

## Overview

### Available Operations

* [List](#list) - List all tracker projects
* [Create](#create) - Create a tracker project
* [Get](#get) - Retrieve a tracker project
* [Delete](#delete) - Delete a tracker project
* [Update](#update) - Update a tracker project

## List

List all tracker projects for the authenticated team.

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

    res, err := s.TrackerProjects.List(ctx, operations.ListTrackerProjectsRequest{
        Cursor: middaygo.String("eyJpZCI6IjEyMyJ9"),
        PageSize: middaygo.Float64(20),
        Q: middaygo.String("website"),
        Start: middaygo.String("2024-04-01"),
        End: middaygo.String("2024-04-30"),
        Status: operations.ListTrackerProjectsStatusInProgress.ToPointer(),
        Customers: []string{
            "customer-1",
            "customer-2",
        },
        Tags: []string{
            "tag-1",
            "tag-2",
        },
        Sort: []string{
            "-createdAt",
            "name",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrackerProjectsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListTrackerProjectsRequest](../../models/operations/listtrackerprojectsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListTrackerProjectsResponse](../../models/operations/listtrackerprojectsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a tracker project for the authenticated team.

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

    res, err := s.TrackerProjects.Create(ctx, &operations.CreateTrackerProjectRequest{
        Name: "New Project",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrackerProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateTrackerProjectRequest](../../models/operations/createtrackerprojectrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.CreateTrackerProjectResponse](../../models/operations/createtrackerprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a tracker project for the authenticated team.

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

    res, err := s.TrackerProjects.Get(ctx, "b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f")
    if err != nil {
        log.Fatal(err)
    }
    if res.TrackerProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTrackerProjectByIDResponse](../../models/operations/gettrackerprojectbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a tracker project for the authenticated team.

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

    res, err := s.TrackerProjects.Delete(ctx, "b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteTrackerProjectResponse](../../models/operations/deletetrackerprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a tracker project for the authenticated team.

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

    res, err := s.TrackerProjects.Update(ctx, "b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f", &operations.UpdateTrackerProjectRequestBody{
        Name: "Website Redesign",
        Description: middaygo.String("Complete redesign of the company website with modern UI/UX and improved performance"),
        Estimate: middaygo.Float64(120),
        Rate: middaygo.Float64(75),
        Currency: middaygo.String("USD"),
        Status: operations.UpdateTrackerProjectStatusInProgress.ToPointer(),
        CustomerID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-1234567890ef"),
        Tags: []operations.UpdateTrackerProjectTag{
            operations.UpdateTrackerProjectTag{
                ID: "f1e2d3c4-b5a6-7890-1234-567890abcdef",
                Value: "Design",
            },
            operations.UpdateTrackerProjectTag{
                ID: "e2d3c4b5-a6f1-7890-1234-567890abcdef",
                Value: "Frontend",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TrackerProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `id`                                                                                                      | *string*                                                                                                  | :heavy_check_mark:                                                                                        | N/A                                                                                                       | b7e6c8e2-1f2a-4c3b-9e2d-1a2b3c4d5e6f                                                                      |
| `requestBody`                                                                                             | [*operations.UpdateTrackerProjectRequestBody](../../models/operations/updatetrackerprojectrequestbody.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |                                                                                                           |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*operations.UpdateTrackerProjectResponse](../../models/operations/updatetrackerprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |