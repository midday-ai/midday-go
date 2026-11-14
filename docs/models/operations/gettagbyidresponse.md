# GetTagByIDResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `HTTPMeta`                                                            | [components.HTTPMetadata](../../models/components/httpmetadata.md)    | :heavy_check_mark:                                                    | N/A                                                                   |
| `TagResponse`                                                         | [*components.TagResponse](../../models/components/tagresponse.md)     | :heavy_minus_sign:                                                    | Retrieve a tag by ID for the authenticated team.                      |
| `ErrorResponse`                                                       | [*components.ErrorResponse](../../models/components/errorresponse.md) | :heavy_minus_sign:                                                    | An error occurred                                                     |