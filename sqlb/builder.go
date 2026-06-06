package sqlb

import (
	"strconv"
	"strings"

	"github.com/calqs/gopkg/dt"
)

type Builder struct {
	node    Node
	columns []string
	from    []string
	joins   Node
	limit   *int
	offset  *int
	order   *Order
	groupBy []string
}

func (wb *Builder) Clone() *Builder {
	var limit *int
	if wb.limit != nil {
		limit = dt.Ptr(*wb.limit)
	}
	var offset *int
	if wb.offset != nil {
		offset = dt.Ptr(*wb.offset)
	}
	var order *Order
	if wb.order != nil {
		order = &Order{
			direction: wb.order.direction,
		}
		if wb.order.column != nil {
			order.column = dt.Ptr(*wb.order.column)
		}
	}
	return &Builder{
		node:    cloneNodeChain(wb.node),
		columns: append([]string{}, wb.columns...),
		from:    append([]string{}, wb.from...),
		joins:   cloneNodeChain(wb.joins),
		limit:   limit,
		offset:  offset,
		order:   order,
		groupBy: append([]string{}, wb.groupBy...),
	}
}

func (wb *Builder) pushJoin(node Node) *Builder {
	if wb.joins == nil {
		wb.joins = node
		return wb
	}
	for ; wb.joins.Next() != nil; wb.joins = wb.joins.Next() {
	}
	wb.joins.SetNext(node)
	node.SetPrev(wb.joins)
	return wb
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
		b.WriteRune(' ')
	}
	if len(wb.from) > 0 {
		b.WriteString("FROM ")
		b.WriteString(strings.Join(wb.from, ", "))
		b.WriteRune(' ')
	}
}

func (wb *Builder) buildJoins(b *strings.Builder) {
	if wb.joins == nil {
		return
	}
	for ; wb.joins.Prev() != nil; wb.joins = wb.joins.Prev() {
	}
	for wb.joins != nil {
		sql, _ := wb.joins.ToSQL(0)
		b.WriteString(strings.TrimSpace(sql))
		b.WriteRune(' ')
		wb.joins = wb.joins.Next()
	}
}

func (wb *Builder) buildWhere(b *strings.Builder) ([]any, error) {
	values := []any{}
	if wb.node == nil {
		return values, nil
	}
	fnode := wb.node
	for fnode.Prev() != nil {
		fnode = fnode.Prev()
	}
	if wb.node != nil {
		b.WriteString("WHERE ")
	}
	for it := 1; fnode != nil; {
		sql, val := fnode.ToSQL(it)
		b.WriteString(strings.TrimSpace(sql))
		b.WriteRune(' ')
		for _, v := range val {
			values = append(values, v)
			it++
		}
		fnode = fnode.Next()
	}
	return values, nil
}

func (wb *Builder) GroupBy(cols ...string) *Builder {
	wb.groupBy = append(wb.groupBy, cols...)
	return wb
}

func (wb *Builder) BuildSQL() (string, []any, error) {
	var query strings.Builder
	wb.buildSelect(&query)
	wb.buildJoins(&query)
	values, err := wb.buildWhere(&query)
	if err != nil {
		return "", nil, err
	}
	if len(wb.groupBy) > 0 {
		query.WriteString("GROUP BY ")
		query.WriteString(strings.Join(wb.groupBy, ", "))
		query.WriteRune(' ')
	}
	if wb.order != nil && wb.order.column != nil {
		query.WriteString("ORDER BY ")
		query.WriteString(*wb.order.column)
		query.WriteRune(' ')
		query.WriteString(wb.order.direction)
		query.WriteRune(' ')
	}
	if wb.limit != nil {
		query.WriteString("LIMIT ")
		query.WriteString(strconv.Itoa(*wb.limit))
		query.WriteRune(' ')
	}
	if wb.offset != nil {
		query.WriteString("OFFSET ")
		query.WriteString(strconv.Itoa(*wb.offset))
	}
	return strings.TrimSpace(query.String()), values, nil
}
