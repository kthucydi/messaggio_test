package kafkahandle

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	logging "github.com/kthucydi/bs_go_logrus"
	"messaggio_test/internal/config"
	"messaggio_test/internal/pgstorage"
)

var (
	PG    = &pgstorage.PGDB
	Log   = &logging.Log
	Kafka = &KafkaData{}
	cfg   = config.Cfg.Kafka
)

type KafkaData struct {
	Producer      *kafka.Producer
	Consumer      *kafka.Consumer
	TopicProduser string
	TopicConsumer string
	SendTimes     int
}
