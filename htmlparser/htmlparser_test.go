package htmlparser

import (
	"fmt"
	"strings"
	"testing"
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

func TestParseTagName(t * testing.T){
	tagName := "1897fadhusdas9fha4afbwj"
	nonTagName := "$&(@*#("
	
	testString := fmt.Sprintf(tagName, nonTagName)
	
	input := strings.NewReader(testString)
	p := newParser(input)
	
	tagNameResult := p.parseTagName()
	
	if tagNameResult != tagName{
		t.Errorf("%s is not the same as %s", tagNameResult, tagName)
	}
}

func TestParseText(t * testing.T){
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