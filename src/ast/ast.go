package ast

import "reboot/token"

type Node interface {
	TokenLiteral() string
}

// Statement represents a single unit of
// execution that does not produce a value.
// Implementers of this interface define
// specific types of statements.
type Statement interface {
	Node
	statementNode()
}

// Expression represents a single unit of
// computation that produces a value.
// Implementers of this interface define
// specific types of expressions.
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
