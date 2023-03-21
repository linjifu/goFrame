package Jobs

import "fmt"

type TestJob struct {
	Params   map[string]interface{}
	Function func(map[string]interface{}) (bool, error)
}

func NewTestJob(params map[string]interface{}) *TestJob {
	return &TestJob{
		Params: params,
	}
}

func (t *TestJob) Run() (bool, error) {
	t.Function = func(params map[string]interface{}) (bool, error) {
		fmt.Println("执行了一次测试的队列方法")
		return false, nil
	}
	return t.Function(t.Params)
}
