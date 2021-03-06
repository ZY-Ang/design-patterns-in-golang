package visitor

import "fmt"

const (
	nodeTypeH1 = iota
	nodeTypeP
)

type Operation interface {
	apply(node HtmlNode)
}
type OperationHighlight struct {}
func (* OperationHighlight) apply(node HtmlNode) {
	nodeTypeHandler := map[int]func(){
		nodeTypeH1: func() {
			fmt.Println("highlighted h1")
		},
		nodeTypeP: func() {
			fmt.Println("highlighted p")
		},
	}
	if handler, ok := nodeTypeHandler[node.nodeType()]; ok {
		handler()
		return
	}
	panic(fmt.Sprintf(
		"Node type %d is invalid",
		node.nodeType(),
	))
}
type OperationToMarkdown struct {}
func (*OperationToMarkdown) apply(node HtmlNode) {
	nodeTypeHandler := map[int]func(){
		nodeTypeH1: func() {
			fmt.Printf("# %s\n\n", node.text())
		},
		nodeTypeP: func() {
			fmt.Printf("%s\n", node.text())
		},
	}
	if handler, ok := nodeTypeHandler[node.nodeType()]; ok {
		handler()
		return
	}
	panic(fmt.Sprintf(
		"Node type %d is invalid",
		node.nodeType(),
	))
}


type HtmlNode interface {
	execute(o Operation)
	nodeType() int
	text() string
}

type H1Node struct {
	value string
}
func (h *H1Node) execute(o Operation) {
	o.apply(h)
}
func (h *H1Node) nodeType() int {
	return nodeTypeH1
}
func (h *H1Node) text() string {
	return h.value
}

type ParagraphNode struct {
	value string
}
func (h *ParagraphNode) execute(o Operation) {
	o.apply(h)
}
func (h *ParagraphNode) nodeType() int {
	return nodeTypeP
}
func (h *ParagraphNode) text() string {
	return h.value
}

type HtmlDocument struct {
	nodes []HtmlNode
}
func (d *HtmlDocument) append(node HtmlNode) {
	if d.nodes == nil {
		d.nodes = make([]HtmlNode, 0)
	}
	d.nodes = append(d.nodes, node)
}
func (d *HtmlDocument) highlight() {
	for _, node := range d.nodes {
		node.execute(&OperationHighlight{})
	}
}
func (d *HtmlDocument) innerText() {
	for _, node := range d.nodes {
		node.execute(&OperationToMarkdown{})
	}
}

func main() {
	document := &HtmlDocument{}
	document.append(&H1Node{"Hello World"})
	document.append(&ParagraphNode{"Welcome to the"})
	document.append(&ParagraphNode{"last design pattern"})

	document.highlight()
	document.innerText()
}
