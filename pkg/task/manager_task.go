package task

import (
	"fmt"
	"strconv"
	"sync"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/monitor/job"
)

func InitAllSchedulerTask() *SchedulerExecutor {
	schedulerExcutor := NewSchedulerExecutor()

	schedulerExcutor.ScheduleEveryMinute(func() {
		job.MonitorResources()
	})

	return schedulerExcutor
}

func InitQueueWorkerTask() *Workers {
	maxWorkers, err := strconv.ParseInt(config.MaxWorkerQueue, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error when init queue worker task %s", err))
	}

	workers := make(Workers, int(maxWorkers))
	for i := 0; i < int(maxWorkers); i++ {
		wg := &sync.WaitGroup{}
		workerName := fmt.Sprintf("Default Worker-%d", i+1)
		workerExecutor := NewWorker(wg, workerName)
		workers[i] = *workerExecutor
	}

	return &workers
}

func RunAllActiveWorker(workers *Workers) {
	maxWorkers, err := strconv.ParseInt(config.MaxParalelWorkerQueue, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Error when init queue worker task %s", err))
	}

	for _, worker := range *workers {
		worker.RunWorker(int(maxWorkers))
	}
}
