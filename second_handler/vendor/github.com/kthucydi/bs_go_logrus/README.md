# Introduction 

release v0.1.0

Package "logging" tune logrus logger for use with environment variable.
Package provide all logrus functional.
During work, package try create directory by LOG_FILE_PATH.

# For Use
Set environment variable:

"LOG_LEVEL"  = {0..6} - logging level: 0 - panic .. 6 - trace, default - 4

"LOG_FILE_PATH"  = Path with fileName for writing log, default logs/all.log
```
example:
LOG_LEVEL=4
LOG_FILE_PATH=logs/all.log
```

import:
```
logging import git@github.com:kthucydi/bs_go_logrus.git
```
registration in code (recommend as global variable):
```
Logger = &logging.Log
```

Use in code:
```
Logger.Error("message")
Logger.Info("message")
```