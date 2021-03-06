package command

import "fmt"

type Order struct {
	ID uint64
	UserID uint64
	ShopID uint64
}

// Command is the contract that
//	governs how requests should
//	be executed
type Command interface {
	execute(interface{})
}

type UndoableCommand interface {
	execute(interface{})
	unexecute(interface{})
}

type UndoableApiProcessor struct {
	UndoableCommand
}
func (p *UndoableApiProcessor) do(data interface{}) {
	p.execute(data)
}
func (p *UndoableApiProcessor) undo(data interface{}) {
	p.unexecute(data)
}

// Action is the command executor
type ApiProcessor struct {
	Command
	name string
}
func (b *ApiProcessor) do(data interface{}) {
	// ApiProcessorAction.do simply executes
	//	the command without knowing the
	//	context of the command being executed
	b.execute(data)
}

type CompositeCommand struct {
	commands []Command
}
func (c *CompositeCommand) Add(cmd Command) {
	if c.commands == nil {
		c.commands = make([]Command, 0)
	}
	c.commands = append(c.commands, cmd)
}
func (c *CompositeCommand) execute(data interface{}) {
	for _, cmd := range c.commands {
		cmd.execute(data)
	}
}

// AddOrderCommand is a concrete command with
//	the context to add an order
type AddOrderCommand struct {
	system *OrderManagementSystem
}
func (c *AddOrderCommand) execute(data interface{}) {
	order := data.(*Order)
	c.system.AddOrderToSystem(order)
}
func (c *AddOrderCommand) unexecute(data interface{}) {
	order := data.(*Order)
	c.system.RemoveOrderFromSystem(order)
}

type AddItemToOrderCommand struct {
	system *OrderManagementSystem
}
func (c *AddItemToOrderCommand) execute(_ interface{}) {
	fmt.Println("Mock add item to order")
}

type ChangeUserIDForOrderIDCommand struct {
	system *OrderManagementSystem
}
func (c *ChangeUserIDForOrderIDCommand) execute(_ interface{}) {
	fmt.Println("Mock change shop ID for Order ID")
}

// OrderManagementSystem is the receiver
type OrderManagementSystem struct {
	orders map[uint64]*Order // map[orderID]*Order
}
func (o *OrderManagementSystem) AddOrderToSystem(order *Order) {
	fmt.Printf("Adding order %d to system...\n", order.ID)
	if o.orders == nil {
		o.orders = make(map[uint64]*Order)
	}
	o.orders[order.ID] = order
}
func (o *OrderManagementSystem) RemoveOrderFromSystem(order *Order) {
	if o.orders == nil {
		o.orders = make(map[uint64]*Order)
		return
	}
	if od, ok := o.orders[order.ID]; ok {
		fmt.Printf("Removing order %d from system...\n", order.ID)
		delete(o.orders, od.ID)
	}
}

func main() {
	oms := &OrderManagementSystem{}
	addOrderApi := &ApiProcessor{
		&AddOrderCommand{oms},
		"add_order_api",
	}
	// Only the caller, concrete AddOrderCommand and
	//	OrderManagementSystem knows the context of the
	//	command being executed. ApiProcessor doesn't
	//	actually know it's doing an add order command
	addOrderApi.do(&Order{
		123,
		456,
		789,
	})

	giftOrderCommand := &CompositeCommand{}
	giftOrderCommand.Add(&AddOrderCommand{oms})
	giftOrderCommand.Add(&AddItemToOrderCommand{oms})
	giftOrderCommand.Add(&ChangeUserIDForOrderIDCommand{oms})
	giftOrderApi := &ApiProcessor{
		giftOrderCommand,
		"gift_order_api",
	}
	giftOrderApi.do(&Order{111,222,333})

	undoableAddOrderApi := &UndoableApiProcessor{
		&AddOrderCommand{oms},
	}
	order := &Order{987,654,321}
	undoableAddOrderApi.do(order)
	undoableAddOrderApi.undo(order)
}
