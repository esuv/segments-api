package scheduler

import (
	"context"
	"sync"
	"time"
)

type Job func(ctx context.Context)

type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func New() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

func (s *Scheduler) Add(ctx context.Context, job Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	s.wg.Add(1)
	go s.process(ctx, job, interval)
}

func (s *Scheduler) Shutdown() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, job Job, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			job(ctx)
		case <-ctx.Done():
			s.wg.Done()
			ticker.Stop()
			return
		}
	}
}
