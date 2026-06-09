package sqlb

type Node interface {
	ToSQL(int) (string, []any)
	Next() Node
	Prev() Node
	SetNext(Node)
	SetPrev(Node)
	Clone() Node
}

type NodeRoutine struct {
	NextNode Node
	PrevNode Node
}

func (n *NodeRoutine) Clone() *NodeRoutine {
	return &NodeRoutine{
		NextNode: nil,
		PrevNode: nil,
	}
}

func (n *NodeRoutine) Next() Node {
	return n.NextNode
}

func (n *NodeRoutine) Prev() Node {
	return n.PrevNode
}

func (n *NodeRoutine) SetNext(node Node) {
	n.NextNode = node
}

func (n *NodeRoutine) SetPrev(node Node) {
	n.PrevNode = node
}
