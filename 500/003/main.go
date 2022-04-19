package main

import (
	"ddd/domain"
	mem "ddd/infrastructure/memory"
)

func itermInit() {
	its := new(mem.ItemsMemory)
	its.Init()
	domain.ItemsRepo = its
}

func catalogInit() {
	cats := new(mem.CatalogsMemory)
	cats.Init()
	domain.CatalogsRepo = cats
}
func main() {

}
