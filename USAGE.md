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
			Oauth2: middaygo.String(os.Getenv("MIDDAY_OAUTH2")),
		}),
	)

	res, err := s.OAuth.PostOAuthRegister(ctx, operations.PostOAuthRegisterRequest{
		ClientName: "ChatGPT",
		RedirectUris: []string{
			"https://chatgpt.com/connector/oauth/callback",
		},
		GrantTypes: []string{
			"authorization_code",
			"refresh_token",
		},
		Scope:     middaygo.String("transactions.read invoices.read"),
		LogoURI:   middaygo.String("https://example.com/logo.png"),
		ClientURI: middaygo.String("https://example.com"),
		ResponseTypes: []string{
			"code",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.TwoHundredApplicationJSONObject != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->