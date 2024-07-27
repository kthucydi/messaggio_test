package kafkahandle

import (
	"strconv"
)

func init() {
	var err error

	// Create Producer and Consumer (with fatalExit on error)
	Kafka.createProduser()
	Kafka.createConsumer()

	// Set data from config
	Kafka.TopicProduser = (cfg["KAFKA_TOPIC_PRODUCER"])
	Kafka.TopicConsumer = (cfg["KAFKA_TOPIC_CONSUMER"])
	if Kafka.SendTimes, err = strconv.Atoi(cfg["KAFKA_SEND_TIMES"]); err != nil {
		Log.Warnf("Failed to set KafkaSendTimes for producer (value set to default = 3): %s\n", err)
		Kafka.SendTimes = 3
	}
}
