# Documents
(*Documents*)

## Overview

### Available Operations

* [List](#list) - List all documents
* [Get](#get) - Retrieve a document
* [Delete](#delete) - Delete a document
* [GetPreSignedURL](#getpresignedurl) - Generate pre-signed URL for document

## List

Retrieve a list of documents for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="listDocuments" method="get" path="/documents" -->
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

    res, err := s.Documents.List(ctx, operations.ListDocumentsRequest{
        Cursor: middaygo.Pointer("20"),
        PageSize: middaygo.Pointer[float64](20),
        Q: middaygo.Pointer("invoice"),
        Tags: []string{
            "tag1",
            "tag2",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListDocumentsRequest](../../models/operations/listdocumentsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListDocumentsResponse](../../models/operations/listdocumentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a document by its unique identifier for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDocumentById" method="get" path="/documents/{id}" -->
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

    res, err := s.Documents.Get(ctx, middaygo.Pointer("<id>"))
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDocumentByIDResponse](../../models/operations/getdocumentbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a document by its unique identifier for the authenticated team.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteDocument" method="delete" path="/documents/{id}" -->
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

    res, err := s.Documents.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteDocumentResponse](../../models/operations/deletedocumentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPreSignedURL

Generate a pre-signed URL for accessing a document. The URL is valid for 60 seconds and allows secure temporary access to the document file.

### Example Usage

<!-- UsageSnippet language="go" operationID="getDocumentPreSignedUrl" method="post" path="/documents/{id}/presigned-url" -->
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

    res, err := s.Documents.GetPreSignedURL(ctx, "b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4", middaygo.Pointer(true))
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | b3b7c1e2-4c2a-4e7a-9c1a-2b7c1e24c2a4                     |
| `download`                                               | **bool*                                                  | :heavy_minus_sign:                                       | N/A                                                      | true                                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetDocumentPreSignedURLResponse](../../models/operations/getdocumentpresignedurlresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.GetDocumentPreSignedURLBadRequestError     | 400                                                  | application/json                                     |
| apierrors.GetDocumentPreSignedURLNotFoundError       | 404                                                  | application/json                                     |
| apierrors.GetDocumentPreSignedURLInternalServerError | 500                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |