package dom

import (
	"testing"
)

func TestNodes(t *testing.T) {
	PrettyPrint(Elem("root", nil, []*Node{
		Text("Test01"),
		Text("Test02"),
	}))
	PrettyPrint(Elem("root", nil, []*Node{
		Elem("Node01", nil, []*Node{
			Text("Test01"),
			Text("Test02"),
		}),
		Text("Test03"),
	}))
}
