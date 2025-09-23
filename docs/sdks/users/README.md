# Users
(*Users*)

## Overview

### Available Operations

* [Get](#get) - Retrieve the current user
* [Update](#update) - Update the current user

## Get

Retrieve the current user for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCurrentUser" method="get" path="/users/me" -->
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
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Users.Get(ctx)
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

**[*operations.GetCurrentUserResponse](../../models/operations/getcurrentuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the current user for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCurrentUser" method="patch" path="/users/me" -->
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
            Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
        }),
    )

    res, err := s.Users.Update(ctx, &operations.UpdateCurrentUserRequest{
        FullName: middaygo.Pointer("Jane Doe"),
        TeamID: middaygo.Pointer("team-abc123"),
        Email: middaygo.Pointer("jane.doe@acme.com"),
        AvatarURL: middaygo.Pointer("https://cdn.midday.ai/avatars/jane-doe.jpg"),
        Locale: middaygo.Pointer("en-US"),
        WeekStartsOnMonday: middaygo.Pointer(true),
        Timezone: middaygo.Pointer("America/New_York"),
        TimezoneAutoSync: middaygo.Pointer(true),
        TimeFormat: middaygo.Pointer[float64](24),
        DateFormat: operations.DateFormatRequestYyyyDashMmDashdd.ToPointer(),
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
| `request`                                                                                  | [operations.UpdateCurrentUserRequest](../../models/operations/updatecurrentuserrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateCurrentUserResponse](../../models/operations/updatecurrentuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |