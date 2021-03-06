package state

import "fmt"

const (
	StateNew = iota
	StateShipped
	StateCompleted
	StateCancelled
)

type Order struct {
	State
}
func (o *Order) update(newState State) {
	o.State = newState
	newState.SetOrder(o)
	if o.State != nil {
		o.Handle()
	}
}

type State interface {
	Handle()
	SetOrder(order *Order)
}

type ShippingStateHandler struct {
	*Order
}
func (h *ShippingStateHandler) Handle() {
	fmt.Println("Shipping state")
}
func (h *ShippingStateHandler) SetOrder(order *Order) {
	h.Order = order
}

type CancelledStateHandler struct {
	*Order
}
func (h *CancelledStateHandler) Handle() {
	fmt.Println("Cancelled state")
}
func (h *CancelledStateHandler) SetOrder(order *Order) {
	h.Order = order
}

func getStateMap(state int) State {
	switch state {
	case StateShipped:
		return &ShippingStateHandler{}
	case StateCancelled:
		return &CancelledStateHandler{}
	default:
		return nil
	}
}

func main() {
	order := &Order{}
	order.update(getStateMap(StateNew))
	order.update(getStateMap(StateShipped))
	order.update(getStateMap(StateCancelled))
}

