package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and still avaliable %v\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here\n")

}
