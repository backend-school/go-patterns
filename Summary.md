## Design patterns summary

### Creational

| Number | Name             | Summary                                           | Examples | Comment |
| ------ | ---------------- | ------------------------------------------------- | -------- | ------- |
| 1      | Singleton        | Create a single object                            |          |         |
| 2      | Factory          | Create many objects                               |          |         |
| 3      | Abstract Factory | Create another factories                          |          |         |
| 4      | Prototype        | Create obj in compile time. Clone in runtime      |          |         |
| 5      | Builder          | Reuse an algo to create many impl of an interface |          |         |

### Structural

| Number | Name      | Summary                                                     | Examples                                | Comment |
| ------ | --------- | ----------------------------------------------------------- | --------------------------------------- | ------- |
| 1      | Proxy     | Wrap object to hide some of its functionality               | Auth, remote, db, protection            |         |
| 2      | Adapter   | Wrap one interface into another                             | Wrap old interface into a new one       |         |
| 3      | Facade    | Hide or restrict complexity with a simpler API or interface |                                         |         |
| 4      | Composite | Compose multiple interfaces to get a richer functionality   |                                         |         |
| 5      | Flyweight | Share a state among many instances                          | Pointers, Double buffering, ring buffer |         |
| 6      | Decorator | Add new features to an existing interface                   |                                         |         |
| 7      | Bridge    | Decouple an abstraction from impl                           |                                         |         |

### Behavioral

| Number | Name                    | Summary                                        | Examples             | Comment                     |
| ------ | ----------------------- | ---------------------------------------------- | -------------------- | --------------------------- |
| 1      | Template                | Create a reusable part and import to many inst |                      |                             |
| 2      | State                   | Implement FSM for complex state transitions    |                      |                             |
| 3      | Strategy                | Focus on changing algorithms                   | io.Writer, io.Reader |                             |
| 4      | Command                 | Container with a single action                 | UI button            | Similar to Strategy         |
| 5      | Chain of Responsibility | Task seggregation and a SRP                    | HTTP/GRPS middleware |                             |
| 6      | Memento                 | An object to remind about something            |                      | Memento, Producer, Consumer |
| 7      | Interpreter             | Create a domain specific lang to ease a task   | SQL                  |                             |
| 8      | Mediator                | Impl loose coupling between algorithms         | Kafka, NATS, Redis   |                             |
| 9      | Observer                | Subscribe for a particular event               | One-to-many compute  |                             |
| 10     | Visitor                 | Control an outside object or algorithm         | io.Writer            |                             |

### Concurrent and event driven

| Number | Name        | Summary                                                | Examples                            | Comment                          |
| ------ | ----------- | ------------------------------------------------------ | ----------------------------------- | -------------------------------- |
| 1      | Worker pool | Limit server resources for a task                      | OS thread pool, goroutine scheduler |                                  |
| 2      | CQRS        | Decouple a producer from consumers with a buffer/queue | Payout service                      | Use Kafka or NATS for decoupling |
