package kafkahandle

// consumer_example implements a consumer using the non-channel Poll() API
// to retrieve messages and events.

import (
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
			if e.Headers != nil {
				Log.Infof("%% Headers: %v\n", e.Headers)
			}

			// We can store the offsets of the messages manually or let
			// the library do it automatically based on the setting
			// enable.auto.offset.store. Once an offset is stored, the
			// library takes care of periodically committing it to the broker
			// if enable.auto.commit isn't set to false (the default is true).
			// By storing the offsets manually after completely processing
			// each message, we can ensure atleast once processing.
			_, err := kfk.Consumer.StoreMessage(e)
			if err != nil {
				Log.Errorf("%% Error storing offset after message %s:\n", e.TopicPartition)
			}
		case kafka.Error:
			// Errors should generally be considered
			// informational, the client will try to
			// automatically recover.
			// But in this example we choose to terminate
			// the application if all brokers are down.
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

func CreateKafkaConsumer() {
	var err error

	Kafka.Consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg["KAFKA_BOOTSTRAP_SERVERS"],
		"broker.address.family":    "v4",
		"group.id":                 "messaggio-consumer-back",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": true,
	})
	if err != nil {
		Log.Fatalf("Failed to create consumer: %s\n", err)
	}

	Log.Infof("Created Consumer %v\n", Kafka.Consumer)
}

func (kfk *KafkaData) CloseConsumer() {
	if err := kfk.Consumer.Close(); err != nil {
		Log.Errorf("can not close kafka consumer: %v", err)
	}
	Log.Info("kafka consumer closed successfuly")
}
