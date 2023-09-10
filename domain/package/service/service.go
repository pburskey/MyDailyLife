package service

import (
	"burskey/mydailylife/domain/package/dao/redis"
	"burskey/mydailylife/domain/package/domain"
)

type TaskService struct {
	dao *redis.DAO
}

func (me *TaskService) SaveTask(task *domain.Task) (err error) {
	if err := me.dao.SaveTask(task); err != nil {
		return err
	}

	return nil
}

func (me *TaskService) GetTask(uuid string) (*domain.Task, error) {
	var task *domain.Task
	var err error
	if task, err = me.dao.GetTask(uuid); err != nil {
		return nil, err
	}

	return task, nil
}

func (me *TaskService) SaveParty(party *domain.Person) (err error) {
	if err := me.dao.SaveParty(party); err != nil {
		return err
	}

	return nil
}

func (me *TaskService) GetParty(uuid string) (*domain.Person, error) {
	var person *domain.Person
	var err error
	if person, err = me.dao.GetParty(uuid); err != nil {
		return nil, err
	}

	return person, nil
}

func (me *TaskService) GetTasksByParty(uuid string) (task []*domain.Task, err error) {
	return nil, nil
}

func (me *TaskService) SaveTaskInProgress(tip *domain.TaskInProgress) error {
	if err := me.dao.SaveTaskInProgress(tip); err != nil {
		return err
	}

	return nil
}

func Factory(aDAO *redis.DAO) *TaskService {
	me := &TaskService{dao: aDAO}
	return me
}
