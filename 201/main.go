package main

import (
	"fmt"
	"version/domain"
)

var Version = "Development"
var version = "development"
var date string

func main() {
	fmt.Println("Version:\t", Version)
	fmt.Println("version:\t", version)
	fmt.Println("date:\t", date)
	fmt.Println("domain version:\t", domain.DomainVersion)
}
