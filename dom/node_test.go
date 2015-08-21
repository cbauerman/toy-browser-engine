package dom

import (
	"testing"
)

func TestNodes(t *testing.T) {
	PrettyPrint(NewElementNode("root", nil, []Node{
		NewTextNode("Test01"),
		NewTextNode("Test02"),
	}))
	PrettyPrint(NewElementNode("root", nil, []Node{
		NewElementNode("Node01", nil, []Node{
			NewTextNode("Test01"),
			NewTextNode("Test02"),
		}),
		NewTextNode("Test03"),
	}))
}
