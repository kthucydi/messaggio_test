package kafkahandle

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (kfk *KafkaData) SendMessage(message string, id int) {

	for i := 0; i < kfk.SendTimes; i++ {
		err := kfk.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &kfk.TopicProduser, Partition: kafka.PartitionAny},
			Value:          []byte(message),
			Headers:        []kafka.Header{{Key: "id", Value: []byte(fmt.Sprint("987"))}},
		}, nil)

		if err != nil {
			if err.(kafka.Error).Code() == kafka.ErrQueueFull {
				// Producer queue is full, wait 1s for messages
				// to be delivered then try again.
				Log.Warnf("Failed to produce message: %v\n", err)
				time.Sleep(time.Second)
				continue
			}
			Log.Warnf("Failed to produce message: %v\n", err)
		}
		break
	}
}

func CreateKafkaProduser() {
	var err error
	Kafka.Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg["KAFKA_BOOTSTRAP_SERVERS"]})
	if err != nil {
		Log.Fatalf("Failed to create Kafka producer: %s\n", err)
	}
}

// Listen to all the events on the default events channel
func (kfk *KafkaData) RunProducer() {

	defer kfk.CloseProducer()

	for e := range kfk.Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			// The message delivery report, indicating success or
			// permanent failure after retries have been exhausted.
			// Application level retries won't help since the client
			// is already configured to do that.
			m := ev
			if m.TopicPartition.Error != nil {
				Log.Warnf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				Log.Infof("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		case kafka.Error:
			// Generic client instance-level errors, such as
			// broker connection failures, authentication issues, etc.
			//
			// These errors should generally be considered informational
			// as the underlying client will automatically try to
			// recover from any errors encountered, the application
			// does not need to take action on them.
			Log.Warnf("Error: %v\n", ev)
		default:
			Log.Infof("Ignored event: %s\n", ev)
		}
	}
}

func (kfk *KafkaData) CloseProducer() {
	// Flush and close the producer and the events channel
	for kfk.Producer.Flush(10000) > 0 {
		Log.Info("Still waiting to flush outstanding messages\n")
	}
	kfk.Producer.Close()
	Log.Info("kafka produser closed")
}
