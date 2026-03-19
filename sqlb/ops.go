package sqlb

type Operation string

const (
	OperationAnd  Operation = "AND"
	OperationOr   Operation = "OR"
	OperationNone Operation = ""
)

type Comparison string

const (
	ComparisonEq    Comparison = "="
	ComparisonNeq   Comparison = "<>"
	ComparisonGt    Comparison = ">"
	ComparisonLt    Comparison = "<"
	ComparisonGeq   Comparison = ">="
	ComparisonLeq   Comparison = "<="
	ComparisonLike  Comparison = "LIKE"
	ComparisonILike Comparison = "ILIKE"
	ComparisonOn    Comparison = "ON"
)
