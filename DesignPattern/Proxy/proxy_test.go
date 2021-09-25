package Proxy

import "testing"

func TestAgentTask_RentHouse(t *testing.T) {
	agent:=NewAgentTask()
	agent.RentHouse("浙江省",10099)
}