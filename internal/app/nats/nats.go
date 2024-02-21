package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/pu4mane/NATSOrderViewer/internal/app/model"
)

func ConnectToNatsStreaming(clusterID, clientID string) stan.Conn {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming: %v", err)
	}

	return sc
}

func PublishOrder(sc stan.Conn, order *model.Order) {
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Error marshalling Order: %v", err)
	}

	err = sc.Publish("order", orderData)
	if err != nil {
		log.Fatalf("Error publishing to channel: %v", err)
	}
	log.Println("Successfully published to channel")
}

func SubscribeToOrder(sc stan.Conn) stan.Subscription {
	sub, err := sc.Subscribe("order", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	log.Println("Successfully subscribed to channel")
	return sub
}
