package main

import (
	"fmt"
	"time"
)

type booking struct {
	customerId int64
	eventId int64
	datePurchase time.Time
}

type bookings []booking

func loadBookings() bookings {
	currentBookings := []booking{}

	return currentBookings
}


func toString(b bookings) string {
	bookingsMade := ""

	fmt.Printf("Number of bookings %d\n", len(b))

	for i, booking := range b {
		fmt.Println(i, booking)
	}

	//fmt.Println(bookingsMade)

	return bookingsMade
}