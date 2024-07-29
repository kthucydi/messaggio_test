package main

import (
	api "messaggio_test/internal/api"        // set Routes and handlers for server
	config "messaggio_test/internal/config"  // get config data from .env and environment - .env priority is highter
	kh "messaggio_test/internal/kafkahandle" // init main kafka struct, create produser and consumer
	migrations "messaggio_test/migrations"   // set migrations
	"time"

	_ "github.com/kthucydi/bs_go_daemon"      // for run program in daemon mode
	server "github.com/kthucydi/bs_go_server" // http server
)

var cfg = *config.Cfg

func main() {
	migrations.Run(cfg["MIGRATIONS"], cfg["DB_CONN"])
	time.Sleep(time.Second * 40)
	go kh.Kafka.RunProducer()
	go kh.Kafka.RunConsumer()
	server.BackServer.Run(cfg["SERVER"], api.API)
	// server.BackServer.RunGracefullShutdown(cfg["SERVER"], api.API)
}
