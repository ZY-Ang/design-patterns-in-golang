package factorymethod

import (
	"fmt"
)

const (
	TypeCessna = "Cessna"
)

type planeInterface interface {
	setName(name string)
	Name() string
	PrintPlane()
}
type plane struct {
	name string
}
func (g *plane) setName(name string) {
	g.name = name
}
func (g *plane) Name() string {
	return g.name
}
func (g *plane) PrintPlane() {
	fmt.Printf("Plane: %s\n", g.name)
}

type cessna struct {
	plane
}

type planeFactory interface {
	makePlane() planeInterface
}
type cessnaFactory struct {}
func newCessnaFactory() planeFactory {
	return &cessnaFactory{}
}

// makePlane is the factory method of the cessna factory
func (p *cessnaFactory) makePlane() planeInterface {
	return &cessna{plane{TypeCessna}}
}

func main() {
	cessnaFactory := newCessnaFactory()
	cessna := cessnaFactory.makePlane()
	cessna.PrintPlane()
}
