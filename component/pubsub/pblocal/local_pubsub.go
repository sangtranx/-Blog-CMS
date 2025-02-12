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

		// Thêm timeout để tránh block vô hạn
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Printf("Attempting to publish message to topic %s. Queue status: %d/%d",
			topic, len(ps.messageQueue), cap(ps.messageQueue))

		select {
		case ps.messageQueue <- data:
			log.Printf("Successfully published message to topic %s: %v", topic, data.Data())
		case <-timeoutCtx.Done():
			log.Printf("ERROR: Timeout publishing message to topic %s: Queue might be full", topic)
		}
		//ps.messageQueue <- data
		//log.Println("New event published", data.String(), data.Data())
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

		mess := <-ps.messageQueue

		log.Println("message dequeue : ", mess)

		if subs, ok := ps.mapChannel[mess.Channel()]; ok {
			for i := range subs {
				go func(c chan *pubsub.Message) {
					defer common.AppRecover()
					c <- mess
				}(subs[i])
			}
		}
	}()

	return nil
}
