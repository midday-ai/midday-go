# Notifications
(*Notifications*)

## Overview

Manage notifications

### Available Operations

* [List](#list) - List all notifications
* [UpdateStatus](#updatestatus) - Update notification status
* [UpdateAllStatus](#updateallstatus) - Update status of all notifications

## List

Retrieve a list of notifications for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="listNotifications" method="get" path="/notifications" -->
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

    res, err := s.Notifications.List(ctx, operations.ListNotificationsRequest{
        Cursor: middaygo.String("20"),
        PageSize: middaygo.Float64(20),
        Status: middaygo.Pointer(operations.CreateStatusArrayOfListNotificationsStatusEnum2(
            []operations.ListNotificationsStatusEnum2{
                operations.ListNotificationsStatusEnum2Unread,
                operations.ListNotificationsStatusEnum2Read,
            },
        )),
        UserID: middaygo.String("a1b2c3d4-e5f6-7890-abcd-ef1234567890"),
        Priority: middaygo.Int64(5),
        MaxPriority: middaygo.Int64(3),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NotificationsResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListNotificationsRequest](../../models/operations/listnotificationsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListNotificationsResponse](../../models/operations/listnotificationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateStatus

Update the status of a specific notification.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateNotificationStatus" method="patch" path="/notifications/{notificationId}/status" -->
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

    res, err := s.Notifications.UpdateStatus(ctx, "b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2", operations.UpdateNotificationStatusRequestBody{
        Status: operations.UpdateNotificationStatusStatusRead,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NotificationResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |                                                                                                                  |
| `notificationID`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The ID of the notification to update                                                                             | b3b6e2c2-1f2a-4e3b-9c1d-2a4b6e2c21f2                                                                             |
| `requestBody`                                                                                                    | [operations.UpdateNotificationStatusRequestBody](../../models/operations/updatenotificationstatusrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |                                                                                                                  |

### Response

**[*operations.UpdateNotificationStatusResponse](../../models/operations/updatenotificationstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAllStatus

Update the status of all notifications for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAllNotificationsStatus" method="post" path="/notifications/update-all-status" -->
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

    res, err := s.Notifications.UpdateAllStatus(ctx, components.UpdateAllNotificationsStatusSchema{
        Status: components.UpdateAllNotificationsStatusSchemaStatusRead,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAllNotificationsStatusResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [components.UpdateAllNotificationsStatusSchema](../../models/components/updateallnotificationsstatusschema.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateAllNotificationsStatusResponse](../../models/operations/updateallnotificationsstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |