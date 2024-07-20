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

## Notes

All our code must belong to a package
The first statement in Go file must be "package...."

The `main` function is the entrypoint of a Go program and also A program can only have 1 main function because you can only have 1 entry point

## Go Packages
A package is a collection of source files - Go programs are organized into packages.
