package htmlparser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"toy-browser-engine/dom"
)

func TestConsumeWhile(t *testing.T) {
	stringStart := "----------"
	stringEnd := "This is a test"
	testString := fmt.Sprint(stringStart, stringEnd)

	input := strings.NewReader(testString)
	p := newParser(input)

	resultStart := p.consumeWhile(func(r rune) bool {
		return (r == '-')
	})

	resultEnd := p.consumeWhile(func(r rune) bool {
		return true
	})

	if resultStart != stringStart {
		t.Errorf("%s is not the same as %s", string(resultStart), stringStart)
	}

	if resultEnd != stringEnd {
		t.Errorf("%s is not the same as %s", string(resultEnd), stringEnd)
	}

}

func TestParseTagName(t *testing.T) {
	tagName := "1897fadhusdas9fha4afbwj"
	nonTagName := "$&(@*#("

	testString := fmt.Sprintf(tagName, nonTagName)

	input := strings.NewReader(testString)
	p := newParser(input)

	tagNameResult := p.parseTagName()

	if tagNameResult != tagName {
		t.Errorf("%s is not the same as %s", tagNameResult, tagName)
	}
}

func TestParseText(t *testing.T) {
	text := "1897fadhusdas9fha4afbwj"
	nonText := "<test>"

	testString := fmt.Sprint(text, nonText)

	input := strings.NewReader(testString)
	p := newParser(input)

	textResult := p.parseText()

	if textResult.Value != text {
		t.Errorf("%s is not the same as %s", textResult.Value, text)
	}
}

func TestParseAttributeValue(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"\"Value01\"GARBAGE!@#$", "Value01"},
		{"'Value02'MOREGARBAGE!*@$(*21312", "Value02"},
	}

	for _, c := range cases {
		p := newParser(strings.NewReader(c.in))
		got := p.parseAttributeValue()
		if got != c.want {
			t.Errorf("parseAttribute value on %s == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestParseAttribute(t *testing.T) {
	cases := []struct {
		in, want1, want2 string
	}{
		{"Var01=\"Value01\"GARBAGE!@#$", "Var01", "Value01"},
		{"Var02='Value02'MOREGARBAGE!*@$(*21312", "Var02", "Value02"},
	}

	for _, c := range cases {
		p := newParser(strings.NewReader(c.in))
		got1, got2 := p.parseAttribute()
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("parseAttribute on %s == (%s, %s) want (%s, %s)", c.in, got1, got2, c.want1, c.want2)
		}
	}
}

func TestParseAttributes(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]string
	}{
		{"Var01=\"Value01\" Var02=\"Value02\">", map[string]string{
			"Var01": "Value01",
			"Var02": "Value02",
		}},
	}

	for _, c := range cases {
		p := newParser(strings.NewReader(c.in))
		got := p.parseAttributes()
		eq := reflect.DeepEqual(c.want, got)
		if !eq {
			t.Logf("parseAttributes on %s failed!\n", c.in)
			t.Logf("Expected:\n")
			for k, v := range c.want {
				t.Logf("%s : %s\n", k, v)
			}
			t.Logf("Got:\n")
			for k, v := range got {
				t.Logf("%s : %s\n", k, v)
			}
			t.Fail()
		}
	}
}

func TestParseElement(t *testing.T) {
	cases := []struct {
		in string
		want *dom.Element
	}{
		{"<tag name=\"name\"></tag>", dom.NewElementNode("tag", map[string]string{
			"name" : "name",
		}, nil )},
	}
	
	for _, c := range cases {
		p := newParser(strings.NewReader(c.in))
		got := p.parseElement()
		
		if got != c.want {
			t.Logf("Failure of parserElement on %s", c.in)
			t.Logf("Expected:\n")
			dom.PrettyPrint(c.want)
			t.Logf("Got:\n")
			dom.PrettyPrint(got)
			t.Fail()
		}
	}
}