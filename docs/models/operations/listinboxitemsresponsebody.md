# ListInboxItemsResponseBody

Retrieve a list of inbox items for the authenticated team.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Meta`                                                                           | [operations.ListInboxItemsMeta](../../models/operations/listinboxitemsmeta.md)   | :heavy_check_mark:                                                               | Pagination metadata for the inbox list response.                                 |
| `Data`                                                                           | [][operations.ListInboxItemsData](../../models/operations/listinboxitemsdata.md) | :heavy_check_mark:                                                               | List of inbox items                                                              |