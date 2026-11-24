# PostOAuthRegisterRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ClientName`                                       | *string*                                           | :heavy_check_mark:                                 | Human-readable name of the client                  | ChatGPT                                            |
| `RedirectUris`                                     | []*string*                                         | :heavy_check_mark:                                 | Array of redirect URIs                             | [<br/>"https://chatgpt.com/connector/oauth/callback"<br/>] |
| `GrantTypes`                                       | []*string*                                         | :heavy_minus_sign:                                 | Requested grant types                              | [<br/>"authorization_code",<br/>"refresh_token"<br/>] |
| `TokenEndpointAuthMethod`                          | **string*                                          | :heavy_minus_sign:                                 | Token endpoint authentication method               | none                                               |
| `Scope`                                            | **string*                                          | :heavy_minus_sign:                                 | Space-separated requested scopes                   | transactions.read invoices.read                    |
| `LogoURI`                                          | **string*                                          | :heavy_minus_sign:                                 | URL of the client logo                             | https://example.com/logo.png                       |
| `ClientURI`                                        | **string*                                          | :heavy_minus_sign:                                 | URL of the client homepage                         | https://example.com                                |
| `ResponseTypes`                                    | []*string*                                         | :heavy_minus_sign:                                 | Requested response types                           | [<br/>"code"<br/>]                                 |