package config

import (
	"github.com/joho/godotenv"
	logging "github.com/kthucydi/bs_go_logrus"
	"os"
)

const FILE_ENV = ".env"

type ConfigData map[string]map[string]string

var Log = &logging.Log
var Cfg *ConfigData = (&ConfigData{}).setData()

// main config function - create and validate flag, file, env
func (cfg *ConfigData) setData() *ConfigData {

	cfg.loadEnv(FILE_ENV) // load config from file to environment
	cfg.initConfigBase()  // set config data list
	cfg.configFill()      // fill config data from environment
	cfg.configValidate()  // validate config data for non empty

	return cfg
}

func (cfg *ConfigData) initConfigBase() {

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

}

// Loading variable from file if it exist
func (cfg *ConfigData) loadEnv(fileName string) {
	err := godotenv.Load(fileName)
	if err != nil {
		Log.Warnf("Can not loading environment from %s", fileName)
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
