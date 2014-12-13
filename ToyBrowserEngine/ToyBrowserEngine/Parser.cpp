#include "Parser.h"

#include <cwctype>
#include <cctype>
#include <cassert>

char Parser::nextChar()
{
	return input.peek();
}
bool Parser::startsWith(std::string str)
{
	std::string in = input.str().substr(0, str.length());
	int result = str.compare(in);
	return (0 == result);
}

bool Parser::eof()
{
	/* operation to trigger error bits*/
	input.peek();
	return input.eof();
}

char Parser::consumeChar()
{
	char result;
	input >> result;
	return result;
}

std::string Parser::consumeWhile(bool(*f)(char))
{
	std::stringstream result;
	char ch;
	while ((*f)(result.peek())){
		input >> ch;
		result << ch;
	}
	return result.str();
}

void Parser::consumeWhitespace()
{
	consumeWhile(
		[](char c)
	{
		return (bool)iswspace(c);
	});
}

std::string Parser::parseTagName()
{
	return consumeWhile(
		[](char c)
	{
		return (bool)(isalnum(c));
	});
}

dom::Node *Parser::parseNode()
{
	if (nextChar() == '<'){
		return parseElement();
	} else {
		return parseText();
	}
}

dom::Text *Parser::parseText()
{
	return new dom::Text(
		consumeWhile(
		[](char c){ return c != '<'; }
		));
}

dom::Element *Parser::parseElement()
{
	assert(consumeChar() == '<');
	std::string tagName = parseTagName();
	dom::AttrMap attrs = parseAttributes();
	assert(consumeChar() == '>');

	std::vector<dom::Node*> children = parseNodes();

	/* closing tag */
	assert(consumeChar() == '<');
	assert(consumeChar() == '/');
	assert(parseTagName() == tagName);
	assert(consumeChar() == '>');

	return new dom::Element(tagName, attrs, children);
}

dom::attr Parser::parseAttr()
{
	std::string name = parseTagName();
	assert(consumeChar() == '=');
	std::string value = parseAttrValue();
	return dom::attr(name, value);
}

std::string Parser::parseAttrValue()
{
	char openQuote = consumeChar();
	assert(openQuote == '"');
	std::string value = consumeWhile(
		[](char c) -> bool{return (c != '"'); });
	assert(consumeChar() == '"');
	return value;
}

dom::AttrMap Parser::parseAttributes()
{
	dom::AttrMap attributes = std::hash_map < std::string, std::string>();

	while (true) {
		consumeWhitespace();
		if (nextChar() == '>'){
			break;
		}

		dom::attr attribute = parseAttr();
		attributes.insert(attribute.name, attribute.value);
	}

	return attributes;
}

std::vector<dom::Node*> Parser::parseNodes()
{
	std::vector<dom::Node*> nodes;

	while (true) {
		consumeWhitespace();

		if (eof() || startsWith("</")){
			nodes.push_back(parseNode());
		}
	}

	return nodes;

}

dom::Node *parse(std::string source){
	Parser parser = Parser(source);

	std::vector<dom::Node*> nodes = parser.parseNodes();

	if (nodes.size() == 1) {
		return nodes[0];
	}
	else {
		std::string html("html");
		return &dom::Element(html,
			std::hash_map < std::string, std::string>(),
			nodes); 
	}
}