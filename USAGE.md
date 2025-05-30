<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
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

	res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
		Cursor: middaygo.String("eyJpZCI6IjEyMyJ9"),
		Sort: []string{
			"date",
			"desc",
		},
		PageSize: middaygo.Float64(50),
		Q:        middaygo.String("office supplies"),
		Categories: []string{
			"office-supplies",
			"travel",
		},
		Tags: []string{
			"tag-1",
			"tag-2",
		},
		Start: middaygo.String("2024-04-01T00:00:00.000Z"),
		End:   middaygo.String("2024-04-30T23:59:59.999Z"),
		Accounts: []string{
			"account-1",
			"account-2",
		},
		Assignees: []string{
			"user-1",
			"user-2",
		},
		Statuses: []string{
			"pending",
			"completed",
		},
		Recurring: []string{
			"monthly",
			"annually",
		},
		Attachments: operations.AttachmentsInclude.ToPointer(),
		AmountRange: []*float64{
			middaygo.Float64(100),
			middaygo.Float64(1000),
		},
		Amount: []string{
			"150.75",
			"299.99",
		},
		Type: operations.ListTransactionsTypeExpense.ToPointer(),
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