package CronJobs

import "fmt"

type TestCron struct {
}

func NewTestCron() *TestCron {
	return &TestCron{}
}

func (t *TestCron) Run() {
	fmt.Println("运行了一次测试定时器任务")
}
