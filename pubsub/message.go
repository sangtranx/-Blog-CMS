package pubsub

import (
	"fmt"
	"time"
)

type Topic string

type Message struct {
	id        int
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now()
	return &Message{
		id:        fmt.Sprint(now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}
