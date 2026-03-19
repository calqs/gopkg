package sqlb

type TokenType string

const (
	TokenTypeOpenParenthesis  TokenType = "("
	TokenTypeCloseParenthesis TokenType = ")"
)

type Token struct {
	NodeRoutine
	Type TokenType
}

func (t *Token) ToSQL(_ int) (string, []any) {
	return string(t.Type), []any{}
}

func (t *Token) Eq(column string, value any) *EqNode {
	return &EqNode{
		column: column,
		value:  value,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func (t *Token) Gt(column string, value any) *EqNode {
	return &EqNode{
		column: column,
		value:  value,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func (t *Token) Lt(column string, value any) *EqNode {
	return &EqNode{
		column: column,
		value:  value,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func (t *Token) IsNull(column string) *IsNullNode {
	return &IsNullNode{
		column: column,
		NodeRoutine: NodeRoutine{
			PrevNode: nil,
			NextNode: nil,
		},
	}
}

func OpenParenthesis() *Token {
	return &Token{
		Type: TokenTypeOpenParenthesis,
	}
}

func CloseParenthesis() *Token {
	return &Token{
		Type: TokenTypeCloseParenthesis,
	}
}
