package marker

type Node interface {
	String() string
	ChildNodes() []Node
}
