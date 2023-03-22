package Jobs

import (
	"fmt"
	"time"
)

type TestJob2 struct {
	Id     string
	Name   string
	Params map[string]interface{}
}

func NewTestJob2(name string, params map[string]interface{}) *TestJob2 {
	return &TestJob2{
		Id:     fmt.Sprintf("TestJob:%d", time.Now().UnixMicro()),
		Name:   name,
		Params: params,
	}
}

func (t *TestJob2) Run() (bool, error) {
	fmt.Printf("Id2:%v，Name2:%v，Params2:%v", t.Id, t.Name, t.Params)
	fmt.Println()
	return false, nil
}
