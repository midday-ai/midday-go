# PostOAuthTokenResponseBody

Token exchange successful


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AccessToken`                                                | *string*                                                     | :heavy_check_mark:                                           | Access token for API requests                                | mid_access_token_abcdef123456789                             |
| `TokenType`                                                  | [operations.TokenType](../../models/operations/tokentype.md) | :heavy_check_mark:                                           | Token type, always 'Bearer'                                  | Bearer                                                       |
| `ExpiresIn`                                                  | *float64*                                                    | :heavy_check_mark:                                           | Token expiration time in seconds                             | 3600                                                         |
| `RefreshToken`                                               | *string*                                                     | :heavy_check_mark:                                           | Refresh token for obtaining new access tokens                | mid_refresh_token_abcdef123456789                            |
| `Scope`                                                      | *string*                                                     | :heavy_check_mark:                                           | Space-separated list of granted scopes                       | transactions.read invoices.read                              |