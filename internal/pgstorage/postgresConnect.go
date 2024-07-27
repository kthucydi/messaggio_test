package pgstorage

import (
	"context"
	"fmt"
	"time"

	pgxscan "github.com/georgysavva/scany/v2/pgxscan"
	// pgx "github.com/jackc/pgx/v5"
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

type Statistic struct {
	All       int64 `db:"all_messages" json:"all_messages"`
	Processed int64 `db:"processed_messages" json:"processed_messages"`
}

func GetStatistic() (*Statistic, error) {
	stmt := `SELECT ( SELECT COUNT(*) FROM messages ) as all_messages, 
			( SELECT COUNT(*) FROM messages 
			WHERE processed = true) as processed_messages;`
	stat := &Statistic{}

	err := pgxscan.Get(context.TODO(), PGDB, stat, stmt)
	if err != nil {
		Log.Errorf("postgres database error during Read Statistic process: %v ", err)
	}

	Log.Infof("stat getting success: %v", stat)
	return stat, err
}

func UpdateProcessed(array []int) error {
	stmt := `UPDATE messages SET processed = true
			WHERE id = ANY ($1) AND processed = false;`
	array = []int{4, 5, 2} // ARRAY!!!
	// stmt := `UPDATE messages SET processed = true
	// 		WHERE id = ANY (@ids) AND processed = false;`
	// array := []int{4, 3, 2, 1}

	// batch := &pgx.Batch{}
	// args := pgx.NamedArgs{"ids": array}
	// batch.Queue(stmt, args)

	// results := PGDB.SendBatch(context.TODO(), batch)
	// _, err := results.Exec()
	// if err != nil {
	// 	Log.Errorf("postgres database error during update processed flag: %v ", err)
	// 	return err
	// }
	// defer results.Close()
	_, err := PGDB.Exec(context.TODO(), stmt, array)
	if err != nil {
		Log.Errorf("postgres database error during update processed flag: %v ", err)
	}

	Log.Info("update processed success")
	return err
}

// UPDATE messages
// SET processed = true
// WHERE id = ANY (ARRAY[1,6]) AND processed = false;
