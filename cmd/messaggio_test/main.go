package main

import (
	_ "github.com/kthucydi/bs_go_daemon"      // for run program in daemon mode
	server "github.com/kthucydi/bs_go_server" // http server
	api "messaggio_test/internal/api"         // set Routes for server
	config "messaggio_test/internal/config"   // get config data from .env and environment - .env priority is highter
	kh "messaggio_test/internal/kafkahandle"  // init main kafka struct, create produser and consumer
	migrations "messaggio_test/migrations"    // set migrations
)

func main() {
	migrations.Run(config.Cfg.Migration, config.Cfg.DBConnectionData)
	go kh.Kafka.RunProducer()
	go kh.Kafka.RunConsumer()
	server.BackServer.Run(config.Cfg.Data, api.API)
}
