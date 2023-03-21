package Tools

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
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
	var job Job
	err := json.Unmarshal([]byte(data), &job)
	if err != nil {
		return nil, err
	}
	return job, nil
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
		fmt.Println(err)
		return err
	}

	_, err = conn.Do("LPUSH", q.name, data)
	if err != nil {
		return err
	}

	return nil
}

func (q *AsyncQueue) Pop() (Job, error) {
	fmt.Println("监听到了123")
	redisPool, _ := GetRedisPool()
	conn := redisPool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("BLPOP", q.name, 0))
	if err != nil {
		return nil, err
	}

	job, err := q.deserialize(data)
	if err != nil {
		return nil, err
	}
	job.Run()
	fmt.Printf("执行了队列%T，%t", job, job)

	return job, nil
}
