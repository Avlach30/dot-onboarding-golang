package singleton

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/task/domain"
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
func ExecuteJobTaskByDB() {
	for {
		time.Sleep(5 * time.Second)

		jobEntities := &[]domain.JobEntity{}
		dbUtil.Clauses(clause.Locking{Strength: "UPDATE"}).Where("booked = false").Order("created_at ASC").Limit(5).Find(&jobEntities)

		countJobEntities := len(*jobEntities)
		if countJobEntities < 1 {
			continue
		}

		jobEntityIds := make([]string, countJobEntities)
		for i, jobEntity := range *jobEntities {
			jobEntityIds[i] = jobEntity.ID.String()
			payload := jobEntity.Payload
			var convertedPayload interface{}

			// Try converting to int
			if intVal, err := strconv.Atoi(payload); err == nil {
				convertedPayload = intVal
			} else if floatVal, err := strconv.ParseFloat(payload, 64); err == nil {
				// Try converting to float64
				convertedPayload = floatVal
			} else if boolVal, err := strconv.ParseBool(payload); err == nil {
				// Try converting to bool
				convertedPayload = boolVal
			} else {
				// If all conversions fail, try unmarshaling to interface{}
				var unmarshaledVal interface{}
				if err := json.Unmarshal([]byte(payload), &unmarshaledVal); err == nil {
					// Check if the unmarshaled value is a map[string]interface{}
					if mapVal, ok := unmarshaledVal.(map[string]interface{}); ok {
						// Try to convert to map[string]string
						stringMap := make(map[string]string)
						for k, v := range mapVal {
							if strVal, ok := v.(string); ok {
								stringMap[k] = strVal
							} else {
								// If any value is not a string, use the original unmarshaledVal
								convertedPayload = unmarshaledVal
								goto end
							}
						}
						convertedPayload = stringMap
					} else {
						convertedPayload = unmarshaledVal
					}
				} else {
					// If unmarshaling fails, keep it as a string
					convertedPayload = payload
				}
			end:
			}

			// Now convertedPayload can be of type bool, int, float64, or string
			utilsSingleton.Workers.ExecuteJobTask(jobEntity.TaskName, convertedPayload, utilsSingleton.ListRegisteredJob)

		}

		dbUtil.Model(&domain.JobEntity{}).Where("id in ?", jobEntityIds).Update("booked", true)
	}
}

func AddJobDictionary(jobDictionary domain.JobDictionary) {
	utilsSingleton.ListRegisteredJob.RegisterJob(jobDictionary)
}
