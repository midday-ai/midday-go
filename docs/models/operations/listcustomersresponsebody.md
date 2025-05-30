# ListCustomersResponseBody

Retrieve a list of customers for the authenticated team.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Meta`                                                                         | [operations.ListCustomersMeta](../../models/operations/listcustomersmeta.md)   | :heavy_check_mark:                                                             | Pagination metadata for the customers response                                 |
| `Data`                                                                         | [][operations.ListCustomersData](../../models/operations/listcustomersdata.md) | :heavy_check_mark:                                                             | Array of customers matching the query criteria                                 |