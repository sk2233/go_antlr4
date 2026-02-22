// Code generated from /Users/sky/Documents/go/ANTLR_test/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // Expr
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by ExprParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ExprParser#varStatement.
	VisitVarStatement(ctx *VarStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#expresstionStatement.
	VisitExpresstionStatement(ctx *ExpresstionStatementContext) interface{}

	// Visit a parse tree produced by ExprParser#Eq.
	VisitEq(ctx *EqContext) interface{}

	// Visit a parse tree produced by ExprParser#AddInEq.
	VisitAddInEq(ctx *AddInEqContext) interface{}

	// Visit a parse tree produced by ExprParser#Add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by ExprParser#MulInAdd.
	VisitMulInAdd(ctx *MulInAddContext) interface{}

	// Visit a parse tree produced by ExprParser#PriInMul.
	VisitPriInMul(ctx *PriInMulContext) interface{}

	// Visit a parse tree produced by ExprParser#Mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by ExprParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by ExprParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by ExprParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ExprParser#Ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by ExprParser#Gro.
	VisitGro(ctx *GroContext) interface{}

	// Visit a parse tree produced by ExprParser#FunInUnary.
	VisitFunInUnary(ctx *FunInUnaryContext) interface{}

	// Visit a parse tree produced by ExprParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by ExprParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by ExprParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by ExprParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}
}
