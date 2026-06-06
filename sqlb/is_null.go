package sqlb

import "fmt"

type IsNullNode struct {
	column string
	NodeRoutine
}

func (n *IsNullNode) ToSQL(depth int) (string, []any) {
	return fmt.Sprintf("%s IS NULL", n.column), []any{}
}

func (n *IsNullNode) Clone() Node {
	return &IsNullNode{
		column:      n.column,
		NodeRoutine: *n.NodeRoutine.Clone(),
	}
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

type NotNullNode struct {
	column string
	NodeRoutine
}

func (n *NotNullNode) Clone() Node {
	return &NotNullNode{
		column:      n.column,
		NodeRoutine: *n.NodeRoutine.Clone(),
	}
}

func (n *NotNullNode) ToSQL(depth int) (string, []any) {
	return fmt.Sprintf("%s IS NOT NULL", n.column), []any{}
}

func NotNull(column string) *NotNullNode {
	return &NotNullNode{
		column: column,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}
