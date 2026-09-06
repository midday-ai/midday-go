# Integrations
(*Integrations*)

## Overview

Integration endpoints

### Available Operations

* [SlackOAuthCallback](#slackoauthcallback) - Slack OAuth callback
* [GetSlackInstallURL](#getslackinstallurl) - Get Slack install URL
* [SlackWebhook](#slackwebhook) - Slack webhook handler
* [SlackInteractions](#slackinteractions) - Slack interactions handler
* [GmailOAuthCallback](#gmailoauthcallback) - Gmail OAuth callback
* [GetGmailInstallURL](#getgmailinstallurl) - Get Gmail install URL
* [OutlookOAuthCallback](#outlookoauthcallback) - Outlook OAuth callback
* [GetOutlookInstallURL](#getoutlookinstallurl) - Get Outlook install URL
* [XeroOAuthCallback](#xerooauthcallback) - Xero OAuth callback
* [GetXeroInstallURL](#getxeroinstallurl) - Get Xero install URL
* [QuickBooksOAuthCallback](#quickbooksoauthcallback) - QuickBooks OAuth callback
* [GetQuickBooksInstallURL](#getquickbooksinstallurl) - Get QuickBooks install URL
* [FortnoxOAuthCallback](#fortnoxoauthcallback) - Fortnox OAuth callback
* [GetFortnoxInstallURL](#getfortnoxinstallurl) - Get Fortnox install URL

## SlackOAuthCallback

Handles OAuth callback from Slack after user authorization. Exchanges authorization code for access token and creates app integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="slackOAuthCallback" method="get" path="/apps/slack/oauth-callback" -->
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

    res, err := s.Integrations.SlackOAuthCallback(ctx, "<value>", "North Carolina")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `code`                                                   | *string*                                                 | :heavy_check_mark:                                       | OAuth authorization code from Slack                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth state parameter for CSRF protection                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SlackOAuthCallbackResponse](../../models/operations/slackoauthcallbackresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.SlackOAuthCallbackBadRequestError     | 400                                             | application/json                                |
| apierrors.SlackOAuthCallbackInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## GetSlackInstallURL

Generates OAuth install URL for Slack integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSlackInstallUrl" method="get" path="/apps/slack/install-url" -->
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

    res, err := s.Integrations.GetSlackInstallURL(ctx)
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

**[*operations.GetSlackInstallURLResponse](../../models/operations/getslackinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SlackWebhook

Handles incoming webhook events from Slack. Verifies request signature and processes events.

### Example Usage

<!-- UsageSnippet language="go" operationID="slackWebhook" method="post" path="/apps/slack/webhook" -->
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

    res, err := s.Integrations.SlackWebhook(ctx)
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

**[*operations.SlackWebhookResponse](../../models/operations/slackwebhookresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SlackInteractions

Handles interactive component actions from Slack (button clicks, etc.)

### Example Usage

<!-- UsageSnippet language="go" operationID="slackInteractions" method="post" path="/apps/slack/interactions" -->
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

    res, err := s.Integrations.SlackInteractions(ctx)
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

**[*operations.SlackInteractionsResponse](../../models/operations/slackinteractionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GmailOAuthCallback

Handles OAuth callback from Google after user authorization. Exchanges authorization code for access token and creates inbox account.

### Example Usage

<!-- UsageSnippet language="go" operationID="gmailOAuthCallback" method="get" path="/apps/gmail/oauth-callback" -->
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

    res, err := s.Integrations.GmailOAuthCallback(ctx, "Delaware", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | Encrypted OAuth state parameter                          |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from Google                     |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GmailOAuthCallbackResponse](../../models/operations/gmailoauthcallbackresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.GmailOAuthCallbackBadRequestError     | 400                                             | application/json                                |
| apierrors.GmailOAuthCallbackInternalServerError | 500                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## GetGmailInstallURL

Generates OAuth install URL for Gmail integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getGmailInstallUrl" method="get" path="/apps/gmail/install-url" -->
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

    res, err := s.Integrations.GetGmailInstallURL(ctx)
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

**[*operations.GetGmailInstallURLResponse](../../models/operations/getgmailinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## OutlookOAuthCallback

Handles OAuth callback from Microsoft after user authorization. Exchanges authorization code for access token and creates inbox account.

### Example Usage

<!-- UsageSnippet language="go" operationID="outlookOAuthCallback" method="get" path="/apps/outlook/oauth-callback" -->
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

    res, err := s.Integrations.OutlookOAuthCallback(ctx, "New Hampshire", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | Encrypted OAuth state parameter                          |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from Microsoft                  |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OutlookOAuthCallbackResponse](../../models/operations/outlookoauthcallbackresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.OutlookOAuthCallbackBadRequestError     | 400                                               | application/json                                  |
| apierrors.OutlookOAuthCallbackInternalServerError | 500                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## GetOutlookInstallURL

Generates OAuth install URL for Outlook integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutlookInstallUrl" method="get" path="/apps/outlook/install-url" -->
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

    res, err := s.Integrations.GetOutlookInstallURL(ctx)
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

**[*operations.GetOutlookInstallURLResponse](../../models/operations/getoutlookinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## XeroOAuthCallback

Handles OAuth callback from Xero after user authorization. Exchanges authorization code for access token and creates app integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="xeroOAuthCallback" method="get" path="/apps/xero/oauth-callback" -->
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

    res, err := s.Integrations.XeroOAuthCallback(ctx, "Maryland", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth state parameter for CSRF protection                |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from Xero                       |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.XeroOAuthCallbackResponse](../../models/operations/xerooauthcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetXeroInstallURL

Generates OAuth install URL for Xero integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getXeroInstallUrl" method="get" path="/apps/xero/install-url" -->
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

    res, err := s.Integrations.GetXeroInstallURL(ctx)
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

**[*operations.GetXeroInstallURLResponse](../../models/operations/getxeroinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## QuickBooksOAuthCallback

Handles OAuth callback from QuickBooks after user authorization. Exchanges authorization code for access token and creates app integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="quickBooksOAuthCallback" method="get" path="/apps/quickbooks/oauth-callback" -->
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

    res, err := s.Integrations.QuickBooksOAuthCallback(ctx, "Georgia", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth state parameter for CSRF protection                |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from QuickBooks                 |
| `realmID`                                                | **string*                                                | :heavy_minus_sign:                                       | QuickBooks company/realm ID                              |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.QuickBooksOAuthCallbackResponse](../../models/operations/quickbooksoauthcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetQuickBooksInstallURL

Generates OAuth install URL for QuickBooks integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getQuickBooksInstallUrl" method="get" path="/apps/quickbooks/install-url" -->
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

    res, err := s.Integrations.GetQuickBooksInstallURL(ctx)
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

**[*operations.GetQuickBooksInstallURLResponse](../../models/operations/getquickbooksinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## FortnoxOAuthCallback

Handles OAuth callback from Fortnox after user authorization. Exchanges authorization code for access token and creates app integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="fortnoxOAuthCallback" method="get" path="/apps/fortnox/oauth-callback" -->
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

    res, err := s.Integrations.FortnoxOAuthCallback(ctx, "South Dakota", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `state`                                                  | *string*                                                 | :heavy_check_mark:                                       | OAuth state parameter for CSRF protection                |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | OAuth authorization code from Fortnox                    |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | OAuth error code if authorization failed                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FortnoxOAuthCallbackResponse](../../models/operations/fortnoxoauthcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFortnoxInstallURL

Generates OAuth install URL for Fortnox integration. Requires authentication.

### Example Usage

<!-- UsageSnippet language="go" operationID="getFortnoxInstallUrl" method="get" path="/apps/fortnox/install-url" -->
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

    res, err := s.Integrations.GetFortnoxInstallURL(ctx)
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

**[*operations.GetFortnoxInstallURLResponse](../../models/operations/getfortnoxinstallurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |