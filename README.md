# Notes: Golang Advanced Concurrency Concepts

The tutorial consist of the following modules (in sense of Go Modules):
* `context/server`
* `context/cancellation`
* `error_group`
* `errors_in_waitgroup`
* `gosched`
* `once`
* `pipelines`.

To run/build any of them:
```
cd <module>
go run main.go
```

## `context`
This package provides tools to signal cancellation/completion of potentially long-running processes.
For example, we would like to signal a process cancellation of a long-running HTTP call or a DB session if there are errors or other custom events in order to free the resources of the application.

There are two sides to the context cancellation, that can be implemented: 
* Listening for the cancellation event, 
* Emitting the cancellation event.


### `context/listen`
In this example an HTTP server on port 9000 is set up. It processes a request within 2 seconds.
If there is no interruption from the client to the server, then after 2 seconds, the server returns 200 OK. Otherwise (when the client interrupts HTTP request), the server uses a `context` object obtained from the HTTP request object to signal the cancellation (write a message to STDOUT).

The `context.Done()` method returns a channel that receives an instance of `struct{}` whenever the `context` receives a cancellation event. Listening for a cancellation event is just waiting for `<-ctx.Done()`.

### `context/emit`
In this example we have two long-running jobs. They both take a `context` parameter - it is a child context created with a cancellation functions obtained from the roor context. Listening for a cancellation event is implemented again as waiting for `<-ctx.Done()`.

## `errors_in_waitgroup`
The module contains an approach to synchronize a pool of goroutines in case of errors.
It uses two channels (for done-signals and error-signals) and a `select` clause that "listens" to these channels. The done-channel is never sent to, it is merely closed after successful completion of all goroutines. If there is an error, then it the first error is passed to the error-channel. As this `happens-before` calling `wg.Done()`, the `select` statement will chose 
the error-channel output before the nil value from the done-channel. If there is no error, then the `select` clause choses the done-channel's output because the error-channel is never ready (no elements in it).

## `runtime.Gosched`
This function tells the Golang runtime to pause executing the current goroutine and instead start another waiting goroutine. It is called `yielding` and can be used to switch between goroutines and ensure no goroutine takes too much time. Goroutine switching is most influenced by this instruction when `GOMAXPROCS == 1`. On a machine with multiple cores there is less chance that goroutines will have to wait for each other because their execution is parallel not sequential.

## `sync.Once`
In the given example, an instruction (a function expression) passed to `sync.Once.Do()` will be run only once, regardless of where it is called.

If once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different value in each invocation. A new instance of `sync.Once` is required for each function to execute.

## `pipelines`
This module presents a very common pattern of joining multiple operations using channels.
Initially data are sent to a channel which is passed to a function that performs first transformation and sends its results to a channel which is then taken by another operation function and so on..