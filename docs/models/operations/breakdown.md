# Breakdown


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `Currency`                          | *string*                            | :heavy_check_mark:                  | Original currency of the invoices   | EUR                                 |
| `OriginalAmount`                    | *float64*                           | :heavy_check_mark:                  | Total amount in original currency   | 15000.5                             |
| `ConvertedAmount`                   | *float64*                           | :heavy_check_mark:                  | Amount converted to base currency   | 16250.75                            |
| `Count`                             | *float64*                           | :heavy_check_mark:                  | Number of invoices in this currency | 5                                   |