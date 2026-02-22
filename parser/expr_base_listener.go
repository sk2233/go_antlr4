// Code generated from /Users/sky/Documents/go/ANTLR_test/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr4-go/antlr/v4"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseExprListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseExprListener) ExitInt(ctx *IntContext) {}

// EnterParen is called when production Paren is entered.
func (s *BaseExprListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BaseExprListener) ExitParen(ctx *ParenContext) {}
