package bridge

import "fmt"

type order interface {
	flushToDB()
	setDBFlusher(db)
}

type db interface {
	flushOrder(orderID string)
}

type physicalOrder struct {
	db db
	orderID string
}
func (m *physicalOrder) flushToDB() {
	fmt.Println("flushToDB request for physicalOrder")
	m.db.flushOrder(m.orderID)
}
func (m *physicalOrder) setDBFlusher(p db) {
	m.db = p
}

type foodOrder struct {
	db db
	orderID string
}
func (w *foodOrder) flushToDB() {
	fmt.Println("flushToDB request for foodOrder")
	w.db.flushOrder(w.orderID)
}
func (w *foodOrder) setDBFlusher(p db) {
	w.db = p
}

type mysql struct {}
func (p *mysql) flushOrder(orderID string) {
	fmt.Printf("mysql db flushing order %s\n", orderID)
}

type nosql struct {}
func (p *nosql) flushOrder(orderID string) {
	fmt.Printf("nosql db flushing order %s\n", orderID)
}

func newOrder(orderType, orderID string) order {
	if orderType == "physical" {
		return &physicalOrder{orderID: orderID}
	}
	return &foodOrder{orderID: orderID}
}

func main() {
	nosqlDB := &nosql{}
	mysqlDB := &mysql{}
	phoneCover := newOrder("physical", "12345")
	kaleSalad := newOrder("food", "67890")

	phoneCover.setDBFlusher(nosqlDB)
	phoneCover.flushToDB()

	phoneCover.setDBFlusher(mysqlDB)
	phoneCover.flushToDB()

	kaleSalad.setDBFlusher(nosqlDB)
	kaleSalad.flushToDB()

	kaleSalad.setDBFlusher(mysqlDB)
	kaleSalad.flushToDB()
}
