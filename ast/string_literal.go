package ast

import (
	"fmt"
	"strings"
)

type StringLiteral struct {
	Address  string
	Position string
	Type     string
	Value    string
	Lvalue   bool
	Children []Node
}

func parseStringLiteral(line string) *StringLiteral {
	groups := groupsFromRegex(
		`<(?P<position>.*)> '(?P<type>.*)' lvalue "(?P<value>.*)"`,
		line,
	)

	return &StringLiteral{
		Address:  groups["address"],
		Position: groups["position"],
		Type:     groups["type"],
		Value:    unescapeString(groups["value"]),
		Lvalue:   true,
		Children: []Node{},
	}
}

func (n *StringLiteral) render(ast *Ast) (string, string) {
	src := fmt.Sprintf("\"%s\"", strings.Replace(n.Value, "\n", "\\n", -1))
	return src, "const char *"
}
