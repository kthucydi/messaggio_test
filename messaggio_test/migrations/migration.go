// This is custom goose binary with sqlite3 support only.

package migrations

import (
	"context"
	"embed"
	"fmt"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	logging "github.com/kthucydi/bs_go_logrus"
)

//go:embed *.sql
var embedFS embed.FS
var Log = &logging.Log

func Run(mgCfg map[string]string, DBData map[string]string) {

	// Check for migration necessity
	if mgCfg["WITHOUT_MIGRATIONS"] == "true" {
		Log.Info("goose: run without migration")
		return
	}

	// Set embed filesystem for goose
	goose.SetBaseFS(embedFS)
	dir := "."

	//Set our Logger for message from goose
	goose.SetLogger(Log)

	//connect to DB
	pgxConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		DBData["POSTGRES_USER"],
		DBData["POSTGRES_PASSWORD"],
		DBData["POSTGRES_HOST"],
		DBData["POSTGRES_PORT"],
		DBData["POSTGRES_DATABASE"])

	Log.Debug("migration pgxConn: ", pgxConn)

	db, err := goose.OpenDBWithDriver("pgx", pgxConn)
	if err != nil {
		Log.Errorf("goose error: failed to open DB: %v\n", err)
	} else {
		Log.Infof("goose connection to postgress database %s successfull", DBData["POSTGRES_DATABASE"])
	}

	// close DB on exit
	defer func() {
		if err := db.Close(); err != nil {
			Log.Errorf("goose error: failed to close DB %s: %v\n", DBData["POSTGRES_DATABASE"], err)
		}
	}()

	//Run goose migrations command
	args := strings.Split(mgCfg["MIGRATIONS_ARGS"], ";")
	err = goose.RunContext(context.Background(), mgCfg["MIGRATIONS_COMMAND"], db, dir, args...)
	if err != nil {
		Log.Errorf("goose error: command:%v: %v", mgCfg["MIGRATIONS_COMMAND"], err)
	}

	// Exit from program if we need only migrations without running program
	if mgCfg["MIGRATIONS_ONLY"] == "true" {
		Log.Infof("Exit by flag 'm-only'")
		os.Exit(0)
	}
}
