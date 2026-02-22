// Code generated from /Users/sky/Documents/go/ANTLR_test/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr4-go/antlr/v4"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterParen is called when entering the Paren production.
	EnterParen(c *ParenContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitParen is called when exiting the Paren production.
	ExitParen(c *ParenContext)
}
