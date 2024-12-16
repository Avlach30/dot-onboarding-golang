package job

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/job/task"
	globalTask "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
)

type JobDictionary struct {
	listJob map[string]globalTask.JobTask
}

// AddJob implements globalTask.JobDictionary.
func (jobDictionary *JobDictionary) AddJob(taskName string, job globalTask.JobTask) {
	jobDictionary.listJob[taskName] = job
}

// GetAllJob implements globalTask.JobDictionary.
func (jobDictionary *JobDictionary) GetAllJob() map[string]globalTask.JobTask {
	return jobDictionary.listJob
}

func NewAuthJobDictionary(listJob map[string]globalTask.JobTask) globalTask.JobDictionary {
	return &JobDictionary{
		listJob: listJob,
	}
}

func InitJob() globalTask.JobDictionary {
	jobDictionary := NewAuthJobDictionary(map[string]globalTask.JobTask{})
	jobDictionary.AddJob(task.CompleteTaskName, task.AuthCompleteJobTask())
	jobDictionary.AddJob(task.StartTaskName, task.AuthStartJobTask())

	return jobDictionary
}
