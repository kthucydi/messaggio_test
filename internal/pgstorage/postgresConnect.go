package pgstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"messaggio_test/internal/config"
)

var cfg = config.Cfg.DBConnectionData

// init function create connect to DB
func init() {
	var err error

	Log.Debug("pgxconn: ", cfg["PGX_CONN"])
	PGDB.Database, err = pgxpool.New(context.Background(), cfg["PGX_CONN"])
	if err != nil {
		Log.Fatalf("database init failed: %v", err)
	}
	Log.Print("postgress database connection successfull")
	Log.Debugf("FROM postgresConnect GetNow: %s", fmt.Sprint(PGDB.GetNow()))
}

func (PGDB DB) GetNow() (string, error) {
	var time time.Time

	err := PGDB.Database.QueryRow(context.Background(), "select NOW()").Scan(&time)
	if err != nil {
		Log.Error("can not receive responce from db:", err)
		return "", err
	}
	return time.String(), err
}

// // init function create connect to DB
// func init() {
// 	var err error

// 	if PGDB.dataConnect, err = getConnectionData(); err != nil {
// 		Log.Errorf("database init failed: %v", err)
// 		os.Exit(1)
// 	}

// 	Log.Debug("pgxconn: ", PGDB.dataConnect["pgxconn"])
// 	PGDB.Database, err = pgxpool.New(context.Background(), PGDB.dataConnect["pgxconn"])
// 	if err != nil {
// 		Log.Errorf("database init failed: %v", err)
// 		os.Exit(1)
// 	}
// 	Log.Print("postgress database connection successfull")
// 	Log.Debugf("FROM postgresConnect GetNow: %s", fmt.Sprint(PGDB.GetNow()))
// }

// // getConnectionData load connection data from file or environment
// func getConnectionData() (dataConnect map[string]string, err error) {
// 	configFile := os.Getenv("CONFIG_FILE")
// 	if err := godotenv.Load(configFile); err != nil {
// 		Log.Warningf("postgres config file %s not found", configFile)
// 	}
// 	Log.Debugf("postgres config file %s loaded", configFile)

// 	// Fill struct with nessesory variable
// 	dataConnect = make(map[string]string)
// 	dataConnect["PGHOST"] = os.Getenv("PGHOST")
// 	dataConnect["PGPORT"] = os.Getenv("PGPORT")
// 	dataConnect["PGUSER"] = os.Getenv("PGUSER")
// 	dataConnect["PGPASSWORD"] = os.Getenv("PGPASSWORD")
// 	dataConnect["PGDATABASE"] = os.Getenv("PGDATABASE")

// 	Log.Debugf("postgres config file %s copyed", configFile)

// 	// Check exist all nessesory variable for connection
// 	for key, value := range PGDB.dataConnect {
// 		if value == "" {
// 			Log.Printf("env variable %s not exported", key)
// 			return dataConnect, errors.New("Cannot get nessesory environment variable")
// 		}
// 	}
// 	Log.Debugf("postgres config file %s valid", configFile)

// 	// Create connection string
// 	dataConnect["pgxconn"] = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
// 		dataConnect["PGUSER"],
// 		dataConnect["PGPASSWORD"],
// 		dataConnect["PGHOST"],
// 		dataConnect["PGPORT"],
// 		dataConnect["PGDATABASE"])

// 	return
// }

// func (PGDB DB) GetNow() (string, error) {
// 	var time time.Time

// 	err := PGDB.Database.QueryRow(context.Background(), "select NOW()").Scan(&time)
// 	if err != nil {
// 		Log.Error("can not receive responce from db:", err)
// 		return "", err
// 	}
// 	return time.String(), err
// }
