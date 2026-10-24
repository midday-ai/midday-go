# Files
(*Files*)

## Overview

File operations

### Available Operations

* [Proxy](#proxy) - Proxy file from storage
* [DownloadFile](#downloadfile) - Download file from vault
* [DownloadInvoice](#downloadinvoice) - Download invoice PDF

## Proxy

Proxies a file from storage. Requires team file key (fk) query parameter for access.

### Example Usage

<!-- UsageSnippet language="go" operationID="proxyFile" method="get" path="/files/proxy" -->
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

    res, err := s.Files.Proxy(ctx, "vault/documents/2024/invoice.pdf", "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                              |
| `filePath`                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | Path to the file in storage. Can include or exclude 'vault/' prefix.                                                                                                                                                                                                                         | vault/documents/2024/invoice.pdf                                                                                                                                                                                                                                                             |
| `fk`                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | Team file key for proxy/download access to team files. This key is returned in the user data response (GET /users/me) as the `fileKey` field. It is team-scoped and deterministic - all members of the same team share the same file key. Use this key to authenticate file access requests. | a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6                                                                                                                                                                                                                                                             |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                              |

### Response

**[*operations.ProxyFileResponse](../../models/operations/proxyfileresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ProxyFileBadRequestError     | 400                                    | application/json                       |
| apierrors.ProxyFileNotFoundError       | 404                                    | application/json                       |
| apierrors.ProxyFileInternalServerError | 500                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## DownloadFile

Downloads a file from the vault storage bucket. Requires team file key (fk) query parameter for access.

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadFile" method="get" path="/files/download/file" -->
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

    res, err := s.Files.DownloadFile(ctx, "vault/documents/2024/invoice.pdf", "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6", middaygo.String("invoice.pdf"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                              |
| `path`                                                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | Path to the file in storage. Can include or exclude 'vault/' prefix.                                                                                                                                                                                                                         | vault/documents/2024/invoice.pdf                                                                                                                                                                                                                                                             |
| `fk`                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | Team file key for proxy/download access to team files. This key is returned in the user data response (GET /users/me) as the `fileKey` field. It is team-scoped and deterministic - all members of the same team share the same file key. Use this key to authenticate file access requests. | a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6                                                                                                                                                                                                                                                             |
| `filename`                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | Optional filename for the Content-Disposition header.                                                                                                                                                                                                                                        | invoice.pdf                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |                                                                                                                                                                                                                                                                                              |

### Response

**[*operations.DownloadFileResponse](../../models/operations/downloadfileresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.DownloadFileBadRequestError     | 400                                       | application/json                          |
| apierrors.DownloadFileUnauthorizedError   | 401                                       | application/json                          |
| apierrors.DownloadFileNotFoundError       | 404                                       | application/json                          |
| apierrors.DownloadFileInternalServerError | 500                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## DownloadInvoice

Downloads an invoice as a PDF. Can be accessed with an invoice ID (requires team file key via fk query parameter) or invoice token (public access).

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadInvoice" method="get" path="/files/download/invoice" -->
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

    res, err := s.Files.DownloadInvoice(ctx, operations.DownloadInvoiceRequest{
        ID: middaygo.String("b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4"),
        Fk: middaygo.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."),
        Token: middaygo.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DownloadInvoiceRequest](../../models/operations/downloadinvoicerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.DownloadInvoiceResponse](../../models/operations/downloadinvoiceresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.DownloadInvoiceBadRequestError     | 400                                          | application/json                             |
| apierrors.DownloadInvoiceUnauthorizedError   | 401                                          | application/json                             |
| apierrors.DownloadInvoiceNotFoundError       | 404                                          | application/json                             |
| apierrors.DownloadInvoiceInternalServerError | 500                                          | application/json                             |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |