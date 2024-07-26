// This is custom goose binary with sqlite3 support only.

package migrations

import (
	"context"
	"embed"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	logging "github.com/kthucydi/bs_go_logrus"
)

//go:embed *.sql
var embedFS embed.FS
var Log = &logging.Log

type Data struct {
	Only    bool
	Exit    bool
	Command string
	Args    []string
}

func Run(migrationData Data, DBData map[string]string) {

	// Check for migration necessity
	if migrationData.Exit {
		Log.Info("goose: run without migration")
		return
	}

	// Set embed filesystem for goose
	goose.SetBaseFS(embedFS)
	dir := "."

	//Set our Logger for message from goose
	goose.SetLogger(Log)

	//connect to DB
	db, err := goose.OpenDBWithDriver("pgx", DBData["pgxconn"])
	if err != nil {
		Log.Errorf("goose error: failed to open DB: %v\n", err)
	} else {
		Log.Infof("goose connection to postgress database %s successfull", DBData["dbPGStringPath"])
	}

	// close DB on exit
	defer func() {
		if err := db.Close(); err != nil {
			Log.Errorf("goose error: failed to close DB at %s: %v\n", DBData["dbPGStringPath"], err)
		}
	}()

	//Run goose migrations command
	if err := goose.RunContext(context.Background(), migrationData.Command, db, dir, migrationData.Args...); err != nil {
		Log.Errorf("goose error: command:%v: %v", migrationData.Command, err)
	}

	// Exit from program if we need only migrations without running program
	if migrationData.Only {
		Log.Printf("Exit by flag 'm-only'")
		os.Exit(0)
	}
}
