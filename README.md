# About Go

## 1. Introduction & Motivation (5 min)

Go is a **statically typed, compiled, imperative** language with strong support for **procedural** and **concurrent** programming. It borrows a few functional ideas (first-class functions, closures), but its core design is **not** functional.

Why Go?
– Origin story at Google
– Target domains: cloud services, microservices, CLI tools

## 2. Init Project & go.mod (10 min)
Modules vs GOPATH
– Pre-1.11 GOPATH baggage: global workspace, import path headaches
– go mod init, go.mod file anatomy (module path, go 1.x, require, replace)

Dependency management
– go get, semantic versioning, tidy/verify (go mod tidy, go mod verify)
– Version upgrades, minimal version selection

Practical demo snippet

```bash
cd ~/projects/myapp
go mod init github.com/you/myapp
go get github.com/gin-gonic/gin@v1.9.0
go mod tidy
```

## 2. Unusual Go Features

1. **Block-scoped variables in `if`/`for`**  
   - You can declare a variable **inside** an `if` or `for` header:  
     ```go
     if n, err := strconv.Atoi(s); err != nil {
         // n and err exist only in this if-block
         log.Println("bad number:", err)
         return
     }
     // n and err are out of scope here
     ```  
   - **Contrast with PHP/JS**: those languages hoist or leak variables to the outer scope; Go keeps them local to the block.

2. **Multiple return values**  
   - Functions can return as many results as you like.  
   - Common pattern: return a value **and** an `error`:  
     ```go
     func ReadConfig(path string) (cfg Config, err error) { … }
     ```  
   - Callers write  
     ```go
     cfg, err := ReadConfig("config.yaml")
     if err != nil {
         log.Fatal(err)
     }
     ```
   - No tuples or destructuring—just positional returns.

3. **Structs & their definitions**  
   - Go’s primary composite type is the `struct`. It groups named fields:  
     ```go
     // Person — defines a person with a name and age.
     type Person struct {
         Name string
         Age  int
     }
     ```  
   - You can instantiate and use it by value or pointer:  
     ```go
     p := Person{Name: "Alice", Age: 30}
     p.Age = 31

     // or
     p2 := &Person{"Bob", 25}
     p2.Name = "Robert"
     ```  
   - Structs support embedding for simple composition instead of inheritance:
     ```go
     type Employee struct {
         Person      // embeds all Person fields and methods
         Position string
     }
     ```

4. **Explicit pointers and value vs. reference semantics**  
   - You must use `*`/`&` to work with pointers; there’s no hidden reference magic.  
   - Example:  
     ```go
     type User struct { Name string }
     func Rename(u *User, newName string) {
         u.Name = newName
     }
     u := User{"Alice"}
     Rename(&u, "Bob")
     ```  
   - This explicitness avoids surprises from implicit references in other languages.

5. **Explicit error handling**  
   - Go has **no exceptions**—errors are just values you check explicitly:  
     ```go
     result, err := DoSomething()
     if err != nil {
         // handle or return
     }
     ```
   - Keeps control flow clear and pushes you to consider errors at every step.

6. **Zero values & explicit initialization**  
   - Every type has a well-defined “zero” default—no `null` or `undefined`:  
     ```go
     var s string   // ""  
     var i int      // 0  
     var m map[string]int // nil, but safe to compare  
     ```
   - You only allocate when you need to:  
     ```go
     m := make(map[string]int)  
     ```

 7. **Coding Style: Short, Contextual Variable Names**
    - **Idiomatic brevity**  
  In Go it’s perfectly fine—and common—to use very short variable names when the context makes their meaning clear.  

  ```go
  for i := 0; i < len(users); i++ {
      u := users[i]
      fmt.Println(u.Name)
  }

  // vs. a more verbose but redundant approach
  for index := 0; index < len(users); index++ {
      user := users[index]
      fmt.Println(user.Name)
  }
  ```
   - Why short names are OK
     - Locality of scope
        - When `i` or `u` appear inside a small loop or function, you immediately see their role.
     - Readability over verbosity
        - Overly long names can clutter the code and distract from its logic.
     - Convention
        - `i`, `j`, `k` for loop indices
        - `r`, `w` for Reader/Writer
        - `ctx` for context.Context
   - When to choose longer names
     - In wide scopes (package-level variables, function parameters), prefer descriptive names:
        ```go
        // OK: short in tight scope
        ch := make(chan int)

        // Better at package level
        var requestCount int
        ```
   - Go tooling enforces style
      - `go` `fmt` and linters won’t complain about short names in small scopes—but will flag unused variables (_ if intentionally unused).
      - Follow the community style: concise, clear, and idiomatic.

Learn more:
[Effective Go](https://go.dev/doc/effective_go)
[Google Go Style Guide](https://google.github.io/styleguide/go/)
[Go Naming Conventions](https://www.gotutorial.org/go-tutorial/go-naming-conventions)


## 4. Extended language features (15 min)
 ### 1. Methods & Exported vs. Unexported (5 min)

  - **Methods on structs**  
  In Go you attach functions to types by declaring methods:
  ```go
  type User struct {
    Name string
  } 

  // Method on *User
  func (u *User) Greet() string {
    return "Hello, " + u.Name
  }
  ```

  - Exported vs. unexported
    - Uppercase identifiers (types, fields, methods) are exported—visible to other packages.
    - Lowercase identifiers are unexported—private to their package.

```go
    // Exported:
    type Account struct {
    ID     string // exported field
    secret string // unexported field
    }

    func (a *Account) Balance() float64 { … }  // exported method
    func (a *Account) updateSecret(s string) { … } // unexported method
```
   - Why it matters
     - Enforces encapsulation—only what you choose is part of your package’s public API.
     - Keeps implementation details hidden.

### 2. Duck-Typing via Interfaces (7 min)
Implicit interfaces
– No “implements” keyword: satisfaction by method set
– Example:

```go
type Reader interface { Read(p []byte) (n int, err error) }
func Copy(dst Writer, src Reader) { … }
```
**Best practice**
- Small, single-method interfaces
- Use interfaces in API boundaries, not everywhere

Use case: testing/mocking by substituting implementations

## 5. Brief Mention: Concurrency (Goroutines & Channels) (5 min)
Only a quick recap: “Go’s built-in concurrency model uses goroutines (lightweight threads) and channels (typed pipes). We won’t dive deep today, but remember: channels are not the only way—prefer sync or context for coordination when appropriate.”

## 6. Ecosystem Highlights (Gin & Testing with Testify) (10 min)
### 1. Gin Web Framework (5 min)
**Why Gin?**

– Fast HTTP multiplexer, middleware support, JSON validation, routing groups

Sample “Hello, Gin”

```go
r := gin.Default()
r.GET("/ping", func(c *gin.Context) {
  c.JSON(200, gin.H{"message": "pong"})
  })
r.Run() // :8080
```
Middleware & grouping
– Authentication, logging, CORS

### 2. Testing with Testify (5 min)
Built-in testing package
– go test, table-driven tests

Testify for assertions & mocks

```go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)
func TestAdd(t *testing.T) {
  assert.Equal(t, 4, Add(2,2), "2+2 should be 4")
}
```

## 7. Go vs Classical OOP Languages (10 min)
**No classes or inheritance**
- Composition over inheritance: embed types

**Interfaces vs abstract base classes**
- Implicit satisfaction vs explicit extension

**Generics**
- Brief mention: type parameters, when they help

**Memory model & GC**
- Simpler object layout, pointer-driven; Go GC tuned for latency

## 8. Live Q&A / Hands-On Challenge (5 min)
Offer a small coding challenge
– “Add a new endpoint to our Gin app that reads from an interface-based service and returns JSON.”

Answer questions

## 9. Resources (5 min)

Further learning
 - [Official Tour ](https://go.dev/doc/tutorial/getting-started)
 - [Effective Go](https://go.dev/doc/effective_go), [Go by Example](https://gobyexample.com/)
