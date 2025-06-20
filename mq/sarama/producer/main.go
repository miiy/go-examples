package main

import (
	"flag"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	brokers = flag.String("brokers", "", "The Kafka brokers to connect to, as a comma separated list")
	verbose = flag.Bool("verbose", false, "Turn on Sarama logging")
)

func main() {
	flag.Parse()

	if *verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	if *brokers == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	testSyncProducer(brokerList)

}

func testSyncProducer(brokers []string) {
	producer, cleanUp := newSyncProducer(brokers)
	defer cleanUp()

	topic := "test"

	i := 0
	for true {
		msg := fmt.Sprintf("test_message: %d", i)

		// We are not setting a message key, which means that all messages will
		// be distributed randomly over the different partitions.
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(msg),
		})
		if err != nil {
			log.Printf("failed to send message to %s: %s", topic, err)
		}
		i++
		log.Printf("wrote message at partition: %d, offset: %d", partition, offset)
		time.Sleep(time.Second * 2)
	}
}

func testAsyncProducer(brokers []string) {
	producer, cleanUp := newAsyncProducer(brokers)
	defer cleanUp()

	topic := "test"

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			producer.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder("test_message " + strconv.Itoa(i)),
			}
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	producer.Successes()
}

func newSyncProducer(brokers []string) (sarama.SyncProducer, func()) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	//// 分区器
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//// 重试次数
	//// 开启幂等性
	//config.Producer.Idempotent = true
	//config.Net.MaxOpenRequests = 1

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	cleanUp := func() {
		if err = producer.Close(); err != nil {
			log.Println("Failed to shut down data kafka producer cleanly", err)
		}
	}
	return producer, cleanUp
}

func newAsyncProducer(brokers []string) (sarama.AsyncProducer, func()) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	cleanUp := func() {
		if err = producer.Close(); err != nil {
			log.Println("Failed to shut down data kafka producer cleanly", err)
		}
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer, cleanUp
}
