package subscriber

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/asyncjob"
	"Blog-CMS/component/pubsub"
	"context"
	"log"
)

type GroupJob interface {
	Run(ctx context.Context) error
}

type comsumerJob struct {
	Title   string
	Handler func(ctx context.Context, msg *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (e *consumerEngine) StartSubTopic(topic pubsub.Topic, isConcurrent bool, comsumerJobs ...comsumerJob) error {

	c, _ := e.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, job := range comsumerJobs {
		log.Println(job.Title)
	}

	getJobHld := func(job *comsumerJob, msg *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			return job.Handler(ctx, msg)
		}
	}

	go func() {
		for {
			msg := <-c
			log.Printf("Consumer processing message: %v", msg.Data())

			// Chỉ tạo và chạy job nếu có message hợp lệ
			jobHldArr := make([]asyncjob.Job, len(comsumerJobs))

			for i, job := range comsumerJobs {
				jobhld := getJobHld(&job, msg)
				jobHldArr[i] = asyncjob.NewJob(jobhld)
			}

			g := asyncjob.NewGroup(isConcurrent, jobHldArr...)

			if err := g.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}

// add new subscribers
func (e *consumerEngine) Start() error {
	e.StartSubTopic(
		common.TopicUserLikePost,
		true,
		IncreasePostLikeCount(e.appCtx))

	e.StartSubTopic(
		common.TopicUserDisLikePost,
		true,
		DecreasePostLikeCount(e.appCtx))
	return nil
}
