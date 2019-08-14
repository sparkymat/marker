package marker

type Attribute uint8

const (
	Bold Attribute = 1 << iota
	Italic
)

type AttrNode struct {
	Attributes Attribute
	Children   []Node
}

func (a AttrNode) IsBold() bool {
	return a.Attributes&Bold != 0
}

func (a AttrNode) IsItalic() bool {
	return a.Attributes&Italic != 0
}

func (a AttrNode) String() string {
	return ""
}

func (a AttrNode) ChildNodes() []Node {
	return a.Children
}
