package Proxy

import "fmt"

// 代理模式

type ITask interface {
	RentHouse(desc string, price int)
}

type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Sprintln(fmt.Printf("房子地址为%s，价格为%d，中介费为%d", desc, price, price))
}

type AgentTask struct {
	task *Task
}

func NewAgentTask() *AgentTask {
	return &AgentTask{
		task: &Task{},
	}
}
func (a *AgentTask) RentHouse(desc string, price int) {
	a.task.RentHouse(desc, price)
}
