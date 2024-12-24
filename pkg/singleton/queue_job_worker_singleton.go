package singleton

import (
	"encoding/json"
	"fmt"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm/clause"
)

// Singleton struct to hold a slice of RSA key pairs
type UtilsSingleton struct {
	Workers           *task.Workers
	ListRegisteredJob *task.ListRegisteredJob
}

// GetKeyPairs returns the singleton instance with a slice of key pairs
func GetGlobalUtils() *UtilsSingleton {
	return utilsSingleton
}

// GetKeyPairs returns the singleton instance with a slice of key pairs
func Delegate(taskName string, payload interface{}) {

	convertedPayload := func() string {
		switch v := payload.(type) {
		case map[string]any:
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				// If marshaling fails, fall back to string representation
				return fmt.Sprintf("%v", v)
			}
			return string(jsonBytes)
		default:
			return fmt.Sprintf("%v", v)
		}
	}()

	jobEntity := &domain.JobEntity{
		Booked:   false,
		Payload:  convertedPayload,
		TaskName: taskName,
	}

	dbUtil.Save(&jobEntity)
}

// GetKeyPairs returns the singleton instance with a slice of key pairs
func DelegateStandalone(taskName string, payload interface{}, jobTask task.JobTask) {

	utilsSingleton.ListRegisteredJob.RegisterSingleJob(taskName, jobTask)

	convertedPayload := func() string {
		switch v := payload.(type) {
		case map[string]any:
			jsonBytes, err := json.Marshal(v)
			if err != nil {
				// If marshaling fails, fall back to string representation
				return fmt.Sprintf("%v", v)
			}
			return string(jsonBytes)
		default:
			return fmt.Sprintf("%v", v)
		}
	}()

	jobEntity := &domain.JobEntity{
		Booked:   false,
		Payload:  convertedPayload,
		TaskName: taskName,
	}

	dbUtil.Save(&jobEntity)
}

// ExecuteJobTask processes unbooked job entities and executes their associated tasks
func ExecuteJobTask() {
	for {
		time.Sleep(5 * time.Second)

		jobEntities := &[]domain.JobEntity{}
		dbUtil.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("booked = false").
			Order("created_at ASC").
			Limit(5).
			Find(&jobEntities)

		if len(*jobEntities) < 1 {
			continue
		}

		for _, jobEntity := range *jobEntities {
			convertedPayload := utils.StringToInterface(jobEntity.Payload)
			utilsSingleton.Workers.ExecuteJobTask(jobEntity.TaskName, convertedPayload, utilsSingleton.ListRegisteredJob)
		}

		var ids []string
		for _, entity := range *jobEntities {
			ids = append(ids, entity.ID.String())
		}
		dbUtil.Model(&domain.JobEntity{}).Where("id IN ?", ids).Update("booked", true)
	}
}

func AddJobDictionary(jobDictionary task.JobDictionary) {
	utilsSingleton.ListRegisteredJob.RegisterJob(jobDictionary)
}
