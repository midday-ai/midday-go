# ListTransactionsResponseBody

Retrieve a list of transactions for the authenticated team.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Meta`                                                                             | [operations.ListTransactionsMeta](../../models/operations/listtransactionsmeta.md) | :heavy_check_mark:                                                                 | Pagination metadata for the transactions response                                  |
| `Data`                                                                             | [][components.TransactionResponse](../../models/components/transactionresponse.md) | :heavy_check_mark:                                                                 | Array of transactions matching the query criteria                                  |