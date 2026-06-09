package sqlb

type AndNode struct {
	NodeRoutine
}

func (a *AndNode) Clone() Node {
	return &AndNode{
		NodeRoutine: *a.NodeRoutine.Clone(),
	}
}

func (a *AndNode) ToSQL(depth int) (string, []any) {
	return string(OperationAnd), []any{}
}

func And() *AndNode {
	return &AndNode{
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func AndBlock(node Node, rest ...Node) *Token {
	head := OpenParenthesis()
	var tail Node = head

	nHead := getHead(node)
	nTail := getTail(node)
	tail.SetNext(nHead)
	nHead.SetPrev(tail)
	tail = nTail

	for _, r := range rest {
		andNode := And()
		tail.SetNext(andNode)
		andNode.SetPrev(tail)
		tail = andNode

		rHead := getHead(r)
		rTail := getTail(r)
		tail.SetNext(rHead)
		rHead.SetPrev(tail)
		tail = rTail
	}

	closeNode := CloseParenthesis()
	tail.SetNext(closeNode)
	closeNode.SetPrev(tail)

	return head
}

func (eq *EqNode) And() *AndNode {
	and := And()
	eq.NextNode = and
	and.PrevNode = eq
	return and
}

func (n *IsNullNode) And() *AndNode {
	and := And()
	n.NextNode = and
	and.PrevNode = n
	return and
}
