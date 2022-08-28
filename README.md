## Dependency Injection & Service Provider Design Pattern Pattern

### This is a simple example to demonstrate 

#### (1) How to abstract the interaction with an external service
#### (2) Dependency Injection

<span style="font-size: 1.3em;">

#### What is dependency injection ?

```
Dependency injection is a software engineering technique where an object or struct receives its 
dependencies at compile time. Wikipedia defines dependency injection as such:

Dependency injection is a technique in which an object receives other objects that it depends on, 
called dependencies. Typically, the receiving object is called a client and the 
passed-in (‘injected’) object is called a service.

To get a better view of this, let’s analyze an example. Take a look at the following code:
```

```go
package main

import (
   "fmt"
)

type Message string

type Greeter struct {
   Message Message
}

type Event struct {
   Greeter Greeter
}

func GetMessage() Message {
   return Message("Hello world!")
}

func GetGreeter(m Message) Greeter {
   return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
   return g.Message
}

func GetEvent(g Greeter) Event {
   return Event{Greeter: g}
}

func (e Event) Start() {
   msg := e.Greeter.Greet()
   fmt.Println(msg)
}

func main() {
   message := GetMessage()
   greeter := GetGreeter(message)
   event := GetEvent(greeter)
   event.Start()
}
```

> If you take a look at the code above, we have a message, a greeter, and an event. 
> There is also a GetMessage function that returns a message; 
> a GetGreeter function that takes in a message and returns a greeter; 
> and a GetEvent function that accepts a greeter and returns an event. 
> The event also has a method called Start that prints out the message.

> If you take a look at our main method, we first create a message, 
> then we pass in the message as a dependency to the greeter and finally pass that to the event. 
> Run the code by running the command go run . in the terminal.

> As you can see, it prints “Hello, world!” to the console. 
> This is a very shallow dependency graph, but you can already see the complexity that comes 
> with this when implementing this in a large codebase. 
> That’s where dependency injection tools like Wire come in.

<hr />

#### Another Example of Dependency Injection

> In order to remove the dependence of business logic on external packages, 
> dependency injection is used.

> For example, through the New constructor, we inject the dependency into the structure of the 
> business logic. This makes the business logic independent (and portable). 
> We can override the implementation of the interface without making changes to the `usecase` package

```go
package usecase

import (
    // Nothing!
)

type Repository interface {
    Get()
}

type UseCase struct {
    repo Repository
}

func New(r Repository) *UseCase{ 
    return &UseCase{
        repo: r,
    }
}

func (uc *UseCase) Do()  {
    uc.repo.Get()
}
```

```
New Constructor -> Is Used To Inject Dependency into the struct of Business Logic

We can override the implementation of the interface,
That is Get() method -> without making changes to the usecase package
```

<hr />

#### Sample Run

`Build Binary`

```bash
go build -o service-provider
```

#### Get Output From Command Line

```bash
[~/git/goworkspace/src/dependency-injection]$ ./service-provider -operation cli -city "Paris" -apikey "<INSERT_APIKEY>"
```

`Output`

```bash
2022/08/28 13:11:55 @ provider : path : /weather?q=Paris&appid=<API_KEY_REDACTED>&units=metric
2022/08/28 13:11:55 @ provider : completeURL : https://api.openweathermap.org/data/2.5/weather?q=Paris&appid=<API_KEY_REDACTED>&units=metric
2022/08/28 13:11:56 WeatherUsecase >

{
    "Temp": 22.72,
    "Pressure": 1017,
    "MinTemp": 21.11,
    "MaxTemp": 23.41
}
```

#### Run HTTP Server : Implement POST Handler Which Returns Weather Data

`Start HTTP Server`

```bash
[~/git/goworkspace/src/dependency-injection]$ ./service-provider -operation api
Mux HTTP server running on port :8181
```

`Make HTTP POST Call`

```bash
[~/git/goworkspace/src/dependency-injection]$ curl -X POST http://localhost:8181/weather -d '{"city":"London","apikey":"<API_KEY>"}' 2>/dev/null | python -m json.tool
````

`Output`

```bash
{
    "Temp": 18.35,
    "Pressure": 1022,
    "MinTemp": 16.59,
    "MaxTemp": 19.82
}
```

</span>
