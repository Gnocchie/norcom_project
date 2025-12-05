package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/Gnocchie/norcom_project/internal/config"
	"github.com/Gnocchie/norcom_project/internal/model"
	"github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	Publish(event model.FileEvent) error
	Close()
}

type RabbitPublisher struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	queue   *amqp091.Queue
}

func NewRabbitPublisher(cfg config.Config) (*RabbitPublisher, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s/",
		cfg.RabbitmqUser, cfg.RabbitmqPass, cfg.RabbitmqHost)

	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("dial error: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel error: %w", err)
	}

	q, err := ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, fmt.Errorf("queue declare error: %w", err)
	}
	return &RabbitPublisher{conn: conn, channel: ch, queue: &q}, nil
}

func (p *RabbitPublisher) Publish(event model.FileEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp091.Persistent,
		},
	)
}

func (p *RabbitPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}
