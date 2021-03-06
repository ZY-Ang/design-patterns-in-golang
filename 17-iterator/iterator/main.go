package iterator

import "fmt"

// Iterator interface
type Iterator interface {
	HasNext() bool
	Current() interface{}
	Next()
}

// OhOrder is an order of OrderHistory
type OhOrder struct {
	ID uint64
	next *OhOrder
	prev *OhOrder
}

// OrderHistory handles initialization and
//	management of order history
type OrderHistory struct {
	head *OhOrder
	tail *OhOrder
}
func (oh *OrderHistory) push(orderID uint64) {
	if oh.head == nil {
		oh.head = &OhOrder{orderID,nil,nil}
		oh.tail = oh.head
	} else {
		prev := oh.tail
		oh.tail = &OhOrder{orderID, nil, prev}
		prev.next = oh.tail
	}
}
func (oh *OrderHistory) pop() uint64 {
	panic("implement")
}
func (oh *OrderHistory) getIterator() Iterator {
	return &OhLinkedListIterator{oh.head}
}

// OhLinkedListIterator is a nested class of OrderHistory
type OhLinkedListIterator struct {
	current *OhOrder
}
func (i *OhLinkedListIterator) HasNext() bool {
	return i.current != nil
}
func (i *OhLinkedListIterator) Current() interface{} {
	if i.current != nil {
		return i.current.ID
	}
	return 0
}
func (i *OhLinkedListIterator) Next() {
	if i.current != nil {
		i.current = i.current.next
	}
}

func main() {
	oh := &OrderHistory{}
	oh.push(1)
	oh.push(2)
	oh.push(3)
	oh.push(4)
	oh.push(5)

	ohIterator := oh.getIterator()
	for ; ohIterator.HasNext(); ohIterator.Next() {
		currentID := ohIterator.Current().(uint64)
		fmt.Printf("cur: %d\n", currentID)
	}
}
