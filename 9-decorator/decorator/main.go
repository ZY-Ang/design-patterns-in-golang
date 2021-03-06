package decorator

import "fmt"

func hello(unwrapped func()) func() {
	return func() {
		fmt.Println("Hello")
		unwrapped()
	}
}

func _name() {
	fmt.Println("Alice")
}

var name = hello(_name)

func main() {
	name()

	// Nested wrapping
	hello(hello(hello(_name)))()
}
