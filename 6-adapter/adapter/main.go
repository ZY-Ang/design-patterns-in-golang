package adapter

type Computer interface {
	insertUSB()
}

type MyMacBookWithUsbAdapter struct {MacBook}
func (c *MyMacBookWithUsbAdapter) insertUSB() {
	c.insertUSB_C()
}

type MacBook struct {}
func (c *MacBook) insertUSB_C() {} // oh no

type PC struct {}
func (c *PC) insertUSB() {}

func connectUsbMouse(c Computer) {
	c.insertUSB()
}

func main() {
	pc := &PC{}
	connectUsbMouse(pc) // Nice. PC is the best

	mac := &MacBook{}
	//connectUsbMouse(mac) // oh no
	macWithAdapter := &MyMacBookWithUsbAdapter{*mac}
	connectUsbMouse(macWithAdapter)
}
