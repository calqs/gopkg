package sqlb

type AndNode struct {
	NodeRoutine
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

func (a *AndNode) Eq(column string, value any) *EqNode {
	eq := Eq(column, value)
	a.NextNode = eq
	eq.PrevNode = a
	return eq
}

func (a *AndNode) IsNull(column string) *IsNullNode {
	isNull := IsNull(column)
	a.NextNode = isNull
	isNull.PrevNode = a
	return isNull
}

func (a *AndNode) Like(column, value string, wildcards Wildcards) *LikeNode {
	like := Like(column, value, wildcards)
	a.NextNode = like
	like.PrevNode = a
	return like
}

func (a *AndNode) ILike(column, value string, wildcards Wildcards) *LikeNode {
	like := ILike(column, value, wildcards)
	a.NextNode = like
	like.PrevNode = a
	return like
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
