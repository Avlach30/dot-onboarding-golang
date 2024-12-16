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
 migrate create -ext sql -dir migration/files -seq [migration name]
```
3. copy .env.example to .env.[folder name] and change based yours
```cli
cp .env.example .env
```
4. do migrate up or down:
```cli
go run main.go --migration true --exec [up,down]
```

## 🍰 The Layer

| Layer                | Directory                  |
| -------------------- |----------------------------|
| Frameworks & Drivers | /repository                |
| Interface            | /dto & /handler            |
| Usecases             | /usecase             |
| Entities             | /domain              |

- ### Repository  🥕
The **Repository** layer is responsible for data persistence and interaction with data sources, such as databases. It defines methods to retrieve and store data.

- **Interface**: Typically defined in the domain layer.
- **Implementation**: Found in the infrastructure layer.

**Responsibilities**:
- Abstracting data access.
- Providing a consistent interface for data operations.

- ### Usecase  🥕

The **Usecase** layer contains the business logic and application-specific rules. It orchestrates the flow of data between the repository and the handler.

- **Role**: Defines operations that the application can perform based on the domain model.
- **Interaction**: Converts incoming requests into actions and handles application-specific logic.

**Responsibilities**:
- Implementing business rules.
- Coordinating data flow between the repository and handler.

- ### Handler  🥕

The **Handler** layer (also known as Controller or API) is responsible for managing incoming requests. It converts these requests into a format suitable for the usecase layer and returns responses.

- **Role**: Interfaces with external systems (like web servers or APIs).
- **Interaction**: Converts requests to usecase actions and formats responses.

**Responsibilities**:
- Handling incoming requests.
- Returning responses to the client.

- ### DTO (Data Transfer Object)  🥕

**DTOs** (Data Transfer Objects) are used to transfer data between layers, especially between the handler and usecase layers. They help in structuring and validating data.

- **Role**: Facilitate data transfer and ensure data integrity.
- **Use**: Define and binding the structure of data being exchanged.

**Responsibilities**:
- Defining data structure.
- Validating data integrity.

- ### Domain  🥕

The **Domain** layer contains the core business logic and domain entities. It is independent of other layers and focuses on business rules and domain-specific logic.

- **Components**: Entities, value objects, and domain services.
- **Role**: Encapsulates the core logic and rules of the application.

**Responsibilities**:
- Defining business rules.
- Implementing domain-specific logic.

The **Scheduler** running in diff goroutine, set your schedule of task inside pkg/task/manager_task.go
- e.g
    ```
    // ... existing code ...

    schedulerExcutor.ScheduleEveryMinute(func() {
        job.MonitorResources()
    })

    // ... existing code ...
    ```

The **Queue** running in diff goroutine, set your task global using this example code
1. create your tasks `app/[domain]/job/your_job/task.go`
2. create your job to queue `[domain]_job.go`
3. register your job in `main.go`
4. now you can use your job anywhere
    ```
    singleton.Delegate(taskName, payload)
    ```
Note: Remember the payload type, example can be found at `app/auth/job`
