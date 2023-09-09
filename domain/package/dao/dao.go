package dao

import "burskey/mydailylife/domain/package/domain"

type DAO interface {
	SaveTask(task *domain.Task) error
	GetTask(guid string) (*domain.Task, error)
	SaveParty(party *domain.Party) error
	GetParty(guid string) (*domain.Party, error)
	SaveTaskInProgress(tip *domain.TaskInProgress) error
}
