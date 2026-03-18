package sqlb

import "fmt"

type EqNode struct {
	column string
	value  any
	NodeRoutine
}

func (eq *EqNode) ToSQL(depth int) (string, []any) {
	return fmt.Sprintf("%s %s $%d", eq.column, ComparisonEq, depth), []any{eq.value}
}

func (eq *EqNode) And() *AndNode {
	and := And()
	eq.NextNode = and
	and.PrevNode = eq
	return and
}

func Eq(column string, value any) *EqNode {
	return &EqNode{
		column: column,
		value:  value,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}
