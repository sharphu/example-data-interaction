package pkg

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func NewProducer() (rocketmq.Producer, error) {
	opts := []producer.Option{
		producer.WithNameServer([]string{"10.0.102.10:9876"}),
		//producer.WithGroupName(""),
	}

	p, err := rocketmq.NewProducer(opts...)

	if err != nil {
		return nil, err
	}
	err = p.Start()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func NewConsumer() (rocketmq.PushConsumer, error) {
	opts := []consumer.Option{
		consumer.WithNameServer([]string{"10.0.102.10:9876"}),
		consumer.WithGroupName("wsgroup"),
	}

	c, err := rocketmq.NewPushConsumer(opts...)
	if err != nil {
		return nil, err
	}
	return c, err
}
