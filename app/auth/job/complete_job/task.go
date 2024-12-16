package complete_job

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
)

const TaskName = "AuthCompletedTask"

type JobTask struct{}

// Handle implements domain.JobTask.
func (jobTask *JobTask) Handle(payload interface{}) error {
	log.Println(payload)
	return nil
}

func AuthCompleteJobTask() domain.JobTask {
	return &JobTask{}
}
