package sqlb

import (
	"fmt"
	"strings"
)

type EqNode struct {
	column     string
	value      any
	comparison Comparison
	NodeRoutine
}

func (eq *EqNode) ToSQL(depth int) (string, []any) {
	if v, ok := eq.value.(string); ok {
		if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) {
			return fmt.Sprintf("%s %s %s", eq.column, eq.comparison, v), []any{}
		}
	}
	return fmt.Sprintf("%s %s $%d", eq.column, eq.comparison, depth), []any{eq.value}
}

func Eq(column string, value any) *EqNode {
	return &EqNode{
		column:     column,
		value:      value,
		comparison: ComparisonEq,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func Gt(column string, value any) *EqNode {
	return &EqNode{
		column:     column,
		value:      value,
		comparison: ComparisonGt,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func Lt(column string, value any) *EqNode {
	return &EqNode{
		column:     column,
		value:      value,
		comparison: ComparisonLt,
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

func (a *AndNode) Gt(column string, value any) *EqNode {
	eq := Gt(column, value)
	a.NextNode = eq
	eq.PrevNode = a
	return eq
}

func (a *AndNode) Lt(column string, value any) *EqNode {
	eq := Lt(column, value)
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

func (o *OrNode) IsNull(column string) *IsNullNode {
	isNull := IsNull(column)
	o.NextNode = isNull
	isNull.PrevNode = o
	return isNull
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
