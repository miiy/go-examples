package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main()  {
	brokers := flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
	topic := flag.String("topic", "", "The Kafka topic to use")
	offset := flag.Int64("offset", sarama.OffsetNewest, "The kafka offset")
	flag.Parse()

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s\n", strings.Join(brokerList, ", "))

	if *topic == "" {
		log.Fatalln("topic is required")
	}

	err := testConsumer(brokerList, *topic, *offset)
	if err != nil {
		log.Fatalln(err)
	}
}

func testConsumer(brokerList []string, topic string, offset int64) error {
	conf := sarama.NewConfig()
	partition := 0

	consumer, err := sarama.NewConsumer(brokerList, conf)
	if err != nil {
		return err
	}
	log.Println("consumer created")

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("commence consuming")
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), offset)
	if err != nil {
		return err
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
	ConsumerLoop: for  {
		select {
		case msg := <-partitionConsumer.Messages():
			logString := fmt.Sprintf(
				"Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s, Timestamp: %s",
				msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value, msg.Timestamp,
			)
			log.Println(logString)
			consumed ++
		case <-signals:
			break ConsumerLoop
		}
	}
	log.Printf("Consumed: %d\n", consumed)

	return nil
}