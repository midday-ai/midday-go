# SearchRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `SearchTerm`                                              | **string*                                                 | :heavy_minus_sign:                                        | The term to search for across all data sources.           | Acme                                                      |
| `Language`                                                | **string*                                                 | :heavy_minus_sign:                                        | Language code to use for search relevance and results.    | en                                                        |
| `Limit`                                                   | **float64*                                                | :heavy_minus_sign:                                        | Maximum number of results to return.                      | 30                                                        |
| `ItemsPerTableLimit`                                      | **float64*                                                | :heavy_minus_sign:                                        | Maximum number of results to return per table/entity.     | 5                                                         |
| `RelevanceThreshold`                                      | **float64*                                                | :heavy_minus_sign:                                        | Minimum relevance score threshold for including a result. | 0.01                                                      |