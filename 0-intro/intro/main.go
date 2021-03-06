package intro

import (
	"fmt"
)

// Order is an interface for orders
type Order interface {
	Sn()
	SetSn()
}
// FakeOrder represents a fake order
type FakeOrder struct {
	Id uint64 // public instance var
	sn string // private instance var
}
// Sn returns serial number
func (o *FakeOrder) Sn() string {
	if o != nil {
		return o.sn
	}
	return ""
}
// SetSn sets sn to the specified value
func (o *FakeOrder) SetSn(sn string) {
	if o != nil {
		o.sn = sn
	} else {
		panic("whoops")
	}
}
// FoodOrder inherits FakeOrder
type FoodOrder struct {
	FakeOrder
	FoodCategory string
	FoodName     string
}
// Print is a unique method of FoodOrder
func (o *FoodOrder) Print() {
	fmt.Println("I am a " + o.FoodCategory + " of type " + o.FoodName)
}
// Sn overrides FakeOrder Sn() and returns
//	serial number of a FoodOrder
func (o *FoodOrder) Sn() string {
	return fmt.Sprintf("%s.%s.%s",
		o.FoodCategory,
		o.FoodName,
		o.FakeOrder.Sn(),
	)
}
func main() {
	kaleSalad := &FoodOrder{
		FakeOrder{},
		"salad",
		"kale",
	}
	kaleSalad.Print()
	kaleSalad.SetSn("someuuid")
	fmt.Println(kaleSalad.Sn())
	fmt.Println(kaleSalad.sn)
}
