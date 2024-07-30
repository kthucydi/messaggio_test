package kafkahandle

import (
	"messaggio_test/internal/getter"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (kfk *KafkaData) RunConsumer() {

	topics := []string{kfk.TopicConsumer}
	if err := kfk.Consumer.SubscribeTopics(topics, nil); err != nil {
		Log.Fatalf("can not suscribe kafka consumer to topics: %v", err)
	}

	defer kfk.CloseConsumer()

	run := true
	for run {
		ev := Kafka.Consumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:

			// Process the message received.
			Log.Infof("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// Send getted message to chanel fo processed
			getter.ResponseCh <- string(e.Value)

			// Proceessed header
			if e.Headers != nil {
				Log.Infof("%% Headers: %v\n", e.Headers)
			}
			// Storing kafka offset
			_, err := kfk.Consumer.StoreMessage(e)
			if err != nil {
				Log.Errorf("%% Error storing offset after message %s:\n", e.TopicPartition)
			}
		case kafka.Error:
			Log.Errorf("%% Error: %v: %v\n", e.Code(), e)
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		case kafka.OffsetsCommitted:
			Log.Infof("OffsetsCommitted %v\n", e)
		default:
			Log.Infof("Ignored %v\n", e)
		}
	}
}

func (kfk *KafkaData) createConsumer() {
	var err error

	kfk.Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg["KAFKA_BOOTSTRAP_SERVERS"],
		"broker.address.family":    "v4",
		"group.id":                 "messaggio-consumer-back",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})

	if err != nil {
		Log.Fatalf("Failed to create consumer: %s\n", err)
	} else {
		Log.Info("created kafka consumer")
	}
}

func (kfk *KafkaData) CloseConsumer() {
	if err := kfk.Consumer.Close(); err != nil {
		Log.Errorf("can not close kafka consumer: %v", err)
	} else {
		Log.Info("kafka consumer closed successfuly")
	}
}
