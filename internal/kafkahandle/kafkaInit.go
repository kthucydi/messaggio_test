package kafkahandle

import (
	"strconv"
)

func init() {
	var err error

	// Create Producer and Consumer (with fatalExit on error)
	CreateKafkaProduser()
	CreateKafkaConsumer()

	// Set data from config
	if Kafka.SendTimes, err = strconv.Atoi(cfg["KAFKA_SEND_TIMES"]); err != nil {
		Log.Warnf("Failed to set KafkaSendTimes for producer (value set to default = 3): %s\n", err)
		Kafka.SendTimes = 3
	}
	Log.Printf("Created Kafka Producer %v\n", Kafka.Producer)
	Kafka.SendTimes, err = strconv.Atoi(cfg["KAFKA_SEND_TIMES"])
	Kafka.TopicProduser = (cfg["KAFKA_TOPIC_PRODUCER"])
	Kafka.TopicConsumer = (cfg["KAFKA_TOPIC_CONSUMER"])

	// Run goroutine for receiving kafka event about sending messages
	// go Run(Kafka)

}
