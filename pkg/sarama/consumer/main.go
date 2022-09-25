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

var (
	brokers  = flag.String("brokers", "", "The Kafka brokers to connect to, as a comma separated list")
	group    = flag.String("group", "", "Kafka consumer group definition")
	topics   = flag.String("topics", "", "The Kafka topic to use")
	assignor = flag.String("assignor", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	oldest   = flag.Bool("oldest", true, "Kafka consumer consume initial offset from oldest")
	verbose  = flag.Bool("verbose", false, "Turn on Sarama logging")
)

func main() {
	flag.Parse()

	if len(*brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(*topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(*group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s\n", strings.Join(brokerList, ", "))

	log.Println("Starting a new Sarama consumer")

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	err := testConsumer(brokerList, *topic, *offset)
	if err != nil {
		log.Fatalln(err)
	}
}

func testConsumer(brokerList []string, topic string, offset int64) error {
	conf := sarama.NewConfig()
	// 手动提交
	conf.Consumer.Offsets.AutoCommit.Enable = false
	partition := 1

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
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			logString := fmt.Sprintf(
				"Topic: %s, Partition: %d, Offset: %d, Key: %s, Value: %s, Timestamp: %s",
				msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value, msg.Timestamp,
			)
			log.Println(logString)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
	log.Printf("Consumed: %d\n", consumed)

	return nil
}
