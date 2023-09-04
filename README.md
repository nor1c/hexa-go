### Hexagonal Architecture

## Run
```
$ air
```

## Directory Structure
- `internal` - used to denote code that is specific to your application and not meant to be imported or used externally. It usually contains the core logic of your application.
- `internal/application` - contain the higher-level logic of your application. This could include files related to the business logic, interactions between different parts of the application, and orchestrating use cases.
- `internal/application/repository` - to deals with data storage and retrieval. This could involve database interactions or other forms of data storage.
- `internal/application/usecase` - define and implement the operations that your application can perform.
- `pk` - contains code that could potentially be shared between different projects or modules. It means to hold reusable components.
- `pkg/domain` - where you define your application's domain models, which represent the core concepts of your application.
- `pkg/http` -  to handling HTTP requests and responses. This could include routing, middleware, and other HTTP-related logic.
- `pkg/ports` - contain interfaces or contracts that define how different parts of your application interact with each other. It serves as a boundary between different layers, enforcing separation of concerns.