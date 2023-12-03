package main

import (
	"testing"
)

func TestBookTicketsInputs(t *testing.T) {
	tests := []struct {
		fName         string
		lName         string
		email         string
		ticketsBooked uint
		expected      bool
	}{
		{"Abdul", "Yunus", "abdul@gmail.com", 10, true},
		{"Tahir", "Y", "taheer@omail.com", 55, true},
	}

	for _, test := range tests {
		validName, validEmail, validTickets := validateInput(test.fName, test.lName, test.email, test.ticketsBooked)
		t.Run("Test valid name", func(t *testing.T) {
			if !validName {
				t.Errorf("Test failed for invalid name: %v", test.fName+" "+test.lName)
			}
		})
		t.Run("Test email", func(t *testing.T) {
			if !validEmail {
				t.Errorf("Test failed for invalid email: %v", test.email)
			}
		})
		t.Run("Test valid tickets", func(t *testing.T) {
			if !validTickets {
				t.Errorf("Test failed for invalid ticket: %v was selected but only %v left", test.ticketsBooked, remainingTickets)
			}
		})
	}
}
