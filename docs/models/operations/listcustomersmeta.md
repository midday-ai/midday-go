# ListCustomersMeta

Pagination metadata for the customers response


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Cursor`                                                        | *string*                                                        | :heavy_check_mark:                                              | Cursor for the next page of results, null if no more pages      | eyJpZCI6IjQ1NiJ9                                                |
| `HasPreviousPage`                                               | *bool*                                                          | :heavy_check_mark:                                              | Whether there are more customers available on the previous page | false                                                           |
| `HasNextPage`                                                   | *bool*                                                          | :heavy_check_mark:                                              | Whether there are more customers available on the next page     | true                                                            |