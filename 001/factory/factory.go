package factory

import "fmt"

/*************************** 对外暴露 ***********************************/
// 接口
type IGun interface {
	GetName() string
	GetPower() int

	setName(name string)
	setPower(power int)
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

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) GetPower() int {
	return g.power
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

type ak47 struct {
	gun
}

func newAk47() *ak47 {
	return &ak47{
		gun{name: "AK47 gun", power: 4}}
}

type musket struct {
	gun
}

func newMusket() *musket {
	return &musket{gun{name: "Musket gun", power: 1}}

}
