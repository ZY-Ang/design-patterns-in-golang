package abstractfactory

import (
	"fmt"
	"os"
)

const (
	BrandBurpee = "burpee"
	BrandLadida = "ladida"
)

// ECommercePlatform is the "abstract factory"
type ECommercePlatform interface {
	makeShoe() Shoe
	makeShirt() Shirt
}
// Burpee is the burpee ecommerce platform
type Burpee struct {}
func (s *Burpee) makeShoe() Shoe {
	return &burpeeShoe{shoe{BrandBurpee}}
}
func (s *Burpee) makeShirt() Shirt {
	return &burpeeShirt{shirt{BrandBurpee}}
}
// Ladida is the ladida ecommerce platform
type Ladida struct {}
func (s *Ladida) makeShoe() Shoe {
	return &ladidaShoe{shoe{BrandLadida}}
}
func (s *Ladida) makeShirt() Shirt {
	return &ladidaShirt{shirt{BrandLadida}}
}

// Shoe is the base type of a shoe
type Shoe interface {
	Print()
	SetBrand(string)
}
type shoe struct {
	brand string
}
func (s *shoe) Print() {
	fmt.Printf("%s shoe\n", s.brand)
}
func (s *shoe) SetBrand(brand string) {
	s.brand = brand
}
// burpeeShoe and ladidaShoe inherits shoe
type burpeeShoe struct {
	shoe
}
type ladidaShoe struct {
	shoe
}

// Shirt is the base type of a shirt
type Shirt interface {
	Print()
	SetBrand(string)
}
type shirt struct {
	brand string
}
func (s *shirt) Print() {
	fmt.Printf("%s shirt\n", s.brand)
}
func (s *shirt) SetBrand(brand string) {
	s.brand = brand
}
// burpeeShirt and ladidaShirt inherits shoe
type burpeeShirt struct {
	shirt
}
type ladidaShirt struct {
	shirt
}

// newPlatform creates a new ECommercePlatform
func newPlatform(brand string) (ECommercePlatform, error) {
	brandMap := map[string]func()ECommercePlatform{
		BrandBurpee: func() ECommercePlatform {return &Burpee{}},
		BrandLadida: func() ECommercePlatform {return &Ladida{}},
	}
	if brandFunc, ok := brandMap[brand]; ok {
		return brandFunc(), nil
	}
	return nil, fmt.Errorf("brand %s is invalid", brand)
}

func main() {
	if len(os.Args[1:]) != 1 {
		panic("Invalid number of arguments")
	}
	brand := os.Args[1]
	platform, err := newPlatform(brand)
	if err != nil {
		panic(err)
	}
	myNewShoes := platform.makeShoe()
	myNewShirt := platform.makeShirt()
	myNewShoes.Print()
	myNewShirt.Print()
}
