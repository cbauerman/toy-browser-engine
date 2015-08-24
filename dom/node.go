package dom

import (
	"fmt"
)

type NodeType int

const (
	TextNode NodeType = iota
	ElementNode
)

type Node interface {
	GetNodeType() NodeType
}

type Text struct {
	Children []Node
	Value    string
}

func (t Text) GetNodeType() NodeType {
	return TextNode
}

type Element struct {
	Children   []Node
	TagName    string
	Attributes map[string]string
}

func (e Element) GetNodeType() NodeType {
	return ElementNode
}

func NewTextNode(data string) *Text {
	return &Text{
		Children: nil,
		Value:    data,
	}
}

func NewElementNode(name string, attributes map[string]string, children []Node) *Element {
	return &Element{
		Children:   children,
		TagName:    name,
		Attributes: attributes,
	}
}

func PrettyPrint(tree Node) {
	prettyPrintHelper(tree, 0, false)
}

func prettyPrintHelper(tree Node, indent int, last bool) {
	switch tree.GetNodeType() {

	case TextNode:
		text := tree.(*Text)
		printIndent(fmt.Sprintf("Text Node : %s\n", text.Value), indent, last)
	case ElementNode:
		elem := tree.(*Element)
		printIndent(fmt.Sprintf("Elem Node with tag : %s\n", elem.TagName), indent, last)
		for k, c := range elem.Attributes {
			printIndent(fmt.Sprintf("%s : %s\n", k, c), indent + 1, last)
		}
		if elem.Children != nil {
			for index, child := range elem.Children {
				prettyPrintHelper(child, indent+1, (index == len(elem.Children)-1))
			}
		}
	default:
		printIndent(fmt.Sprintf("Unknown Node: %d", tree.GetNodeType()), indent, last)
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
