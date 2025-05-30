# ListInvoicesResponseBody

Response containing a list of invoices and pagination metadata


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Meta`                                                                       | [operations.ListInvoicesMeta](../../models/operations/listinvoicesmeta.md)   | :heavy_check_mark:                                                           | Pagination metadata                                                          |
| `Data`                                                                       | [][operations.ListInvoicesData](../../models/operations/listinvoicesdata.md) | :heavy_check_mark:                                                           | Array of invoice objects                                                     |