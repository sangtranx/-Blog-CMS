package asyncjob

import (
	"Blog-CMS/common"
	"context"
	"log"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	g := group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           &sync.WaitGroup{},
	}
	return &g
}

func (g *group) runJob(ctx context.Context, j Job) error {

	if err := j.Execute(ctx); err != nil {
		for {

			log.Fatal(err)

			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}

func (g *group) Run(ctx context.Context) error {

	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {

		if g.isConcurrent {

			defer common.AppRecover()

			go func(aj Job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}

		job := g.jobs[i]

		err := g.runJob(ctx, job)

		if err != nil {
			return err
		}

		errChan <- err
		g.wg.Done()
	}

	g.wg.Wait()

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}

	return err
}
