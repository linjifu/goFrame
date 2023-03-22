package Queue

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cobra"
	"goFrame/app/Tools"
	"time"
)

type Queue struct {
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q Queue) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "queue",
		Short: "队列模块",
		Long:  `队列模块的启动监听`,
		Run: func(cmd *cobra.Command, args []string) {
			q.run()
		},
	}
}

func (q *Queue) run() {
	go func() {
		queueListRedisKey := "queueList"
		queueList := make(map[string]int8, 0)
		for {
			//获取连接池
			redisPool, _ := Tools.GetRedisPool()
			//获取连接
			c := redisPool.Get()

			tempQueueList, getErr := redis.Strings(c.Do("SMEMBERS", queueListRedisKey))
			if getErr == nil {
				for _, queueName := range tempQueueList {
					if _, ok := queueList[queueName]; !ok {
						//启动队列
						go func(qName string) {
							AsyncQueue := Tools.NewAsyncQueue(qName)
							for {
								AsyncQueue.Pop()
								//延时1秒钟执行一次
								time.Sleep(1 * time.Second)
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
			//延时三分钟执行一次
			time.Sleep(30 * time.Minute)
		}
	}()
	fmt.Println("队列模块启动成功")
	select {}
}
