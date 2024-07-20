package main

import "fmt"

//  Entry point of our application
func main() {
	conferenceName := "Go Conference"
	const conferenceTicket = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	var bookings [50]string // Array declaration

	// ask user for their name
	// &userName is pointer that references the variable userName

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole Array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array size: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v has %v tickets left", conferenceName, remainingTickets)

}
