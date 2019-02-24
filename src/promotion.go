package main

import (
	"fmt"
	"strings"
)

type promotion struct {
	id int64
	name string
	rule string
	status bool
}

type promotions []promotion

func loadPromotions() promotions {
	currentPromotions := []promotion{
		{1, "Any", "Buy 5, Get 20% off the 5th experience", true},
		{2, "Team Building", "Buy 2 for $1500", true},
		{3, "Wine Tour and Picnic", "Buy 4, ONLY Pay for 3", true},
		{2, "Team Building", "Buy 2, get 1 free", true},
	}
	return currentPromotions
}

func (p promotions) lookupPromotion(nameOfPromotion string) promotion {
	var foundPromotion promotion

	for _, promotion := range p {
		if (strings.Compare(promotion.name, nameOfPromotion) == 0 ) {
			fmt.Println("Found promotion")
			foundPromotion = promotion;
			break
		}
	}

	return foundPromotion
}

func Buy5Get20Percentoff5thOne(b bookings, e events) (rawTotal float64, disTotal float64) {

	calRawTotal := 0.0
	calDisTotal := 0.0
	currentCost := 0.0

	if (len(b) >= 5 ) {

		for i, booking := range b {
			currentCost = e.lookupEventCost(booking.eventId)

			if (i == (len(b) - 1)) {
				// Discount 20% of last item
				calDisTotal = calDisTotal + currentCost * 0.80
			} else {
				calDisTotal = calDisTotal + currentCost
			}
			calRawTotal = calRawTotal + currentCost

			fmt.Printf("RawTotal: %v\n", calRawTotal)
			fmt.Printf("DisTotal: %v\n", calDisTotal)
		}
	}

	return calRawTotal, calDisTotal
}

func Buy2TeamEventsFor1500(b bookings, e events) (rawTotal float64, disTotal float64) {
	calRawTotal := 0.0
	calDisTotal := 0.0
	currentCost := 0.0
	countTeamBuilding := 0;

	if (len(b) == 2 ) {

		for _, booking := range b {
			currentCost = e.lookupEventCost(booking.eventId)
			currentEvent := e.lookupEvent("Team Building")

			if ( strings.Compare( currentEvent.name, "Team Building") == 0 ) {
				countTeamBuilding += 1
			}

			calDisTotal = calDisTotal + currentCost
			calRawTotal = calRawTotal + currentCost

			if (countTeamBuilding == 2) {
				fmt.Println("Discount 1500 applied to 2 Team Event Bookings")
				calDisTotal = 1500;
			}

			fmt.Printf("RawTotal: %v\n", calRawTotal)
			fmt.Printf("DisTotal: %v\n", calDisTotal)
		}
	}

	return calRawTotal, calDisTotal
}

func Buy4WineToursAndPicnicsPayfor3(b bookings, e events) (rawTotal float64, disTotal float64) {
	calRawTotal := 0.0
	calDisTotal := 0.0
	currentCost := 0.0
	countTeamBuilding := 0;

	if (len(b) == 4 ) {
		for i, booking := range b {
			currentCost = e.lookupEventCost(booking.eventId)
			currentEvent1 := e.lookupEvent("Wine Tour")
			currentEvent2 := e.lookupEvent("Picnic")

			if ( ( strings.Compare( currentEvent1.name, "Wine Tour") == 0 ) ||
				 ( strings.Compare( currentEvent1.name, "Picnic") == 0 ) ) ||
			   ( ( strings.Compare( currentEvent2.name, "Wine Tour") == 0 ) ||
				 ( strings.Compare( currentEvent2.name, "Picnic") == 0 ) ) {
				countTeamBuilding += 1
			}

			if (countTeamBuilding == 4) && (i == (len(b)-1) ) {
				fmt.Println("ONLY Pay for 3!")
			} else {
				calDisTotal = calDisTotal + currentCost
			}

			calRawTotal = calRawTotal + currentCost

			fmt.Printf("RawTotal: %v\n", calRawTotal)
			fmt.Printf("DisTotal: %v\n", calDisTotal)
		}
	}

	return calRawTotal, calDisTotal
}

func Buy2PicnicsGet1free(b bookings, e events) (rawTotal float64, disTotal float64) {
	calRawTotal := 0.0
	calDisTotal := 0.0
	currentCost := 0.0
	countPicnics := 0;

	if (len(b) == 3 ) {
		for i, booking := range b {
			currentCost = e.lookupEventCost(booking.eventId)
			currentEvent := e.lookupEvent("Picnic")

			if (strings.Compare( currentEvent.name, "Picnic") == 0) {
				countPicnics += 1
			}

			if (countPicnics == 3) && (i == (len(b)-1) ) {
				fmt.Println("Buy 2, get 1 free!")
			} else {
				calDisTotal = calDisTotal + currentCost
			}

			calRawTotal = calRawTotal + currentCost

			fmt.Printf("RawTotal: %v\n", calRawTotal)
			fmt.Printf("DisTotal: %v\n", calDisTotal)
		}
	}

	return calRawTotal, calDisTotal
}