package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	brokers := kafka.TCP("localhost:9092")
	c := kafka.Client{
		Addr: brokers,
	}

	// resources := []kafka.DescribeConfigRequestResource{
	// 	kafka.DescribeConfigRequestResource{
	// 		ResourceType: 2,
	// 		ResourceName: "testtopic1",
	// 	},
	// }

	// configRequest := kafka.DescribeConfigsRequest{
	// 	Addr:      brokers,
	// 	Resources: resources,
	// }

	// resp, err := c.DescribeConfigs(context.TODO(), &configRequest)
	// if err != nil {
	// 	log.Fatalf("Failed : %+v \n", err)
	// }
	// fmt.Println(resp)

	rsp, err := c.Metadata(context.TODO(), &kafka.MetadataRequest{Addr: brokers})
	if err != nil {
		log.Fatalf("Failed : %+v \n", err)
	}

	// list brokers
	fmt.Println(rsp.Brokers)
	// fmt.Println(rsp.Controller.Host)
	// list all topics
	for _, topic := range rsp.Topics {
		fmt.Println(topic.Name)
	}
}
