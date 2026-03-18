package sqlb

type Order struct {
	column    *string
	direction string
}

func (wb *Builder) OrderBy(column string) *Builder {
	wb.order = &Order{
		column:    &column,
		direction: defaultDirection.String(),
	}
	return wb
}

const defaultDirection = Asc

type Direction string

const (
	Asc  Direction = "ASC"
	Desc Direction = "DESC"
)

func (d Direction) String() string {
	switch d {
	case Asc:
		return "ASC"
	case Desc:
		return "DESC"
	default:
		return defaultDirection.String()
	}
}

func (wb *Builder) OrderDir(direction Direction) *Builder {
	if wb.order == nil {
		wb.order = &Order{
			direction: direction.String(),
		}
	} else {
		wb.order.direction = direction.String()
	}
	return wb
}
