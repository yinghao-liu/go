package main

import (
	"fmt"
	"net"
)

func ipConnect(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		// handle error
		fmt.Printf("%#v\n", err)
		fmt.Printf("error is %s\n", err.Error())
		return err
	}
	defer ln.Close()
	return nil
}
