package pblocal

import (
	"Blog-CMS/common"
	"Blog-CMS/component/pubsub"
	"context"
	"log"
	"sync"
	"time"
)

// A pubsub run locally
// It has a queue (buffer channel) at it's core and many group of  subscribers
// Because we want to send a message with a specific topic for many subscribers in a group

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubsub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 100),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)
	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published", data.String(), data.Data())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)
	ps.locker.Lock()

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("unsubscribe topic", topic)

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {

				if chans[i] == c {
					// remove element at index in chans

					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {

	log.Println(("Starting local pubsub"))

	go func() {

		defer common.AppRecover()

		for {
			select {
			case mess := <-ps.messageQueue: // recieved message from queue
				log.Println("Message dequeued:", mess)
				ps.dispatchMessage(mess)
			default:
				time.Sleep(1 * time.Second) // reduce loading CPU
			}
		}
	}()

	return nil
}

func (ps *localPubSub) dispatchMessage(msg *pubsub.Message) {
	ps.locker.RLock()
	subs, ok := ps.mapChannel[msg.Channel()]
	ps.locker.RUnlock()

	if ok {
		for _, sub := range subs {
			go func(c chan *pubsub.Message) {
				defer common.AppRecover()
				select {
				case c <- msg:
					log.Println("Message sent to subscriber")
				default:
					log.Println("Warning: Subscriber channel is full, dropping message")
				}
			}(sub)
		}
	}
}
