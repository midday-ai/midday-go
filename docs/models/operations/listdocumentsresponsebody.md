# ListDocumentsResponseBody

Response containing a list of documents and pagination metadata.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Meta`                                                                         | [operations.ListDocumentsMeta](../../models/operations/listdocumentsmeta.md)   | :heavy_check_mark:                                                             | Pagination metadata for the documents list.                                    |
| `Data`                                                                         | [][operations.ListDocumentsData](../../models/operations/listdocumentsdata.md) | :heavy_check_mark:                                                             | Array of document objects.                                                     |