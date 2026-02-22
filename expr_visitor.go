package main

import (
	"ANTLR_test/ast"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// functionDef 表示函数定义
type functionDef struct {
	params []string
	body   antlr.ParseTree
}

type ExprVisitor struct {
	ast.BaseExprVisitor
	// 符号表：变量名 -> 值（含函数）
	scope map[string]interface{}
	// 返回值（用于 return 语句）
	returnValue interface{}
	// 控制流标志
	breakFlag    bool
	continueFlag bool
}

func NewExprVisitor() *ExprVisitor {
	return &ExprVisitor{
		scope: make(map[string]interface{}),
	}
}

// Visit 必须覆盖该方法，否则默认传递的是嵌入对象，下面的 VisitXxx 不会被调用
func (v *ExprVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// toInt 将值转换为 int
func (v *ExprVisitor) toInt(val interface{}) (int, bool) {
	switch x := val.(type) {
	case int:
		return x, true
	case int64:
		return int(x), true
	default:
		return 0, false
	}
}

// toBool 将值转换为 bool（用于 if/for 条件）
func (v *ExprVisitor) toBool(val interface{}) bool {
	if val == nil {
		return false
	}
	if b, ok := val.(bool); ok {
		return b
	}
	if i, ok := v.toInt(val); ok {
		return i != 0
	}
	return false
}

func (v *ExprVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	statements := ctx.AllStatement()
	for _, stmt := range statements {
		stmt.Accept(v)
		if v.returnValue != nil {
			return v.returnValue
		}
	}
	return nil
}

func (v *ExprVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	for _, tmp := range ctx.GetChildren() {
		switch tar := tmp.(type) {
		case *ast.VarStatementContext:
			return v.VisitVarStatement(tar)
		case *ast.AssignStatementContext:
			return v.VisitAssignStatement(tar)
		case *ast.ReturnStatementContext:
			return v.VisitReturnStatement(tar)
		case *ast.BlockStatementContext:
			return v.VisitBlockStatement(tar)
		case *ast.IfStatementContext:
			return v.VisitIfStatement(tar)
		case *ast.ForStatementContext:
			return v.VisitForStatement(tar)
		case *ast.BreakStatementContext:
			return v.VisitBreakStatement(tar)
		case *ast.ContinueStatementContext:
			return v.VisitContinueStatement(tar)
		case *ast.ExpresstionStatementContext:
			return v.VisitExpresstionStatement(tar)
		}
	}
	panic("invalid statement")
}

func (v *ExprVisitor) VisitVarStatement(ctx *ast.VarStatementContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	expr := ctx.Expression()
	result := v.Visit(expr)
	v.scope[name] = result
	return result
}

func (v *ExprVisitor) VisitAssignStatement(ctx *ast.AssignStatementContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	expr := ctx.Expression()
	result := v.Visit(expr)
	v.scope[name] = result
	return result
}

func (v *ExprVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	v.returnValue = v.Visit(ctx.Expression())
	return v.returnValue
}

func (v *ExprVisitor) VisitBlockStatement(ctx *ast.BlockStatementContext) interface{} {
	statements := ctx.AllStatement()
	for _, stmt := range statements {
		stmt.Accept(v)
		if v.returnValue != nil {
			return v.returnValue
		}
		if v.breakFlag || v.continueFlag {
			return nil
		}
	}
	return nil
}

func (v *ExprVisitor) VisitIfStatement(ctx *ast.IfStatementContext) interface{} {
	condVal := v.Visit(ctx.GetCond())
	if v.toBool(condVal) {
		return ctx.GetIfBody().Accept(v)
	}
	if elseBody := ctx.GetElseBody(); elseBody != nil {
		return elseBody.Accept(v)
	}
	return nil
}

func (v *ExprVisitor) VisitForStatement(ctx *ast.ForStatementContext) interface{} {
	// 执行 init
	if init := ctx.GetInit(); init != nil {
		init.Accept(v)
	}
	for {
		if v.breakFlag {
			v.breakFlag = false
			break
		}
		// 检查 cond（cond 是 statement，可能是 expression; 或 ;）
		if cond := ctx.GetCond(); cond != nil {
			if exprStmt := cond.ExpresstionStatement(); exprStmt != nil {
				condVal := v.Visit(exprStmt.Expression())
				if !v.toBool(condVal) {
					break
				}
			}
		}
		// 执行 body
		ctx.GetBody().Accept(v)
		if v.returnValue != nil {
			return v.returnValue
		}
		if v.breakFlag {
			v.breakFlag = false
			break
		}
		if v.continueFlag {
			v.continueFlag = false
		}
		// 执行 step
		if step := ctx.GetStep(); step != nil {
			step.Accept(v)
		}
	}
	return nil
}

func (v *ExprVisitor) VisitBreakStatement(ctx *ast.BreakStatementContext) interface{} {
	v.breakFlag = true
	return nil
}

func (v *ExprVisitor) VisitContinueStatement(ctx *ast.ContinueStatementContext) interface{} {
	v.continueFlag = true
	return nil
}

func (v *ExprVisitor) VisitExpresstionStatement(ctx *ast.ExpresstionStatementContext) interface{} {
	return ctx.Expression().Accept(v)
}

// 表达式求值：== 或 !=
func (v *ExprVisitor) VisitEq(ctx *ast.EqContext) interface{} {
	leftExpr := ctx.GetLeft()
	rightExpr := ctx.GetRight()
	if leftExpr == nil || rightExpr == nil {
		return v.VisitChildren(ctx)
	}
	left := v.Visit(leftExpr)
	right := v.Visit(rightExpr)
	return v.binaryOp(left, right, ctx.GetOp().GetTokenType())
}

func (v *ExprVisitor) binaryOp(left interface{}, right interface{}, tk int) interface{} {
	switch tk {
	case ast.ExprLexerEQ:
		return left == right
	case ast.ExprLexerNEQ:
		return left != right
	case ast.ExprLexerPLUS:
		return left.(int) + right.(int)
	case ast.ExprLexerMINUS:
		return left.(int) - right.(int)
	case ast.ExprLexerMULTIPLY:
		return left.(int) * right.(int)
	case ast.ExprLexerDIVIDE:
		return left.(int) / right.(int)
	}
	return nil
}

func (v *ExprVisitor) VisitAddInEq(ctx *ast.AddInEqContext) interface{} {
	return v.Visit(ctx.AdditionExpression())
}

// 表达式求值：+ 或 -
func (v *ExprVisitor) VisitAdd(ctx *ast.AddContext) interface{} {
	// 递归规则中 left/right 通过 GetLeft/GetRight 获取，不一定在 GetChildren 中
	leftExpr := ctx.GetLeft()
	rightExpr := ctx.GetRight()
	if leftExpr == nil || rightExpr == nil {
		return v.VisitChildren(ctx)
	}
	left := v.Visit(leftExpr)
	right := v.Visit(rightExpr)
	return v.binaryOp(left, right, ctx.GetOp().GetTokenType())
}

func (v *ExprVisitor) VisitMulInAdd(ctx *ast.MulInAddContext) interface{} {
	return v.Visit(ctx.MultiplicationExpression())
}

func (v *ExprVisitor) VisitPriInMul(ctx *ast.PriInMulContext) interface{} {
	return v.Visit(ctx.Primary())
}

// 表达式求值：* 或 /
func (v *ExprVisitor) VisitMul(ctx *ast.MulContext) interface{} {
	leftExpr := ctx.GetLeft()
	rightExpr := ctx.GetRight()
	if leftExpr == nil || rightExpr == nil {
		return v.VisitChildren(ctx)
	}
	left := v.Visit(leftExpr)
	right := v.Visit(rightExpr)
	return v.binaryOp(left, right, ctx.GetOp().GetTokenType())
}

// 函数调用
func (v *ExprVisitor) VisitCallExpr(ctx *ast.CallExprContext) interface{} {
	// left 是函数（Ident 解析为变量，变量值可能是 functionDef）
	callable := v.Visit(ctx.GetLeft())
	if fn, ok := callable.(*functionDef); ok {
		args := ctx.Arguments().Accept(v).([]any)
		return v.callFunction(fn, args)
	}
	panic(fmt.Sprintf("not a function: %v", callable))
}

func (v *ExprVisitor) callFunction(fn *functionDef, args []any) interface{} {
	oldScope := v.scope
	v.scope = make(map[string]interface{})
	defer func() { v.scope = oldScope }()

	oldReturn := v.returnValue
	v.returnValue = nil
	defer func() { v.returnValue = oldReturn }()

	for i, name := range fn.params {
		v.scope[name] = args[i]
	}
	return v.Visit(fn.body)
}

func (v *ExprVisitor) VisitUnaryExpr(ctx *ast.UnaryExprContext) interface{} {
	return v.Visit(ctx.Unary())
}

func (v *ExprVisitor) VisitNumber(ctx *ast.NumberContext) interface{} {
	text := ctx.NUMBER().GetText()
	if i, err := strconv.Atoi(text); err == nil {
		return i
	}
	panic(fmt.Sprintf("invalid number: %s", text))
}

func (v *ExprVisitor) VisitIdent(ctx *ast.IdentContext) interface{} {
	name := ctx.IDENTIFIER().GetText()
	if val, ok := v.scope[name]; ok {
		return val
	}
	panic(fmt.Sprintf("undefined variable: %s", name))
}

func (v *ExprVisitor) VisitFunInUnary(ctx *ast.FunInUnaryContext) interface{} {
	return &functionDef{
		params: ctx.Params().Accept(v).([]string),
		body:   ctx.BlockStatement(),
	}
}

func (v *ExprVisitor) VisitGroup(ctx *ast.GroupContext) interface{} {
	// group: '(' expression ')'，返回表达式结果
	return v.Visit(ctx.Expression())
}

func (v *ExprVisitor) VisitParams(ctx *ast.ParamsContext) interface{} {
	res := make([]string, 0)
	for _, tmp := range ctx.AllIDENTIFIER() {
		res = append(res, tmp.GetText())
	}
	return res
}

func (v *ExprVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	res := make([]any, 0)
	for _, tmp := range ctx.AllExpression() {
		res = append(res, tmp.Accept(v))
	}
	return res
}
