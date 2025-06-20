package main

import (
	"errors"
	"flag"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

var (
	brokers = flag.String("brokers", "", "The Kafka brokers to connect to, as a comma separated list")
	topics  = flag.String("topics", "", "The Kafka topic to use")
	verbose = flag.Bool("verbose", false, "Turn on Sarama logging")
)

func main() {
	flag.Parse()

	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s\n", strings.Join(brokerList, ", "))

	err := testSyncProducer(brokerList, *topics)
	if err != nil {
		log.Fatalln(err)
	}
	//
	//err = testAsyncProducer(brokerList, *topic)
	//if err != nil {
	//	log.Fatalln(err)
	//}
}

func testSyncProducer(brokers []string, topic string) error {
	conf := sarama.NewConfig()
	// ack: -1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	// 分区器
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	// 压缩
	conf.Producer.Compression = sarama.CompressionSnappy
	// 重试次数
	conf.Producer.Retry.Max = 3
	// 开启幂等性
	conf.Producer.Idempotent = true
	conf.Net.MaxOpenRequests = 1
	conf.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("test message: " + time.Now().Format("2006-01-02 13:04:05")),
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
				Topic: topic,
				Value: sarama.StringEncoder("test message " + strconv.Itoa(i)),
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
