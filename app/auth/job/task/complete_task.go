package task

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
)

const CompleteTaskName = "AuthCompletedTask"

type CompleteJobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *CompleteJobTask) Handle(payload interface{}) error {
	log.Println(payload)
	return nil
}

func AuthCompleteJobTask() task.JobTask {
	return &CompleteJobTask{}
}
