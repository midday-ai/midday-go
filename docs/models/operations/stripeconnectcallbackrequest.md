# StripeConnectCallbackRequest


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Code`                                    | **string*                                 | :heavy_minus_sign:                        | OAuth authorization code from Stripe      |
| `State`                                   | *string*                                  | :heavy_check_mark:                        | OAuth state parameter for CSRF protection |
| `Error`                                   | **string*                                 | :heavy_minus_sign:                        | OAuth error code if authorization failed  |
| `ErrorDescription`                        | **string*                                 | :heavy_minus_sign:                        | OAuth error description                   |