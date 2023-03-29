package Tools

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goFrame/app/Jobs"
	"time"
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

	//阻塞10秒，防止连接超时
	fmt.Println("["+q.name+"]", time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second())
	data, err := redis.Strings(conn.Do("BLPOP", q.name, 10))
	if err != nil {
		return false, err
	}

	job, err := q.deserialize(data[1])
	if err != nil {
		return false, err
	}

	return job.Run()
}

func QueueStart() {
	go func() {
		queueListRedisKey := "queueList"
		queueList := make(map[string]int8, 0)
		for {
			//获取连接池
			redisPool, _ := GetRedisPool()
			//获取连接
			c := redisPool.Get()

			tempQueueList, getErr := redis.Strings(c.Do("SMEMBERS", queueListRedisKey))
			if getErr == nil {
				for _, queueName := range tempQueueList {
					if _, ok := queueList[queueName]; !ok {
						//启动队列
						go func(qName string) {
							AsyncQueue := NewAsyncQueue(qName)
							for {
								AsyncQueue.Pop()
							}
						}(queueName)
						//放入已启动列表
						queueList[queueName] = 1
						fmt.Println("队列【" + queueName + "】已启动监听")
					}
				}
			}

			//关闭连接
			c.Close()
			//延时30分钟执行一次
			time.Sleep(30 * time.Minute)
		}
	}()
	fmt.Println("队列模块启动成功")
	select {}
}
