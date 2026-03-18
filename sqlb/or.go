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

func (o *OrNode) Eq(column string, value any) *EqNode {
	return &EqNode{
		column: column,
		value:  value,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func (o *OrNode) IsNull(column string) *IsNullNode {
	return &IsNullNode{
		column: column,
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
