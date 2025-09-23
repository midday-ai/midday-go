<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/components"
	"github.com/midday-ai/midday-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := middaygo.New(
		middaygo.WithSecurity(components.Security{
			Oauth2: middaygo.Pointer(os.Getenv("MIDDAY_OAUTH2")),
		}),
	)

	res, err := s.OAuth.GetOAuthAuthorization(ctx, operations.GetOAuthAuthorizationRequest{
		ResponseType:  operations.ResponseTypeCode,
		ClientID:      "mid_client_abcdef123456789",
		RedirectURI:   "https://myapp.com/callback",
		Scope:         "transactions.read invoices.read",
		State:         "abc123xyz789_secure-random-state-value-with-sufficient-entropy",
		CodeChallenge: middaygo.Pointer("E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->