package task

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
)

type StandaloneJobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *StandaloneJobTask) Handle(payload interface{}) error {
	log.Println(payload.(map[string]any))
	return nil
}

// Handle implements domain.JobTask.
func (jobTask *StandaloneJobTask) Name() string {
	return "AuthStandaloneJobTask"
}

func AuthStandaloneJobTask() task.JobTask {
	return &StandaloneJobTask{}
}
