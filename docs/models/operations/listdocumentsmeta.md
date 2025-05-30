# ListDocumentsMeta

Pagination metadata for the documents list.


## Fields

| Field                             | Type                              | Required                          | Description                       | Example                           |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `Cursor`                          | **string*                         | :heavy_minus_sign:                | Cursor for pagination.            | 20                                |
| `HasPreviousPage`                 | *bool*                            | :heavy_check_mark:                | Whether there is a previous page. | false                             |
| `HasNextPage`                     | *bool*                            | :heavy_check_mark:                | Whether there is a next page.     | true                              |