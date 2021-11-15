package main

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())
}
