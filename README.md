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

## `runtime.Gosched`
This function tells the Golang runtime to pause executing the current goroutine and instead start another waiting goroutine. It is called `yielding` and can be used to switch between goroutines and ensure no goroutine takes too much time. Goroutine switching is most influenced by this instruction when `GOMAXPROCS == 1`. On a machine with multiple cores there is less chance that goroutines will have to wait for each other because their execution is parallel not sequential.

## `sync.Once`
In the given example, an instruction (a function expression) passed to `sync.Once.Do()` will be run only once, regardless of where it is called.

If once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different value in each invocation. A new instance of `sync.Once` is required for each function to execute.
