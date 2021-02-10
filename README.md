# GOLANG - Advanced studies

This is a project to learn about some Golang concurrency concepts.

## Go routines

To create a go routine just adds the **go** keyword before the function.

Example:

```go
... 
    go something()
...
```

## Race conditions

To check if there are race conditions in the program, run the code with command below.

```shell
go run -race [file.go]
```

To solve the race conditions problem, you can use the Mutex concept. The Mutex ensure that a variable not will be modified by 2 or more processes in the same time.

Example:

```go
var m sync.Mutex
... 
    m.Lock()
    anyChangedVariable = newValue
    m.Unlock()
...
```

## Channels

To create a channel use the code below

```go
newChannel := make(chan [type])
```

To write value in the channel

```go
newChannel <- value
```

To read channel value

```go
anyVariable <- newChannel
```