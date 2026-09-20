# GmailOAuthCallbackRequest


## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `Code`                                   | **string*                                | :heavy_minus_sign:                       | OAuth authorization code from Google     |
| `State`                                  | *string*                                 | :heavy_check_mark:                       | Encrypted OAuth state parameter          |
| `Error`                                  | **string*                                | :heavy_minus_sign:                       | OAuth error code if authorization failed |