package sqlb

import (
	"strings"
)

type Operation string

const (
	OperationAnd  Operation = "AND"
	OperationOr   Operation = "OR"
	OperationNone Operation = ""
)

type Comparison string

const (
	ComparisonEq  Comparison = "="
	ComparisonNeq Comparison = "<>"
	ComparisonGt  Comparison = ">"
	ComparisonLt  Comparison = "<"
	ComparisonGeq Comparison = ">="
	ComparisonLeq Comparison = "<="
)

type Node interface {
	ToSQL(int) (string, []any)
	Next() Node
	Prev() Node
	SetNext(node Node)
	SetPrev(node Node)
}

type NodeRoutine struct {
	NextNode Node
	PrevNode Node
}

func (n *NodeRoutine) Next() Node {
	return n.NextNode
}

func (n *NodeRoutine) Prev() Node {
	return n.PrevNode
}

func (n *NodeRoutine) SetNext(node Node) {
	n.NextNode = node
}

func (n *NodeRoutine) SetPrev(node Node) {
	n.PrevNode = node
}

type Builder struct {
	node    Node
	columns []string
	from    []string
}

func Where(nodes ...Node) *Builder {
	if len(nodes) == 0 {
		return &Builder{}
	}
	node := nodes[0]
	for _, n := range nodes[1:] {
		node.SetNext(n)
		n.SetPrev(node)
		node = n
	}
	return &Builder{
		node: node,
	}
}

func (wb *Builder) pushNode(node Node) *Builder {
	if wb.node == nil {
		wb.node = node
		return wb
	}
	for ; wb.node.Next() != nil; wb.node = wb.node.Next() {
	}
	wb.node.SetNext(node)
	node.SetPrev(wb.node)
	return wb
}

func (wb *Builder) Eq(column string, value any) *Builder {
	return wb.pushNode(Eq(column, value))
}

func (wb *Builder) IsNull(column string) *Builder {
	return wb.pushNode(IsNull(column))
}

func (wb *Builder) And() *Builder {
	return wb.pushNode(And())
}

func (wb *Builder) Or() *Builder {
	return wb.pushNode(Or())
}

func (wb *Builder) OrBlock(node Node, rest ...Node) *Builder {
	return wb.pushNode(OrBlock(node, rest...))
}

func (wb *Builder) Select(columns ...string) *Builder {
	wb.columns = columns
	return wb
}

func (wb *Builder) From(from ...string) *Builder {
	wb.from = from
	return wb
}

func (wb *Builder) BuildSQL() (string, []any, error) {
	if wb.node == nil {
		return "", []any{}, ErrNilPointer
	}
	fnode := wb.node
	for fnode.Prev() != nil {
		fnode = fnode.Prev()
	}

	values := []any{}
	var query strings.Builder
	if len(wb.columns) > 0 {
		query.WriteString("SELECT ")
		query.WriteString(strings.Join(wb.columns, ", "))
	}
	if len(wb.from) > 0 {
		query.WriteString(" FROM ")
		query.WriteString(strings.Join(wb.from, ", "))
	}
	if wb.node != nil {
		query.WriteString(" WHERE ")
	}
	for it := 1; fnode != nil; {
		sql, val := fnode.ToSQL(it)
		query.WriteString(strings.TrimSpace(sql))
		query.WriteRune(' ')
		for _, v := range val {
			values = append(values, v)
			it++
		}
		fnode = fnode.Next()
	}
	return strings.TrimSpace(query.String()), values, nil
}
