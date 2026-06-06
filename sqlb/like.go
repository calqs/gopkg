package sqlb

import (
	"strconv"
	"strings"
)

type LikeNode struct {
	EqNode
	wildcards Wildcards
}

type Wildcards int

const (
	WildcardStart Wildcards = 1 << iota
	WildcardEnd
	WildcardBoth = WildcardStart | WildcardEnd
)

func (l *LikeNode) ToSQL(depth int) (string, []any) {
	var sb strings.Builder
	sb.WriteString(l.column)
	sb.WriteRune(' ')
	sb.WriteString(string(l.comparison))
	sb.WriteRune(' ')
	if l.wildcards&WildcardStart != 0 {
		sb.WriteString("'%' || ")
	}
	sb.WriteRune('$')
	sb.WriteString(strconv.Itoa(depth))
	if l.wildcards&WildcardEnd != 0 {
		sb.WriteString(" || '%' ")
	}
	return sb.String(), []any{l.value}
}

func (l *LikeNode) Clone() Node {
	eqClone := l.EqNode.Clone().(*EqNode)
	return &LikeNode{
		EqNode:    *eqClone,
		wildcards: l.wildcards,
	}
}

func (wb *Builder) Like(column, value string, wildcards Wildcards) *Builder {
	return wb.pushNode(&LikeNode{
		EqNode: EqNode{
			column:     column,
			value:      value,
			comparison: ComparisonLike,
		},
		wildcards: wildcards,
	})
}

func Like(column, value string, wildcards Wildcards) *LikeNode {
	return &LikeNode{
		EqNode: EqNode{
			column:     column,
			value:      value,
			comparison: ComparisonLike,
			NodeRoutine: NodeRoutine{
				PrevNode: nil,
				NextNode: nil,
			},
		},
		wildcards: wildcards,
	}
}

func (wb *Builder) ILike(column, value string, wildcards Wildcards) *Builder {
	return wb.pushNode(&LikeNode{
		EqNode: EqNode{
			column:     column,
			value:      value,
			comparison: ComparisonILike,
		},
		wildcards: wildcards,
	})
}

func ILike(column, value string, wildcards Wildcards) *LikeNode {
	return &LikeNode{
		EqNode: EqNode{
			column:     column,
			value:      value,
			comparison: ComparisonILike,
			NodeRoutine: NodeRoutine{
				PrevNode: nil,
				NextNode: nil,
			},
		},
		wildcards: wildcards,
	}
}

func (a *AndNode) Like(column, value string, wildcards Wildcards) *LikeNode {
	like := Like(column, value, wildcards)
	a.NextNode = like
	like.PrevNode = a
	return like
}

func (a *AndNode) ILike(column, value string, wildcards Wildcards) *LikeNode {
	like := ILike(column, value, wildcards)
	a.NextNode = like
	like.PrevNode = a
	return like
}

func (o *OrNode) Like(column, value string, wildcards Wildcards) *LikeNode {
	like := Like(column, value, wildcards)
	o.NextNode = like
	like.PrevNode = o
	return like
}

func (o *OrNode) ILike(column, value string, wildcards Wildcards) *LikeNode {
	like := ILike(column, value, wildcards)
	o.NextNode = like
	like.PrevNode = o
	return like
}
