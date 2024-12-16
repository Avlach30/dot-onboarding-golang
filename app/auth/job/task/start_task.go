package task

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
)

const StartTaskName = "AuthStartedTask"

type StartJobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *StartJobTask) Handle(payload interface{}) error {
	log.Println(payload.(map[string]any))
	return nil
}

func AuthStartJobTask() task.JobTask {
	return &StartJobTask{}
}
