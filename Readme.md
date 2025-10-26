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

| Feature         | Array `[N]int`      | Slice `[]int`                |
| --------------- | ------------------- | ---------------------------- |
| Size            | Fixed               | Dynamic                      |
| Can append?     | ❌ No               | ✅ Yes                       |
| Memory          | Contiguous          | Points to underlying array   |
| Passing to func | Copies entire array | Passes reference (efficient) |
| Type uniqueness | `[4]int] ≠ [5]int`  | `[]int` same type always     |

| Feature         | `new(Type)`                   | `make(Type, ...)`                          |
| --------------- | ----------------------------- | ------------------------------------------ |
| Returns         | Pointer (`*Type`)             | Initialized value (not pointer)            |
| Purpose         | Allocate memory               | Allocate + initialize                      |
| Zero value      | Yes                           | Not needed, already usable                 |
| Types supported | Any type (struct, int, array) | Only reference types (slice, map, channel) |
| Usage           | `p := new(int)`               | `s := make([]int, 5)`                      |

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
