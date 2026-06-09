package sqlb

type VoidNode struct {
	NodeRoutine
	node Node
}

func Void() *VoidNode {
	return &VoidNode{}
}

func Placeholder() *VoidNode {
	return Void()
}

func V() *VoidNode {
	return Void()
}

func (p *VoidNode) ToSQL(i int) (string, []any) {
	if p.node != nil {
		return p.node.ToSQL(i)
	}
	return "", nil
}

func (p *VoidNode) Clone() Node {
	return Void()
}

func (p *VoidNode) Gte(column string, value any) *EqNode {
	p.node = Gte(column, value)
	return p.node.(*EqNode)
}

func (p *VoidNode) Gt(column string, value any) *EqNode {
	p.node = Gt(column, value)
	return p.node.(*EqNode)
}

func (p *VoidNode) Lt(column string, value any) *EqNode {
	p.node = Lt(column, value)
	return p.node.(*EqNode)
}

func (p *VoidNode) Lte(column string, value any) *EqNode {
	p.node = Lte(column, value)
	return p.node.(*EqNode)
}

func (p *VoidNode) Eq(column string, value any) *EqNode {
	p.node = Eq(column, value)
	return p.node.(*EqNode)
}

func (p *VoidNode) In(column string, values ...any) *InNode {
	p.node = In(column, values...)
	return p.node.(*InNode)
}

func (p *VoidNode) IsNull(column string) *IsNullNode {
	p.node = IsNull(column)
	return p.node.(*IsNullNode)
}

func (p *VoidNode) NotNull(column string) *NotNullNode {
	p.node = NotNull(column)
	return p.node.(*NotNullNode)
}

func (p *VoidNode) And() *AndNode {
	p.node = And()
	return p.node.(*AndNode)
}

func (p *VoidNode) Or() *OrNode {
	p.node = Or()
	return p.node.(*OrNode)
}

func (p *VoidNode) AndBlock(node Node, rest ...Node) *AndNode {
	p.node = AndBlock(node, rest...)
	return p.node.(*AndNode)
}

func (p *VoidNode) OrBlock(node Node, rest ...Node) *OrNode {
	p.node = OrBlock(node, rest...)
	return p.node.(*OrNode)
}
