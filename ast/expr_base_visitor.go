// Code generated from /Users/wepie/Documents/github/go_antlr4/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // Expr
import "github.com/antlr4-go/antlr/v4"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitVarStatement(ctx *VarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExpresstionStatement(ctx *ExpresstionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitEq(ctx *EqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAddInEq(ctx *AddInEqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitMulInAdd(ctx *MulInAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitPriInMul(ctx *PriInMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitFunInUnary(ctx *FunInUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
