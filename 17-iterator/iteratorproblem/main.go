package iteratorproblem

import "fmt"

// OrderHistory is list of order history for say, a user
type OrderHistory struct {
	orderIDs []uint64
}
func (oh *OrderHistory) push(orderID uint64) {
	if oh.orderIDs == nil {
		oh.orderIDs = make([]uint64, 0)
	}
	oh.orderIDs = append(oh.orderIDs, orderID)
}
func (oh *OrderHistory) pop() uint64 {
	if len(oh.orderIDs) > 0 {
		lastIndex := len(oh.orderIDs) - 1
		lastOrder := oh.orderIDs[lastIndex]
		oh.orderIDs = oh.orderIDs[:lastIndex]
		return lastOrder
	}
	return 0
}

func main() {
	oh := &OrderHistory{}
	oh.push(1)
	oh.push(2)
	oh.push(3)
	oh.push(4)
	oh.push(5)

	for i := 0; i < len(oh.orderIDs); i++ {
		orderID := oh.orderIDs[i]
		fmt.Printf("%d\n", orderID)
		// What if we decide to change our data
		//	structure that stores order history?
		// 	this tight coupling of array type
		//	disallows us to change the underlying
		//	implementation without affecting
		//	clients.
	}
}
