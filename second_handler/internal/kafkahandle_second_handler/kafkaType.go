package kafkahandle

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	logging "github.com/kthucydi/bs_go_logrus"
	config "second_handler/internal/config_second_handler"
)

var (
	Log   = &logging.Log
	Kafka = &KafkaData{}
	cfg   = (*config.Cfg)["KAFKA"]
)

type KafkaData struct {
	Producer      *kafka.Producer
	Consumer      *kafka.Consumer
	TopicProduser string
	TopicConsumer string
	SendTimes     int
}
