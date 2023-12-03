package main

import (
	"fmt"
	"strconv"
	"strings"
)

const EVENT_NAME = "BitDev"
const TOTAL_TICKETS = 100

var count = 0
var remainingTickets uint = 100
var users = []map[string]string{}

func main() {

	var ticketsBooked uint
	var firstName string
	var lastName string
	var email string

	for {
		var response string

		greetUser()
		fmt.Scan(&response)

		if strings.ToLower(response) == "q" {
			break
		}

		isBooking := strings.ToLower(response) == "y" || strings.ToLower(response) == "n"

		if isBooking && strings.ToLower(response) == "y" {
			fmt.Println("Enter your firstname: ")
			fmt.Scan(&firstName)

			fmt.Println("Enter your lastname: ")
			fmt.Scan(&lastName)

			fmt.Println("Enter your email: ")
			fmt.Scan(&email)

			fmt.Println("Enter number of tickets: ")
			fmt.Scan(&ticketsBooked)

			validName, validEmail, validTickets := validateInput(firstName, lastName, email, ticketsBooked)
			if validName && validTickets && validEmail {
				remainingTickets -= ticketsBooked

				commitUser(firstName, lastName, email, ticketsBooked)

				count++

			} else {
				fmt.Print("Provide valid data please\n\n")
			}
		} else if isBooking && strings.ToLower(response) == "n" {
			println("")
			continue
		}

	}

}

func greetUser() {
	fmt.Printf("%v people(s) are going so far\n", count)
	fmt.Printf("They are: %v\n", users)
	fmt.Printf("Welcome to %v yearly event\n", EVENT_NAME)
	fmt.Printf("Book your ticket with us, there are %v out of %v left\n", remainingTickets, TOTAL_TICKETS)

	fmt.Println("Would you like to book now? (y/n) or q to quit")
}

func commitUser(firstName string, lastName string, email string, ticketsBooked uint) {
	fullname := firstName + " " + lastName
	user := map[string]string{"name": fullname, "email": email, "tickets": strconv.FormatUint(uint64(ticketsBooked), 10)}
	users = append(users, user)
	fmt.Printf("Thank you %v, you've booked %v tickets\n\n", firstName, ticketsBooked)
}
