package main

import (
	"time"

	"github.com/qastub/ultron"
)

type (
	benchmark struct{}
)

func (b benchmark) Name() string {
	return "benchmark"
}

func (b benchmark) Fire() error {
	time.Sleep(time.Millisecond * 10)
	return nil
}

func main() {
	task := ultron.NewTask()
	t := benchmark{}
	task.Add(t, 1)

	ultron.LocalEventHook.Concurrency = 200
	ultron.LocalRunner.WithTask(task)
	ultron.LocalRunner.Config.Concurrence = 10000
	ultron.LocalRunner.Config.HatchRate = 200
	ultron.LocalRunner.Config.Duration = time.Minute * 3
	ultron.LocalRunner.Config.MaxWait = ultron.ZeroDuration
	ultron.LocalRunner.Config.MinWait = ultron.ZeroDuration

	ultron.LocalRunner.Start()
}
