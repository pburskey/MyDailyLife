package service

import (
	"burskey/mydailylife/domain/package/dao"
	"burskey/mydailylife/domain/package/domain"
)

type TaskService struct {
	dao *dao.DAO
}

func (me *TaskService) SaveTask(task *domain.Task) (err error) {
	return nil
}

func (me *TaskService) GetTasksByParty(uuid string) (task []*domain.Task, err error) {
	return nil, nil
}
