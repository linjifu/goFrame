package Tools

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"goFrame/app/Jobs"
)

type Job interface {
	Run() (bool, error)
}

type AsyncQueue struct {
	name string
}

func NewAsyncQueue(name string) *AsyncQueue {
	return &AsyncQueue{
		name,
	}
}

func (q *AsyncQueue) serialize(job Job) (string, error) {
	data, err := json.Marshal(&job)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (q *AsyncQueue) deserialize(data string) (Job, error) {
	queueList := q.queueList()
	jobMap, ok := queueList[q.name]
	if ok {
		err := json.Unmarshal([]byte(data), jobMap)
		if err != nil {
			return nil, err
		}
	}
	return jobMap, nil
}

func (q *AsyncQueue) queueList() map[string]Job {
	queueMap := make(map[string]Job, 1)

	queueMap["a"] = new(Jobs.TestJob)
	queueMap["b"] = new(Jobs.TestJob2)

	return queueMap
}

func (q *AsyncQueue) PushBack(job Job) error {
	redisPool, _ := GetRedisPool()
	conn := redisPool.Get()
	defer conn.Close()

	data, err := q.serialize(job)
	if err != nil {
		return err
	}

	_, err = conn.Do("RPUSH", q.name, data)
	if err != nil {
		return err
	}

	return nil
}

func (q *AsyncQueue) PushFront(job Job) error {
	redisPool, _ := GetRedisPool()
	conn := redisPool.Get()
	defer conn.Close()

	data, err := q.serialize(job)
	if err != nil {
		return err
	}

	_, err = conn.Do("LPUSH", q.name, data)
	if err != nil {
		return err
	}

	return nil
}

func (q *AsyncQueue) Pop() (bool, error) {
	redisPool, _ := GetRedisPool()
	conn := redisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("LPOP", q.name))
	if err != nil {
		return false, err
	}

	job, err := q.deserialize(data)
	if err != nil {
		return false, err
	}
	return job.Run()
}
