package payments

import (
	"encoding/json"
	"errors"
	"golang-mono-micro/pkg/common/price"
	"golang-mono-micro/pkg/orders/domain/orders"
	payments_amqp_interface "golang-mono-micro/pkg/payments/interfaces/amqp"
	"log"
	"github.com/streadway/amqp"
)

type AMQPService struct {
	queue amqp.Queue
    channel *amqp.Channel
}

func NewAMQPService(url string, queueName string) (AMQPService, error) {
    conn, err := amqp.Dial(url)
    if err != nil {
        return AMQPService{}, errors.New("unable to setup connection");
    }

    ch, err := conn.Channel()
    if err != nil {
        return AMQPService{}, errors.New("unable to setup server channel");
    }

    q, err := ch.QueueDeclare(
        queueName,
        true,
        false,
        false,
        false,
        nil,
    );
    if err != nil {
        return AMQPService{}, errors.New("unable to setup queue");
    }

    return AMQPService{
        queue: q,
        channel: ch,
    }, nil
}

func (a AMQPService) InitializeOrderPayment(id orders.ID, price price.Price) error {
    order := payments_amqp_interface.OrderToProcessView{
        ID: string(id),
        Price: payments_amqp_interface.PriceView{
            Cents: price.Cents(),
            Currency: price.Currency(),
        },
    }

    b, err := json.Marshal(order);
    if err != nil {
        return errors.New("cannot marshal order for amqp");
    }

    err = a.channel.Publish(
        "",
        a.queue.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body: b,
        },
    )

    if err != nil {
        return errors.New("cannot send orders to amqp");
    }

    log.Printf("Sent order %s to amqp", id);
    return nil
}

func (a AMQPService) Close() error {
    return a.channel.Close();
}