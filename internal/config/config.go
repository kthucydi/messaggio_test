package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	logging "github.com/kthucydi/bs_go_logrus"
	"messaggio_test/migrations"
)

const fileEnv = ".env"

type ConfigStructure struct {
	Data             map[string]string
	DBConnectionData map[string]string
	Migration        migrations.Data
	Kafka            map[string]string
}

var Log = &logging.Log
var Cfg *ConfigStructure = newConfig()

func newConfig() *ConfigStructure {
	LoadEnv(fileEnv)

	cfg := ConfigStructure{}
	cfg.flagHandle()
	cfg.initConfigBase()
	cfg.configFill()
	cfg.configValidate()
	cfg.getPostgresConfig()

	return &cfg
}

func (config *ConfigStructure) flagHandle() {
	//Flags and args handle

	MigrationArgsNotExist := len(os.Args) < 2
	if MigrationArgsNotExist {
		config.Migration.Exit = true
		return
	} else {

		flags := flag.NewFlagSet("goose", flag.ContinueOnError)
		migrationCommand := flags.String("m", "", "migration command")
		migrationOnly := flags.String("m-only", "", "migration Only")
		// migrationOnly := flags.Bool("m-only", false, "migration Only")
		err := flags.Parse(os.Args[1:])
		if err != nil {
			Log.Errorf("error during parse flags: %v", err)
			return
		}

		config.Migration.Command = *migrationCommand
		if *migrationOnly == "" {
			config.Migration.Only = false
			config.Migration.Command = *migrationCommand
		} else {
			config.Migration.Only = true
			config.Migration.Command = *migrationOnly
		}

		config.Migration.Args = flag.Args()
	}
}

// Here to edit environment variable
func (config *ConfigStructure) initConfigBase() {
	config.Data = map[string]string{
		"JWT_TOKEN_DURATION":        "",
		"JWT_SECRET":                "",
		"BACKEND_SERVER_PORT":       "",
		"BACKEND_SERVER_URL_PREFIX": "",
		"USE_INNER_CORS":            "",
		"USE_INNER_LOGGER":          "",
		"LOG_LEVEL":                 "",
		"LOG_FILE_PATH":             "",
		// "DAEMON_PID_FILE_NAME": "",
		// "DAEMON_LOG_FILE_NAME": "",
		"DAEMON_MODE":             "",
		"KAFKA_BOOTSTRAP_SERVERS": "",
		"KAFKA_TOPIC_PRODUCER":    "",
		"KAFKA_TOPIC_CONSUMER":    "",
		"KAFKA_SEND_TIMES":        "",
	}
}

// Loading variable from file if it exist
func LoadEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		Log.Warnf("Can not loading environment from %s", fileName)
	}
}

// Fill config value from environment
func (config *ConfigStructure) configFill() {
	for key := range config.Data {
		if value := os.Getenv(key); value != "" {
			config.Data[key] = value
		}
	}
}

// Validate config variable for not empty exist
func (config *ConfigStructure) configValidate() {
	exitFlag := false
	for key, value := range config.Data {
		if value == "" {
			Log.Errorf("environment variable %s not exported", key)
			exitFlag = true
		}
	}
	if exitFlag {
		Log.Fatal("the necessary variables have not been exported")
	}
}

func (config *ConfigStructure) getPostgresConfig() (err error) {
	data := make(map[string]string)

	// Fill struct with nessesory variable
	data["POSTGRES_HOST"] = os.Getenv("POSTGRES_HOST")
	data["POSTGRES_PORT"] = os.Getenv("POSTGRES_PORT")
	data["POSTGRES_USER"] = os.Getenv("POSTGRES_USER")
	data["POSTGRES_PASSWORD"] = os.Getenv("POSTGRES_PASSWORD")
	data["POSTGRES_DATABASE"] = os.Getenv("POSTGRES_DATABASE")

	// Check exist all nessesory variable for connection
	for key, value := range data {
		if value == "" {
			Log.Printf("config error: env variable %s not exported", key)
			return errors.New("config error: cannot get nessesory environment variable")
		}
	}

	// Create connection string
	data["PGX_CONN"] = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		data["POSTGRES_USER"],
		data["POSTGRES_PASSWORD"],
		data["POSTGRES_HOST"],
		data["POSTGRES_PORT"],
		data["POSTGRES_DATABASE"])

	data["POSTGRES_STRING"] = fmt.Sprintf("%s:%s %s",
		data["POSTGRES_HOST"],
		data["POSTGRES_PORT"],
		data["POSTGRES_DATABASE"])

	config.DBConnectionData = data
	return
}

// import (
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	logging "github.com/kthucydi/bs_go_logrus"
// 	"messaggio_test/migrations"
// )

// const fileEnv = ".env"

// type ConfigStructure struct {
// 	Data             map[string]string
// 	DBConnectionData map[string]string
// 	Migration        migrations.Data
// }

// var Log = &logging.Log
// var Cfg *ConfigStructure = newConfig()

// func newConfig() *ConfigStructure {
// 	LoadEnv(fileEnv)

// 	cfg := ConfigStructure{}
// 	cfg.flagHandle()
// 	cfg.initConfigBase()
// 	cfg.configFill()
// 	cfg.configValidate()

// 	return &cfg
// }

// func (config *ConfigStructure) flagHandle() {
// 	//Flags and args handle

// 	if len(os.Args) < 2 {
// 		config.Migration.Exit = true
// 		return
// 	} else {

// 		flags := flag.NewFlagSet("goose", flag.ContinueOnError)
// 		migrationCommand := flags.String("m", "", "migration command")
// 		migrationOnly := flags.Bool("m-only", false, "migration Only")
// 		err := flags.Parse(os.Args[1:])
// 		if err != nil {
// 			Log.Errorf("error during parse flags: %v", err)
// 			return
// 		}

// 		config.Migration.Command = *migrationCommand
// 		config.Migration.Only = *migrationOnly
// 		config.Migration.Args = flag.Args()

// 		// Skip migration if flags not set
// 		if len(os.Args) < 2 {
// 			return
// 		}
// 	}
// }

// // Here to edit environment variable
// func (config *ConfigStructure) initConfigBase() {
// 	config.Data = map[string]string{
// 		"JWT_TOKEN_DURATION":        "",
// 		"JWT_SECRET":                "",
// 		"BACKEND_SERVER_PORT":       "",
// 		"BACKEND_SERVER_URL_PREFIX": "",
// 		"USE_INNER_CORS":            "",
// 		"USE_INNER_LOGGER":          "",
// 		"LOG_LEVEL":                 "",
// 		"LOG_FILE_PATH":             "",
// 		"MAIL_FROM":                 "",
// 		"MAIL_FROM_PASSWORD":        "",
// 		"MAIL_SERVER":               "",
// 		"MAIL_PORT":                 "",
// 		// "DAEMON_PID_FILE_NAME": "",
// 		// "DAEMON_LOG_FILE_NAME": "",
// 		"DAEMON_MODE": "",
// 	}
// }

// // Loading variable from file if it exist
// func LoadEnv(fileName string) {
// 	err := godotenv.Load(fileName)
// 	if err != nil {
// 		Log.Warnf("Can not loading environment from %s", fileName)
// 	}
// }

// // Fill config value from environment
// func (config *ConfigStructure) configFill() {
// 	for key := range config.Data {
// 		if value := os.Getenv(key); value != "" {
// 			config.Data[key] = value
// 		}
// 	}
// }

// // Validate config variable for not empty exist
// func (config *ConfigStructure) configValidate() {
// 	exitFlag := false
// 	for key, value := range config.Data {
// 		if value == "" {
// 			Log.Errorf("environment variable %s not exported", key)
// 			exitFlag = true
// 		}
// 	}
// 	if exitFlag {
// 		Log.Fatal("the necessary variables have not been exported")
// 	}
// }

// func (config *ConfigStructure) getDBPGConnectionData() (err error) {
// 	data := make(map[string]string)

// 	// Fill struct with nessesory variable
// 	data["PGHOST"] = os.Getenv("PGHOST")
// 	data["PGPORT"] = os.Getenv("PGPORT")
// 	data["PGUSER"] = os.Getenv("PGUSER")
// 	data["PGPASSWORD"] = os.Getenv("PGPASSWORD")
// 	data["PGDATABASE"] = os.Getenv("PGDATABASE")

// 	// Check exist all nessesory variable for connection
// 	for key, value := range data {
// 		if value == "" {
// 			Log.Printf("config error: env variable %s not exported", key)
// 			return errors.New("config error: cannot get nessesory environment variable")
// 		}
// 	}

// 	// Create connection string
// 	data["pgxconn"] = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
// 		data["PGUSER"],
// 		data["PGPASSWORD"],
// 		data["PGHOST"],
// 		data["PGPORT"],
// 		data["PGDATABASE"])

// 	data["dbPGStringPath"] = fmt.Sprintf("%s:%s %s",
// 		data["PGHOST"],
// 		data["PGPORT"],
// 		data["PGDATABASE"])

// 	config.DBConnectionData = data
// 	return
// }
