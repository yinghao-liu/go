package main

import (
	"factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	ak47.Fire()
	musket.Fire()
}
