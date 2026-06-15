package sqlb

import (
	"fmt"
	"strings"
)

type InsertBuilder struct {
	table   string
	columns []string
	values  []any
}

func Insert(table string) *InsertBuilder {
	return &InsertBuilder{
		table: table,
	}
}

func (ib *InsertBuilder) Columns(cols ...string) *InsertBuilder {
	ib.columns = append(ib.columns, cols...)
	return ib
}

func (ib *InsertBuilder) Values(vals ...any) *InsertBuilder {
	ib.values = append(ib.values, vals...)
	return ib
}

func (ib *InsertBuilder) Set(col string, val any) *InsertBuilder {
	ib.columns = append(ib.columns, col)
	ib.values = append(ib.values, val)
	return ib
}

func (ib *InsertBuilder) BuildSQL() (string, []any, error) {
	if len(ib.columns) != len(ib.values) {
		return "", nil, fmt.Errorf("InsertBuilder: columns length (%d) does not match values length (%d)", len(ib.columns), len(ib.values))
	}

	var query strings.Builder
	query.WriteString("INSERT INTO ")
	query.WriteString(ib.table)

	if len(ib.columns) > 0 {
		query.WriteString(" (")
		query.WriteString(strings.Join(ib.columns, ", "))
		query.WriteString(")")
	}

	query.WriteString(" VALUES (")
	for i := range ib.values {
		if i > 0 {
			query.WriteString(", ")
		}
		query.WriteString(fmt.Sprintf("$%d", i+1))
	}
	query.WriteString(")")

	return query.String(), ib.values, nil
}
