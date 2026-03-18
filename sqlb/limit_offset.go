package sqlb

func (wb *Builder) Limit(limit int) *Builder {
	wb.limit = &limit
	return wb
}

func (wb *Builder) Offset(offset int) *Builder {
	wb.offset = &offset
	return wb
}
