package scheduler

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestScheduler_Run(t *testing.T) {
	scheduler := New()
	scheduler.Add(context.Background(), myJob, 1*time.Second)

	time.Sleep(3 * time.Second)
	scheduler.Shutdown()
	time.Sleep(10 * time.Second)
}

func myJob(ctx context.Context) {
	fmt.Println("hello from scheduler -> myJob func executing")
}
