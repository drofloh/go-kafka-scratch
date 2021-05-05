package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := kafka.ConfigMap{}
	config.SetKey("bootstrap.servers", "localhost:9092")

	adm, err := kafka.NewAdminClient(&config)
	if err != nil {
		fmt.Println("Failed to create admin client: ", err)
	}
	defer adm.Close()
	// fmt.Println(adm)
	topics, err := adm.GetMetadata(nil, true, 100)
	if err != nil {
		fmt.Println("Failed to list topics: ", err)
	}
	fmt.Println("-------------------------------------------")

	for _, topic := range topics.Topics {
		fmt.Printf("%+v\n", topic.Topic)
	}
	fmt.Println("-------------------------------------------")

	for _, broker := range topics.Brokers {
		fmt.Printf("%+v\n", broker.Host)
		fmt.Printf("%+v\n", broker.ID)
		fmt.Printf("%+v\n", broker.Port)
	}

	fmt.Println("-------------------------------------------")

	// Contexts are used to abort or limit the amount of time
	// the Admin call blocks waiting for a result.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dur, _ := time.ParseDuration("20s")
	resourceType, err := kafka.ResourceTypeFromString("topic")
	if err != nil {
		fmt.Printf("Invalid resource type: %s\n", "topic")
		os.Exit(1)
	}
	resourceName := "testtopic3"
	// Ask cluster for the resource's current configuration
	results, err := adm.DescribeConfigs(
		ctx,
		[]kafka.ConfigResource{
			{
				Type: resourceType,
				Name: resourceName,
			},
		},
		kafka.SetAdminRequestTimeout(dur),
	)
	if err != nil {
		fmt.Printf("Failed to DescribeConfigs(%s, %s): %s\n", resourceType, resourceName, err)
		os.Exit(1)
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s %s: %s:\n", result.Type, result.Name, result.Error)
		for _, entry := range result.Config {
			// Truncate the value to 60 chars, if needed, for nicer formatting.
			fmt.Printf("%60s = %-60.60s   %-20s Read-only:%v Sensitive:%v\n",
				entry.Name, entry.Value, entry.Source,
				entry.IsReadOnly, entry.IsSensitive)
		}

	}
}
