# GetTimerStatusData


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `IsRunning`                                                        | *bool*                                                             | :heavy_check_mark:                                                 | Whether there is currently a running timer                         | true                                                               |
| `CurrentEntry`                                                     | [operations.CurrentEntry](../../models/operations/currententry.md) | :heavy_check_mark:                                                 | Current running timer details, null if not running                 |                                                                    |
| `ElapsedTime`                                                      | *float64*                                                          | :heavy_check_mark:                                                 | Elapsed time in seconds for the current running timer              | 1800                                                               |