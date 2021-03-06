package prototype

import (
	"fmt"
)

// itemListing is cloneable Prototype interface
//	imagine this is what appears on ecommerce platforms
type itemListing interface {
	duplicate() itemListing
	print(indentation string)
}

// item is a base concrete itemListing
type item struct {
	title string
	price uint64
}
func (i *item) duplicate() itemListing {
	return &item{i.title + " copy", i.price}
}
func (i *item) print(indentation string) {
	fmt.Println(indentation + i.title)
}

// bundle is concrete "folder" for itemListing
//	that can be bundle deals
type bundle struct {
	title string
	price uint64
	children []itemListing
}
func (i *bundle) duplicate() itemListing {
	cloneBundle := &bundle{title: i.title + " copy", price: i.price}
	children := make([]itemListing, len(i.children))
	for j, child := range i.children {
		children[j] = child.duplicate()
	}
	cloneBundle.children = children
	return cloneBundle
}
func (i *bundle) print(indentation string) {
	fmt.Println(indentation + i.title)
	for _, child := range i.children {
		child.print(indentation + indentation)
	}
}
func (i *bundle) addListing(listing itemListing) {
	i.children = append(i.children, listing)
}

func main() {
	fakePods := &item{"- orange fakepods", 239}
	doodelBuds := &item{"- doodel buds", 150}

	fakePodsBundle := &bundle{
		title: "Fakepods buy-1-get-1-free bundle",
		price: 239,
	}
	fakePodsBundle.addListing(fakePods)
	fakePodsBundle.addListing(fakePods.duplicate())

	fmt.Println("======================================")
	fakePodsBundle.print("  ")
	// Fakepods buy-1-get-1-free bundle
	//   - orange fakepods
	//   - orange fakepods copy

	fmt.Println("======================================")
	superBundle := &bundle{
		title: "Super tws earbuds bundle",
		price: 300,
	}
	superBundle.addListing(doodelBuds)
	superBundle.addListing(fakePodsBundle.duplicate())
	superBundle.print("  ")
	// Super tws earbuds bundle
	//   - doodel buds
	//   Fakepods buy-1-get-1-free bundle copy
	//       - orange fakepods copy
	//       - orange fakepods copy copy
}
