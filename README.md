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

Above commands compiles the code and generated 'service-provider' binary


#### Get Output From Command Line

```bash
*[master][~/git/goworkspace/src/dependency-injection]$ ./service-provider -version "v1" -operation "cli" -city "London" -apikey "<API_KEY>"

2022/08/28 16:41:10 @ provider : path : /weather?q=London&appid=<API_KEY>&units=metric
2022/08/28 16:41:10 @ provider : completeURL : https://api.openweathermap.org/data/2.5/weather?q=London&appid=<API_KEY>&units=metric
2022/08/28 16:41:11 WeatherUsecase (1) >

{
    "Temp": 16.28,
    "Pressure": 1022,
    "MinTemp": 14.2,
    "MaxTemp": 17.64
}

*[master][~/git/goworkspace/src/dependency-injection]$ ./service-provider -version "v2" -operation "cli" -city "London" -apikey "<API_KEY>"

2022/08/28 16:41:14 @ provider : path : /weather?q=London&appid=<API_KEY>&units=metric
2022/08/28 16:41:14 @ provider : completeURL : https://api.openweathermap.org/data/2.5/weather?q=London&appid=<API_KEY>&units=metric
2022/08/28 16:41:15 WeatherUsecase (2) >

{
    "Temp": 16.28,
    "Pressure": 1022,
    "MinTemp": 14.09,
    "MaxTemp": 17.64
}
```

#### Run HTTP Server : Implement POST Handler Which Returns Weather Data

`Start HTTP Server`

```bash
$ ./service-provider -operation api
Mux HTTP server running on port :8181
```

`Make HTTP POST Call`

```bash
$ curl -X POST http://localhost:8181/weather -d '{"city":"London","apikey":"<API_KEY>"}' 2>/dev/null | python -m json.tool
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
