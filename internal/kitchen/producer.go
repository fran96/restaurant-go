package kitchen

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	contracts "github.com/fran96/restaurant-go/contracts/avro"
	"github.com/fran96/restaurant-go/internal/util"
)

func produce(orderID string) error {
	flag.Parse()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.KafkaServerAddress})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {

		for e := range p.Events() {
			fmt.Printf("produce..")
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "orderCompleted"
	orderCompleted := contracts.OrderCompleted{
		OrderID: orderID,
	}

	val, err := json.Marshal(orderCompleted)
	if err != nil {
		return err
	}

	// add integration with schema registry
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(val),
	}, nil)
	fmt.Println("ORDER SUCCESS, ORDERID: ", orderID)

	p.Flush(15 * 1000)

	return nil
}
