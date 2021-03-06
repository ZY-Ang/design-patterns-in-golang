package factorymethod

import (
	"fmt"
)

type food interface {
	Eat()
}

type taco struct {}
func (g *taco) Eat() {
	fmt.Println("Crunch")
}

type foodFactory interface {
	make() food
}
type tacoFactory struct {}
func newTacoFactory() foodFactory {
	return &tacoFactory{}
}

// make is the factory method of the taco factory
func (p *tacoFactory) make() food {
	return &taco{}
}

func main() {
	tacoFactory := newTacoFactory()
	foods := make([]food, 10)
	for i := 0; i < 10; i++ {
		foods[i] = tacoFactory.make()
	}
	for i := 0; i < 10; i++ {
		foods[i].Eat()
	}
}

