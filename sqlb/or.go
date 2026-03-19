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
	eq := Eq(column, value)
	o.NextNode = eq
	eq.PrevNode = o
	return eq
}

func (o *OrNode) Gt(column string, value any) *EqNode {
	eq := Gt(column, value)
	o.NextNode = eq
	eq.PrevNode = o
	return eq
}

func (o *OrNode) Like(column, value string, wildcards Wildcards) *LikeNode {
	like := Like(column, value, wildcards)
	o.NextNode = like
	like.PrevNode = o
	return like
}

func (o *OrNode) ILike(column, value string, wildcards Wildcards) *LikeNode {
	like := ILike(column, value, wildcards)
	o.NextNode = like
	like.PrevNode = o
	return like
}

func (o *OrNode) IsNull(column string) *IsNullNode {
	isNull := IsNull(column)
	o.NextNode = isNull
	isNull.PrevNode = o
	return isNull
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
