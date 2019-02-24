package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNewBookings(t *testing.T) {

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

	if (len(lstBookings) != 2) {
		t.Errorf("Number of expected bookings was 2, but we got %v", len(lstBookings))
	}


}

func TestMoreThan5AnyEvents20Discount(t *testing.T) {
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
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})

	rawTotal, disTotal := Buy5Get20Percentoff5thOne(lstBookings, lstEvents)

	fmt.Println("Raw Total: ",rawTotal)
	fmt.Println("Dicounted Total: ",disTotal)


	if (len(lstBookings) != 5) {
		t.Errorf("Number of expected bookings is greater than 5, but we got %v", len(lstBookings))
	}
}


func TestBuy2TeamBuildEventsFor1500(t *testing.T) {
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

	rawTotal, disTotal := Buy2TeamEventsFor1500(lstBookings, lstEvents)

	fmt.Println("Raw Total: ",rawTotal)
	fmt.Println("Dicounted Total: ",disTotal)


	if (len(lstBookings) != 2) {
		t.Errorf("Number of expected bookings is not 2, but we got %v", len(lstBookings))
	}
}

func TestBuy4AndOnlyPay3(t *testing.T) {
	lstCustomers := loadCustomers()
	lstEvents := loadEvents()
	lstPromotions := loadPromotions()
	lstBookings := loadBookings()

	customerBooking := customers.lookupCustomer(lstCustomers, 2)
	fmt.Println(customerBooking)

	bookingEvent1 := events.lookupEvent(lstEvents,"Wine Tour")
	fmt.Println(bookingEvent1)

	bookingEvent2 := events.lookupEvent(lstEvents,"Picnic")
	fmt.Println(bookingEvent2)

	promotingEvent := promotions.lookupPromotion(lstPromotions, "Wine Tour and Picnic")
	fmt.Println(promotingEvent)

	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent1.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent2.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent1.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent2.id, time.Now()})

	rawTotal, disTotal := Buy4WineToursAndPicnicsPayfor3(lstBookings, lstEvents)

	fmt.Println("Raw Total: ",rawTotal)
	fmt.Println("Dicounted Total: ",disTotal)


	if (len(lstBookings) != 4) {
		t.Errorf("Number of expected bookings is not 4, but we got %v", len(lstBookings))
	}
}


func TestBuy1Get1Free(t *testing.T) {
	lstCustomers := loadCustomers()
	lstEvents := loadEvents()
	lstPromotions := loadPromotions()
	lstBookings := loadBookings()

	customerBooking := customers.lookupCustomer(lstCustomers, 2)
	fmt.Println(customerBooking)

	bookingEvent := events.lookupEvent(lstEvents,"Picnic")
	fmt.Println(bookingEvent)

	promotingEvent := promotions.lookupPromotion(lstPromotions, "Wine Tour and Picnic")
	fmt.Println(promotingEvent)

	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})
	lstBookings = append(lstBookings, booking{customerBooking.id, bookingEvent.id, time.Now()})

	rawTotal, disTotal := Buy2PicnicsGet1free(lstBookings, lstEvents)

	fmt.Println("Raw Total: ",rawTotal)
	fmt.Println("Dicounted Total: ",disTotal)

	if (len(lstBookings) != 3) {
		t.Errorf("Number of expected bookings is not 3, but we got %v", len(lstBookings))
	}
}