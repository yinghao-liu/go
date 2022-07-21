package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	// Get a new client
	// 默认是指向 http://127.0.0.1:8500的client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Get a handle to the agent API
	// agent相关的url
	agent := client.Agent()

	// 相当于 GET http://127.0.0.1:8500/v1/agent/services
	srvs, err := agent.Services()
	for i, j := range srvs {
		fmt.Printf("Services: %s %#v\n", i, j)
	}

	// 相当于 GET http://127.0.0.1:8500/v1/agent/service/web
	web, meta, err := agent.Service("web", nil)
	fmt.Printf("addr:%s, port:%d\n", web.Address, web.Port)
	fmt.Printf("meta is %#v\n", meta)

	RegisterService(agent)
	device2, meta, err := agent.Service("device2", nil)
	fmt.Printf("device2 addr:%s, port:%d\n", device2.Address, device2.Port)

	st, hsrvs, e := agent.AgentHealthServiceByName("device")
	if nil == e {
		fmt.Printf("status is %s\n", st)
		fmt.Printf("services is %#v\n", hsrvs)
	}

	st, hsrvs, e = agent.AgentHealthServiceByName("web")
	if nil == e {
		fmt.Printf("status is %s\n", st)
		fmt.Printf("services is %#v\n", hsrvs)
	}
}
