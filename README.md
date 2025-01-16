## 📖 Contains

- [Migration](#-migration)
- [The Layer](#-the-layer)

## 🌊 Migration

1. install [golang-migrate] (https://github.com/golang-migrate/migrate)
    ```cli
    brew install golang-migrate
    ```
2. create migration by:
    ```cli
    go run main.go --migration true --exec create --fileName [fileName]
    ```
3. copy .env.example to .env.[folder name] and change based yours
    ```cli
    cp .env.example .env
    ```
4. do migrate up or down:
    ```cli
    go run main.go --migration true --exec [up,down,fresh]
    ```
5. do seed your data:
    ```cli
    go run main.go --dbseed true
    ```
    ```cli
    go run main.go --dbseed true --class UserSeeder,RoleSeeder
    ```

## 🍰 The Layer

| Layer                | Directory                  |
| -------------------- |----------------------------|
| Frameworks & Drivers | /repository                |
| Interface            | /interface                 |
| Usecases             | /usecase                   |
| Entities             | /domain                    |

## 🌊 The Utils

The **Scheduler** running in diff goroutine, set your schedule of task inside pkg/task/manager_task.go
1. this example to add MonitorResource run every minute
    ```
    // ... existing code ...

    schedulerExcutor.ScheduleEveryMinute(func() {
        job.MonitorResources()
    })

    // ... existing code ...
    ```

The **Queue** running in diff goroutine, set your task global using this example code
1. create your tasks `app/[domain]/job/task/[action]_task`
2. create your job to queue `[domain]_job.go`
3. register your job in `main.go`
4. now you can use your job anywhere
    ```
    singleton.Delegate(taskName, payload)
    ## or using standalone
    singleton.DelegateStandalone("AuthSomeStandaonle", "string", &authTask.StandaloneJobTask{})
    ```
Note: Remember the payload type, example can be found at `app/auth/job`

The **Logger** log the message whatever u want
1. define log driver inside your .env with LOG_DRIVER, if LOG_DRIVER is file. config the LOG_FOLDER_PATH as u want
4. now you can use your log anywhere with speciic [category], otherwise default
    ```
    log.Println("[activity] some activity")
    log.Println("[integration] some integration")
    log.Println("[etc] some etc")
    ```
