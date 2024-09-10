For the Low-Level Design (LLD) round, you will need to focus on designing scalable, flexible systems with clean and maintainable code.

### Steps to prepare for LLD:
1. **Understand the Requirements:** Break down the problem into smaller components.
2. **Identify Classes and Objects:** Recognize the key entities and their relationships.
3. **Define Methods and Interactions:** Think about how the objects interact (methods and data flow).
4. **Consider Design Principles:** Follow SOLID principles, design patterns, and separation of concerns.
5. **Code Modularity and Reusability:** Ensure classes and methods are modular and reusable.

Let’s go through some common LLD problems:

### Problem 1: Design a Parking Lot System
**Requirements:**
- Multiple parking levels, each level with multiple parking spots.
- Parking spots can be small, medium, or large.
- Vehicles can be motorcycles, cars, or buses.
- Assign the correct spot size to each vehicle.
- Track free and occupied spots.

#### Approach:
1. **Entities:** Vehicle, ParkingSpot, ParkingLot, ParkingLevel.
2. **Classes:**
   - **Vehicle:** Has `type` (motorcycle, car, bus).
   - **ParkingSpot:** Has `size` (small, medium, large).
   - **ParkingLot:** Manages levels and spots.
   - **ParkingLevel:** Manages spots on a level.

Here’s a basic structure in Go:

```go
package main

import "fmt"

// Vehicle interface for different vehicle types
type Vehicle interface {
    GetSize() string
}

type Car struct{}
type Motorcycle struct{}
type Bus struct{}

// Car size method
func (c *Car) GetSize() string {
    return "medium"
}

// Motorcycle size method
func (m *Motorcycle) GetSize() string {
    return "small"
}

// Bus size method
func (b *Bus) GetSize() string {
    return "large"
}

// ParkingSpot struct
type ParkingSpot struct {
    size     string
    occupied bool
}

// ParkingLevel struct, manages parking spots
type ParkingLevel struct {
    spots []ParkingSpot
}

// ParkingLot struct
type ParkingLot struct {
    levels []ParkingLevel
}

// Main function to demonstrate parking
func main() {
    car := &Car{}
    motorcycle := &Motorcycle{}
    fmt.Println("Car Size:", car.GetSize())           // Output: Car Size: medium
    fmt.Println("Motorcycle Size:", motorcycle.GetSize()) // Output: Motorcycle Size: small
}
```

### Problem 2: Design a Library Management System
**Requirements:**
- Manage books, members, and transactions (issuing, returning books).
- Different book categories.
- Track availability and issue limits.

#### Approach:
1. **Entities:** Book, Member, Librarian, Transaction.
2. **Classes:**
   - **Book:** Title, author, ISBN, category.
   - **Member:** Issue and return books.
   - **Transaction:** Records issued and returned books.\
> [!TIP]
> Practice this by yourself using the above approach, As hands-on is a Must!!!

### Problem 3: Design an Online Movie Booking System
**Requirements:**
- Allow users to search for movies, book tickets, and choose seats.
- Theaters have multiple screens and multiple shows per day.
- Track seat availability for each show.

#### Approach:
1. **Entities:** Movie, Screen, Theater, Booking, User.
2. **Classes:**
   - **Movie:** Title, duration, language, rating.
   - **Theater:** Manages screens and shows.
   - **Booking:** Stores booking details for users.
