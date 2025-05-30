# GetTeamByIDResponseBody

Team details


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ID`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | Unique identifier of the team                                            | 123e4567-e89b-12d3-a456-426614174000                                     |
| `Name`                                                                   | *string*                                                                 | :heavy_check_mark:                                                       | Name of the team or organization                                         | Acme Corporation                                                         |
| `LogoURL`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | URL to the team's logo image                                             | https://cdn.midday.ai/logos/acme-corp.png                                |
| `Plan`                                                                   | [operations.GetTeamByIDPlan](../../models/operations/getteambyidplan.md) | :heavy_check_mark:                                                       | Current subscription plan of the team                                    | pro                                                                      |