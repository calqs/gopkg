package sqlb

func Select(columns ...string) *Builder {
	return &Builder{
		columns: columns,
	}
}

func (s *Builder) Where(nodes ...Node) *Builder {
	b := Where(nodes...)
	b.columns = s.columns
	b.from = s.from
	return b
}

func From(from ...string) *Builder {
	return &Builder{
		from: from,
	}
}
