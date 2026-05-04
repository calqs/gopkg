package sqlb

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

func (wb *Builder) Eq(column string, value any) *Builder {
	return wb.pushNode(Eq(column, value))
}

func (wb *Builder) Gt(column string, value any) *Builder {
	return wb.pushNode(Gt(column, value))
}

func (wb *Builder) Gte(column string, value any) *Builder {
	return wb.pushNode(Gte(column, value))
}

func (wb *Builder) Lt(column string, value any) *Builder {
	return wb.pushNode(Lt(column, value))
}

func (wb *Builder) Lte(column string, value any) *Builder {
	return wb.pushNode(Lte(column, value))
}

func (wb *Builder) In(column string, values ...any) *Builder {
	return wb.pushNode(In(column, values...))
}

func (wb *Builder) IsNull(column string) *Builder {
	return wb.pushNode(IsNull(column))
}

func (wb *Builder) NotNull(column string) *Builder {
	return wb.pushNode(NotNull(column))
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

func (wb *Builder) AndBlock(node Node, rest ...Node) *Builder {
	return wb.pushNode(AndBlock(node, rest...))
}

func (wb *Builder) Select(columns ...string) *Builder {
	wb.columns = columns
	return wb
}

func (wb *Builder) From(from ...string) *Builder {
	wb.from = from
	return wb
}

func (s *Builder) Where(nodes ...Node) *Builder {
	b := Where(nodes...)
	b.joins = s.joins
	b.columns = s.columns
	b.from = s.from
	return b
}
