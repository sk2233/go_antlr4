package main

import (
	"ANTLR_test/ast"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type ExprVisitor struct {
	ast.BaseExprVisitor // 持有默认实现，定制化的覆盖默认实现即可
}

func NewExprVisitor() *ExprVisitor {
	return &ExprVisitor{}
}

// Visit 必须覆盖该方法，否者默认传递的是嵌入对象 下面的 VisitXxx 不会被调用
func (v *ExprVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ExprVisitor) visitBinaryOp(tmp1 ast.IExprContext, tmp2 ast.IExprContext, op antlr.Token) interface{} {
	val1 := tmp1.Accept(v).(*Var)
	val2 := tmp2.Accept(v).(*Var)
	switch op.GetTokenType() {
	case ast.ExprLexerMUL:
		return NewIntVar(val1.Int * val2.Int)
	case ast.ExprLexerDIV:
		return NewIntVar(val1.Int / val2.Int)
	case ast.ExprLexerADD:
		return NewIntVar(val1.Int + val2.Int)
	case ast.ExprLexerSUB:
		return NewIntVar(val1.Int - val2.Int)
	default:
		panic(fmt.Sprintf("unknow op %s", op.GetText()))
	}
}

func (v *ExprVisitor) VisitMulDiv(ctx *ast.MulDivContext) interface{} {
	return v.visitBinaryOp(ctx.Expr(0), ctx.Expr(1), ctx.GetOp())
}

func (v *ExprVisitor) VisitAddSub(ctx *ast.AddSubContext) interface{} {
	return v.visitBinaryOp(ctx.Expr(0), ctx.Expr(1), ctx.GetOp())
}

func (v *ExprVisitor) VisitInt(ctx *ast.IntContext) interface{} {
	val, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	return NewIntVar(int(val))
}

func (v *ExprVisitor) VisitParen(ctx *ast.ParenContext) interface{} {
	return ctx.Expr().Accept(v)
}
