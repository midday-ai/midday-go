# CreateBankAccountResponseBody

A single bank account object response.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `ID`                                                 | *string*                                             | :heavy_check_mark:                                   | Unique identifier for the bank account.              | b7e6c2a0-1f2d-4c3b-9a8e-123456789abc                 |
| `Name`                                               | *string*                                             | :heavy_check_mark:                                   | Name of the bank account.                            | Checking Account                                     |
| `Currency`                                           | *string*                                             | :heavy_check_mark:                                   | Currency code of the bank account (e.g., USD, EUR).  | USD                                                  |
| `Type`                                               | *string*                                             | :heavy_check_mark:                                   | Type of the bank account (e.g., depository, credit). | depository                                           |
| `Enabled`                                            | *bool*                                               | :heavy_check_mark:                                   | Whether the bank account is enabled.                 | true                                                 |
| `Balance`                                            | *float64*                                            | :heavy_check_mark:                                   | Current balance of the bank account.                 | 1500.75                                              |
| `Manual`                                             | *bool*                                               | :heavy_check_mark:                                   | Whether the bank account is a manual account.        | false                                                |