# Dependency management

  - `go mod init myapp`  Create a new module
  - `go get` fetches modules by semantic version (v1.2.3) or pseudo-versions.  
  - `go mod tidy` removes unused entries and adds missing ones.  
  - `go mod verify` checks that downloaded modules match their cryptographic checksums.  
  - Upgrades use minimal version selection: projects only upgrade to the lowest version required by any dependency graph.


# Creating a New Local Go Project with Tests

## 1. Install Go

1. Download the installer from [https://golang.org/dl/](https://golang.org/dl/)  
2. Follow the installation instructions for your OS  
3. Verify installation:
   ```bash
   go version
   ```

## 2. Create Your Project Directory
```bash
mkdir ~/go-projects/myapp
cd ~/go-projects/myapp
```
It’s best to work outside of  `$GOPATH` when using modules.

## 3. Initialize a Module
```bash
go mod init github.com/yourusername/myapp
```
This creates a `go.mod` file with your module path.

## 4. Write Your Code

`main.go`
```go
package main

import "fmt"

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println("2 + 3 =", Add(2, 3))
}
```

## 5. Run Your Application

```bash
go run main.go
# Output: 2 + 3 = 5
```


## 5. Add a Basic Test

`main_test.go`

```go
package main

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

## 7. Run Your Tests

```bash

go test ./...
# PASS

```

## 8. Add Testify and Rework Tests

1. Install Testify:

    ```bash
    go get github.com/stretchr/testify@latest
    ```

2. Update `main_test.go` to use 
`assert`:

    ```go
    package main

    import (
        "testing"
        "github.com/stretchr/testify/assert"
    )

    func TestAdd(t *testing.T) {
        assert.Equal(t, 5, Add(2, 3), "2 + 3 should equal 5")
    }
    ```

3. Run tests again:

    ```bash
    go test
    # PASS
    ```
