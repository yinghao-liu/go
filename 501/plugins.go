package main

import (
	"fmt"
	"plugin"
)

func plunginInit() {
	p, err := plugin.Open("plugins.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("PluginValue")
	if err != nil {
		panic(err)
	}
	fmt.Printf("PluginValue:%T\n", v)
	f, err := p.Lookup("PluginPrint")
	if err != nil {
		panic(err)
	}
	fmt.Printf("PluginPrint:%T\n", f)
	*v.(*int) = 7
	f.(func())()
}
