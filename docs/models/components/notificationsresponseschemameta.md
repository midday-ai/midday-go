# NotificationsResponseSchemaMeta

Pagination metadata


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   | Example                                       |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Cursor`                                      | *string*                                      | :heavy_check_mark:                            | Cursor for pagination (null if no more pages) | 40                                            |
| `HasPreviousPage`                             | *bool*                                        | :heavy_check_mark:                            | Whether there are previous pages available    | true                                          |
| `HasNextPage`                                 | *bool*                                        | :heavy_check_mark:                            | Whether there are more pages available        | false                                         |