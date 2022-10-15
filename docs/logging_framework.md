We need to research the logging framework space in go

Criteria for logging 

|                            | logrus                                             | zap                                        | zerolog                                                     |
|----------------------------|----------------------------------------------------|--------------------------------------------|-------------------------------------------------------------|
| structured logging formats | yes                                                | yes                                        | yes                                                         |
| Error handling behavior    | not documented                                     |                                            | https://github.com/rs/zerolog#error-logging-with-stacktrace |
| performance                |                                                    | https://github.com/uber-go/zap#performance | https://github.com/rs/zerolog#benchmarks                    |
| maintenance                | only security patching no v2- last release 07/2019 | last release 08/2022                       | last release 07/2017                                        |
| data types                 |                                                    |                                            | https://github.com/rs/zerolog#benchmarks                    |

Performance metrics can be found here :
https://github.com/rs/zerolog#benchmarks    
https://github.com/uber-go/zap#performance

conflicting maybes?


Decision:
Zerolog as it provided better example, especially on stack trace