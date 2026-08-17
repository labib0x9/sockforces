package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labib0x9/sockforces/config"
	queue_domain "github.com/labib0x9/sockforces/internal/domain/queue"
	"github.com/rabbitmq/amqp091-go"
	mq "github.com/rabbitmq/amqp091-go"
)

var (
	exchange = "sockforces"
	eType    = "direct"
	route    = "sockforces.action.push"
	queue    = "submissions.test"
)

type delivery struct {
	d mq.Delivery
}

func (d *delivery) Body() []byte {
	return d.d.Body
}

func (d *delivery) Ack() error {
	return d.d.Ack(false)
}

func (d *delivery) Nack(requeue bool) error {
	return d.d.Nack(false, requeue)
}

type consumer struct {
	msg  <-chan mq.Delivery
	ch   *mq.Channel
	name string
}

func (c *consumer) Name() string {
	return c.name
}

// set prefetch, semaphore to handle concurrent request upto concurrency
// executes the given handler function
func (c *consumer) Run(ctx context.Context, concurrency int, handler func(ctx context.Context, msg queue_domain.Delivery)) error {
	if err := c.ch.Qos(concurrency, 0, false); err != nil {
		return err
	}
	sem := make(chan struct{}, concurrency)
	for {
		select {
		case <-ctx.Done():
			return nil
		case d, ok := <-c.msg:
			if !ok {
				return fmt.Errorf("channel closed!")
			}
			sem <- struct{}{}
			go func(ctx context.Context, msg queue_domain.Delivery) {
				defer func() { <-sem }()
				handler(ctx, msg)
			}(ctx, &delivery{d: d})
		}
	}
}

func (c *consumer) Close() error {
	return c.ch.Close()
}

type client struct {
	conn *amqp091.Connection
}

func NewClient(cnf *config.RabbittMQ) queue_domain.Queue {
	conn, err := mq.DialConfig(
		fmt.Sprintf("amqp://%s:%s@%s", cnf.User, cnf.Pass, cnf.Addr),
		mq.Config{
			ChannelMax: 25,
		},
	)
	if err != nil {
		panic(err)
	}
	return &client{conn: conn}
}

func (c *client) Setup() error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(exchange, eType, true, false, false, false, nil); err != nil {
		return err
	}

	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	if err := ch.QueueBind(q.Name, route, exchange, false, nil); err != nil {
		return err
	}

	return nil
}

// name = consumer identifier
// close consumer or channel leak
func (c *client) NewConsumer(ctx context.Context, name string) (queue_domain.Consumer, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}

	delivery, err := ch.ConsumeWithContext(ctx, queue, name, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	_ = delivery
	return &consumer{
		msg:  delivery,
		ch:   ch,
		name: name,
	}, nil
}

func (c *client) Publish(ctx context.Context, payload any) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, exchange, route, false, false, mq.Publishing{
		ContentType: "application/json",
		Timestamp:   time.Now(),
		Body:        body,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Close() error {
	return c.conn.Close()
}
