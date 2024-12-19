# Go Interview Preparation Notes
#### [OOP Concepts in Golang](#oop-concepts-in-golang-1)
#### [Reading Resources](#sources)
#### [Unique features of Go](#go-unique-features)
#### [Most Common Interview Topics](#most-common-interview-topics-1)
#### [Some advanced Go concepts](#some-advanced-go-concepts-1)
---

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
if still not clear, checkout this blog about interface why to use them : https://dev.to/githaiga22/how-i-stopped-worrying-and-learned-to-love-go-interfaces-3m7p?ref=dailydev
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

### 8. **Embedding (Composition over Inheritance)**

Instead of inheritance, Go uses embedding to achieve similar behaviour. A struct can include another struct, and it will inherit its fields and methods.
```go
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedding Person struct
	Salary int
}

func main() {
	e := Employee{Person: Person{Name: "John", Age: 30}, Salary: 50000}
	fmt.Println(e.Name)   // Accessing embedded struct's field
	fmt.Println(e.Salary) // Accessing Employee's own field
}
```

## Abstract Class vs Interface in Golang
- Go uses **interfaces** for abstraction. There are no abstract classes in Go, but embedding and interfaces can be used to extend behaviour.

### Summary

- **Structs** in Go are used instead of classes and are the way you define your types.
- **Methods** are associated with types and can be defined on structs to encapsulate behavior.
- **Interfaces** are used to achieve polymorphism. They define a set of methods that a type must implement.
- **Polymorphism** is achieved by writing functions that accept interfaces rather than concrete types.
- **Embedding** allows for a form of inheritance, where one struct can embed another and gain its fields and methods.
- **Method overloading** is not supported directly in Go, but similar behavior can be achieved by using different method names or leveraging interfaces.

## Go Unique features
### Concurrency in Go: Goroutines, Channels, and Synchronization

Concurrency is one of Go's standout features. It is achieved through **goroutines** and **channels**, providing efficient handling of multiple tasks in parallel. Synchronization can be managed using the **sync** package (e.g., `WaitGroup`, `Mutex`).

1. **Goroutines**:
   - More memory-efficient lightweight threads managed by the Go runtime, runs concurrently with other functions.
   - Started using the `go` keyword, making them simpler and more efficient than traditional threads.
   - They use a concept called "goroutine stack," which grows and shrinks as needed, making them more lightweight.
   - Goroutines are **non-blocking**.
   - By default, Go uses **one logical processor**. This can be increased using `runtime.GOMAXPROCS()`.
   ```go
   func sayHello() {
       fmt.Println("Hello, Goroutine!")
   }

   func main() {
       go sayHello()  // Starts a new goroutine
       fmt.Println("Main function")
   }
   ```
> **Note**: Goroutines do not run in order; they execute independently. We use `time.Sleep` here to simulate a delay, but real-world cases should use proper synchronization mechanisms.
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
> **Note**: Channels block until both sending and receiving sides are ready, ensuring proper synchronization.
3. **WaitGroup**:
   - Used to wait for multiple goroutines to finish and ensure the main function doesn't exit before they complete.
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
> **Note**: `wg.Add(n)` specifies how many goroutines you're waiting for. Each `wg.Done()` decrements the counter. `wg.Wait()` waits until the counter reaches zero.

4. **Mutex** (Mutual Exclusion for Data Synchronization)

The **Mutex** is a locking mechanism to ensure **only one goroutine accesses shared data** at a time, preventing race conditions.

**Key Points**:
- A **mutex** locks the critical section of the code.
- Only one goroutine can acquire the lock at a time.

**Example**:
```go
var (
	counter int
	mu      sync.Mutex
)

func increment() {
	// Locking the critical section
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	wg.Add(2)
	go func() {  // Start goroutines that access shared data
		for i := 0; i < 1000; i++ {
			increment()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			increment()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
```

> **Note**: Without `Mutex`, race conditions can occur when multiple goroutines try to modify the same shared variable.
---
### Best Use Cases:

1. **Concurrency**:
   - Efficiently managing multiple tasks (e.g., handling multiple network connections).
   - Performing CPU-bound or I/O-bound tasks without blocking.

2. **Channels**:
   - Coordinating work between multiple goroutines.
   - Communicating results of asynchronous operations.

3. **WaitGroups**:
   - Waiting for multiple concurrent operations to complete.
   - Avoiding using `time.Sleep` or busy waiting to synchronize.

4. **Interfaces**:
   - Defining reusable, flexible code by focusing on behavior instead of concrete types.
   - Enabling polymorphic behavior (writing functions that can accept different types).

5. **Mutex**:
   - Ensuring data integrity when multiple goroutines modify shared data concurrently.
   - Preventing race conditions in critical sections.

### Goroutine Management: Avoiding Memory Leaks, Proper Use of Channels

In Go, **goroutines** are an essential feature for writing concurrent code. However, if not managed properly, they can lead to **memory leaks** or resource exhaustion. Let's explore how to avoid common pitfalls.

---

#### 1. **Avoiding Memory Leaks with Goroutines**

Memory leaks happen when goroutines are left hanging or blocked indefinitely. Some common causes include:

- **Unclosed channels**: Goroutines may be blocked forever waiting on a channel.
- **Blocked operations**: If a goroutine tries to send/receive from a channel and no other goroutine is available to handle the operation.

**Example of Memory Leak**:
```go
package main

import (
	"fmt"
	"time"
)

func leakyGoroutine(ch chan int) {
	time.Sleep(2 * time.Second)
	// Trying to send to a channel that's never read from
	ch <- 10 // Goroutine blocked here forever!
}

func main() {
	ch := make(chan int)
	go leakyGoroutine(ch)
	fmt.Println("Main function exits without reading from channel.")
}
```

In the above example, the goroutine is waiting to send to the channel, but no one is reading from it, resulting in a **blocked goroutine**.

---

#### 2. **Proper Use of Channels**

To avoid memory leaks and deadlocks, ensure that channels are correctly managed.

- **Always close channels** when you're done sending values. This ensures that the receiving goroutine knows when no more data is coming.
- **Avoid blocking goroutines** by making sure every send operation has a corresponding receive, and vice versa.

**Example: Proper Channel Usage**:
```go
package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	defer close(ch) // Close the channel when done

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	ch := make(chan int)

	// Start the worker goroutine
	go worker(ch)

	// Receive from the channel until it's closed
	for value := range ch {
		fmt.Println("Received:", value)
	}
	fmt.Println("Channel closed, no more data.")
}
```

> **Note**: Closing the channel signals to the receiver that no more values will be sent, avoiding potential memory leaks or deadlocks.

---

### Errors and Error Handling in Go

In Go, **error handling** is done explicitly using the `error` type. Go doesn't use exceptions like many other languages, making error handling predictable and consistent.

#### 1. **Basic Error Handling**
Go functions return multiple values, with the last return value often being an `error` to signal if something went wrong.

**Example**:
```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

> **Note**: This approach forces the programmer to handle errors at every point where something can go wrong, ensuring that errors are not silently ignored.

#### 2. **Custom Error Types**

Go allows you to define **custom error types** to give more context about the error.

**Example**:
```go
package main

import (
	"fmt"
)

// Custom error type
type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Custom Error: %s", e.Msg)
}

func main() {
	err := &MyError{Msg: "Something went wrong!"}
	fmt.Println(err)
}
```

> **Note**: Custom errors allow you to embed more information about what went wrong, making debugging easier.

---

### `context` Package: Managing Timeouts and Cancellation

The `context` package provides a way to manage **timeouts**, **deadlines**, and **cancellation** of operations across goroutines.

#### 1. **Why Use `context`?**
- **Timeouts**: Context can set a deadline for operations that should not exceed a certain time.
- **Cancellation**: Context can signal when an operation should be canceled (e.g., user cancellation or request timeout).
- **Propagation**: The same context can be passed down through multiple functions and goroutines, allowing for consistent cancellation and timeout behavior.

#### 2. **Example: Timeout using `context.WithTimeout`**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a 1-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Simulate a long-running operation
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}
```

> **Note**: The context's `Done` channel is used to signal that the operation has either completed or been canceled.

#### 3. **Example: Cancellation with `context.WithCancel`**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work canceled:", ctx.Err())
	}
}

func main() {
	// Create a cancelable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start the work
	go doWork(ctx)

	// Simulate user cancellation after 2 seconds
	time.Sleep(2 * time.Second)
	cancel() // Cancels the context

	// Give the goroutine time to finish
	time.Sleep(1 * time.Second)
}
```

> **Note**: The `cancel()` function is used to stop the goroutine early, preventing unnecessary work.

---

### `net/http`: Building HTTP Servers in Go

Go's `net/http` package is a powerful standard library for building HTTP servers and clients.

#### 1. **Building a Simple HTTP Server**

Creating a basic HTTP server with a few endpoints is straightforward in Go.

**Example: Basic HTTP Server**:
```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	// Define route and handler
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
```

> **Note**: `http.HandleFunc()` associates a URL path with a handler function. `http.ListenAndServe()` starts the server and listens for incoming requests.

#### 2. **Routing in `net/http`**

To implement more advanced routing, you can use third-party packages like **`gorilla/mux`** or handle routing logic manually.

**Example: Advanced Routing with `gorilla/mux`**:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", r)
}
```

> **Note**: `mux.NewRouter()` provides a flexible way to handle multiple routes with URL parameters, methods, and middleware.

#### 3. **Middleware in `net/http`**

Middleware allows you to add functionality (e

.g., logging, authentication) to your HTTP server in a clean and reusable way. Middleware wraps around the core logic of your HTTP handler functions.

**Example: Middleware for Logging**:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Logger is a middleware that logs requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Pass on the request to the actual handler
		next.ServeHTTP(w, r)
		// Log after the handler has finished
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	// Attach the middleware to the hello handler
	mux.Handle("/", Logger(http.HandlerFunc(helloHandler)))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
```

> **Note**: The `Logger` function wraps the handler logic, allowing you to execute additional code (in this case, logging) before and after the handler processes the request.

#### 4. **HTTP Client in `net/http`**

In addition to building servers, Go’s `net/http` package provides a robust client for making HTTP requests.

**Example: Making HTTP Requests with `net/http`**:
```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make a GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", string(body))
}
```

> **Note**: This example shows how to perform a basic GET request. You can also use `http.Client` for more advanced usage like setting timeouts, headers, and custom configurations.

---

### Summary of Best Practices:

1. **Goroutine Management**:
   - Always ensure goroutines are **unblocked** and channels are **closed** when finished.
   - Use `context` for **timeouts** and **cancellation** to prevent memory leaks from long-running goroutines.

2. **Error Handling**:
   - Handle errors **explicitly** and propagate them up the stack where necessary.
   - Create **custom error types** for more context and clarity in your errors.

3. **Context**:
   - Use `context` for managing **timeouts**, **deadlines**, and **cancellations** in long-running operations.
   - Pass `context` through function signatures for better control over request lifecycle.

4. **`net/http`**:
   - Build simple, performant HTTP servers with Go's standard library.
   - Use middleware to **separate concerns** (e.g., logging, authentication).
   - Make HTTP requests with the `net/http` client for efficient API consumption.
---


## Important Packages to Know
- **fmt**: For formatted I/O (Println, Scanf, etc.).
- **errors**: For error creation and handling.
- **sync**: For synchronization primitives like `WaitGroup`, `Mutex`.
- **context**: For managing request-scoped values and cancellation signals.
- **net/http**: For building HTTP servers and clients.
---

## Sources:
- go-concurrency: goroutines https://antonz.org/go-concurrency/goroutines/

## Most Common Interview Topics
1. **Concurrency**: Understanding goroutines, channels, and synchronization is crucial.
2. **Interfaces & Polymorphism**: The cornerstone of Go’s abstraction mechanism.
3. **Error Handling**: Idiomatic Go error handling using error values and the `errors` package.
4. **Structs & Composition**: How Go achieves behavior extension and "inheritance" through composition.
5. **Goroutine Management**: Proper goroutine usage, channel communication, and avoiding memory leaks.
---

## Some advanced Go concepts

### 1. **Goroutine Leaks and Managing Concurrency**

#### **Common Issues Leading to Goroutine Leaks**
- **Unbuffered Channels**: If a goroutine is waiting on an unbuffered channel and no other goroutine is ready to consume the message, it will block indefinitely.
- **Forgotten goroutines**: Goroutines that run in the background and are never terminated.

#### **Preventing Goroutine Leaks**:
- Use **`sync.WaitGroup`** to ensure that all spawned goroutines complete before exiting.
- Properly handle channel closing, cancellation, and timeouts using the **`context`** package.

**Example: Using `WaitGroup` to avoid goroutine leaks**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the goroutine as done when it exits
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the waitgroup counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Block until all goroutines have finished
	fmt.Println("All workers done")
}
```

> **Advanced Question**: How would you handle a scenario where a goroutine might get stuck waiting for a long-running operation to complete?  
**Answer**: Use a **`context.WithTimeout`** or **`context.WithCancel`** to cancel a goroutine if the operation takes too long.

---

### 2. **Memory Management & Escape Analysis**

Go automatically manages memory using **garbage collection**, but understanding **escape analysis** helps optimize performance by controlling heap vs. stack allocations.

- **Escape Analysis** determines whether a variable can be allocated on the stack or if it needs to "escape" to the heap.
- Stack allocations are faster, and heap allocations require garbage collection, so minimizing heap allocations can improve performance.

#### **Heap vs Stack Example**:
```go
package main

import "fmt"

func escapeToHeap() *int {
	x := 42
	return &x // x escapes to the heap because its address is returned
}

func stackOnly() int {
	y := 42
	return y // y stays on the stack
}

func main() {
	fmt.Println(*escapeToHeap()) // Allocated on heap
	fmt.Println(stackOnly())     // Allocated on stack
}
```

> **Advanced Question**: How do you check if a variable escapes to the heap?  
**Answer**: Use the `go build -gcflags="-m"` command to see where escape analysis moves variables to the heap.

---

### 3. **Channel Buffering and Synchronization Patterns**

#### **Buffered vs Unbuffered Channels**
- **Buffered Channels** allow sending values without blocking until the buffer is full.
- **Unbuffered Channels** block until the message is received.

#### **Best Practices for Channel Buffering**:
- **Buffered Channels** are ideal when you have multiple producers and consumers, and you don’t want producers to be blocked waiting for consumers.
- **Unbuffered Channels** should be used when you want strict synchronization between sender and receiver.

**Example: Using Buffered Channels**
```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Buffered channel with capacity of 2

	ch <- 1 // Doesn't block because buffer is empty
	ch <- 2 // Doesn't block because buffer has space
	fmt.Println(<-ch) // Reads 1
	fmt.Println(<-ch) // Reads 2
}
```

> **Advanced Question**: What happens if you try to send to a full buffered channel?  
**Answer**: The goroutine will block until space is available in the buffer.

---

### 4. **Select Statement**

The `select` statement allows a goroutine to wait on multiple channel operations, enabling **non-blocking communication** and **timeout handling**.

**Example: Using `select` for Timeouts**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
```

> **Advanced Question**: What happens if multiple cases in a `select` block are ready?  
**Answer**: Go will randomly choose one of the ready cases to proceed.

---

### 5. **Error Handling Best Practices**

Go's idiomatic error handling is based on returning errors rather than throwing exceptions.

#### **Best Practices for Error Handling**:
- **Wrap errors** with additional context using `fmt.Errorf` or third-party libraries like `pkg/errors`.
- Handle errors immediately and **propagate them upwards** if needed.

**Example: Wrapping Errors**
```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division error: %w", errors.New("division by zero"))
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

> **Advanced Question**: How do you check the root cause of an error wrapped with `fmt.Errorf`?  
**Answer**: Use `errors.Is` or `errors.Unwrap` to unwrap and inspect errors.

---

### 6. **Context for Cancellation, Timeouts, and Deadlines**

The `context` package is essential for **controlling goroutine lifecycles**, especially in **web servers**, **long-running operations**, and **API requests**.

#### **Context Types**:
- `context.Background()`: Base context with no deadline or cancellation.
- `context.WithTimeout()`: Adds a timeout to a context.
- `context.WithCancel()`: Allows manual cancellation.

**Example: Using `context.WithTimeout`**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(2 * time.Second)
}
```

> **Advanced Question**: Why is it important to always call `cancel()` when using `context.WithCancel` or `context.WithTimeout`?  
**Answer**: To avoid **resource leaks** by ensuring that context resources are freed once they are no longer needed.

---

### 7. **Advanced `net/http` Features**

#### **Custom HTTP Handlers and Middleware**

Middleware allows you to modify requests and responses, such as logging, authentication, and rate limiting, by wrapping around the core HTTP handlers.

**Example: Logging Middleware**
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Request: %s %s took %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.Handle("/", logger(http.HandlerFunc(helloHandler)))
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
```

> **Advanced Question**: How do you handle long-running requests in `net/http` without blocking the server?  
**Answer**: Use **goroutines** and handle **cancellation** via the **context** passed through `http.Request.Context()`.

---

### 8. **Reflection in Go**

Reflection allows you to inspect the type and value of variables at runtime using the `reflect` package. This is powerful for **building frameworks**, **serialization**, and **dynamic dispatch**.

**Example: Using Reflection to Inspect Types**
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("Type:", reflect.TypeOf(x))
	fmt.Println("Value:", reflect.ValueOf(x))
}
```

> **Advanced Question**: When would you use reflection in production code?  
**Answer**: Reflection should be used sparingly due to performance costs, but it can be useful for **serialization**, **JSON marshalling**, and **dynamic routing** in frameworks.

---

