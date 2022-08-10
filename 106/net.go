package main

import (
	"fmt"
	"net"
)

func hostPort() {

	//hostport := "[2001::2222]:8080"
	//hostport := "[2001::2222]"
	// hostport := "127.0.0.1:8080"
	hostport := "127.0.0.1"
	host, port, err := net.SplitHostPort(hostport)
	fmt.Printf("host:%s, port:%s, err:%v\n", host, port, err)
	fmt.Printf("JoinHostPort:%s\n", net.JoinHostPort(host, port))
}
