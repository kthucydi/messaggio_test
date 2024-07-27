package pgstorage

import (
	"context"
	"fmt"
	"time"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	logging "github.com/kthucydi/bs_go_logrus"
	config "messaggio_test/internal/config"
)

var Log = &logging.Log
var PGDB *pgxpool.Pool
var configDB = (*config.Cfg)["DB_CONN"]

// init function create connect to DB
func init() {
	var err error

	pgxConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		configDB["POSTGRES_USER"],
		configDB["POSTGRES_PASSWORD"],
		configDB["POSTGRES_HOST"],
		configDB["POSTGRES_PORT"],
		configDB["POSTGRES_DATABASE"])

	Log.Debugf("pgxconn: %s", pgxConn)

	PGDB, err = pgxpool.New(context.Background(), pgxConn)
	if err != nil {
		Log.Fatalf("database init failed: %v", err)
	}
	Log.Info("postgress database connection successfull")
	Log.Debugf("FROM postgresConnect GetNow: %s", fmt.Sprint(GetNow(PGDB)))
}

func GetNow(PGDB *pgxpool.Pool) (string, error) {
	var time time.Time

	err := PGDB.QueryRow(context.Background(), "select NOW()").Scan(&time)
	if err != nil {
		Log.Error("can not receive responce from db:", err)
		return "", err
	}
	return time.String(), err
}

func MessageCreate(message string) (id int64, err error) {
	stmt := `INSERT INTO messages (message) VALUES ($1) RETURNING id`

	err = PGDB.QueryRow(context.TODO(), stmt, message).Scan(&id)
	if err != nil {
		Log.Errorf("postgres database error during writeMessage process: %v ", err)
	}
	Log.Infof("message id=%d inserted to DB", id)
	return
}

// var name string
// var weight int64
// err := conn.QueryRow("select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
