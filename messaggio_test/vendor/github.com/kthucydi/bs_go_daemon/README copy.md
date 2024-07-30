
This package is fork by go-daemon.

It's adapted to silence running and working with environment variable.

For import:
```
import (
	_ "dev.azure.com/PluhSoft/Office/_git/ps-go-deamon.git"
)
```
This method creat new process, and close parent process.

For kill process use:
```
pkill <name of process>
```
name of process usually = bynary name

## Environment:

Package try to load environment from ".env", if unsuccess, write warinig to log, and then try to load variables from environment.

Use next variables:
 - DAEMON_PID_FILE_NAME - filename for pid number (file's content - pid, for example - 34567), default = ""
 - DAEMON_LOG_FILE_NAME - filename for inner package log, default = ""

 dont stay it empty! if pid file or lof-file unnessesory, keep it commented

 - DAEMON_MODE - "true" if demon_mode on, else "false"