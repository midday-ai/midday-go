# Desktop
(*Desktop*)

## Overview

Desktop app endpoints

### Available Operations

* [CheckUpdate](#checkupdate) - Check for desktop app updates
* [DownloadUpdate](#downloadupdate) - Download desktop app update artifact

## CheckUpdate

Returns the latest desktop app version info in Tauri updater format. Download URLs are rewritten to proxy through this API.

### Example Usage

<!-- UsageSnippet language="go" operationID="checkDesktopUpdate" method="get" path="/desktop/update" -->
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

    res, err := s.Desktop.CheckUpdate(ctx)
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

**[*operations.CheckDesktopUpdateResponse](../../models/operations/checkdesktopupdateresponse.md), error**

### Errors

| Error Type                                  | Status Code                                 | Content Type                                |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| apierrors.CheckDesktopUpdateBadGatewayError | 502                                         | application/json                            |
| apierrors.APIError                          | 4XX, 5XX                                    | \*/\*                                       |

## DownloadUpdate

Proxies the download of a desktop app update artifact from GitHub releases. Only URLs pointing to the midday-ai/midday repository are accepted.

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadDesktopUpdate" method="get" path="/desktop/update/download" -->
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

    res, err := s.Desktop.DownloadUpdate(ctx, "https://github.com/midday-ai/midday/releases/download/midday-v1.0.0/Midday.app.tar.gz")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `url_`                                                                                | *string*                                                                              | :heavy_check_mark:                                                                    | The artifact download URL to proxy                                                    | https://github.com/midday-ai/midday/releases/download/midday-v1.0.0/Midday.app.tar.gz |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.DownloadDesktopUpdateResponse](../../models/operations/downloaddesktopupdateresponse.md), error**

### Errors

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.DownloadDesktopUpdateBadRequestError | 400                                            | application/json                               |
| apierrors.DownloadDesktopUpdateBadGatewayError | 502                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |