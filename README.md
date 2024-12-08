# Booking-app
A simple CLI app built using Go

## Introduction To Go

```
go mod init <module path>
```

This cmd initializes a go.md file which describes the module

creates a new module

module path can correspond to a repository you plan to publish
your module to  (e.g. github.com/utuedey/booking-app)

## Variables in Go

- Variables are used to store values and how to declare variables below

`var conferenceName = "Go Conference"`

- fmt.Printf is used to print formatted output

```
fmt.Printf("welcome to %v booking application\n", conferenceName)
```

## Data Types in Go
- Strings
- Integers
- Booleans
- Maps
- Arrays

Go is a statically typed language - You need to tell Go compiler, the data type when declaring the variable

But, Go can infer the type when you assign a value

## what is a Pointer?

- A pointer is a variable that points to the memory address of another variable

## Slices in Go
- Slice is an abstraction of an Array
- Slices are more flexible and powerful: variable-length or get an sub-array of its own
- slices are also index-based and have a size, but is resized when needed

*syntax for creating a slice*

`var bookings []string` or `var bookings = []string{}`

*package level variables*

- Defined at the top outside all functions
- They can be accessed inside any of the functions

*maps*
- Declaration: map[string][string]

*struct*

```Go
type userData struct {
    firstName string
    lastName string
    email string
    numberOfTickets uint

}
```

*Concurrency*
- go...starts a new goroutine
- A goroutine is a lightweight thread managed by the Go runtime.

*Waitgroup*
- Waits for the launched goroutine to finish
- Package "sync" provides basic synchronization functionality
- Add: Sets the number of goroutines to wait for 
- Wait: Blocks until the WaitGroup counter is 0

## Notes

All our code must belong to a package
The first statement in Go file must be "package...."

The `main` function is the entrypoint of a Go program and also A program can only have 1 main function because you can only have 1 entry point

- Blank identifier is use to ignore a variable that is not in use (for _, booking)

## Go Packages
A package is a collection of source files - Go programs are organized into packages.
