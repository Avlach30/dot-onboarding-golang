package domain

type JobTask interface {
	Handle(payload interface{}) error
}

type JobDictionary interface {
	GetAllJob() map[string]JobTask
	AddJob(taskName string, job JobTask)
}
