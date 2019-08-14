package marker

type TextNode string

func (t TextNode) String() string {
	return string(t)
}

func (t TextNode) ChildNodes() []Node {
	return []Node{}
}
