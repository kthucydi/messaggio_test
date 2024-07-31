package main

import (
	_ "github.com/kthucydi/bs_go_daemon"                          // for run program in daemon mode
	khandler "second_handler/internal/kafkahandle_second_handler" // init main kafka struct, create produser and consumer
)

func main() {

	go khandler.Kafka.RunProducer()
	go khandler.RunGetter()
	khandler.Kafka.RunConsumer()

}
