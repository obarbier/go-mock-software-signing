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
https://betterstack.com/community/guides/logging/zerolog/

Understand log level [LINK](https://sematext.com/blog/logging-levels/)

| Log Level |                                                                                       Importance                                                                                      |   |
|:---------:|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|---|
|   Fatal   | One or more key business functionalities are not working and the whole system doesn’t fulfill the business functionalities.                                                           |   |
|   Error   | One or more functionalities are not working, preventing some functionalities from working correctly.                                                                                  |   |
|    Warn   | Unexpected  behavior happened inside the application, but it is continuing its work  and the key business features are operating as expected.                                         |   |
|    Info   | An event happened, the event is purely informative and can be ignored during normal operations.                                                                                       |   |
|   Debug   | A log level used for events considered to be useful during software debugging when more granular information is needed.                                                               |   |
|   Trace   | A  log level describing events showing step by step execution of your code  that can be ignored during the standard operation, but may be useful  during extended debugging sessions. |   |   


### hierarchy

TRACE -> DEBUG -> INFO -> WARN -> ERROR -> FATAL

If you set your logging framework to have the root logging level to WARN you will only get log events with WARN, ERROR, and FATAL levels


# Reference

* https://www.dataset.com/blog/the-10-commandments-of-logging/