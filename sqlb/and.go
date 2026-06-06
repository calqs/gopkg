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

func AndBlock(node Node, rest ...Node) *AndNode {
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
	and := And()
	and.SetNext(n)
	and.NextNode.SetPrev(and)
	return and
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
