// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/midday-ai/midday-go/models/components"
)

type ListBankAccountsRequest struct {
	Enabled *bool `queryParam:"style=form,explode=true,name=enabled"`
	Manual  *bool `queryParam:"style=form,explode=true,name=manual"`
}

func (o *ListBankAccountsRequest) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ListBankAccountsRequest) GetManual() *bool {
	if o == nil {
		return nil
	}
	return o.Manual
}

// ListBankAccountsData - A single bank account object response.
type ListBankAccountsData struct {
	// Unique identifier for the bank account.
	ID string `json:"id"`
	// Name of the bank account.
	Name *string `json:"name"`
	// Currency code of the bank account (e.g., USD, EUR).
	Currency *string `json:"currency"`
	// Type of the bank account (e.g., depository, credit).
	Type *string `json:"type"`
	// Whether the bank account is enabled.
	Enabled bool `json:"enabled"`
	// Current balance of the bank account.
	Balance *float64 `json:"balance"`
	// Whether the bank account is a manual account.
	Manual *bool `json:"manual"`
}

func (o *ListBankAccountsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListBankAccountsData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListBankAccountsData) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ListBankAccountsData) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListBankAccountsData) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *ListBankAccountsData) GetBalance() *float64 {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *ListBankAccountsData) GetManual() *bool {
	if o == nil {
		return nil
	}
	return o.Manual
}

// ListBankAccountsResponseBody - Response containing a list of bank accounts.
type ListBankAccountsResponseBody struct {
	// Array of bank account objects.
	Data []ListBankAccountsData `json:"data"`
}

func (o *ListBankAccountsResponseBody) GetData() []ListBankAccountsData {
	if o == nil {
		return []ListBankAccountsData{}
	}
	return o.Data
}

type ListBankAccountsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a list of bank accounts
	Object *ListBankAccountsResponseBody
}

func (o *ListBankAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListBankAccountsResponse) GetObject() *ListBankAccountsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
