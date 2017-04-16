package ast

type ModeAttr struct {
	Address  string
	Position string
	Name     string
	Children []Node
}

func parseModeAttr(line string) *ModeAttr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> (?P<name>.+)",
		line,
	)

	return &ModeAttr{
		Address:  groups["address"],
		Position: groups["position"],
		Name:     groups["name"],
		Children: []Node{},
	}
}

func (n *ModeAttr) render(ast *Ast) (string, string) {
	return "", ""
}
