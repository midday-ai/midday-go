# PostOAuthRegisterResponseBody2

Client registered successfully


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ClientID`                                         | *string*                                           | :heavy_check_mark:                                 | Assigned client ID                                 | mid_client_abcdef123456789                         |
| `ClientName`                                       | *string*                                           | :heavy_check_mark:                                 | Human-readable name of the client                  | ChatGPT                                            |
| `RedirectUris`                                     | []*string*                                         | :heavy_check_mark:                                 | Registered redirect URIs                           | [<br/>"https://chatgpt.com/connector/oauth/callback"<br/>] |
| `GrantTypes`                                       | []*string*                                         | :heavy_check_mark:                                 | Granted grant types                                | [<br/>"authorization_code",<br/>"refresh_token"<br/>] |
| `TokenEndpointAuthMethod`                          | *string*                                           | :heavy_check_mark:                                 | Token endpoint authentication method               | none                                               |
| `ResponseTypes`                                    | []*string*                                         | :heavy_check_mark:                                 | Supported response types                           | [<br/>"code"<br/>]                                 |