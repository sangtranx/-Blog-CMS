package asyncjob

import (
	"context"
	"time"
)

// job requirement
// 1. job can do something (handler)
// 2. job can retry (config retry times and duration)
// 3. Should be stateful
// 4. we should have job manger to manage jobs(*)

type JobState int

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDuration(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 5
	defaultMaxRetries = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}

	return &j
}

func (j *job) Execute(ctx context.Context) error {

	j.state = StateRunning

	var err error

	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return nil
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {

	if j.retryIndex == len(j.config.Retries)-1 {
		return nil
	}

	j.retryIndex++

	time.Sleep(j.config.Retries[j.retryIndex])
	err := j.Execute(ctx)
	if err != nil {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed

	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDuration(times []time.Duration) {

	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}
