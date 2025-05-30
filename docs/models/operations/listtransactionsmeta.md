# ListTransactionsMeta

Pagination metadata for the transactions response


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Cursor`                                                           | **string*                                                          | :heavy_minus_sign:                                                 | Cursor for the next page of results, undefined if no more pages    | eyJpZCI6IjQ1NiJ9                                                   |
| `HasPreviousPage`                                                  | *bool*                                                             | :heavy_check_mark:                                                 | Whether there are more transactions available on the previous page | false                                                              |
| `HasNextPage`                                                      | *bool*                                                             | :heavy_check_mark:                                                 | Whether there are more transactions available on the next page     | true                                                               |