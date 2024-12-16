package task

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

type SchedulerExecutor struct {
	Executor *gocron.Scheduler
}

func NewSchedulerExecutor() *SchedulerExecutor {
	return &SchedulerExecutor{
		Executor: gocron.NewScheduler(),
	}
}

func (schedulerExecutor *SchedulerExecutor) RunScheduler() {
	executor := schedulerExecutor.Executor

	// Get next running time
	_, nextRunTime := executor.NextRun()
	fmt.Println(nextRunTime)

	// Start the scheduler
	<-executor.Start()
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEverySecond(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Second().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoSecond(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Second().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryFiveSecond(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(5).Second().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoSeconds(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Seconds().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryMinute(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Minute().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoMinutes(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Minutes().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryHour(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Hour().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoHours(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Hours().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryDay(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Day().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoDays(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Days().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryWeek(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Week().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryTwoWeeks(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(2).Weeks().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEverySecondWithParams(callback func(int, *interface{}), a int, b string) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Second().Do(callback, a, b)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryMonday(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Monday().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleEveryThursday(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Thursday().Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleDailyAt(time string, callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Day().At(time).Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleMondayAt(time string, callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Monday().At(time).Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleTuesdayAt(time string, callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Tuesday().At(time).Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleHourlyFromNextTick(callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Hour().From(gocron.NextTick()).Do(callback)
}

func (schedulerExecutor *SchedulerExecutor) ScheduleHourlyFromDate(t *time.Time, callback func()) {
	executor := schedulerExecutor.Executor
	executor.Every(1).Hour().From(t).Do(callback)
}
