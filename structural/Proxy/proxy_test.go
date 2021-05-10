package Proxy

import "testing"

func TestAgentTask_RentHouse(t *testing.T) {
	agent := NewAgentTask()
	agent.RentHouse("望京", 9000)
}
