# Go Interview Preparation Notes

## OOP Concepts in Golang
Go is not a traditional Object-Oriented Programming (OOP) language, However, Go still supports some key OOP concepts, such as encapsulation and polymorphism, 
using structs, interfaces, and methods.\
These features make Go a powerful language that supports some aspects of OOP, while also embracing simplicity and a more functional approach.
### Basic Concepts
### 1. **Structs**

Structs in Go are similar to classes in other languages, but they are more lightweight. A struct is a composite data type that groups together variables under a single name.

#### Example:

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p.Name) // Accessing struct fields
}
```

### 2. **Methods**

In Go, you can define methods on types (including structs) which allows for encapsulation.

#### Example:

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Method on Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	p.Greet() // Calling the method
}
```

### 3. **Interfaces**

An interface in Go is a collection of method signatures. Any type that implements all the methods in the interface is said to implement that interface.

#### Example:

```go
package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	var a Animal

	dog := Dog{}
	a = dog
	fmt.Println(a.Speak()) // "Woof!"

	cat := Cat{}
	a = cat
	fmt.Println(a.Speak()) // "Meow!"
}
```
### 4. **Abstraction**:
   Hiding the implementation details from the user and showing only the functionality of the programming to the user.
   - In Go, abstraction is achieved using **interfaces**. 
   - Interfaces define behaviour but don't provide implementation details.
   - Types implement interfaces by defining the required methods without explicitly stating the relationship.

   ```go
   // Interface example - Abstraction
   type Shape interface {
       Area() float64
   }

   type Circle struct {
       Radius float64
   }

   // Implementing the Area method of the Shape interface for Circle
   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }

   func main() {
       var shape Shape = Circle{Radius: 5}
       fmt.Println("Area of Circle:", shape.Area())
   }
   ```

---

### 5. **Inheritance (via Composition)**: In which one object acquires all the properties from the parent object.
   - Go doesn’t support classical inheritance like Java but promotes **composition**.
   - A type can include another type and use its fields and methods.

   ```go
   // Composition - An alternative to inheritance in Go
   type Person struct {
       Name string
   }

   type Employee struct {
       Person    // Embedding Person struct
       Position string
   }

   func main() {
       emp := Employee{Person: Person{Name: "John"}, Position: "Engineer"}
       fmt.Println("Employee Name:", emp.Name) // Accessing the Name field from the embedded Person struct
   }
   ```

---

### 6. **Polymorphism**: If a task can be done in multiple ways.
   - Go achieves polymorphism through **interfaces**.
   - You can write functions that work with any type that implements a given interface.

   ```go
   // Polymorphism - Go interfaces
   type Animal interface {
       Speak() string
   }

   type Dog struct {}

   func (d Dog) Speak() string {
       return "Woof!"
   }

   type Cat struct {}

   func (c Cat) Speak() string {
       return "Meow!"
   }

   func makeNoise(a Animal) {
       fmt.Println(a.Speak())  // Polymorphic behavior - different implementations of Speak
   }

   func main() {
       dog := Dog{}
       cat := Cat{}
       makeNoise(dog)  // Outputs: Woof!
       makeNoise(cat)  // Outputs: Meow!
   }
   ```

   **Types of Polymorphism**:
   - **Runtime Polymorphism**: Achieved via method overriding using interfaces.
   - **Compile-Time Polymorphism**: Go does not support method overloading, so it doesn't have compile-time polymorphism. You can simulate overloading by using different function names.

   ```go
   // Function overloading (not directly supported in Go)
   func addInt(a int, b int) int {
       return a + b
   }

   func addFloat(a float64, b float64) float64 {
       return a + b
   }
   ```

---

### 7. **Encapsulation**: fundamental concept that governs how data and methods are organized and accessed within different packages.
   - languages like Java, it typically involve classes and access modifiers (e.g., private, protected, public). In Go, encapsulation is more
     about controlling visibility and organizing code within and across packages.
   - If a field or method starts with a capital letter, it’s public (exported); otherwise, it’s private (unexported).

   ```go
   // Encapsulation - Unexported (private) and Exported (public) fields/methods
 
func privateFunction() {   // Unported function as name starts with small p
    // Code here
}

type Person struct { // Exported type
    Name string // Exported field
    age  int    // Unexported field
}

// Constructor function (often exported)
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, age: age}
}
   ```
Packages are the primary unit of encapsulation in Go.
you can expose only the parts of your code that you want to be public(func name starts with capital), while keeping the internal implementation details 
hidden(place code in func keeping name starts with a small letter).
---

## Abstract Class vs Interface in Golang

- Go uses **interfaces** for abstraction. There are no abstract classes in Go, but embedding and interfaces can be used to extend behaviour.

---
### 8. Method Overloading
- Go does not support method overloading natively. You can use different function names to achieve similar functionality.

   ```go
   // Simulating method overloading by using different function names
   func addInt(a int, b int) int {
       return a + b
   }

   func addFloat(a float64, b float64) float64 {
       return a + b
   }
   ```

---

## Go-Routines and Concurrency

1. **Goroutines**:
   - Lightweight threads managed by Go’s runtime.
   - Started using the `go` keyword, making them simpler and more efficient than traditional threads.

   ```go
   func sayHello() {
       fmt.Println("Hello, Goroutine!")
   }

   func main() {
       go sayHello()  // Starts a new goroutine
       fmt.Println("Main function")
   }
   ```

2. **Channels**:
   - Used for communication between goroutines.
   - Channels are typed and allow goroutines to send and receive values in a synchronized manner.

   ```go
   func main() {
       ch := make(chan int)  // Create a channel of int type
       go func() {
           ch <- 42  // Send value into channel
       }()
       fmt.Println(<-ch)  // Receive value from channel
   }
   ```

3. **WaitGroup**:
   - Used to wait for multiple goroutines to finish.
   - Helps in synchronizing goroutines.

   ```go
   var wg sync.WaitGroup

   func worker(id int) {
       defer wg.Done()  // Notify that the goroutine is done
       fmt.Printf("Worker %d done\n", id)
   }

   func main() {
       wg.Add(2)  // Add two goroutines to the WaitGroup
       go worker(1)
       go worker(2)
       wg.Wait()  // Wait for all goroutines to finish
   }
   ```

---

## Important Packages to Know
- **fmt**: For formatted I/O (Println, Scanf, etc.).
- **errors**: For error creation and handling.
- **sync**: For synchronization primitives like `WaitGroup`, `Mutex`.
- **context**: For managing request-scoped values and cancellation signals.
- **net/http**: For building HTTP servers and clients.

---

## Most Common Interview Topics
1. **Concurrency**: Understanding goroutines, channels, and synchronization is crucial.
2. **Interfaces & Polymorphism**: The cornerstone of Go’s abstraction mechanism.
3. **Error Handling**: Idiomatic Go error handling using error values and the `errors` package.
4. **Structs & Composition**: How Go achieves behavior extension and "inheritance" through composition.
5. **Goroutine Management**: Proper goroutine usage, channel communication, and avoiding memory leaks.
