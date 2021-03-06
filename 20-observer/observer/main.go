package observer

import "fmt"

type Order struct {
	// ID will be used as pointer to a flyweight pool somewhere else
	ID int
	state int
	observers []OrderStateChangeObserver
}
func (ds *Order) State() int {
	return ds.state
}
func (ds *Order) SetState(newState int) {
	// When setting the state, we need to
	//	notify listeners that this state
	//	has changed.
	// However, we don't know which objects
	//	are subscribed to this change in
	//	this setter method, which may even
	//	change at runtime.
	oldState := ds.state

	// We can have different hooks
	ds.notifyWillUpdate(oldState, newState)
	ds.state = newState
	ds.notifyDidUpdate(oldState, newState)
}
func (ds *Order) addObserver(observer OrderStateChangeObserver) {
	if ds.observers == nil {
		ds.observers = make([]OrderStateChangeObserver, 0)
	}
	ds.observers = append(ds.observers, observer)
	observer.SetOrder(ds)
}
func (ds *Order) removeObserver(observer OrderStateChangeObserver) {
	for i, obs := range ds.observers {
		if obs.ObserverID() == observer.ObserverID() {
			ds.observers = append(ds.observers[:i], ds.observers[i+1:]...)
		}
	}
}

// Notification with push style communication causing coupling on the observer end.
func (ds *Order) notifyWillUpdate(old, new int) {
	for _, obs := range ds.observers {
		obs.WillUpdate(old, new)
	}
}
func (ds *Order) notifyDidUpdate(old, new int) {
	for _, obs := range ds.observers {
		obs.DidUpdate(old, new)
	}
}

// OrderStateChangeObserver is an observer interface
type OrderStateChangeObserver interface {
	ObserverID() int
	// SetOrder allows for pull-type communication which causes coupling on the concrete implementation instead of observer interface.
	SetOrder(order *Order)
	WillUpdate(old, new int)
	DidUpdate(old, new int)
}
const (
	observerShipping = iota + 1
	observerDownstream
)

// ShippingTeam is a concrete listener
//	to order state change
type ShippingTeam struct {
	*Order
}
func (*ShippingTeam) ObserverID() int {
	return observerShipping
}
func (s *ShippingTeam) SetOrder(order *Order) {
	s.Order = order
}
func (*ShippingTeam) WillUpdate(old, new int) {
	fmt.Printf("ShippingTeam handler: Order will update from %d to %d\n", old, new)
}
func (*ShippingTeam) DidUpdate(old, new int) {
	fmt.Printf("ShippingTeam handler: Order did update from %d to %d\n", old, new)
}
// DownstreamHandler is a concrete listener
//	to order state change
type DownstreamHandler struct {
	*Order
}
func (*DownstreamHandler) ObserverID() int {
	return observerDownstream
}
func (s *DownstreamHandler) SetOrder(order *Order) {
	s.Order = order
}
func (*DownstreamHandler) WillUpdate(old, new int) {
	fmt.Printf("ShippingTeam handler: Order will update from %d to %d\n", old, new)
}
func (*DownstreamHandler) DidUpdate(old, new int) {
	fmt.Printf("ShippingTeam handler: Order did update from %d to %d\n", old, new)
}
