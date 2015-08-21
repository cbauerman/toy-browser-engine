package htmlparser

import (
	"bufio"
	"fmt"
	"io"
	"toy-browser-engine/dom"
	"unicode"
)

type parser struct {
	input      *bufio.Reader
	error      func(p *parser, msg string)
	errorCount uint
}

func newParser(input io.Reader) *parser {

	p := &parser{
		input:      bufio.NewReader(input),
		errorCount: 0,
		error: func(p *parser, msg string) {
			fmt.Printf("Error in Parser: %s", msg)
		},
	}
	return p
}

const eof = -2

func (p *parser) next() rune {
	next, _, err := p.input.ReadRune()

	if err != nil && err == io.EOF {
		return eof
	} else if err != nil {
		p.errorCount++
		if p.error != nil {
			p.error(p, err.Error())
		}
	}
	return next
}

func (p *parser) peek() rune {
	next := p.next()
	p.input.UnreadRune()
	return next
}

func (p *parser) startsWith(str string) bool {
	return true
}

type testRune func(rune) bool

func (p *parser) consumeWhile(test testRune) string {
	result := make([]rune, 0, 100)

	for p.peek() != eof && test(p.peek()) {
		//fmt.Printf("Reading %q\n", p.peek())
		result = append(result, p.next())
	}

	//p.input.UnreadRune()

	return string(result)
}

func (p *parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}

func (p *parser) parseText() *dom.Text {
	return dom.NewTextNode(p.consumeWhile(func(r rune) bool {
		return (r != '<')
	}))
}

func (p *parser) parseTagName() string {
	return p.consumeWhile(func(r rune) bool {
		return (unicode.IsLetter(r) || unicode.IsDigit(r))
	})
}

func (p *parser) parseAttributeValue() string {
	openQuote := p.next()
	if openQuote != '"' || openQuote != '\'' {
		//error here
	}
	value := p.consumeWhile(func(r rune) bool {
		return r != openQuote
	})
	if p.next() != openQuote {
		//error here
	}
	return value
}

func (p *parser) parseAttribute() (string, string) {
	name := p.parseTagName()

	if p.next() != '=' {
		//need an error here
	}

	value := p.parseAttributeValue()

	return name, value
}

func (p *parser) parseAttributes() map[string]string {
	attributes := make(map[string]string)

	for  {
		p.consumeWhitespace() 
		if p.peek() == '>' {
			break
		}
		
		name, value := p.parseAttribute()
		attributes[name] = value
	}

	return attributes
}

func (p *parser) parseElement() dom.Node {
	if p.next() != '<' {
		//error on missing '<'
	}

	tagName := p.parseTagName()
	attributes := p.parseAttributes()

	if p.next() != '>' {
		//error on missing '>'
	}

	children := p.parseNodes()

	if p.next() != '<' {
		//errors a plenty
	}
	if p.next() != '/' {
		//errors a plenty
	}
	if p.parseTagName() != tagName {
		//errors a plenty
	}
	if p.next() != '>' {
		//errors a plenty
	}

	return dom.NewElementNode(tagName, attributes, children)
}


func (p *parser) parseNode() dom.Node {
	switch p.peek() {
	case '<':
		return p.parseElement()
	default:
		return p.parseText()
	}
}

func (p *parser) parseNodes() []dom.Node {
	nodes := make([]dom.Node, 0)
	for {
		 p.consumeWhitespace()
		 if p.peek() == eof || p.startsWith("</") {
			 break
		 }
		nodes = append(nodes, p.parseNode())
	}
	return nodes
}

func Parse(source io.Reader) dom.Node {
	nodes := newParser(source).parseNodes()

	if len(nodes) == 1 {
		return nodes[0]
	} else {
		return dom.NewElementNode("html", make(map[string]string), nodes)
	}
}
