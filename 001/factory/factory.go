package factory

import "fmt"

/*************************** 对外暴露 ***********************************/
// 接口
type IGun interface {
	Fire()
}

// 工厂
func GetGun(gunType string) (IGun, error) {

	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}

/****************************** 以下为内部参数 **********************************/

type gun struct {
	name  string
	power int
}

type ak47 struct {
	gun
}

func (g *ak47) Fire() {
	fmt.Println("AK47 fire power:", g.power)
}

func newAk47() *ak47 {
	return &ak47{
		gun{name: "AK47 gun", power: 4}}
}

type musket struct {
	gun
}

func (g *musket) Fire() {
	fmt.Println("Musket fire power:", g.power)
}

func newMusket() *musket {
	return &musket{gun{name: "Musket gun", power: 1}}
}
