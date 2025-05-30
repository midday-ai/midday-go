# ListInvoicesRequest


## Fields

| Field                                    | Type                                     | Required                                 | Description                              | Example                                  |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Cursor`                                 | **string*                                | :heavy_minus_sign:                       | N/A                                      | 25                                       |
| `Sort`                                   | []*string*                               | :heavy_minus_sign:                       | N/A                                      | [<br/>"createdAt",<br/>"desc"<br/>]      |
| `PageSize`                               | **float64*                               | :heavy_minus_sign:                       | N/A                                      | 25                                       |
| `Q`                                      | **string*                                | :heavy_minus_sign:                       | N/A                                      | Acme                                     |
| `Start`                                  | **string*                                | :heavy_minus_sign:                       | N/A                                      | 2024-01-01                               |
| `End`                                    | **string*                                | :heavy_minus_sign:                       | N/A                                      | 2024-01-31                               |
| `Statuses`                               | []*string*                               | :heavy_minus_sign:                       | N/A                                      | [<br/>"paid",<br/>"unpaid"<br/>]         |
| `Customers`                              | []*string*                               | :heavy_minus_sign:                       | N/A                                      | [<br/>"customer-uuid-1",<br/>"customer-uuid-2"<br/>] |