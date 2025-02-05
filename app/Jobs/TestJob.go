package Jobs

import (
	"fmt"
	"time"
)

type TestJob struct {
	Id     string
	Name   string
	Params map[string]interface{}
}

func NewTestJob(name string, params map[string]interface{}) *TestJob {
	return &TestJob{
		Id:     fmt.Sprintf("TestJob:%d", time.Now().UnixMicro()),
		Name:   name,
		Params: params,
	}
}

func (t *TestJob) Run() (bool, error) {
	fmt.Printf("Id:%v，Name:%v，Params:%v", t.Id, t.Name, t.Params)
	fmt.Println()
	return false, nil
}
