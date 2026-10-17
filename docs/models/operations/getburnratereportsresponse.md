# GetBurnRateReportsResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `HTTPMeta`                                                            | [components.HTTPMetadata](../../models/components/httpmetadata.md)    | :heavy_check_mark:                                                    | N/A                                                                   |
| `GetBurnRateResponseSchema`                                           | [][components.BurnRateItem](../../models/components/burnrateitem.md)  | :heavy_minus_sign:                                                    | Burn rate reports for the authenticated team.                         |
| `ErrorResponse`                                                       | [*components.ErrorResponse](../../models/components/errorresponse.md) | :heavy_minus_sign:                                                    | An error occurred                                                     |