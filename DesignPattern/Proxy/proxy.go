package Proxy

import "fmt"

type ITask interface {
	RentHouse(desc string, price int)
}

type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Println(fmt.Sprintf("租房子的地址%s,价格%d,中介费%d", desc, price, price))
}

type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{task: &Task{}}
}

func (t *AgentTask) RentHouse(desc string, price int) {
	t.task.RentHouse(desc, price)
}
