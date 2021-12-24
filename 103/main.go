package main

import (
	ass "gormtest/association"
	inf "gormtest/infrastructure"
)

func main() {
	inf.GormInit()
	ass.Association()
}
