package sqlb

type VoidNode struct {
	NodeRoutine
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

func (p *VoidNode) Gte(column string, value any) *EqNode {
	return Gte(column, value)
}

func (p *VoidNode) Gt(column string, value any) *EqNode {
	return Gt(column, value)
}

func (p *VoidNode) Lt(column string, value any) *EqNode {
	return Lt(column, value)
}

func (p *VoidNode) Lte(column string, value any) *EqNode {
	return Lte(column, value)
}

func (p *VoidNode) Eq(column string, value any) *EqNode {
	return Eq(column, value)
}

func (p *VoidNode) In(column string, values ...any) *InNode {
	return In(column, values...)
}

func (p *VoidNode) IsNull(column string) *IsNullNode {
	return IsNull(column)
}

func (p *VoidNode) NotNull(column string) *NotNullNode {
	return NotNull(column)
}

func (p *VoidNode) And() *AndNode {
	return And()
}

func (p *VoidNode) Or() *OrNode {
	return Or()
}

func (p *VoidNode) AndBlock(node Node, rest ...Node) *AndNode {
	return AndBlock(node, rest...)
}

func (p *VoidNode) OrBlock(node Node, rest ...Node) *OrNode {
	return OrBlock(node, rest...)
}
