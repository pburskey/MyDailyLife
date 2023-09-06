package main

import (
	"burskey/mydailylife/domain/package/domain"
	"github.com/google/uuid"
)

func main() {

	party := &domain.Person{
		ID:    uuid.New().String(),
		First: "p",
		Last:  "b",
	}

	taska := domain.NewTask("a", "a")
	tip := taska.Start(party.ID)
	if err := tip.Start(); err != nil {
		panic("Unable to start a task")
	}

	if err := tip.Complete(); err != nil {
		panic("Unable to complete a task")
	}

}
