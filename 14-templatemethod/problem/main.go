package problem

import "fmt"

type UpdateOrderStatusTask struct {
	auditTrail *AuditTrail
}
func (t *UpdateOrderStatusTask) execute() {
	t.auditTrail.record()

	fmt.Println("Update Order Status")
}

type UpdateLogisticStatusTask struct {
	auditTrail *AuditTrail
}
func (t *UpdateLogisticStatusTask) execute() {
	t.auditTrail.record()

	fmt.Println("Update Logistics Status")
}

type AuditTrail struct {}
func (at *AuditTrail) record() {
	fmt.Println("Audit")
}
