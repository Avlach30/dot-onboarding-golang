package task

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
)

type SampleJobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *SampleJobTask) Handle(payload interface{}) error {
	log.Println(payload)
	return nil
}

// Handle implements domain.JobTask.
func (jobTask *SampleJobTask) Name() string {
	return "AuthSampleTask"
}

func AuthSampleJobTask() task.JobTask {
	return &SampleJobTask{}
}
