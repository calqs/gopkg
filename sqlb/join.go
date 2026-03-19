package sqlb

import "strings"

type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"
	LeftJoin  JoinType = "LEFT JOIN"
	RightJoin JoinType = "RIGHT JOIN"
	FullJoin  JoinType = "FULL JOIN"
)

type JoinNode struct {
	table string
	*EqNode
	joinType JoinType
}

func (jn *JoinNode) ToSQL(it int) (string, []any) {
	sql, val := jn.EqNode.ToSQL(it)
	if len(val) > 0 {
		if v, ok := val[0].(string); ok {
			sql = strings.Replace(sql, "$0", v, 1)
		}
	}
	return string(jn.joinType) + " " + jn.table + " ON " + sql, nil
}

func join(table string, eq *EqNode, joinType JoinType) *JoinNode {
	return &JoinNode{
		table:    table,
		EqNode:   eq,
		joinType: joinType,
	}
}

func (wb *Builder) Join(table string, eq *EqNode) *Builder {
	return wb.pushJoin(join(table, eq, InnerJoin))
}

func (wb *Builder) LeftJoin(table string, eq *EqNode) *Builder {
	return wb.pushJoin(join(table, eq, LeftJoin))
}

func (wb *Builder) RightJoin(table string, eq *EqNode) *Builder {
	return wb.pushJoin(join(table, eq, RightJoin))
}

func (wb *Builder) FullJoin(table string, eq *EqNode) *Builder {
	return wb.pushJoin(join(table, eq, FullJoin))
}
