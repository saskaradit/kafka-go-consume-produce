package messagebus

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(ctx context.Context) {
	i := 0

	// initialize writer with the address
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address, broker2Address},
		Topic:   topic,
	})

	for {
		// each kafka message will have a key
		// the key will be used to decide which partioion
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// Create an arbitray message payload
			Value: []byte("this is the mssage " + strconv.Itoa(i)),
		})

		if err != nil {
			panic("could not write message" + err.Error())
		}

		// Logs once the messages has been written
		fmt.Println("writes:", i)
		i++

		time.Sleep(time.Second)
	}
}
