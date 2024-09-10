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
> [!TIP]
> Practice this by yourself using the above approach, As hands-on is a Must!!!
[ans](#answer-of-online-movie-booking-system)
#### Answer of Library Management System
```go
package main

import (
	"fmt"
	"time"
)

// Book struct with more attributes
type Book struct {
	title       string
	author      string
	isbn        int
	category    string
	isAvailable bool
}

// Member struct with ID and borrowed books
type Member struct {
	memberID      int
	name          string
	borrowedBooks []*Book
}

// Transaction struct to track issue and return dates
type Transaction struct {
	book       *Book
	member     *Member
	issueDate  time.Time
	returnDate *time.Time // Pointer to handle no return date initially
}

// Library struct to manage transactions
type Library struct {
	transactions []*Transaction
}

// Method for a member to issue a book and create a transaction
func (l *Library) IssueBook(member *Member, book *Book) {
	if book.isAvailable {
		member.borrowedBooks = append(member.borrowedBooks, book)
		book.isAvailable = false
		transaction := &Transaction{
			book:      book,
			member:    member,
			issueDate: time.Now(),
		}
		l.transactions = append(l.transactions, transaction)
		fmt.Printf("Book '%s' issued to %s on %s.\n", book.title, member.name, transaction.issueDate)
	} else {
		fmt.Println("Book is not available.")
	}
}

// Method for a member to return a book and update the transaction
func (l *Library) ReturnBook(member *Member, book *Book) {
	for i, borrowedBook := range member.borrowedBooks {
		if borrowedBook.isbn == book.isbn {
			member.borrowedBooks = append(member.borrowedBooks[:i], member.borrowedBooks[i+1:]...)
			book.isAvailable = true
			for _, transaction := range l.transactions {
				if transaction.book.isbn == book.isbn && transaction.member.memberID == member.memberID && transaction.returnDate == nil {
					returnDate := time.Now()
					transaction.returnDate = &returnDate
					fmt.Printf("Book '%s' returned by %s on %s.\n", book.title, member.name, *transaction.returnDate)
					return
				}
			}
		}
	}
	fmt.Println("Book not found in borrowed list.")
}

// Main function to demonstrate the library system
func main() {
	// Create a library instance
	library := &Library{}

	// Create a new book
	book := &Book{
		title:       "Harry Potter",
		author:      "J.K. Rowling",
		isbn:        123456,
		category:    "Fantasy",
		isAvailable: true,
	}

	// Create a new member
	member := &Member{
		memberID: 1,
		name:     "John Doe",
	}

	// Member issues a book
	library.IssueBook(member, book)

	// Member returns a book
	library.ReturnBook(member, book)
}
```

This implementation tracks both book borrowing and returning along with their respective dates, making it more realistic for a library system.

#### Answer of Online Movie Booking System
##### Key Points:
1. **Encapsulation:** Avoid exposing internal fields like `remaining_seats`, and use getter/setter methods instead.
2. **Missing Methods:** Add methods to manage bookings, seat availability, and user interaction.
3. **Additional Entities:** Consider adding `User`, `Show`, and `Payment` entities for a more complete system.
4. **Relationships:** Associate `Theatre` with `Show` instead of directly associating with `Movie`. Multiple shows for the same movie can run at different times on different screens.
```go
package main

import "fmt"

// Movie entity
type Movie struct {
	title    string
	id       int
	duration int // in minutes
	language string
	rating   float32
}

// Screen entity
type Screen struct {
	id              int
	totalSeats      int
	remainingSeats  int
	shows           []*Show
}

// Show entity, linking Movie and Screen
type Show struct {
	movie *Movie
	time  string
	screen *Screen
}

// Theatre entity
type Theatre struct {
	name    string
	screens []*Screen
}

// Booking entity
type Booking struct {
	show  *Show
	seats int
}

// Method to book seats for a show
func (s *Screen) bookSeats(seats int) bool {
	if s.remainingSeats >= seats {
		s.remainingSeats -= seats
		return true
	}
	return false
}

// Create a new booking
func (b *Booking) createBooking(seats int) bool {
	if b.show.screen.bookSeats(seats) {
		b.seats = seats
		fmt.Printf("Successfully booked %d seats for movie '%s' at %s\n",
			seats, b.show.movie.title, b.show.time)
		return true
	}
	fmt.Println("Failed to book seats")
	return false
}

// Main function to demonstrate
func main() {
	// Create Movies
	movie1 := &Movie{title: "Inception", id: 1, duration: 148, language: "English", rating: 8.8}
	movie2 := &Movie{title: "Titanic", id: 2, duration: 195, language: "English", rating: 7.8}

	// Create Screen
	screen1 := &Screen{id: 1, totalSeats: 100, remainingSeats: 100}
	screen2 := &Screen{id: 2, totalSeats: 200, remainingSeats: 200}

	// Create Shows
	show1 := &Show{movie: movie1, time: "7:00 PM", screen: screen1}
	show2 := &Show{movie: movie2, time: "8:00 PM", screen: screen2}
	screen1.shows = append(screen1.shows, show1)
	screen2.shows = append(screen2.shows, show2)

	// Create Theatre
	theatre := &Theatre{name: "Cineplex", screens: []*Screen{screen1, screen2}}

	// Create Booking
	booking1 := &Booking{show: show1}
	booking1.createBooking(5)

	booking2 := &Booking{show: show2}
	booking2.createBooking(210) // Fail due to lack of seats
}
```
This makes the system more modular, extensible, and easier to maintain. You can now expand this by adding user authentication, payment systems, etc.


