package templatemethod

import "fmt"

type iTask interface {
	doExecute()
}
type Task struct {
	iTask
	auditTrail *AuditTrail
}
func (t *Task) execute() {
	t.auditTrail.record()
	t.doExecute()
}

type UpdateOrderStatusTask struct {}
func (t *UpdateOrderStatusTask) doExecute() {
	fmt.Println("Update Order Status")
}


type UpdateLogisticStatusTask struct {}
func (t *UpdateLogisticStatusTask) doExecute() {
	fmt.Println("Update Logistics Status")
}

type AuditTrail struct {}
func (at *AuditTrail) record() {
	fmt.Println("I am an audit log")
}

func main() {
	updateOrderStatus := &Task{
		&UpdateOrderStatusTask{},
		&AuditTrail{},
	}
	updateOrderStatus.execute()

	updateLogisticStatus := &Task{
		&UpdateLogisticStatusTask{},
		&AuditTrail{},
	}
	updateLogisticStatus.execute()
}
