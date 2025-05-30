# UpdateTransactionsResponseBody

Transactions updated


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Meta`                                                                                 | [operations.UpdateTransactionsMeta](../../models/operations/updatetransactionsmeta.md) | :heavy_check_mark:                                                                     | Pagination metadata for the transactions response                                      |
| `Data`                                                                                 | [][components.TransactionResponse](../../models/components/transactionresponse.md)     | :heavy_check_mark:                                                                     | Array of transactions matching the query criteria                                      |