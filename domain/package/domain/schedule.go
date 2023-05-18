package domain

import (
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	ID       uuid.UUID `json:"id" bson:"id,omitempty"`
	Name     string
	Weekdays []time.Weekday
	Start    time.Time
	End      time.Time
}

type ScheduleTask struct {
	ScheduleID uuid.UUID `json:"scheduleID" bson:"scheduleID,omitempty"`
	TaskID     uuid.UUID `json:"taskId" bson:"taskId,omitempty"`
}

func main() {
	schedule := &Schedule{
		ID:       uuid.New(),
		Name:     "",
		Weekdays: []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
		Start:    time.Now(),
		End:      time.Now(),
	}

	if schedule != nil {

	}

}
