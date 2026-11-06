# QuickBooksOAuthCallbackRequest


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Code`                                    | **string*                                 | :heavy_minus_sign:                        | OAuth authorization code from QuickBooks  |
| `State`                                   | *string*                                  | :heavy_check_mark:                        | OAuth state parameter for CSRF protection |
| `RealmID`                                 | **string*                                 | :heavy_minus_sign:                        | QuickBooks company/realm ID               |
| `Error`                                   | **string*                                 | :heavy_minus_sign:                        | OAuth error code if authorization failed  |