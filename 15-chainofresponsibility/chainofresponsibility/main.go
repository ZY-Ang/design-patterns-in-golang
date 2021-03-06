package chainofresponsibility

import (
	"fmt"
)

type HttpRequest struct {
	username string
	password string
}
func (h *HttpRequest) Username() string {
	return h.username
}
func (h *HttpRequest) Password() string {
	return h.password
}


type WebServer struct {
	rootHandler *Handler
}
func (s *WebServer) handle(request *HttpRequest) {
	err := s.rootHandler.Handle(request)
	if err != nil {
		// e.g. in real code: res.SetError(err.Error())
		fmt.Printf("Oh no request failed with error %s\n", err.Error())
	}
}

// Handler is the abstract class for a handler in the chain
type Handler struct {
	iHandler
	next *Handler
}
type iHandler interface {
	doHandle(request *HttpRequest) error
}
// Handle is the template method of a Handler
func (h *Handler) Handle(request *HttpRequest) error {
	err := h.doHandle(request)
	if err != nil {
		// This handler failed. We can stop here.
		return err
	}

	// If there are still handlers in the chain
	if h.next != nil {
		return h.next.Handle(request)
	}
	// No more handlers in the chain. We end.
	return nil
}

type Authenticator struct {}
func (a *Authenticator) doHandle(request *HttpRequest) error {
	fmt.Println("Authenticate invoked")
	isValid := request.Username() == "admin" && request.Password() == "123456"
	if !isValid {
		return fmt.Errorf("invalid user %s", request.Username())
	}
	return nil
}

type Compressor struct {}
func (c *Compressor) doHandle(_ *HttpRequest) error {
	fmt.Println("Compress invoked, do some low level compression stuff on the response")
	return nil
}

type Metrics struct {}
func (l *Metrics) doHandle(_ *HttpRequest) error {
	fmt.Println("Collect invoked, send to prometheus or buffer")
	return nil
}

type Business struct {}
func (b *Business) doHandle(req *HttpRequest) error {
	fmt.Println("Do some stuff with user " + req.Username())
	return nil
}

func main() {
	// Say we want the flow of requests to go through the flow:
	// authenticator -> business -> compressor -> metrics

	// The chain is instantiated in reverse
	myApiHandlerMetric := &Handler{&Metrics{}, nil}
	myApiHandlerCompressor := &Handler{&Compressor{}, myApiHandlerMetric}
	myApiHandlerBusiness := &Handler{&Business{}, myApiHandlerCompressor}
	myApiRootHandler := &Handler{&Authenticator{}, myApiHandlerBusiness}

	// Using the chain
	ws := &WebServer{myApiRootHandler}
	ws.handle(&HttpRequest{
		username: "admin",
		password: "123456",
	})
}
