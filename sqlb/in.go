package sqlb

import (
	"fmt"
	"strings"
)

type InNode struct {
	EqNode
	values []any
}

func (i *InNode) ToSQL(depth int) (string, []any) {
	placeholders := make([]string, len(i.values))
	for idx := range i.values {
		placeholders[idx] = fmt.Sprintf("$%d", depth+idx)
	}
	return fmt.Sprintf("%s %s (%s)", i.column, i.comparison, strings.Join(placeholders, ", ")), i.values
}

func (i *InNode) Clone() Node {
	eqClone := i.EqNode.Clone().(*EqNode)
	return &InNode{
		EqNode: *eqClone,
		values: append([]any{}, i.values...), // deep copy slices
	}
}

func In(column string, values ...any) *InNode {
	return &InNode{
		EqNode: EqNode{
			column:     column,
			value:      nil,
			comparison: ComparisonIn,
			NodeRoutine: NodeRoutine{
				PrevNode: nil,
				NextNode: nil,
			},
		},
		values: values,
	}
}

func (a *AndNode) In(column string, values ...any) *InNode {
	in := In(column, values...)
	a.NextNode = in
	in.PrevNode = a
	return in
}

func (o *OrNode) In(column string, values ...any) *InNode {
	in := In(column, values...)
	o.NextNode = in
	in.PrevNode = o
	return in
}
