package main

import (
	"context"

	"github.com/saskaradit/kafka-go-consume-produce.git/messagebus"
)

func main() {
	// Creates a new context
	ctx := context.Background()

	// produce messages in a new go routine
	// so that produce and consume does not block
	go messagebus.Produce(ctx)
	messagebus.Consume(ctx)
}
