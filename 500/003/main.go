package main

import (
	"ddd/domain"
	repomem "ddd/infrastructure/repomemory"
)

func itermInit() {
	its := new(repomem.ItemsMemory)
	its.Init()
	domain.ItemsRepo = its
}

func catalogInit() {
	cats := new(repomem.CatalogsMemory)
	cats.Init()
	domain.CatalogsRepo = cats
}
func main() {

}
