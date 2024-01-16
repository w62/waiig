package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
  input := `=+(){},;`

  tests = [] struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.PLUS, "+"},
  }
}
