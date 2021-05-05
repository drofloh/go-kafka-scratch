package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	config := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "kafkago-group",
		Topic:    "testtopic1",
		MinBytes: 10e3,
		MaxBytes: 10e6,
		// MaxWait:  100 * time.Millisecond,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()
	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
