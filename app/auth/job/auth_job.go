package job

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/job/complete_job"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/job/start_job"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
)

type JobDictionary struct {
	listJob map[string]domain.JobTask
}

// AddJob implements domain.JobDictionary.
func (jobDictionary *JobDictionary) AddJob(taskName string, job domain.JobTask) {
	jobDictionary.listJob[taskName] = job
}

// GetAllJob implements domain.JobDictionary.
func (jobDictionary *JobDictionary) GetAllJob() map[string]domain.JobTask {
	return jobDictionary.listJob
}

func NewAuthJobDictionary(listJob map[string]domain.JobTask) domain.JobDictionary {
	return &JobDictionary{
		listJob: listJob,
	}
}

func InitJob() domain.JobDictionary {
	jobDictionary := NewAuthJobDictionary(map[string]domain.JobTask{})
	jobDictionary.AddJob(complete_job.TaskName, complete_job.AuthCompleteJobTask())
	jobDictionary.AddJob(start_job.TaskName, start_job.AuthStartJobTask())

	return jobDictionary
}
