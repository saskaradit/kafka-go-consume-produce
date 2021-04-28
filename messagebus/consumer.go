package messagebus

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "message-log"
	broker1Address = "localhost:8080"
	broker2Address = "localhost:8081"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the GrouipID identifies the consumer and prevents
	// duplicate messages

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topic,
		GroupID: "rad-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message" + err.Error())
		}

		fmt.Println("received: ", string(msg.Value))
	}
}
