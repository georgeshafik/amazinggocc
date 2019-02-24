package main

import (
	"fmt"
	"time"
)

func main() {

	lstCustomers := loadCustomers()
	lstEvents := loadEvents()
	lstPromotions := loadPromotions()
	lstBookings := loadBookings()

	customerBooking := customers.lookupCustomer(lstCustomers, 2)
	fmt.Println(customerBooking)

	bookingEvent := events.lookupEvent(lstEvents,"Team Building")
	fmt.Println(bookingEvent)

	promotingEvent := promotions.lookupPromotion(lstPromotions, "Wine Tour and Picnic")
	fmt.Println(promotingEvent)

	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	//lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, promotingEvent.id, time.Now()})

	fmt.Println(toString(lstBookings))

}

