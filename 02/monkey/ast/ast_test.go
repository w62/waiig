package ast

import (
  "monkey/token"
  "testing"
)

func TestString (t *testing.T) {
  program := &Program{
    Statements: []Statement{
      &LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name: &Identifier{
        Token: token.Token{Type: token.IDENT, Literal: "myVar"},
          Value: "myVar",
        },
        Value: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
          Value: "anotherVar",
        },
      },
    },
  }

  if program.String() != "let myVar = anotherVar;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }
}


// Test for one return statement
func TestString2 (t *testing.T) {
  program := &Program{
    Statements: []Statement{
      &ReturnStatement{
        Token: token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "Var2"},
          Value: "Var2",
        },
      },
    },
  }

  if program.String() != "return Var2;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }
}


// Test for two return Statements
func TestString3 (t *testing.T) {
  program := &Program{
    Statements: []Statement{
      &ReturnStatement{
        Token: token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "Var2"},
          Value: "Var2",
        },
      },
      &ReturnStatement{
        Token: token.Token{Type: token.RETURN, Literal: "return"},
        ReturnValue: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "Var3"},
          Value: "Var3",
        },
      },

    },
  }

  if program.String() != "return Var2;return Var3;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }
}
