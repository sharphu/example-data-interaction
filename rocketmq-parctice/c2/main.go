package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"

	"awesomeProject1/pkg"
)

func main() {
	fmt.Println("c2...")
	ch := make(chan struct{})
	c, err := pkg.NewConsumer()
	if err != nil {
		panic(err)
	}
	err = c.Subscribe("wstopic1", consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "c2",
	}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, m := range ext {
			info := fmt.Sprintf("c2===topic[%s]===tag:[%s]=======body:[%s]", m.Topic, m.GetTags(), string(m.Body))
			fmt.Println(info)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(err)
	}

	err = c.Start()
	if err != nil {
		panic(err)
	}

	<-ch

}
