package builder

import (
	"fmt"
)

// Let's just say we are never going to ever use the cart
//	class directly to create buyers
type cart struct {
	Name           string
	NumItemsInCart uint64
}

// CartBuilder is the builder of the builder pattern
type CartBuilder struct {
	cart // optional: CartBuilder can compose of a cart
}
// SetName sets the name for the current active cart
func (b *CartBuilder) SetName(cartName string) *CartBuilder {
	if b != nil {
		b.Name = cartName
		return b
	}
	return &CartBuilder{cart{Name: cartName}}
}
// Shopping allows shoppers to shop
func (b *CartBuilder) Shopping(numItems uint64) *CartBuilder {
	if b != nil {
		b.NumItemsInCart = numItems
		return b
	}
	return &CartBuilder{cart{NumItemsInCart: numItems}}
}
// Checkout sends the cart to the user and resets
func (b *CartBuilder) Checkout() *cart {
	savedCart := b.cart
	b.Name = ""
	b.NumItemsInCart = 0
	return &savedCart
}

func main() {
	builder := &CartBuilder{}

	// Get a new shopping cart and throw 90 items inside there
	builder.
		SetName("alex@burpee.com").
		Shopping(90)
	// Actually we don't need all that stuff. Perhaps we can take items out
	builder.Shopping(84)
	// Checkout my cart!
	alex := builder.Checkout()

	// Get a new shopping cart for bob
	builder.
		SetName("bob@burpee.com").
		Shopping(25)
	// Bob checkout
	bob := builder.Checkout()

	fmt.Printf("Checkout: %s has %d items\n", alex.Name, alex.NumItemsInCart)
	fmt.Printf("Checkout: %s has %d items\n", bob.Name, bob.NumItemsInCart)
}
