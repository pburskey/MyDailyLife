package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `json:"id" bson:"id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
}

type TaskInProgress struct {
	ID       uuid.UUID    `json:"id" bson:"id,omitempty"`
	TaskID   uuid.UUID    `json:"taskID" bson:"taskID,omitempty"`
	Creation time.Time    `json:"dateTime" bson:"dateTime,omitempty"`
	Status   *StatusPoint `json:"status" bson:"status,omitempty"`
}

func (t *Task) Create() *TaskInProgress {
	tip := &TaskInProgress{
		ID:       uuid.New(),
		TaskID:   t.ID,
		Creation: time.Now(),
		Status: &StatusPoint{
			ID:        uuid.New(),
			Timestamp: time.Now(),
			Status:    NOT_STARTED,
		},
	}
	return tip
}

func NewTask(aName string, aDescription string) *Task {
	return &Task{ID: uuid.New(), Name: aName, Description: aDescription}
}

func (tip *TaskInProgress) changeStatus(status *StatusPoint) error {
	var err error
	if status != nil {

		if tip.Status.Status == status.Status {
			//message := fmt.Sprintf("Unable to change from status: %s to status: %s", tip.Status.Status, status.Status)
			err = errors.New("Unable to change status")
		}

		tip.Status = status
	}
	return err

}

func (tip *TaskInProgress) Start() error {
	aStatus := &StatusPoint{
		ID:        uuid.New(),
		Timestamp: time.Now(),
		Status:    STARTED,
	}

	err := tip.changeStatus(aStatus)
	return err
}

func (tip *TaskInProgress) Complete() error {
	aStatus := &StatusPoint{
		ID:        uuid.New(),
		Timestamp: time.Now(),
		Status:    COMPLETE,
	}

	err := tip.changeStatus(aStatus)
	return err

}

type TaskSchedule struct {
	TaskID      uuid.UUID `json:"taskID" bson:"taskID,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
}
