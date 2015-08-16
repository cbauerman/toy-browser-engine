package dom

import (
	"fmt"
)

type Node struct {
	Children []*Node
	//for now type is decided by the data
	NodeType interface{}
}

type ElementData struct {
	TagName    string
	Attributes map[string]string
}

func Text(data string) *Node {
	return &Node{
		Children: nil,
		NodeType: data,
	}
}

func Elem(name string, attributes map[string]string, children []*Node) *Node {
	return &Node{
		Children: children,
		NodeType: &ElementData{
			TagName:    name,
			Attributes: attributes,
		},
	}
}

func PrettyPrint(tree *Node) {
	prettyPrintHelper(tree, 0, false)
}

func prettyPrintHelper(tree *Node, indent int, last bool) {
	switch tree.NodeType.(type) {

	case string:
		printIndent(fmt.Sprintf("Text Node : %s\n", tree.NodeType), indent, last)
	case *ElementData:
		elem := tree.NodeType.(*ElementData)
		printIndent(fmt.Sprintf("Elem Node with tag : %s\n", elem.TagName), indent, last)
		if tree.Children != nil {
			for index, child := range tree.Children {
				prettyPrintHelper(child, indent+1, (index == len(tree.Children)-1))
			}
		}
	default:
		printIndent(fmt.Sprintf("Unknown Node of Type %T\n", tree.NodeType), indent, last)
	}
}

func printIndent(text string, indent int, last bool) {

	if indent != 0 {
		for i := 0; i < (indent - 1); i++ {
			fmt.Printf("|   ")
		}
		if last {
			fmt.Printf("`-- ")
		} else {
			fmt.Printf("|-- ")
		}
	}
	fmt.Printf("%s", text)
}
