package main

import (
  "fmt"
  "time"
)

func sayHello(name string) {
  for i := 0; i < 3; i++ {
    fmt.Printf("Hello, %s! (%d)\n", name, i)
    time.Sleep(100 * time.Millisecond)
  }
}

func main() {
  // Start sayHello in a new goroutine
  go sayHello("Alice")

  // Meanwhile, main continues running
  sayHello("Bob")

  // Give the goroutine time to finish
  time.Sleep(500 * time.Millisecond)
  fmt.Println("Done")
}
