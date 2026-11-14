# ListTagsResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `HTTPMeta`                                                            | [components.HTTPMetadata](../../models/components/httpmetadata.md)    | :heavy_check_mark:                                                    | N/A                                                                   |
| `TagsResponse`                                                        | [*components.TagsResponse](../../models/components/tagsresponse.md)   | :heavy_minus_sign:                                                    | Retrieve a list of tags for the authenticated team.                   |
| `ErrorResponse`                                                       | [*components.ErrorResponse](../../models/components/errorresponse.md) | :heavy_minus_sign:                                                    | An error occurred                                                     |