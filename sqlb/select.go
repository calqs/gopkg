package sqlb

func Select(columns ...string) *Builder {
	if len(columns) == 0 {
		columns = []string{"*"}
	}
	return &Builder{
		columns: columns,
	}
}

func From(from ...string) *Builder {
	return &Builder{
		from: from,
	}
}
