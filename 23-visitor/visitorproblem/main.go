package visitorproblem

import "fmt"

type HtmlNode interface {
	highlight()
	// FIXME: can't implement innerText() string
}

type H1Node struct {}
func (h H1Node) highlight() {
	fmt.Println("highlighted h1")
}

type ParagraphNode struct {}
func (h ParagraphNode) highlight() {
	fmt.Println("highlighted p")
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
		node.highlight()
	}
}

func main() {
	document := &HtmlDocument{}
	document.append(&H1Node{})
	document.append(&ParagraphNode{})
}
