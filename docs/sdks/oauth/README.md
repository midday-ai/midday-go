# OAuth
(*OAuth*)

## Overview

### Available Operations

* [GetOAuthAuthorization](#getoauthauthorization) - OAuth Authorization Endpoint
* [PostOAuthAuthorization](#postoauthauthorization) - OAuth Authorization Decision
* [PostOAuthToken](#postoauthtoken) - OAuth Token Exchange
* [PostOAuthRevoke](#postoauthrevoke) - OAuth Token Revocation

## GetOAuthAuthorization

Initiate OAuth authorization flow and get consent screen information

### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthAuthorization" method="get" path="/oauth/authorize" -->
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

    res, err := s.OAuth.GetOAuthAuthorization(ctx, operations.GetOAuthAuthorizationRequest{
        ResponseType: operations.ResponseTypeCode,
        ClientID: "mid_client_abcdef123456789",
        RedirectURI: "https://myapp.com/callback",
        Scope: "transactions.read invoices.read",
        State: "abc123xyz789_secure-random-state-value-with-sufficient-entropy",
        CodeChallenge: middaygo.String("E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetOAuthAuthorizationRequest](../../models/operations/getoauthauthorizationrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetOAuthAuthorizationResponse](../../models/operations/getoauthauthorizationresponse.md), error**

### Errors

| Error Type                                     | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| apierrors.GetOAuthAuthorizationBadRequestError | 400                                            | application/json                               |
| apierrors.APIError                             | 4XX, 5XX                                       | \*/\*                                          |

## PostOAuthAuthorization

Process user's authorization decision (allow/deny)

### Example Usage

<!-- UsageSnippet language="go" operationID="postOAuthAuthorization" method="post" path="/oauth/authorize" -->
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

    res, err := s.OAuth.PostOAuthAuthorization(ctx, &operations.PostOAuthAuthorizationRequest{
        ClientID: "mid_client_abcdef123456789",
        Decision: operations.DecisionAllow,
        Scopes: []operations.Scope{
            operations.ScopeTransactionsRead,
            operations.ScopeInvoicesRead,
        },
        RedirectURI: "https://myapp.com/callback",
        State: "abc123xyz789_secure-random-state-value-with-sufficient-entropy",
        CodeChallenge: middaygo.String("E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"),
        TeamID: "123e4567-e89b-12d3-a456-426614174000",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PostOAuthAuthorizationRequest](../../models/operations/postoauthauthorizationrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostOAuthAuthorizationResponse](../../models/operations/postoauthauthorizationresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.PostOAuthAuthorizationBadRequestError | 400                                             | application/json                                |
| apierrors.UnauthorizedError                     | 401                                             | application/json                                |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## PostOAuthToken

Exchange authorization code for access token or refresh an access token

### Example Usage

<!-- UsageSnippet language="go" operationID="postOAuthToken" method="post" path="/oauth/token" -->
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

    res, err := s.OAuth.PostOAuthToken(ctx, middaygo.Pointer(operations.CreatePostOAuthTokenRequestRefreshToken(
        operations.RefreshToken{
            GrantType: operations.GrantTypeRefreshTokenRefreshToken,
            RefreshToken: "mid_rt_abcdef123456789",
            ClientID: "mid_client_abcdef123456789",
            ClientSecret: middaygo.String("mid_secret_abcdef123456789"),
            Scope: middaygo.String("transactions.read invoices.read"),
        },
    )))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostOAuthTokenRequest](../../models/operations/postoauthtokenrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.PostOAuthTokenResponse](../../models/operations/postoauthtokenresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.PostOAuthTokenBadRequestError | 400                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |

## PostOAuthRevoke

Revoke an access token or refresh token

### Example Usage

<!-- UsageSnippet language="go" operationID="postOAuthRevoke" method="post" path="/oauth/revoke" -->
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

    res, err := s.OAuth.PostOAuthRevoke(ctx, &operations.PostOAuthRevokeRequest{
        Token: "mid_access_token_abcdef123456789",
        TokenTypeHint: operations.TokenTypeHintAccessToken.ToPointer(),
        ClientID: "mid_client_abcdef123456789",
        ClientSecret: middaygo.String("mid_secret_abcdef123456789"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PostOAuthRevokeRequest](../../models/operations/postoauthrevokerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PostOAuthRevokeResponse](../../models/operations/postoauthrevokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |