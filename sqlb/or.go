package sqlb

type OrNode struct {
	NodeRoutine
}

func (o *OrNode) ToSQL(depth int) (string, []any) {
	return string(OperationOr), []any{}
}

func (o *OrNode) Clone() Node {
	return &OrNode{
		NodeRoutine: *o.NodeRoutine.Clone(),
	}
}

func Or() *OrNode {
	return &OrNode{
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func OrBlock(node Node, rest ...Node) *Token {
	head := OpenParenthesis()
	var tail Node = head

	nHead := getHead(node)
	nTail := getTail(node)
	tail.SetNext(nHead)
	nHead.SetPrev(tail)
	tail = nTail

	for _, r := range rest {
		orNode := Or()
		tail.SetNext(orNode)
		orNode.SetPrev(tail)
		tail = orNode

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

func (eq *EqNode) Or() *OrNode {
	or := Or()
	eq.NextNode = or
	or.PrevNode = eq
	return or
}
