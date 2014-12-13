#ifndef _NODE_H_
#define _NODE_H_

#include<vector>
#include<string>
#include<hash_map>

namespace dom
{

	struct attr{

		attr(std::string name_, std::string value_) 
		: name(name_), value(value_) {}

		std::string name;
		std::string value;
	};

	typedef std::hash_map<std::string, std::string> AttrMap;

	class Node 
	{
	public:
		Node(std::vector<Node*> &children = std::vector<Node*>()) : children_(children) {}

	private:
		std::vector<Node*> children_;
	};

	class Element : public Node
	{
	public:
		Element(std::string &name, AttrMap &attributes, std::vector<Node*> &children = std::vector<Node*>())
		: tagName_(name), attributes_(attributes), Node(children) {}

	private:
		std::string tagName_;
		AttrMap attributes_;
	};

	class Text : public Node
	{
	public:
		Text(std::string &text) : text_(text) {}
	private:
		std::string text_;
	};
}

#endif /* _NODE_H_ */