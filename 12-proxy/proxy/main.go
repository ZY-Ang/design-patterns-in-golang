package proxy

import "fmt"

type server interface {
	handleRequest(req, res string) int
}

type application struct {}
func (a *application) handleRequest(_, _ string) int {
	fmt.Println("request is handled. Yay.")
	return 0
}

type Nginx struct {
	server
}
func (a *Nginx) handleRequest(req, res string) int {
	fmt.Println("Do some stuff like rate limiting before handling request")
	return a.server.handleRequest(req, res)
}

func main() {

}
