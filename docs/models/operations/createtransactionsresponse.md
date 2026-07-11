# CreateTransactionsResponse


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md)                 | :heavy_check_mark:                                                                 | N/A                                                                                |
| `TransactionResponses`                                                             | [][components.TransactionResponse](../../models/components/transactionresponse.md) | :heavy_minus_sign:                                                                 | Transactions created                                                               |
| `ErrorResponse`                                                                    | [*components.ErrorResponse](../../models/components/errorresponse.md)              | :heavy_minus_sign:                                                                 | An error occurred                                                                  |