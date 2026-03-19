package sqlb

type OrNode struct {
	NodeRoutine
}

func (o *OrNode) ToSQL(depth int) (string, []any) {
	return string(OperationOr), []any{}
}

func Or() *OrNode {
	return &OrNode{
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func OrBlock(node Node, rest ...Node) *OrNode {
	chain := node
	for _, r := range rest {
		r.SetPrev(chain)
		chain.SetNext(r)
		chain = r
	}
	chain.SetNext(CloseParenthesis())
	chain.Next().SetPrev(chain)
	for chain.Prev() != nil {
		chain = chain.Prev()
	}
	n := OpenParenthesis()
	chain.SetPrev(n)
	n.SetNext(chain)
	or := Or()
	or.SetNext(n)
	or.NextNode.SetPrev(or)
	return or
}

func (eq *EqNode) Or() *OrNode {
	or := Or()
	eq.NextNode = or
	or.PrevNode = eq
	return or
}
