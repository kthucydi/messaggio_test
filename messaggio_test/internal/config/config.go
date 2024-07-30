package config

import (
	"flag"
	"github.com/joho/godotenv"
	logging "github.com/kthucydi/bs_go_logrus"
	"os"
	"strings"
)

const FILE_ENV = ".env"

type ConfigData map[string]map[string]string

var Log = &logging.Log
var Cfg *ConfigData = (&ConfigData{}).setData()

// main config function - create and validate flag, file, env
func (cfg *ConfigData) setData() *ConfigData {

	cfg.loadEnv(FILE_ENV)     // load config from file to environment
	cfg.initConfigBase()      // set config data list
	cfg.configFill()          // fill config data from environment
	cfg.migrationFlagHandle() // set migration config by flag

	cfg.configValidate() // validate config data for non empty

	return cfg
}

func (cfg *ConfigData) initConfigBase() {

	(*cfg)["SERVER"] = map[string]string{
		"JWT_TOKEN_DURATION":        "72",
		"JWT_SECRET":                "secret",
		"BACKEND_SERVER_PORT":       "8080",
		"BACKEND_SERVER_URL_PREFIX": "/",
		"USE_INNER_CORS":            "false",
		"USE_INNER_LOGGER":          "false",
	}

	(*cfg)["LOG"] = map[string]string{
		"LOG_LEVEL":     "4",
		"LOG_FILE_PATH": "logs/all.log",
	}

	(*cfg)["DAEMON"] = map[string]string{
		"DAEMON_MODE": "true",
	}

	(*cfg)["KAFKA"] = map[string]string{
		"KAFKA_BOOTSTRAP_SERVERS": "",
		"KAFKA_TOPIC_PRODUCER":    "",
		"KAFKA_TOPIC_CONSUMER":    "",
		"KAFKA_SEND_TIMES":        "3",
	}

	(*cfg)["MIGRATIONS"] = map[string]string{
		"MIGRATIONS_ONLY":    "",
		"WITHOUT_MIGRATIONS": "",
		"MIGRATIONS_COMMAND": "",
		"MIGRATIONS_ARGS":    "",
	}

	(*cfg)["DB_CONN"] = map[string]string{
		"POSTGRES_USER":     "",
		"POSTGRES_PASSWORD": "",
		"POSTGRES_HOST":     "",
		"POSTGRES_PORT":     "",
		"POSTGRES_DATABASE": "",
	}
}

// Loading variable from file if it exist
func (cfg *ConfigData) loadEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		Log.Warnf("Can not loading environment from %s", fileName)
	}
}

// migration flags handler
func (cfg *ConfigData) migrationFlagHandle() {

	// Set default for validating
	(*cfg)["MIGRATIONS"]["WITHOUT_MIGRATIONS"] = "false"
	(*cfg)["MIGRATIONS"]["MIGRATIONS_ONLY"] = " "
	(*cfg)["MIGRATIONS"]["MIGRATIONS_COMMAND"] = " "
	(*cfg)["MIGRATIONS"]["MIGRATIONS_ARGS"] = " "

	// Check for flags presence
	if len(os.Args) < 2 { // migration args not exist
		(*cfg)["MIGRATIONS"]["WITHOUT_MIGRATIONS"] = "true"
		return
	}

	// Assign flag
	flags := flag.NewFlagSet("goose", flag.ContinueOnError)
	migrationCommand := flags.String("m", "", "migration command")
	migrationOnly := flags.String("m-only", "", "migration-only command")

	// Parse flag
	err := flags.Parse(os.Args[1:])
	if err != nil {
		Log.Errorf("error during parse flags: %v", err)
		return
	}

	// Check for empty flags
	if *migrationOnly == "" && *migrationCommand == "" {
		(*cfg)["MIGRATIONS"]["WITHOUT_MIGRATIONS"] = "true"
		return
	}

	// Set command and m-only flag
	if *migrationOnly == "" {
		(*cfg)["MIGRATIONS"]["MIGRATIONS_ONLY"] = "false"
		(*cfg)["MIGRATIONS"]["MIGRATIONS_COMMAND"] = *migrationCommand
	} else {
		(*cfg)["MIGRATIONS"]["MIGRATIONS_ONLY"] = "true"
		(*cfg)["MIGRATIONS"]["MIGRATIONS_COMMAND"] = *migrationOnly
	}

	// Set added args if it exist
	if temp := strings.Join(flag.Args(), ";"); temp != "" {
		(*cfg)["MIGRATIONS"]["MIGRATIONS_ARGS"] = temp
	} else {
		(*cfg)["MIGRATIONS"]["MIGRATIONS_ARGS"] = " "
	}

}

// Fill config value from environment
func (cfg *ConfigData) configFill() {
	for name, list := range *cfg {
		for key := range list {
			if value := os.Getenv(key); value != "" {
				(*cfg)[name][key] = value
			}
		}
	}
}

// Validate config variable for not empty exist
func (cfg *ConfigData) configValidate() {
	exitFlag := false
	for _, list := range *cfg {
		for key, value := range list {
			if value == "" {
				Log.Errorf("environment variable %s not exported", key)
				exitFlag = true
			}
		}
		if exitFlag {
			Log.Fatal("the necessary variables have not been exported")
		}
	}
}
