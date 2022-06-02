package main

import (
	"ddd/domain"
	"ddd/domain/xlog"
	"ddd/infrastructure/logfmt"
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
func xlogInit() {
	xlog.Log = new(logfmt.Log)
}
func main() {

}
