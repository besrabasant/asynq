package rdb

import (
	"github.com/hibiken/asynq/internal/base"
	"github.com/redis/go-redis/v9"
)

// redisPubSub wraps redis.PubSub to implement base.PubSub interface.
type redisPubSub struct {
	ps *redis.PubSub
}

func newRedisPubSub(ps *redis.PubSub) *redisPubSub {
	return &redisPubSub{ps: ps}
}

func (r *redisPubSub) Channel() <-chan *base.PubSubMessage {
	out := make(chan *base.PubSubMessage)
	ch := r.ps.Channel()
	go func() {
		for m := range ch {
			out <- &base.PubSubMessage{Payload: m.Payload}
		}
		close(out)
	}()
	return out
}

func (r *redisPubSub) Close() error {
	return r.ps.Close()
}
