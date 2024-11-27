package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/payments/application"
	"log"

	"github.com/streadway/amqp"
)

type PriceView struct {
	Cents uint `json:"cents"`
	Currency string `json:"currency"`
}

type OrderToProcessView struct {
	ID string `json:"id"`
	Price PriceView
}

type PaymentsInterface struct {
	conn *amqp.Connection
	queue amqp.Queue
	channel *amqp.Channel
	service application.PaymentsService
}

func NewPaymentsInterface(url string, queueName string, service application.PaymentsService) (PaymentsInterface, error) {
	conn, err := amqp.Dial(url);
	if err != nil {
		return PaymentsInterface{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return PaymentsInterface{}, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return PaymentsInterface{}, err
	}

	return PaymentsInterface{
		conn: conn,
		queue: q,
		channel: ch,
		service: service,
	}, nil
}

func (p PaymentsInterface) Run(ctx context.Context) error {
	msgs, err := p.channel.Consume(
		p.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.New("unable to consume messages");
	}

	done := ctx.Done();
	defer func() {
		if err := p.conn.Close(); err != nil {
			log.Print("Cannot close connection", err)
		}
		
		if err := p.channel.Close(); err != nil {
			log.Print("Cannot close channel", err)
		}
	}()

	for {
		select {
		case msg := <-msgs:
			err := p.processMsg(msg)
			if err != nil {
				log.Printf("cannot process msg: %s, error: %s", msg.Body, err)
			}

		case <-done:
			return nil
		}
	}
}

func (p PaymentsInterface) processMsg(msg amqp.Delivery) error {
	orderView := OrderToProcessView{};
	err := json.Unmarshal(msg.Body, &orderView);
	if err != nil {
		log.Printf("cannot decode msg: %s, error: %s", string(msg.Body), err);
	}

	orderPrice, err := price.NewPrice(orderView.Price.Cents, orderView.Price.Currency);
	if err != nil {
		log.Printf("cannot decode price for msg: %s, %s", string(msg.Body), err);
	}

	return p.service.InitializeOrderPayment(orderView.ID, orderPrice)
}