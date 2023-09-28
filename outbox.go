package main

import (
	"context"
	"database/sql"
	"fmt"
)

type KafkaMessage struct {
	offset      int64
	fromAccount int64
	toAccount   int64
	amount      int64
}

type KafkaConsumer interface {
	Read(ctx context.Context) ([]KafkaMessage, error)
	Ack(ctx context.Context, offset int64)
}

type Antifraud interface {
	IsValid(ctx context.Context, fromAccount int64, toAccount int64, amount int64) (bool, error)
}

type Limits interface {
	IsLimited(ctx context.Context, account int64, amount int64) (bool, error)
	UpdateLimits(ctx context.Context, account int64, amount int64) error
}

type PaymentGateway interface {
	Transfer(ctx context.Context, fromAccount int64, toAccount int64, amount int64) error
}

type Worker struct {
	consumer  KafkaConsumer
	db        *sql.DB
	antifraud Antifraud
	limits    Limits
	payment   PaymentGateway
}

func New(consumer KafkaConsumer, db *sql.DB, antifraud Antifraud, limits Limits, p PaymentGateway) *Worker {
	return &Worker{consumer: consumer, db: db, antifraud: antifraud, limits: limits, payment: p}
}

func (w *Worker) Work(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			messages, err := w.consumer.Read(ctx)
			if err != nil {
				return err
			}
			for _, m := range messages {
				fmt.Printf("Send money from %d to %d, amount: %d\n", m.fromAccount, m.toAccount, m.amount)

				ok, err := w.antifraud.IsValid(ctx, m.fromAccount, m.toAccount, m.amount)
				if err != nil {
					return err
				}
				if !ok {
					// notify caller about antifraud
					continue
				}

				isLimited, err := w.limits.IsLimited(ctx, m.fromAccount, m.amount)
				if err != nil {
					return err
				}
				if isLimited {
					// notify caller about transfer limits
					continue
				}

				err = w.payment.Transfer(ctx, m.fromAccount, m.toAccount, m.amount)
				if err != nil {
					return err
				}

				err = w.limits.UpdateLimits(ctx, m.fromAccount, m.amount)
				if err != nil {
					return err
				}

				// notify caller about success
				w.consumer.Ack(ctx, m.offset)
			}
		}
	}
}
