package memento

import "fmt"

type cart struct {
	state string
}
func (e *cart) createMemento() *memento {
	return &memento{state: e.state}
}
func (e *cart) restoreState(m *memento) {
	e.state = m.getSavedState()
}
func (e *cart) setState(state string) {
	e.state = state
}
func (e *cart) getState() string {
	return e.state
}

type memento struct {
	state string
}
func (m *memento) getSavedState() string {
	return m.state
}

type cartHistory struct {
	mementoArray []*memento
}
func (c *cartHistory) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}
func (c *cartHistory) getMemento(index int) *memento {
	return c.mementoArray[index]
}

func main() {
	cartHistory := &cartHistory{
		mementoArray: make([]*memento, 0),
	}
	// Set state A
	cart := &cart{state: "A"}
	fmt.Printf("Cart Current State: %s\n", cart.getState())
	cartHistory.addMemento(cart.createMemento())

	// Set state B
	cart.setState("B")
	fmt.Printf("Cart Current State: %s\n", cart.getState())
	cartHistory.addMemento(cart.createMemento())

	// Set state C
	cart.setState("C")
	fmt.Printf("Cart Current State: %s\n", cart.getState())
	cartHistory.addMemento(cart.createMemento())

	// Go back to B
	cart.restoreState(cartHistory.getMemento(1))
	fmt.Printf("Restored to State: %s\n", cart.getState())

	// Go back to A
	cart.restoreState(cartHistory.getMemento(0))
	fmt.Printf("Restored to State: %s\n", cart.getState())
}

