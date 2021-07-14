package main

import (
	"factory"
	"fmt"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factory.IGun) {
	fmt.Println("Gun and Power: ", g.GetName(), g.GetPower())
}
