package abstractfactory

import (
	"fmt"
)

const (
	TypeTapioca = "tapioca"
	TypePotato  = "potato"
)

// chipFactoryInterface is the "abstract factory"
type chipFactoryInterface interface {
	makeChip() chipInterface
}

// chip is the base type of an edible chip
type chipInterface interface {
	crunch()
	getType() string
}
type chip struct {
	eaten bool
	_type string
}
func (c *chip) crunch() {
	if !c.eaten {
		c.eaten = true
		fmt.Println("Yum")
	} else {
		fmt.Printf("This %s chip currently being digested!\n", c._type)
	}
}
func (c *chip) getType() string {
	return c._type
}

// potatoChip and tapiocaChip inherits chip
type potatoChip struct {
	chip
}
type tapiocaChip struct {
	chip
}

// potatoChipFactory and tapiocaChipFactory implements chipFactoryInterface.
// These are the abstract factories that make concrete, edible chips
type potatoChipFactory struct {}
type tapiocaChipFactory struct {}
func (c *potatoChipFactory) makeChip() chipInterface {
	return &potatoChip{
		chip{
			eaten: false,
			_type: TypePotato,
		},
	}
}
func (c *tapiocaChipFactory) makeChip() chipInterface {
	return &tapiocaChip{
		chip{
			eaten: false,
			_type: TypeTapioca,
		},
	}
}

func newChipFactory(_type string) (chipFactoryInterface, error) {
	allowedChipFactories := map[string]func()chipFactoryInterface{
		TypePotato:  func()chipFactoryInterface{return &potatoChipFactory{}},
		TypeTapioca: func()chipFactoryInterface{return &tapiocaChipFactory{}},
	}
	if factoryGenerator, exists := allowedChipFactories[_type]; exists {
		return factoryGenerator(), nil
	}
	return nil, fmt.Errorf("%s is not a valid chip factory", _type)
}

func main() {
	potatoFactory, err := newChipFactory(TypePotato)
	if err != nil {
		panic(err.Error())
	}
	tapiocaFactory, err := newChipFactory(TypeTapioca)
	if err != nil {
		panic(err.Error())
	}
	potatoChip := potatoFactory.makeChip()
	tapiocaChip := tapiocaFactory.makeChip()

	fmt.Printf("About to eat %s chip\n", potatoChip.getType())
	potatoChip.crunch()

	fmt.Printf("Now I'm about to eat the %s chip I made\n", tapiocaChip.getType())
	tapiocaChip.crunch()

	potatoChip.crunch()
}
