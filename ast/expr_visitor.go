// Code generated from /Users/sky/Documents/go/ANTLR_test/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // Expr
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by ExprParser#Paren.
	VisitParen(ctx *ParenContext) interface{}
}
