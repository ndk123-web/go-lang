# Go Lang

1. Statically Type Language
2. Strongly Type Language
3. Go is Compiled
4. Fast Compiled Time
5. Built in Concurrency
6. Simplicity (Garbage Collector)

## Intro Of Go Lang

- Go (or Golang) was created by Google engineers (Rob Pike, Ken Thompson, and Robert Griesemer) to solve a major problem:

- “Big companies like Google had huge C++ and Java systems that compiled slowly and were hard to maintain in large-scale systems.”

So they designed Go to be:

- Simple like Python
- Fast like C
- Concurrent like Erlang
- Compiled & Typed like C/C++

| Feature                  | Description                                              |
| ------------------------ | -------------------------------------------------------- |
| **Compiled language**    | Compiles directly to machine code → runs super fast.     |
| **Static typing**        | Catches type errors during compile-time.                 |
| **Garbage Collection**   | Manages memory automatically.                            |
| **Concurrency built-in** | Has native support for goroutines (lightweight threads). |
| **Cross-platform**       | One command builds executables for any OS.               |
| **Easy syntax**          | Clean and minimal — faster learning curve.               |

## Data Types

| Category      | Examples                                                                      |
| ------------- | ----------------------------------------------------------------------------- |
| Numeric Types | int, int8, int16, int32, int64, uint, float32, float64, complex64, complex128 |
| String Type   | string                                                                        |
| Boolean Type  | bool                                                                          |
| Derived Types | array, slice, struct, map, pointer, function, interface, channel              |

## Built in Methods

1. len(array|string)
2. cap(array)

## Commands

1. Before we need init (like npm init -y) it generates go.mod

   ```bash
   go mod init myapp
   ```

2. go run file.go

   - Directly Run go file

   ```bash
   go run file.go
   ```

3. go build file.go

   - Gives .exe and we can run directly

   ```bash
   go build file.go
   ```

4. goos

   - We can generate any type of build (linux , windows , mac (darwin))

   ```bash
   goos="darwin" go build
   goos="windows" go build
   goos="linux" go build
   ```

5. go mod vendor

   - Its like node_modules of npm

   ```bash
   go mod vendor
   ```

6. go get path

   - Install package from web
   - All the files are going in cache (u can see by `go env` and their u can see `GOPATH` and in that location all the files will be there)

   ```bash
   go get https://github.com/gorilla/mux
   ```

7. go mod tidy
   - removes indirect from name inside go.mod
   - in `go.mod` in package u can see `// indirect` command so using this command it removes `// indirect` it means our module is using this new module
   ```bash
   go mod tidy
   ```

## Go Mod Commands

- go mod init — create go.mod, set module name & Go version s
- go get — fetch dependency into module cache and update go.mod/go.sum s
- go mod tidy — add missing, remove unused deps; clean go.mod/go.sum s
- go mod verify — check go.sum hashes vs cache to ensure integrity s
- go list -m all / go list -m -versions — inspect modules and available versions s s
- go mod graph — show dependency graph (who depends on whom) s
- go mod edit -go= / -module= — programmatically edit go.mod fields s
- go mod vendor & -mod=vendor — vendor deps locally and build/run using vendor folder s s
- go env — inspect GOPATH/GOMODCACHE to find cached modules s
- go run / go build — run or compile your module-aware project

## Memory Management

- It has Garbage Collector
- GC happens automatically

- new()

  - Allocate memory but no Initialization
  - you will get a memory address
  - zeroed storage (can't put data)

- make()
  - Allocate memory and Initialization
  - you will get a memory address
  - non-zerores storage (can put data)

## Array Vs Slice

| Feature         | Array `[N]int`      | Slice `[]int`                |
| --------------- | ------------------- | ---------------------------- |
| Size            | Fixed               | Dynamic                      |
| Can append?     | ❌ No               | ✅ Yes                       |
| Memory          | Contiguous          | Points to underlying array   |
| Passing to func | Copies entire array | Passes reference (efficient) |
| Type uniqueness | `[4]int] ≠ [5]int`  | `[]int` same type always     |

## new vs make

| Feature         | `new(Type)`                   | `make(Type, ...)`                          |
| --------------- | ----------------------------- | ------------------------------------------ |
| Returns         | Pointer (`*Type`)             | Initialized value (not pointer)            |
| Purpose         | Allocate memory               | Allocate + initialize                      |
| Zero value      | Yes                           | Not needed, already usable                 |
| Types supported | Any type (struct, int, array) | Only reference types (slice, map, channel) |
| Usage           | `p := new(int)`               | `s := make([]int, 5)`                      |

## Naming Conventions in Go Lang

| Identifier   | Starts with | Accessible Outside Package? | Example                    |
| ------------ | ----------- | --------------------------- | -------------------------- |
| Function     | Uppercase   | ✅ Yes                      | `fmt.Println()`            |
| Function     | Lowercase   | ❌ No                       | `fmt.printf()` (not valid) |
| Struct       | Uppercase   | ✅ Yes                      | `type Book struct`         |
| Struct Field | Uppercase   | ✅ Yes                      | `Book.Title`               |
| Struct Field | Lowercase   | ❌ No                       | `Book.title`               |
| Constant     | Uppercase   | ✅ Yes                      | `const Pi = 3.14`          |
| Constant     | Lowercase   | ❌ No                       | `const pi = 3.14`          |

## Some Points

- Allowed to use walrus operator (:=) inside main() but not allowed globally
- There is No Try Catch Method but has Comma Ok Syntax (input , err) like this
- No Inheritance , No Super / Child / Parent
- There is goto labels
- `defer` concept which is important keyPoint (as soon as defer then LIFO execution)
- `packages should be in lowercase`
- `always field startwith Uppercase can be exported else not exported to other packages`
- `main.go` cant be export / import
- `marshal` & `marshalIndent` is very important for golang json
- `map[string]interface{}` → key is string, value can be any type.
- `Thread` Managed By OS (1MB)
- `Goroutines` managed by runtime(2KB) (Parallelism)

### Context in MongoDB operations

- Why context:
  - MongoDB operations are I/O operations; they might take time.
  - Context allows you to:
  - Timeout operations (stop them after X seconds)
  - Cancel if needed
  - Pass metadata if required

| Concept                 | Meaning                        | Example                                                    |
| ----------------------- | ------------------------------ | ---------------------------------------------------------- |
| `context.Background()`  | Root context                   | in `main()` or server start                                |
| `context.TODO()`        | Temporary placeholder          | prototype code                                             |
| `context.WithTimeout()` | Auto cancel after duration     | `context.WithTimeout(context.Background(), 5*time.Second)` |
| `ctx.Done()`            | Channel closed when cancelled  | `<-ctx.Done()`                                             |
| `ctx.Err()`             | Error explaining why cancelled | `context.DeadlineExceeded` or `context.Canceled`           |

### Goroutines

| Concept            | Explanation                                                 |
| ------------------ | ----------------------------------------------------------- |
| Goroutine          | Lightweight thread managed by Go runtime                    |
| Main Goroutine     | The first goroutine — runs `main()`                         |
| Main exits early   | All other goroutines killed instantly                       |
| To wait for others | Use `sync.WaitGroup` or `channel` synchronization           |
| Runtime scheduler  | Decides which goroutine runs on which OS thread (M:N model) |

### Mutex

- Lock And UnLock the Shared Memory So that at one Time Only One GoRoutine can Access and modifiy that Share Memory

| Function       | Use                                                                                      |
| -------------- | ---------------------------------------------------------------------------------------- |
| `mut.Lock()`   | Current goroutine ko write permission deta hai (exclusive access)                        |
| `mut.Unlock()` | Permission release karta hai (next goroutine likh sakta hai)                             |
| Without it     | Multiple goroutines ek hi variable pe likh kar crash ya inconsistent data laa sakti hain |

- Example
  | Time | Goroutine | Action | Comment |
  | ---- | --------- | ---------------------- | ------------------ |
  | t1 | G1 | `mut.Lock()` | G1 ne lock le liya |
  | t2 | G1 | `append("google.com")` | G1 likh raha hai |
  | t3 | G2 | wait (blocked) | Lock occupied |
  | t4 | G1 | `mut.Unlock()` | Lock release |
  | t5 | G2 | `mut.Lock()` | Ab G2 likhta hai |
  | t6 | G3 | wait (blocked) | G2 ka wait |
  | t7 | G2 | `append("github.com")` | Safe write |
  | t8 | G2 | `mut.Unlock()` | Lock release |
  | t9 | G3 | `mut.Lock()` | G3 likhta hai |

## Channels

- Simple Meaning is , To Send Any Type Of Data Between the Go Routines we need Channels
- Unbuffered channel means — no internal storage, sender aur receiver dono ready hone chahiye → automatic synchronization hota hai.
- Buffered channel means — limited queue hai, synchronization manual ya buffer size pe depend karta hai.
- Channel is a reference type, make(chan int) returns reference, not address → &myCh likhna galat / unnecessary hai.
- chan<- / <-chan bas direction restrict karta hai (send-only / receive-only).

- Multiple Go Routines can Talk with Each other
  | Concept | Description |
  | ------------------- | ----------------------------------------- |
  | `chan int` | Channel of integers |
  | `make(chan int)` | Creates **unbuffered** channel |
  | `make(chan int, N)` | Creates **buffered** channel with size N |
  | `chan<- int` | Send-only channel |
  | `<-chan int` | Receive-only channel |
  | `ch <- 10` | Send 10 to channel |
  | `<-ch` | Receive from channel |
  | `close(ch)` | Close channel (no more sends allowed) |
  | `<-ch` after close | Returns zero value + `ok=false` |
  | Synchronization | Coordination between goroutines |
  | `WaitGroup` | Ensures all goroutines finish before exit |

1. Listeners default value 0 if not sending any routine
2. Senders Directly Gives Deadlock issue if not anyone listening

| Concept         | Unbuffered Channel                     | Buffered Channel                 |
| --------------- | -------------------------------------- | -------------------------------- |
| Communication   | Direct handoff between sender/receiver | Through a queue of capacity `N`  |
| Blocking Rule   | Send blocks until recv ready           | Send blocks only if buffer full  |
| Receive Rule    | Recv blocks until send ready           | Recv blocks only if buffer empty |
| Data stored?    | No                                     | Yes (temporary)                  |
| Synchronization | Strong (strict timing)                 | Loose (decoupled timing)         |
| Use Case        | Real-time signaling                    | Producer-consumer queues         |
