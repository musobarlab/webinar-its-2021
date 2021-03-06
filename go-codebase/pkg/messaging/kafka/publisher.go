package kafka

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	logrus "github.com/sirupsen/logrus"
	"gitlab.com/Wuriyanto/go-codebase/config"
)

//Publisher struct
type Publisher struct {
	producer sarama.SyncProducer
}

//NewPublisher constructor of PublisherImpl
func NewPublisher(appName string) *Publisher {
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	cfg := sarama.NewConfig()
	cfg.ClientID = appName
	cfg.Producer.Retry.Max = 10
	cfg.Producer.Retry.Backoff = 10 * time.Second
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Timeout = 10 * time.Second
	cfg.Producer.Compression = sarama.CompressionSnappy
	cfg.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer(addresses, config)

	// sync producer
	brokers := strings.Split(config.GlobalEnv.Kafka.Brokers, ",")
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		panic(err)
	}

	return &Publisher{producer: producer}
}

//Publish function
func (publisher *Publisher) Publish(topic string, messageKey string, message []byte) error {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(messageKey),
		Value: sarama.ByteEncoder(message),
	}
	p, o, err := publisher.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	logrus.Println("Partition ", p)
	logrus.Println("Offset ", o)

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{
	return nil
}
