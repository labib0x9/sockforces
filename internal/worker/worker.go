package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/labib0x9/sockforces/internal/domain/queue"
	"github.com/labib0x9/sockforces/internal/domain/worker"
)

type testWorker struct {
	rmq queue.Queue
}

func NewWorker(rmq queue.Queue) worker.Worker {
	return &testWorker{rmq: rmq}
}

// test to handle the error
func (t *testWorker) Run(ctx context.Context, name string, concurrency int) error {
	consumer, err := t.rmq.NewConsumer(ctx, name)
	if err != nil {
		return err
	}
	defer consumer.Close()
	return consumer.Run(ctx, concurrency, t.handle)
}

func (t *testWorker) handle(ctx context.Context, msg queue.Delivery) {
	var msgD queue.PushMessage
	if err := json.Unmarshal(msg.Body(), &msgD); err != nil {
		fmt.Println("ERROR", err)
		msg.Nack(false)
		return
	}

	// need to create testing container, clone the repository, run tests, sse output
	fmt.Println("WORKER:::", msgD.Id, msgD.Username)
	msg.Ack()
}
