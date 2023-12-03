package main

import "strings"

func validateInput(fName string, lName string, email string, tBooked uint) (bool, bool, bool) {
	validName := len(fName+lName) > 3
	validEmail := strings.Contains(email, "@") && len(email) > 3
	validTickets := tBooked <= remainingTickets
	return validName, validEmail, validTickets
}
