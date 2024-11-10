package main

import (
	"Booking-app/helper"
	"fmt"
	"strconv"
	"sync"
	"time"
)


var conferenceName = "GO Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var age uint
var nationality string
var bookings = make([]map[string]string,0)
var wg = sync.WaitGroup{}

type verification struct {
	ageVerification uint
	residentOfIndia string
}

func main(){

	greeting()

	for {

	fname, lname, email, userTickets := helper.GetUserInput()
	isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(fname,lname,email,userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTickets{
		remainingTickets = remainingTickets - userTickets

		// structure to store some data
		fmt.Println("Enter your age: ")
		fmt.Scan(&age)
		fmt.Println("Do you live in India: ")
		fmt.Scan(&nationality)

		var verification = verification{
			ageVerification: age,
			residentOfIndia: nationality,
		}

		fmt.Printf("Additional Details: %v\n", verification)

		//map to store user data
		var userData = make(map[string]string)
		userData["fname"] = fname
		userData["lname"] = lname
		userData["email"] = email
		userData["userTickets"] = strconv.FormatUint(uint64(userTickets),10)

		bookings = append(bookings, userData)
		fmt.Printf("List of Booking: %v\n",bookings)

	
		fmt.Printf("Thank you %v %v for booking %v tickets. You'll receive a confirmation mail at %v\n",fname,lname,userTickets,email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
		printFirstName()

		helper.BookingId()

		wg.Add(1)
		go sendTicket(userTickets,fname,lname,email)
	
		if remainingTickets == 0{
			fmt.Println("We are booked out, Come back next year!")
			break
		}

	} else {
		if !isValidName{
			fmt.Println("Firstname or Lastname you entered is too short")
		}
		if !isValidEmail{
			fmt.Println("Check the email you entered!")
		}
		if !isValidUserTickets{
			fmt.Println("Number of tickets entered is invalid!")
		}
	}

	wg.Wait()

	}

}

func greeting(){
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining! \n",conferenceTickets,remainingTickets)
	fmt.Println("Book your tickets here!")
}

func printFirstName(){
	firstName := []string{}
		for _, booking := range bookings {
			firstName = append(firstName, booking["fname"])
		}
	
	fmt.Printf("These are all our bookings : %v\n", firstName)
}

func sendTicket(userTickets uint, fname string, lname string, email string){
	time.Sleep(10*time.Second)
	var Ticket = fmt.Sprintf("%v tickets booked by %v %v",userTickets,fname, lname)
	fmt.Println("####################")
	fmt.Printf("Sent to email: %v\n %v\n", email, Ticket)
	fmt.Println("####################")
	wg.Done()
}

