# ListInboxItemsTransaction

Matched transaction for this inbox item, if any


## Fields

| Field                                | Type                                 | Required                             | Description                          | Example                              |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `ID`                                 | *string*                             | :heavy_check_mark:                   | Transaction ID (UUID)                | a1b2c3d4-5678-4e7a-9c1a-2b7c1e24c2a4 |
| `Amount`                             | *float64*                            | :heavy_check_mark:                   | Transaction amount                   | 123.45                               |
| `Currency`                           | *string*                             | :heavy_check_mark:                   | Transaction currency (ISO 4217)      | USD                                  |
| `Name`                               | *string*                             | :heavy_check_mark:                   | Transaction name or payee            | Acme Corp                            |
| `Date`                               | *string*                             | :heavy_check_mark:                   | Transaction date (ISO 8601)          | 2024-05-01                           |