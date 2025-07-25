// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/midday-ai/midday-go/models/components"
)

type ListInvoicesRequest struct {
	Cursor    *string  `queryParam:"style=form,explode=true,name=cursor"`
	Sort      []string `queryParam:"style=form,explode=true,name=sort"`
	PageSize  *float64 `queryParam:"style=form,explode=true,name=pageSize"`
	Q         *string  `queryParam:"style=form,explode=true,name=q"`
	Start     *string  `queryParam:"style=form,explode=true,name=start"`
	End       *string  `queryParam:"style=form,explode=true,name=end"`
	Statuses  []string `queryParam:"style=form,explode=true,name=statuses"`
	Customers []string `queryParam:"style=form,explode=true,name=customers"`
}

func (o *ListInvoicesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListInvoicesRequest) GetSort() []string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListInvoicesRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *ListInvoicesRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *ListInvoicesRequest) GetStart() *string {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *ListInvoicesRequest) GetEnd() *string {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *ListInvoicesRequest) GetStatuses() []string {
	if o == nil {
		return nil
	}
	return o.Statuses
}

func (o *ListInvoicesRequest) GetCustomers() []string {
	if o == nil {
		return nil
	}
	return o.Customers
}

// ListInvoicesMeta - Pagination metadata
type ListInvoicesMeta struct {
	// Cursor for pagination; null if there is no next page
	Cursor *string `json:"cursor"`
	// Indicates if there is a previous page of results
	HasPreviousPage bool `json:"hasPreviousPage"`
	// Indicates if there is a next page of results
	HasNextPage bool `json:"hasNextPage"`
}

func (o *ListInvoicesMeta) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListInvoicesMeta) GetHasPreviousPage() bool {
	if o == nil {
		return false
	}
	return o.HasPreviousPage
}

func (o *ListInvoicesMeta) GetHasNextPage() bool {
	if o == nil {
		return false
	}
	return o.HasNextPage
}

// ListInvoicesStatus - Current status of the invoice
type ListInvoicesStatus string

const (
	ListInvoicesStatusDraft    ListInvoicesStatus = "draft"
	ListInvoicesStatusOverdue  ListInvoicesStatus = "overdue"
	ListInvoicesStatusPaid     ListInvoicesStatus = "paid"
	ListInvoicesStatusUnpaid   ListInvoicesStatus = "unpaid"
	ListInvoicesStatusCanceled ListInvoicesStatus = "canceled"
)

func (e ListInvoicesStatus) ToPointer() *ListInvoicesStatus {
	return &e
}
func (e *ListInvoicesStatus) UnmarshalJSON(data []byte) error {
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
		*e = ListInvoicesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListInvoicesStatus: %v", v)
	}
}

// ListInvoicesCustomer - Customer details
type ListInvoicesCustomer struct {
	// Unique identifier for the customer
	ID string `json:"id"`
	// Name of the customer
	Name string `json:"name"`
	// Website URL of the customer
	Website *string `json:"website"`
	// Email address of the customer
	Email *string `json:"email"`
}

func (o *ListInvoicesCustomer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListInvoicesCustomer) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListInvoicesCustomer) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *ListInvoicesCustomer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

// ListInvoicesData - Invoice object
type ListInvoicesData struct {
	// Unique identifier for the invoice
	ID string `json:"id"`
	// Current status of the invoice
	Status ListInvoicesStatus `json:"status"`
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
	Customer ListInvoicesCustomer `json:"customer"`
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

func (o *ListInvoicesData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListInvoicesData) GetStatus() ListInvoicesStatus {
	if o == nil {
		return ListInvoicesStatus("")
	}
	return o.Status
}

func (o *ListInvoicesData) GetDueDate() string {
	if o == nil {
		return ""
	}
	return o.DueDate
}

func (o *ListInvoicesData) GetIssueDate() string {
	if o == nil {
		return ""
	}
	return o.IssueDate
}

func (o *ListInvoicesData) GetInvoiceNumber() string {
	if o == nil {
		return ""
	}
	return o.InvoiceNumber
}

func (o *ListInvoicesData) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *ListInvoicesData) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *ListInvoicesData) GetCustomer() ListInvoicesCustomer {
	if o == nil {
		return ListInvoicesCustomer{}
	}
	return o.Customer
}

func (o *ListInvoicesData) GetPaidAt() *string {
	if o == nil {
		return nil
	}
	return o.PaidAt
}

func (o *ListInvoicesData) GetReminderSentAt() *string {
	if o == nil {
		return nil
	}
	return o.ReminderSentAt
}

func (o *ListInvoicesData) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *ListInvoicesData) GetVat() *float64 {
	if o == nil {
		return nil
	}
	return o.Vat
}

func (o *ListInvoicesData) GetTax() *float64 {
	if o == nil {
		return nil
	}
	return o.Tax
}

func (o *ListInvoicesData) GetDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.Discount
}

func (o *ListInvoicesData) GetSubtotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Subtotal
}

func (o *ListInvoicesData) GetViewedAt() *string {
	if o == nil {
		return nil
	}
	return o.ViewedAt
}

func (o *ListInvoicesData) GetCustomerName() *string {
	if o == nil {
		return nil
	}
	return o.CustomerName
}

func (o *ListInvoicesData) GetSentTo() *string {
	if o == nil {
		return nil
	}
	return o.SentTo
}

func (o *ListInvoicesData) GetSentAt() *string {
	if o == nil {
		return nil
	}
	return o.SentAt
}

func (o *ListInvoicesData) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *ListInvoicesData) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

// ListInvoicesResponseBody - Response containing a list of invoices and pagination metadata
type ListInvoicesResponseBody struct {
	// Pagination metadata
	Meta ListInvoicesMeta `json:"meta"`
	// Array of invoice objects
	Data []ListInvoicesData `json:"data"`
}

func (o *ListInvoicesResponseBody) GetMeta() ListInvoicesMeta {
	if o == nil {
		return ListInvoicesMeta{}
	}
	return o.Meta
}

func (o *ListInvoicesResponseBody) GetData() []ListInvoicesData {
	if o == nil {
		return []ListInvoicesData{}
	}
	return o.Data
}

type ListInvoicesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A list of invoices for the authenticated team.
	Object *ListInvoicesResponseBody
}

func (o *ListInvoicesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListInvoicesResponse) GetObject() *ListInvoicesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
