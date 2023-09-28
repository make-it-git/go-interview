package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
}

// TODO: rate limit api calls
// 1. In fly requests limit
// 2. Per second limit
func makeApiCall(ctx context.Context, c Client, log *logrus.Logger, requests []Request) {
	for _, r := range requests {
		err := c.SendRequest(ctx, r)
		if err != nil {
			log.WithError(err).Error("send request")
		}
	}
}
