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
