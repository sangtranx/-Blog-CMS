package pubsub

import "context"

type Pubsub interface {
	Publish(ctx context.Context, channel Topic, data *Message) error
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func())
}
