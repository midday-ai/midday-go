# Category

Category information assigned to the transaction for organization


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ID`                                                       | *string*                                                   | :heavy_check_mark:                                         | Unique identifier of the category                          | office-supplies                                            |
| `Name`                                                     | *string*                                                   | :heavy_check_mark:                                         | Display name of the category                               | Office Supplies                                            |
| `Color`                                                    | *string*                                                   | :heavy_check_mark:                                         | Hex color code associated with the category for UI display | #FF5733                                                    |
| `TaxRate`                                                  | *float64*                                                  | :heavy_check_mark:                                         | Tax rate of the category                                   | 10                                                         |
| `TaxType`                                                  | *string*                                                   | :heavy_check_mark:                                         | Tax type of the category                                   | VAT                                                        |
| `Slug`                                                     | *string*                                                   | :heavy_check_mark:                                         | URL-friendly slug of the category                          | office-supplies                                            |