package main

import "burskey/mydailylife/domain/package/domain"

func main() {

	taska := domain.NewTask("a", "a")
	tip := taska.Create()
	if err := tip.Start(); err != nil {
		panic("Unable to start a task")
	}

	if err := tip.Complete(); err != nil {
		panic("Unable to complete a task")
	}

}
