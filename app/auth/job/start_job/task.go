package start_job

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
)

const TaskName = "AuthStartedTask"

type JobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *JobTask) Handle(payload interface{}) error {
	log.Println(payload.(map[string]any))
	return nil
}

func AuthStartJobTask() domain.JobTask {
	return &JobTask{}
}
