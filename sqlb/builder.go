package sqlb

import (
	"strconv"
	"strings"
)

type Builder struct {
	node    Node
	columns []string
	from    []string
	limit   *int
	offset  *int
	order   *Order
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

func (wb *Builder) buildSelect(b *strings.Builder) {
	if len(wb.columns) > 0 {
		b.WriteString("SELECT ")
		b.WriteString(strings.Join(wb.columns, ", "))
	}
	if len(wb.from) > 0 {
		b.WriteString(" FROM ")
		b.WriteString(strings.Join(wb.from, ", "))
	}
}

func (wb *Builder) BuildSQL() (string, []any, error) {
	if wb.node == nil {
		return "", []any{}, ErrEmptyWhereClause
	}
	fnode := wb.node
	for fnode.Prev() != nil {
		fnode = fnode.Prev()
	}
	values := []any{}
	var query strings.Builder
	wb.buildSelect(&query)
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
	if wb.order != nil && wb.order.column != nil {
		query.WriteString("ORDER BY ")
		query.WriteString(*wb.order.column)
		query.WriteString(" ")
		query.WriteString(wb.order.direction)
	}
	if wb.limit != nil {
		query.WriteString("LIMIT ")
		query.WriteString(strconv.Itoa(*wb.limit))
		query.WriteRune(' ')
	}
	if wb.offset != nil {
		query.WriteString("OFFSET ")
		query.WriteString(strconv.Itoa(*wb.offset))
		query.WriteRune(' ')
	}
	return strings.TrimSpace(query.String()), values, nil
}
