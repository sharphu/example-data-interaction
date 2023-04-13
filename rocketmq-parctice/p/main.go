package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2/primitive"

	"awesomeProject1/pkg"
)

func main() {
	fmt.Println("p...")
	//test1()
	test2()
}

// case1 | case3
func test1() {
	producer, err := pkg.NewProducer()
	if err != nil {
		panic(err)
	}
	cnt := 0
	tag := ""
	for {
		time.Sleep(5 * time.Second)
		cnt++
		msg := primitive.NewMessage("wstopic1", []byte(fmt.Sprintf("p cnt: %d", cnt)))

		if cnt%2 == 0 {
			tag = "c1"
		} else {
			tag = "c2"
		}

		msg.WithTag(tag)
		fmt.Println(msg.String())
		sync, err := producer.SendSync(context.Background(), msg)
		if err != nil {
			panic(err)
		}
		fmt.Println(sync.String())
	}
}

// case2
func test2() {
	producer, err := pkg.NewProducer()
	if err != nil {
		panic(err)
	}
	cnt := 0
	tag := "c1"
	topic := ""
	for {
		time.Sleep(5 * time.Second)
		cnt++

		if cnt%2 == 0 {
			topic = "wstopic1"
		} else {
			topic = "wstopic2"
		}

		msg := primitive.NewMessage(topic, []byte(fmt.Sprintf("p cnt: %d", cnt)))

		msg.WithTag(tag)
		fmt.Println(msg.String())
		sync, err := producer.SendSync(context.Background(), msg)
		if err != nil {
			panic(err)
		}
		fmt.Println(sync.String())
	}
}
