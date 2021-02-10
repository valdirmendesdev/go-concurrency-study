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