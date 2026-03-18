package sqlb

import "fmt"

type IsNullNode struct {
	column string
	NodeRoutine
}

func (n *IsNullNode) ToSQL(depth int) (string, []any) {
	return fmt.Sprintf("%s IS NULL", n.column), []any{}
}

func IsNull(column string) *IsNullNode {
	return &IsNullNode{
		column: column,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func (n *IsNullNode) And() *AndNode {
	and := And()
	n.NextNode = and
	and.PrevNode = n
	return and
}
