package domain

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type Status uint

const (
	NOT_STARTED Status = iota
	STARTED
	PAUSED
	COMPLETE
	INCOMPLETE
)

func StatusFactory(name string) Status {
	var status Status
	if strings.EqualFold(name, "NOT_STARTED") {
		status = NOT_STARTED
	} else if strings.EqualFold(name, "STARTED") {
		status = STARTED
	} else if strings.EqualFold(name, "PAUSED") {
		status = PAUSED
	} else if strings.EqualFold(name, "COMPLETE") {
		status = COMPLETE
	} else if strings.EqualFold(name, "INCOMPLETE") {
		status = INCOMPLETE
	}
	return status
}

type StatusPoint struct {
	SKEY      uint      `json:"skey" bson:"skey,omitempty"`
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	Timestamp time.Time `json:"timeStamp" bson:"timeStamp,omitempty"`
	Status    Status    `json:"status" bson:"status,omitempty"`
}

type StatusError struct {
	Msg string
}

type StatusMachine struct {
	statusPoints *[]StatusPoint
}
