package sqlb

func Select(columns ...string) *Builder {
	return &Builder{
		columns: columns,
	}
}

func From(from ...string) *Builder {
	return &Builder{
		from: from,
	}
}
