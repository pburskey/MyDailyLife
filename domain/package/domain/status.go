package domain

import (
	"github.com/google/uuid"
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

type StatusPoint struct {
	ID        uuid.UUID `json:"id" bson:"id,omitempty"`
	Timestamp time.Time `json:"timeStamp" bson:"timeStamp,omitempty"`
	Status    Status    `json:"status" bson:"status,omitempty"`
}

type StatusError struct {
	Msg string
}
