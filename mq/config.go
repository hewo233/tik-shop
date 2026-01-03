package mq

import (
	"github.com/IBM/sarama"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/joho/godotenv"
	"os"
)

var Brokers []string
var TopicOrderCreate string
var ConsumerGroup string

type config struct {
	Brokers          []string
	TopicOrderCreate string
	ConsumerGroup    string
}

func readConfig() (*config, error) {
	// read from .env
	if err := godotenv.Load(consts.RPCKafkaEnvPath); err != nil {
		return nil, err
	}
	broker := os.Getenv("KAFKA_HOST") + os.Getenv("KAFKA_PORT")
	topicOrderCreate := os.Getenv("KAFKA_TOPIC_ORDER_CREATE")
	consumerGroup := os.Getenv("KAFKA_CONSUMER_GROUP")

	return &config{
		Brokers:          []string{broker},
		TopicOrderCreate: topicOrderCreate,
		ConsumerGroup:    consumerGroup,
	}, nil
}

// InitSaramaConfig 初始化通用的 Sarama 配置
func InitSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 读取最早的
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
