package helper

import (
	"fmt"
	"math/rand"
	"strings"
)

func ValidateUserInput(fname string, lname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool){
	isValidName := len(fname) >=2 && len(lname) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}


func BookingId(){
	random := rand.Intn(50)
	fmt.Printf("Your booking ID: %v\n", random)
}

func GetUserInput()(string, string, string, uint){
	var userTickets uint
	var fname string
	var lname string
	var email string

	fmt.Println("Enter your name: ")
	fmt.Scan(&fname)
	fmt.Scan(&lname)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return fname, lname, email, userTickets
}