package daemon

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		Log.Warnf("Can not loading environment from %s", ".env")
	}
	Env = GetEnvVariable()
	if Env["DAEMON_MODE"] == "true" {
		DarkRun()
	}
}

func DarkRun() {
	daemonObject := createContext()             // Create context for daemon
	processStruct, err := daemonObject.Reborn() // Create child process
	if err != nil {
		Log.Fatal("Unable to run daemon: ", err)
	} else if processStruct != nil { // close parent process
		os.Exit(0) // exit for parent
	}
	Log.Println("Daemon created") //log child run
}

func Run() (*Context, bool) {
	daemonObject := createContext()             // Create context for daemon
	processStruct, err := daemonObject.Reborn() // Create child process
	if err != nil {
		Log.Fatal("Unable to run daemon: ", err)
	} else if processStruct != nil { // close parent process
		return daemonObject, true // return true for parent
	}
	return daemonObject, false // return false for child
}

func createContext() *Context {

	return &Context{
		PidFileName: Env["DAEMON_PID_FILE_NAME"],
		PidFilePerm: 0644,
		LogFileName: Env["DAEMON_LOG_FILE_NAME"],
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		// Args:        []string{},
		// Args:        []string{"[go-daemon sample]"},
		// Args: os.Args,
	}
}

func Desactive(daemon *Context) {
	if err := daemon.Release(); err != nil {
		Log.Warning("error during closing daemon")
		return
	}
}

func GetEnvVariable() map[string]string {
	return map[string]string{
		"DAEMON_PID_FILE_NAME": os.Getenv("DAEMON_PID_FILE_NAME"),
		"DAEMON_LOG_FILE_NAME": os.Getenv("DAEMON_LOG_FILE_NAME"),
		"DAEMON_MODE":          os.Getenv("DAEMON_MODE"),
	}
}
