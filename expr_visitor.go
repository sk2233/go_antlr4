package main

import (
	"ANTLR_test/ast"
	"github.com/antlr4-go/antlr/v4"
)

type ExprVisitor struct {
	ast.ExprVisitor // 持有默认实现，定制化的覆盖默认实现即可
}

func NewExprVisitor() *ExprVisitor {
	return &ExprVisitor{}
}

// Visit 必须覆盖该方法，否者默认传递的是嵌入对象 下面的 VisitXxx 不会被调用
func (v *ExprVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ExprVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitVarStatement(ctx *ast.VarStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitBlockStatement(ctx *ast.BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitIfStatement(ctx *ast.IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitForStatement(ctx *ast.ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitExpresstionStatement(ctx *ast.ExpresstionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitEq(ctx *ast.EqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitAddInEq(ctx *ast.AddInEqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitAdd(ctx *ast.AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitMulInAdd(ctx *ast.MulInAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitPriInMul(ctx *ast.PriInMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitMul(ctx *ast.MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitCallExpr(ctx *ast.CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitUnaryExpr(ctx *ast.UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitNumber(ctx *ast.NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitIdent(ctx *ast.IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitGro(ctx *ast.GroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitFunInUnary(ctx *ast.FunInUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitGroup(ctx *ast.GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitFunction(ctx *ast.FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitParams(ctx *ast.ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
