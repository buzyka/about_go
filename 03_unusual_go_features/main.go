package main

import (
	"encoding/json"
	"fmt"
)

func blockSkope() {
	// This function demonstrates block scope in Go.
	// Variables declared inside this function are not accessible outside of it.
	var x int = 10
	if x > 5 {
		var y int = 20 // y is only accessible within this if block
		x += y
	}
	// fmt.Println(y) // This would cause a compile error because y is not defined here.
}

func multipleReturnValues() (int, int) {
	// This function returns two integers.
	return 1, 2
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedding Person struct
	Position  string
}

type EmployeeWithPointer struct {
	Person *Person // Using a pointer to Person struct
	Position string
}

func main() {
	// ------
	fmt.Println("Example 1: Block Scope and Multiple Return Values")
	blockSkope()
	fmt.Println("Block scope example executed.")

	// ------
	fmt.Println("Example 2: Multiple Return Values")
	a, b := multipleReturnValues()
	fmt.Printf("Multiple return values: a = %d, b = %d\n", a, b)

	// ------
	fmt.Println("Example 3: Structs and Embedding")
	p := Person{Name: "Alice", Age: 30}
	p.Age += 1 // Modifying a field of a struct
	fmt.Printf("Person: Name = %s, Age = %d\n", p.Name, p.Age)

	e := Employee{
		Person:   p,
		Position: "Developer",
	}
	fmt.Printf("Employee: Name = %s, Age = %d, Position = %s\n", e.Name, e.Age, e.Position)

	// ------
	fmt.Println("Example 4: Explicit pointers and values")
	p2 := &Person{Name: "Bob", Age: 25} // Using a pointer to Person struct
	e2 := EmployeeWithPointer{
		Person:   p2,
		Position: "Manager",
	}
	fmt.Printf("EmployeeWithPointer: Name = %s, Age = %d, Position = %s\n", e2.Person.Name, e2.Person.Age, e2.Position)
	
	p2.Age += 1 // Modifying a field of a struct through pointer
	fmt.Printf("After incrementing age: Name = %s, Age = %d\n", e2.Person.Name, e2.Person.Age)


	// ------
	fmt.Println("Example 5: Explicit error handling")
	jsonString := `{"name": "Alice", "age": 30` // Missing closing brace
	//jsonString := `{"name": "Bob", "age": 25}` // Correct JSON
	p3 := &Person{}
	if err := json.Unmarshal([]byte(jsonString), p3); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	} else {
		fmt.Printf("Unmarshalled Person: Name = %s, Age = %d\n", p3.Name, p3.Age)
	}

	// ------
	fmt.Println("Example 6: Zero values and explicit initialization")
	var s string // ""
	var i int     // 0
	var f float64 // 0.0
	fmt.Printf("Zero values: string = '%s', int = %d, float64 = %f\n", s, i, f)
}
