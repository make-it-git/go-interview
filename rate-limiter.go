package main

import (
	"context"
	"sync"

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
	wg := sync.WaitGroup{}
	for _, r := range requests {
		r := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := c.SendRequest(ctx, r)
			if err != nil {
				log.WithError(err).Error("send request")
			}
		}()
	}
	wg.Wait()
}
