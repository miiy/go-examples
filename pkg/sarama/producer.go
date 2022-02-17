package main

import (
	"errors"
	"flag"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main()  {
	brokers := flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
	topic := flag.String("topic", "", "The Kafka topic to use")
	flag.Parse()

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s\n", strings.Join(brokerList, ", "))

	if *topic == "" {
		log.Fatalln("topic is required")
	}

	err := testSyncProducer(brokerList, *topic)
	if err != nil {
		log.Fatalln(err)
	}

	err = testAsyncProducer(brokerList, *topic)
	if err != nil {
		log.Fatalln(err)
	}
}

func testSyncProducer(brokers []string, topic string) error {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewHashPartitioner
	conf.Producer.Return.Successes = true


	syncProducer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder("test message"),
	}

	partition, offset, err := syncProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	log.Printf("wrote message at partition: %d, offset: %d", partition, offset)
	defer func() {
		if err := syncProducer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	return nil
}

func testAsyncProducer(brokers []string, topic string) error {
	conf := sarama.NewConfig()

	asyncProducer, err := sarama.NewAsyncProducer(brokers, conf)
	if err != nil {
		return errors.New("failed to start sarama producer:" + err.Error())
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			asyncProducer.Input() <- &sarama.ProducerMessage{
				Topic:     topic,
				Value:     sarama.StringEncoder("test message " + strconv.Itoa(i)),
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	asyncProducer.Successes()

	if err := asyncProducer.Close(); err != nil {
		log.Println("Failed to shut down access log producer cleanly", err)
	}

	return nil
}