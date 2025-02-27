package task

type JobTask interface {
	Name() string
	Handle(payload interface{}) error
}

type JobDictionary interface {
	GetAllJob() map[string]JobTask
	AddJob(job JobTask)
}
