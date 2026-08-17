package queue

import "context"

type Queue interface {
	Setup() error
	Publish(ctx context.Context, payload any) error
	NewConsumer(ctx context.Context, name string) (Consumer, error)
	Close() error
}

type Consumer interface {
	Close() error
	Run(ctx context.Context, concurrency int, handler func(ctx context.Context, msg Delivery)) error
	Name() string
}

type Delivery interface {
	Body() []byte
	Ack() error
	Nack(requeue bool) error
}
