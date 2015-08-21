package htmlparser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParserConsumeWhile(t *testing.T) {
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

	if string(resultStart) != stringStart {
		t.Errorf("%s is the not the same as %s", string(resultStart), stringStart)
	}

	if string(resultEnd) != stringEnd {
		t.Errorf("%s is the not the same as %s", string(resultEnd), stringEnd)
	}

}
