package redis

import (
	redis_utility "burskey/mydailylife/domain/internal/redis"
	"burskey/mydailylife/domain/package/domain"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
)

type DAO struct {
	redis *redis_utility.RedisConnection
}

func Factory(aRedis *redis_utility.RedisConnection) *DAO {
	return &DAO{redis: aRedis}
}

func (me *DAO) SaveTask(task *domain.Task) error {
	conn := me.redis.GetRedisConnection()
	defer conn.Close()

	bytes, err := json.Marshal(task)
	if err != nil {

	}
	id := task.ID
	if _, err := conn.Do("SET", id, string(bytes)); err != nil {
		log.Fatal("Unable to store data in redis.... ", err)
		return err
	}

	return nil
}
func (me *DAO) GetTask(guid string) (*domain.Task, error) {

	conn := me.redis.GetRedisConnection()
	defer conn.Close()
	values, err := redis.String(conn.Do("GET", guid))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var bytes []byte
	bytes, err = redis.Bytes(values, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	task := &domain.Task{}
	if err = json.Unmarshal(bytes, task); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return task, err
}
func (me *DAO) SaveParty(party *domain.Person) error {
	conn := me.redis.GetRedisConnection()
	defer conn.Close()

	bytes, err := json.Marshal(party)
	if err != nil {

	}
	id := party.ID
	if _, err := conn.Do("SET", id, string(bytes)); err != nil {
		log.Fatal("Unable to store data in redis.... ", err)
		return err
	}

	return nil
}
func (me *DAO) GetParty(guid string) (*domain.Person, error) {

	conn := me.redis.GetRedisConnection()
	defer conn.Close()
	values, err := redis.String(conn.Do("GET", guid))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var bytes []byte
	bytes, err = redis.Bytes(values, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	party := &domain.Person{}
	if err = json.Unmarshal(bytes, party); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return party, err
}

func (me *DAO) SaveTaskInProgress(tip *domain.TaskInProgress) error {
	conn := me.redis.GetRedisConnection()
	defer conn.Close()

	bytes, err := json.Marshal(tip)
	if err != nil {

	}
	id := tip.ID
	if _, err := conn.Do("SET", id, string(bytes)); err != nil {
		log.Fatal("Unable to store data in redis.... ", err)
		return err
	}

	return nil
}

//
//func (me db) GetMetricsByCategory(code *domain.Code) ([]*domain.CovidMetric, error) {
//	panic("implement me")
//}
//
//func (me db) GetMetricsBySchoolAndCategory(school *domain.Code, category *domain.Code) ([]*domain.CovidMetric, error) {
//	panic("implement me")
//}
//
//func (me db) GetSchools() ([]*domain.Code, error) {
//	var data []*domain.Code
//
//	conn := me.redis.GetRedisConnection()
//	defer conn.Close()
//
//	//key := fmt.Sprintf("SCHOOL_%s_DATA", aString)
//	values, err := redis.Values(conn.Do("SMEMBERS", "SCHOOLS"))
//
//	if err != nil {
//		log.Fatal(err)
//		return nil, err
//	}
//
//	for _, aValue := range values {
//		var anObject *domain.Code
//		var bytes []byte
//		if bytes, err = redis.Bytes(aValue, nil); err != nil {
//			log.Fatal(err)
//			return nil, err
//		}
//
//		if err = json.Unmarshal(bytes, &anObject); err != nil {
//			log.Fatal(err)
//			return nil, err
//		}
//
//		data = append(data, anObject)
//	}
//
//	return data, err
//}
//
//func (me db) GetCategories() ([]*domain.Code, error) {
//	var data []*domain.Code
//
//	conn := me.redis.GetRedisConnection()
//	defer conn.Close()
//
//	//key := fmt.Sprintf("SCHOOL_%s_DATA", aString)
//	values, err := redis.Values(conn.Do("SMEMBERS", "CATEGORIES"))
//
//	if err != nil {
//		log.Fatal(err)
//		return nil, err
//	}
//
//	for _, aValue := range values {
//		var anObject *domain.Code
//		var bytes []byte
//		if bytes, err = redis.Bytes(aValue, nil); err != nil {
//			log.Fatal(err)
//			return nil, err
//		}
//
//		if err = json.Unmarshal(bytes, &anObject); err != nil {
//			log.Fatal(err)
//			return nil, err
//		}
//
//		data = append(data, anObject)
//	}
//
//	return data, err
//}
//
//func (me db) GetMetric(u int) (*domain.CovidMetric, error) {
//	panic("implement me")
//}
//
//func (me db) GetMetricsBySchool(school *domain.Code) ([]*domain.CovidMetric, error) {
//	panic("implement me")
//}
//
//func (me db) SaveMetric(metric *domain.CovidMetric) (*domain.CovidMetric, error) {
//	panic("implement me")
//}
//
//func (me db) SaveSchool(school *domain.Code) (*domain.Code, error) {
//	conn := me.redis.GetRedisConnection()
//	defer conn.Close()
//
//	bytes, err := json.Marshal(school)
//
//	if _, err = conn.Do("sadd", "SCHOOLS", bytes); err != nil {
//		log.Fatal("Unable to store data in redis.... ", err)
//		return nil, err
//	}
//
//	return school, nil
//}
//
//func (me db) SaveCategory(category *domain.Code) (*domain.Code, error) {
//	conn := me.redis.GetRedisConnection()
//	defer conn.Close()
//
//	bytes, err := json.Marshal(category)
//
//	if _, err = conn.Do("sadd", "CATEGORIES", bytes); err != nil {
//		log.Fatal("Unable to store data in redis.... ", err)
//		return nil, err
//	}
//
//	return category, nil
//}
