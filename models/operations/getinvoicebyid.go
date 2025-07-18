// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/midday-ai/midday-go/models/components"
)

type GetInvoiceByIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetInvoiceByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetInvoiceByIDStatus - Current status of the invoice
type GetInvoiceByIDStatus string

const (
	GetInvoiceByIDStatusDraft    GetInvoiceByIDStatus = "draft"
	GetInvoiceByIDStatusOverdue  GetInvoiceByIDStatus = "overdue"
	GetInvoiceByIDStatusPaid     GetInvoiceByIDStatus = "paid"
	GetInvoiceByIDStatusUnpaid   GetInvoiceByIDStatus = "unpaid"
	GetInvoiceByIDStatusCanceled GetInvoiceByIDStatus = "canceled"
)

func (e GetInvoiceByIDStatus) ToPointer() *GetInvoiceByIDStatus {
	return &e
}
func (e *GetInvoiceByIDStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft":
		fallthrough
	case "overdue":
		fallthrough
	case "paid":
		fallthrough
	case "unpaid":
		fallthrough
	case "canceled":
		*e = GetInvoiceByIDStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetInvoiceByIDStatus: %v", v)
	}
}

// GetInvoiceByIDCustomer - Customer details
type GetInvoiceByIDCustomer struct {
	// Unique identifier for the customer
	ID string `json:"id"`
	// Name of the customer
	Name string `json:"name"`
	// Website URL of the customer
	Website *string `json:"website"`
	// Email address of the customer
	Email *string `json:"email"`
}

func (o *GetInvoiceByIDCustomer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetInvoiceByIDCustomer) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetInvoiceByIDCustomer) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *GetInvoiceByIDCustomer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

// GetInvoiceByIDResponseBody - Invoice object
type GetInvoiceByIDResponseBody struct {
	// Unique identifier for the invoice
	ID string `json:"id"`
	// Current status of the invoice
	Status GetInvoiceByIDStatus `json:"status"`
	// Due date of the invoice in ISO 8601 format
	DueDate string `json:"dueDate"`
	// Issue date of the invoice in ISO 8601 format
	IssueDate string `json:"issueDate"`
	// Invoice number as shown to the customer
	InvoiceNumber string `json:"invoiceNumber"`
	// Total amount of the invoice
	Amount float64 `json:"amount"`
	// Currency code (ISO 4217) for the invoice amount
	Currency string `json:"currency"`
	// Customer details
	Customer GetInvoiceByIDCustomer `json:"customer"`
	// Timestamp when the invoice was paid (ISO 8601), or null if unpaid
	PaidAt *string `json:"paidAt"`
	// Timestamp when a payment reminder was sent (ISO 8601), or null if never sent
	ReminderSentAt *string `json:"reminderSentAt"`
	// Optional note attached to the invoice
	Note *string `json:"note"`
	// Value-added tax amount, or null if not applicable
	Vat *float64 `json:"vat"`
	// Tax amount, or null if not applicable
	Tax *float64 `json:"tax"`
	// Discount amount applied to the invoice, or null if none
	Discount *float64 `json:"discount"`
	// Subtotal before taxes and discounts, or null if not calculated
	Subtotal *float64 `json:"subtotal"`
	// Timestamp when the invoice was viewed by the customer (ISO 8601), or null if never viewed
	ViewedAt *string `json:"viewedAt"`
	// Name of the customer as shown on the invoice, or null if not set
	CustomerName *string `json:"customerName"`
	// Email address to which the invoice was sent, or null if not sent
	SentTo *string `json:"sentTo"`
	// Timestamp when the invoice was sent (ISO 8601), or null if not sent
	SentAt *string `json:"sentAt"`
	// Timestamp when the invoice was created (ISO 8601)
	CreatedAt string `json:"createdAt"`
	// Timestamp when the invoice was last updated (ISO 8601)
	UpdatedAt string `json:"updatedAt"`
}

func (o *GetInvoiceByIDResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetInvoiceByIDResponseBody) GetStatus() GetInvoiceByIDStatus {
	if o == nil {
		return GetInvoiceByIDStatus("")
	}
	return o.Status
}

func (o *GetInvoiceByIDResponseBody) GetDueDate() string {
	if o == nil {
		return ""
	}
	return o.DueDate
}

func (o *GetInvoiceByIDResponseBody) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *GetInvoiceByIDResponseBody) GetInvoiceNumber() string {
	if o == nil {
		return ""
	}
	return o.InvoiceNumber
}

func (o *GetInvoiceByIDResponseBody) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *GetInvoiceByIDResponseBody) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *GetInvoiceByIDResponseBody) GetCustomer() GetInvoiceByIDCustomer {
	if o == nil {
		return GetInvoiceByIDCustomer{}
	}
	return o.Customer
}

func (o *GetInvoiceByIDResponseBody) GetPaidAt() *string {
	if o == nil {
		return nil
	}
	return o.PaidAt
}

func (o *GetInvoiceByIDResponseBody) GetReminderSentAt() *string {
	if o == nil {
		return nil
	}
	return o.ReminderSentAt
}

func (o *GetInvoiceByIDResponseBody) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *GetInvoiceByIDResponseBody) GetVat() *float64 {
	if o == nil {
		return nil
	}
	return o.Vat
}

func (o *GetInvoiceByIDResponseBody) GetTax() *float64 {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *GetInvoiceByIDResponseBody) GetDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.Discount
}

func (o *GetInvoiceByIDResponseBody) GetSubtotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Subtotal
}

func (o *GetInvoiceByIDResponseBody) GetViewedAt() *string {
	if o == nil {
		return nil
	}
	return o.ViewedAt
}

func (o *GetInvoiceByIDResponseBody) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *GetInvoiceByIDResponseBody) GetSentTo() *string {
	if o == nil {
		return nil
	}
	return o.SentTo
}

func (o *GetInvoiceByIDResponseBody) GetSentAt() *string {
	if o == nil {
		return nil
	}
	return o.SentAt
}

func (o *GetInvoiceByIDResponseBody) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *GetInvoiceByIDResponseBody) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

type GetInvoiceByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a invoice by its unique identifier for the authenticated team.
	Object *GetInvoiceByIDResponseBody
}

func (o *GetInvoiceByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetInvoiceByIDResponse) GetObject() *GetInvoiceByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
