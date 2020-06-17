# Notes: Golang Advanced Concurrency Concepts

The tutorial consist of the following modules (in sense of Go Modules):
* `context`
* `error_group`
* `gosched`
* `once`.

To run/build any of them:
```
cd <module>
go run main.go
```
## `sync.Once`
In the given example, an instruction (a function expression) passed to `sync.Once.Do()` will be run only once, regardless of where it is called.

If once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different value in each invocation. A new instance of `sync.Once` is required for each function to execute.