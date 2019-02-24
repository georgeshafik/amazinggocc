package main

import (
	"fmt"
	"strings"
)

type event struct {
	id int64
	name string
	cost float64
	status bool
}

type events []event

func loadEvents() events {
	currentEvents := []event{
		{1, "Kids Party",220.00, true},
		{2, "Wine Tour",440.00, true},
		{3, "Team Building", 880, true},
		{4, "Picnic", 110, true},
	}
	return currentEvents
}

func (e events) lookupEvent(nameOfEvent string) event {
	var foundEvent event

	for _, event := range e {
		if (strings.Compare(event.name, nameOfEvent) == 0 ) {
			fmt.Println("Found event by name verifying it exists")
			foundEvent = event;
			break
		}
	}

	return foundEvent
}

func (e events) lookupEventCost(idOfEvent int64) float64 {
	var foundEventCost float64

	for _, event := range e {
		if (event.id == idOfEvent) {
			fmt.Println("Found event by id for costings")
			foundEventCost = event.cost;
			break
		}
	}

	return foundEventCost
}