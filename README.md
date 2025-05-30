# github.com/midday-ai/midday-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/midday-ai/midday-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/midday-ai/midday-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/midday/midday). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Midday API: Midday is a platform for Invoicing, Time tracking, File reconciliation, Storage, Financial Overview & your own Assistant.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/midday-ai/midday-go](#githubcommidday-aimidday-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/midday-ai/midday-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type | Scheme      | Environment Variable |
| ------- | ---- | ----------- | -------------------- |
| `Token` | http | HTTP Bearer | `MIDDAY_TOKEN`       |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [List](docs/sdks/bankaccounts/README.md#list) - List all bank accounts
* [Create](docs/sdks/bankaccounts/README.md#create) - Create a bank account
* [Get](docs/sdks/bankaccounts/README.md#get) - Retrieve a bank account
* [Delete](docs/sdks/bankaccounts/README.md#delete) - Delete a bank account
* [Update](docs/sdks/bankaccounts/README.md#update) - Update a bank account

### [Customers](docs/sdks/customers/README.md)

* [List](docs/sdks/customers/README.md#list) - List all customers
* [Create](docs/sdks/customers/README.md#create) - Create customer
* [Get](docs/sdks/customers/README.md#get) - Retrieve a customer
* [Delete](docs/sdks/customers/README.md#delete) - Delete a customer
* [Update](docs/sdks/customers/README.md#update) - Update a customer

### [Documents](docs/sdks/documents/README.md)

* [List](docs/sdks/documents/README.md#list) - List all documents
* [Get](docs/sdks/documents/README.md#get) - Retrieve a document
* [Delete](docs/sdks/documents/README.md#delete) - Delete a document

### [Inbox](docs/sdks/inbox/README.md)

* [List](docs/sdks/inbox/README.md#list) - List all inbox items
* [Get](docs/sdks/inbox/README.md#get) - Retrieve a inbox item
* [Delete](docs/sdks/inbox/README.md#delete) - Delete a inbox item
* [Update](docs/sdks/inbox/README.md#update) - Update a inbox item

### [Invoices](docs/sdks/invoices/README.md)

* [List](docs/sdks/invoices/README.md#list) - List all invoices
* [GetInvoicesPaymentStatus](docs/sdks/invoices/README.md#getinvoicespaymentstatus) - Payment status
* [Summary](docs/sdks/invoices/README.md#summary) - Invoice summary
* [Get](docs/sdks/invoices/README.md#get) - Retrieve a invoice
* [Delete](docs/sdks/invoices/README.md#delete) - Delete a invoice

### [Metrics](docs/sdks/metrics/README.md)

* [Revenue](docs/sdks/metrics/README.md#revenue) - Revenue metrics
* [Profit](docs/sdks/metrics/README.md#profit) - Profit metrics
* [BurnRate](docs/sdks/metrics/README.md#burnrate) - Burn rate metrics
* [Runway](docs/sdks/metrics/README.md#runway) - Runway metrics
* [Expenses](docs/sdks/metrics/README.md#expenses) - Expense metrics
* [Spending](docs/sdks/metrics/README.md#spending) - Spending metrics


### [Search](docs/sdks/search/README.md)

* [Search](docs/sdks/search/README.md#search) - Search

### [Tags](docs/sdks/tags/README.md)

* [List](docs/sdks/tags/README.md#list) - List all tags
* [Create](docs/sdks/tags/README.md#create) - Create a new tag
* [Get](docs/sdks/tags/README.md#get) - Retrieve a tag
* [Delete](docs/sdks/tags/README.md#delete) - Delete a tag
* [Update](docs/sdks/tags/README.md#update) - Update a tag

### [Teams](docs/sdks/teams/README.md)

* [List](docs/sdks/teams/README.md#list) - List all teams
* [Get](docs/sdks/teams/README.md#get) - Retrieve a team
* [Update](docs/sdks/teams/README.md#update) - Update a team
* [Members](docs/sdks/teams/README.md#members) - List all team members

### [Tracker](docs/sdks/tracker/README.md)

* [Delete](docs/sdks/tracker/README.md#delete) - Delete a tracker entry

### [TrackerEntries](docs/sdks/trackerentries/README.md)

* [List](docs/sdks/trackerentries/README.md#list) - List all tracker entries

### [TrackerProjects](docs/sdks/trackerprojects/README.md)

* [List](docs/sdks/trackerprojects/README.md#list) - List all tracker projects
* [Create](docs/sdks/trackerprojects/README.md#create) - Create a tracker project
* [Get](docs/sdks/trackerprojects/README.md#get) - Retrieve a tracker project
* [Delete](docs/sdks/trackerprojects/README.md#delete) - Delete a tracker project
* [Update](docs/sdks/trackerprojects/README.md#update) - Update a tracker project

### [Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - List all transactions
* [Create](docs/sdks/transactions/README.md#create) - Create a transaction
* [Get](docs/sdks/transactions/README.md#get) - Retrieve a transaction
* [Delete](docs/sdks/transactions/README.md#delete) - Delete a transaction
* [Update](docs/sdks/transactions/README.md#update) - Update a transaction
* [CreateMany](docs/sdks/transactions/README.md#createmany) - Bulk create transactions
* [DeleteMany](docs/sdks/transactions/README.md#deletemany) - Bulk delete transactions
* [UpdateMany](docs/sdks/transactions/README.md#updatemany) - Bulk update transactions

### [Users](docs/sdks/users/README.md)

* [Get](docs/sdks/users/README.md#get) - Retrieve the current user
* [Update](docs/sdks/users/README.md#update) - Update the current user

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"github.com/midday-ai/midday-go/retry"
	"log"
	"models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/operations"
	"github.com/midday-ai/midday-go/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := middaygo.New(
		middaygo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| apierrors.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	middaygo "github.com/midday-ai/midday-go"
	"github.com/midday-ai/midday-go/models/apierrors"
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

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		middaygo.WithServerURL("https://api.midday.ai"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/midday-ai/midday-go&utm_campaign=go)
