package mq

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog" // 使用 Kitex 的日志库
)

var Producer *OrderProducer

type OrderProducer struct {
	client sarama.SyncProducer
}

// InitProducer 初始化生产者
func InitProducer() {
	config := InitSaramaConfig()

	// 创建同步生产者
	p, err := sarama.NewSyncProducer(Brokers, config)
	if err != nil {
		klog.Fatalf("Failed to start Sarama producer: %v", err)
	}

	Producer = &OrderProducer{
		client: p,
	}
	klog.Infof("Kafka Producer started successfully")
}

// SendOrderMsg 发送订单消息
func (p *OrderProducer) SendOrderMsg(ctx context.Context, msg OrderMessage) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: TopicOrderCreate,
		Value: sarama.ByteEncoder(bytes),
		// Key:   sarama.StringEncoder(fmt.Sprintf("%d", msg.OrderId)), // 可选：设置Key保证同一订单有序
	}

	// 3. 发送消息
	partition, offset, err := p.client.SendMessage(kafkaMsg)
	if err != nil {
		klog.Errorf("Failed to send message: %v", err)
		return err
	}

	klog.Debugf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}

// Close 关闭连接
func (p *OrderProducer) Close() {
	if p.client != nil {
		p.client.Close()
	}
}
