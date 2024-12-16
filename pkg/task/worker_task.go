package task

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
)

type Job struct {
	ID      uuid.UUID
	Name    string
	Payload interface{}
	Task    domain.JobTask
}

type Worker struct {
	Jobs        chan Job
	Wg          *sync.WaitGroup
	Name        string
	JobsCounter int64
}

type Workers []Worker

type ListRegisteredJob map[string]domain.JobTask

func (listRegisteredJob *ListRegisteredJob) RegisterJob(jobDictionary domain.JobDictionary) {
	listAllJob := jobDictionary.GetAllJob()
	merged := make(map[string]domain.JobTask)
	for key, value := range *listRegisteredJob {
		merged[key] = value
	}

	for key, value := range listAllJob {
		merged[key] = value
	}

	*listRegisteredJob = merged
}

func (workers *Workers) ExecuteJobTask(taskName string, payload interface{}, listRegisteredJob *ListRegisteredJob) {
	// Add a new worker to the pool.
	workerRunnning := *workers

	workerLessJob := workerRunnning[0]
	for _, worker := range workerRunnning {
		if worker.JobsCounter < workerLessJob.JobsCounter {
			workerLessJob = worker
		}
	}

	jobList := *listRegisteredJob
	jobTask := jobList[taskName]
	if jobTask == nil {
		return
	}

	workerLessJob.JobsCounter++
	newJob := Job{
		ID:      uuid.New(),
		Name:    taskName,
		Payload: payload,
		Task:    jobTask,
	}

	go workerLessJob.AddJob(newJob)
}

func NewWorker(wg *sync.WaitGroup, name string) *Worker {
	return &Worker{
		Jobs: make(chan Job),
		Wg:   wg,
		Name: name,
	}
}

func DoWork(id int, job Job) {
	fmt.Printf("Worker-%d : started task [%s]\n", id, job.Name)
	log.Println("==========RUN TASK========")
	maxTries, err := strconv.ParseInt(config.MaxTriesQueue, 10, 64)
	if err != nil {
		log.Println("Parse error maxTries : ", err)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println("==========ERROR TASK========")
			return
		}
	}()

	tries := 1
	for {
		if int(maxTries) < tries {
			log.Printf("Max Tries %d, Error %s", maxTries, err)
			break
		}

		err := job.Task.Handle(job.Payload)
		if err != nil {
			log.Printf("Error %s\n", err.Error())
			log.Printf("Worker-%d : %s failed, retrying ...\n", id, job.Name)
			time.Sleep(3 * time.Second)
		} else {
			break
		}

		tries++
	}

	log.Println("==========END TASK========")

	fmt.Printf("Worker-%d : completed %s!\n", id, job.Name)
}

func (worker *Worker) RunWorker(maxParalelWorkers int) {
	jobs := worker.Jobs
	wg := worker.Wg

	wg.Add(maxParalelWorkers)
	for i := 1; i <= maxParalelWorkers; i++ {
		go func(i int) {
			defer wg.Done()

			for j := range jobs {
				DoWork(i, j)
			}
		}(i)
	}

	// wait for workers to complete
	wg.Wait()
}

func (worker *Worker) AddJob(newJob Job) {
	jobs := worker.Jobs
	jobs <- newJob
}

func (worker *Worker) CloseJob(newJob Job) {
	jobs := worker.Jobs
	close(jobs)
}
