# CreateInvoiceBadRequestError

Bad request. Invalid input data or validation errors.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Message`                                                          | *string*                                                           | :heavy_check_mark:                                                 | Error message describing the validation failure                    | scheduledAt is required for scheduled delivery                     |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |