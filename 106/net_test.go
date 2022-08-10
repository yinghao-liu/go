package main

import "testing"

func TestConnect(t *testing.T) {
	ipConnect(":8080")
}

func TestHostPort(t *testing.T) {
	hostPort()
}
