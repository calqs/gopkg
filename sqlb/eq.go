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

func (eq *EqNode) And() *AndNode {
	and := And()
	eq.NextNode = and
	and.PrevNode = eq
	return and
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
