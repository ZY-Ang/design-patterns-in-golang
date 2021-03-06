package mediator

import (
	"fmt"
	"sync"
	"time"
)

const (
	// Shared states
	stateNew = iota
	stateCancelled
	stateShipping
	stateComplete
)

type Mediator interface {
	// CanFulfil checks if an order can be fulfilled
	CanFulfil(order Order) bool
	// NotifyRiderAvailable notifies that we have
	//	riders available to fulfil the food
	NotifyRiderAvailable(order Order)

	// CanCancel checks if an order can be cancelled
	CanCancel(order Order) bool
	// NotifyCancelled notifies subscribers of cancellation
	NotifyCancelled(order Order)
}

type OrderManager struct {
	isRiderAvailable bool
	riderLock        *sync.Mutex // say we are a startup and we only have one riderLock available :(
	ordersQueue      []Order
}
func (om *OrderManager) CanFulfil(order Order) bool {
	om.riderLock.Lock()
	defer om.riderLock.Unlock()
	if om.isRiderAvailable {
		om.isRiderAvailable = false
		return true
	}
	om.ordersQueue = append(om.ordersQueue, order)
	return false
}
func (om *OrderManager) NotifyRiderAvailable(order Order) {
	om.riderLock.Lock()
	defer om.riderLock.Unlock()
	if !om.isRiderAvailable {
		om.isRiderAvailable = true
	}
	if len(om.ordersQueue) > 0 {
		firstOrderInQueue := om.ordersQueue[0]
		om.ordersQueue = om.ordersQueue[1:]
		firstOrderInQueue.requestFulfil()
	}
}
func (om *OrderManager) CanCancel(order Order) bool {
	if order.orderType() == "physical" {
		return order.orderState() == stateNew

	} else if order.orderType() == "food" {
		var cancellableStates = map[uint64]bool{
			stateNew: true,
			stateShipping: true, // maybe we're foodpanda, idk
		}
		ok, exists := cancellableStates[order.orderState()]
		return exists && ok

	}
	return false
}
func (om *OrderManager) NotifyCancelled(o Order) {
	for i, order := range om.ordersQueue {
		if order.ID() == o.ID() {
			om.ordersQueue = append(om.ordersQueue[:i], om.ordersQueue[i+1:]...)
		}
	}
}


type Order interface {
	ID() uint64
	orderType() string
	orderState() uint64
	requestCancel()
	requestFulfil()
}

type PhysicalOrder struct {
	Mediator
	id, state uint64
}
func (o *PhysicalOrder) ID() uint64 {
	return o.id
}
func (o *PhysicalOrder) requestFulfil() {
	if o.CanFulfil(o) {
		fmt.Printf("Fulfilling physical order %d...\n", o.ID())
		o.state = stateShipping
		time.Sleep(time.Second)
		o.state = stateComplete
		o.NotifyRiderAvailable(o)

	} else {
		fmt.Println("Not possible to fulfil at the moment, will try again later.")
	}
}
func (o *PhysicalOrder) orderState() uint64 {
	return o.state
}
func (o *PhysicalOrder) orderType() string {
	return "physical"
}
func (o *PhysicalOrder) requestCancel() {
	if o.CanCancel(o) {
		o.state = stateCancelled
		o.NotifyCancelled(o)

	} else {
		fmt.Println("Order is not cancellable")
	}
}

type FoodOrder struct {
	Mediator
	id, state uint64
}
func (o *FoodOrder) ID() uint64 {
	return o.id
}
func (o *FoodOrder) requestFulfil() {
	if o.CanFulfil(o) {
		fmt.Printf("Fulfilling food order %d...\n", o.ID())
		o.state = stateShipping
		time.Sleep(time.Second)
		o.state = stateComplete
		o.NotifyRiderAvailable(o)

	} else {
		fmt.Println("Not possible to fulfil at the moment, will try again later.")
	}
}
func (o *FoodOrder) orderState() uint64 {
	return o.state
}
func (o *FoodOrder) orderType() string {
	return "physical"
}
func (o *FoodOrder) requestCancel() {
	if o.CanCancel(o) {
		o.state = stateCancelled
		o.NotifyCancelled(o)

	} else {
		fmt.Println("Order is not cancellable")
	}
}

func newOrder(otype string, id uint64, orderMediator Mediator) Order {
	if otype == "food" {
		return &FoodOrder{
			Mediator: orderMediator,
			id:       id,
			state:    stateNew,
		}
	} else {
		return &PhysicalOrder{
			Mediator: orderMediator,
			id:       id,
			state:    stateNew,
		}
	}
}

func main() {
	orderManager := &OrderManager{
		isRiderAvailable: true,
		riderLock:        &sync.Mutex{},
		ordersQueue:      make([]Order, 0),
	}
	food1 := newOrder("food", 1, orderManager)
	phy1 := newOrder("phys", 2, orderManager)
	food2 := newOrder("food", 3, orderManager)
	phy2 := newOrder("phys", 4, orderManager)
	food3 := newOrder("food", 5, orderManager)
	phy3 := newOrder("phys", 6, orderManager)

	// 1,2,3,4 will be fulfilled in order with delay
	food1.requestFulfil()
	phy1.requestFulfil()
	food2.requestFulfil()
	phy2.requestFulfil()

	// Cancel phy3 - success
	phy3.requestCancel()

	// Start shipping food3 but cancel - fail
	food3.requestFulfil()
	food3.requestCancel()
}
