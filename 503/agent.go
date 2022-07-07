package main

import "github.com/hashicorp/consul/api"

func RegisterService(agent *api.Agent) {
	var service api.AgentServiceRegistration
	service.Address = "127.0.0.1"
	service.Name = "device"
	service.ID = "device1"
	service.Port = 10100

	var check api.AgentServiceCheck
	check.HTTP = "http://127.0.0.1:10100/v1/device/version"
	check.Interval = "10s"
	service.Check = &check
	agent.ServiceRegister(&service)
}
