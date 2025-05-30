# CreateBankAccountRequest

Schema for creating a new bank account.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Name`                                             | *string*                                           | :heavy_check_mark:                                 | The name of the bank account.                      | Checking Account                                   |
| `Currency`                                         | **string*                                          | :heavy_minus_sign:                                 | The currency code for the bank account (ISO 4217). | USD                                                |
| `Manual`                                           | **bool*                                            | :heavy_minus_sign:                                 | Whether the bank account is a manual account.      | false                                              |