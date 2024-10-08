package ast

import (
	"testing"

	"github.com/dhruvdabhi101/interpreter/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.BET, Literal: "bet"},
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
	if program.String() != "bet myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
