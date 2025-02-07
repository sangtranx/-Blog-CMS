package pubsub

import (
	"fmt"
	"time"
)

type Topic string

type Message struct {
	id        string
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now()
	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("Message %s", m.channel)
}

func (m *Message) Channel() Topic {
	return m.channel
}

func (m *Message) SetChannel(channel Topic) {
	m.channel = channel
}

func (m *Message) Data() interface{} {
	return m.data
}
