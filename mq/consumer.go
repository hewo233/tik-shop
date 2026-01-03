package mq

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
)

type OrderConsumerHandler struct {
	// 这里可以注入 Service 层，用于调用实际的下单逻辑
	// ProcessLogic func(msg *OrderMessage) error
}

// Setup 在会话开始前运行
func (h *OrderConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup 在会话结束后运行
func (h *OrderConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

type MqSqlManageImpl struct {

}

type MqSqlManage interface {

}

// ConsumeClaim 核心消费逻辑：循环读取 Channel
func (h *OrderConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		klog.Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic)

		var orderMsg OrderMessage
		err := json.Unmarshal(msg.Value, &orderMsg)
		if err != nil {
			klog.Errorf("Unmarshal error: %v", err)
			session.MarkMessage(msg, "")
			continue
		}

		// create order in db


		session.MarkMessage(msg, "")
	}
	return nil
}

// InitConsumer 启动消费者（通常在 main 中作为 goroutine 启动）
func InitConsumer(ctx context.Context) {
	config := InitSaramaConfig()

	// 创建消费者组
	consumerGroup, err := sarama.NewConsumerGroup(Brokers, ConsumerGroup, config)
	if err != nil {
		klog.Fatalf("Error creating consumer group client: %v", err)
	}

	handler := &OrderConsumerHandler{}

	// 启动消费循环
	go func() {
		defer consumerGroup.Close()
		for {
			// Consume 会阻塞，直到发生错误或 ctx 被取消
			if err := consumerGroup.Consume(ctx, []string{TopicOrderCreate}, handler); err != nil {
				klog.Errorf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	klog.Infof("Kafka Consumer started")
}
