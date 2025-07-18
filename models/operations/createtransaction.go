// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/midday-ai/midday-go/models/components"
)

type CreateTransactionAttachment struct {
	// Path(s) of the attachment file(s).
	Path []string `json:"path"`
	// Name of the attachment file.
	Name string `json:"name"`
	// Size of the attachment file in bytes.
	Size float64 `json:"size"`
	// MIME type of the attachment file.
	Type string `json:"type"`
}

func (o *CreateTransactionAttachment) GetPath() []string {
	if o == nil {
		return []string{}
	}
	return o.Path
}

func (o *CreateTransactionAttachment) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTransactionAttachment) GetSize() float64 {
	if o == nil {
		return 0.0
	}
	return o.Size
}

func (o *CreateTransactionAttachment) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type CreateTransactionRequest struct {
	// Name of the transaction.
	Name string `json:"name"`
	// Amount of the transaction.
	Amount float64 `json:"amount"`
	// Currency of the transaction.
	Currency string `json:"currency"`
	// Date of the transaction (ISO 8601).
	Date string `json:"date"`
	// Bank account ID associated with the transaction.
	BankAccountID string `json:"bankAccountId"`
	// Assigned user ID for the transaction.
	AssignedID *string `json:"assignedId,omitempty"`
	// Category slug for the transaction.
	CategorySlug *string `json:"categorySlug,omitempty"`
	// Note for the transaction.
	Note *string `json:"note,omitempty"`
	// Whether the transaction is internal.
	Internal *bool `json:"internal,omitempty"`
	// Array of attachments for the transaction.
	Attachments []CreateTransactionAttachment `json:"attachments,omitempty"`
}

func (o *CreateTransactionRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTransactionRequest) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *CreateTransactionRequest) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *CreateTransactionRequest) GetDate() string {
	if o == nil {
		return ""
	}
	return o.Date
}

func (o *CreateTransactionRequest) GetBankAccountID() string {
	if o == nil {
		return ""
	}
	return o.BankAccountID
}

func (o *CreateTransactionRequest) GetAssignedID() *string {
	if o == nil {
		return nil
	}
	return o.AssignedID
}

func (o *CreateTransactionRequest) GetCategorySlug() *string {
	if o == nil {
		return nil
	}
	return o.CategorySlug
}

func (o *CreateTransactionRequest) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *CreateTransactionRequest) GetInternal() *bool {
	if o == nil {
		return nil
	}
	return o.Internal
}

func (o *CreateTransactionRequest) GetAttachments() []CreateTransactionAttachment {
	if o == nil {
		return nil
	}
	return o.Attachments
}

type CreateTransactionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Transaction created
	TransactionResponse *components.TransactionResponse
}

func (o *CreateTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTransactionResponse) GetTransactionResponse() *components.TransactionResponse {
	if o == nil {
		return nil
	}
	return o.TransactionResponse
}
