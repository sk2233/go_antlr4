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

// toFloat 将值转换为 float64
func (v *ExprVisitor) toFloat(val interface{}) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int:
		return float64(x), true
	case int64:
		return float64(x), true
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
	if f, ok := v.toFloat(val); ok {
		return f != 0
	}
	return true
}

func (v *ExprVisitor) VisitProgram(ctx *ast.ProgramContext) interface{} {
	statements := ctx.AllStatement()
	var lastResult interface{}
	for _, stmt := range statements {
		lastResult = stmt.Accept(v)
		if v.returnValue != nil {
			return v.returnValue
		}
	}
	return lastResult
}

func (v *ExprVisitor) VisitStatement(ctx *ast.StatementContext) interface{} {
	for _, tmp := range ctx.GetChildren() {
		switch tar := tmp.(type) {
		case *ast.VarStatementContext:
			return v.VisitVarStatement(tar)
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

func (v *ExprVisitor) VisitReturnStatement(ctx *ast.ReturnStatementContext) interface{} {
	v.returnValue = v.Visit(ctx.Expression())
	return v.returnValue
}

func (v *ExprVisitor) VisitBlockStatement(ctx *ast.BlockStatementContext) interface{} {
	statements := ctx.AllStatement()
	var lastResult interface{}
	for _, stmt := range statements {
		lastResult = stmt.Accept(v)
		if v.returnValue != nil {
			return lastResult
		}
		if v.breakFlag || v.continueFlag {
			return lastResult
		}
	}
	return lastResult
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
	op := ctx.GetOp().GetText()
	leftF, leftOk := v.toFloat(left)
	rightF, rightOk := v.toFloat(right)
	if leftOk && rightOk {
		switch op {
		case "==":
			return leftF == rightF
		case "!=":
			return leftF != rightF
		}
	}
	// 其他类型比较
	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	}
	return false
}

func (v *ExprVisitor) VisitAddInEq(ctx *ast.AddInEqContext) interface{} {
	addExpr := ctx.AdditionExpression()
	if addExpr != nil {
		return v.Visit(addExpr)
	}
	return v.VisitChildren(ctx)
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
	leftF, leftOk := v.toFloat(left)
	rightF, rightOk := v.toFloat(right)
	if !leftOk || !rightOk {
		panic(fmt.Sprintf("invalid operands for +/-: %v, %v", left, right))
	}
	switch ctx.GetOp().GetText() {
	case "+":
		return leftF + rightF
	case "-":
		return leftF - rightF
	}
	return 0
}

func (v *ExprVisitor) VisitMulInAdd(ctx *ast.MulInAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitPriInMul(ctx *ast.PriInMulContext) interface{} {
	return v.VisitChildren(ctx)
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
	leftF, leftOk := v.toFloat(left)
	rightF, rightOk := v.toFloat(right)
	if !leftOk || !rightOk {
		panic(fmt.Sprintf("invalid operands for *//: %v, %v", left, right))
	}
	switch ctx.GetOp().GetText() {
	case "*":
		return leftF * rightF
	case "/":
		if rightF == 0 {
			panic("division by zero")
		}
		return leftF / rightF
	}
	return 0
}

// 函数调用
func (v *ExprVisitor) VisitCallExpr(ctx *ast.CallExprContext) interface{} {
	// left 是函数（Ident 解析为变量，变量值可能是 functionDef）
	callable := v.Visit(ctx.GetLeft())
	if fn, ok := callable.(*functionDef); ok {
		return v.callFunction(fn, ctx.Arguments())
	}
	panic(fmt.Sprintf("not a function: %v", callable))
}

func (v *ExprVisitor) callFunction(fn *functionDef, argsCtx ast.IArgumentsContext) interface{} {
	oldScope := v.scope
	v.scope = make(map[string]interface{})
	defer func() { v.scope = oldScope }()

	oldReturn := v.returnValue
	v.returnValue = nil
	defer func() { v.returnValue = oldReturn }()

	argValues := []interface{}{}
	if argsCtx != nil {
		for _, expr := range argsCtx.AllExpression() {
			argValues = append(argValues, v.Visit(expr))
		}
	}
	for i, name := range fn.params {
		if i < len(argValues) {
			v.scope[name] = argValues[i]
		}
	}
	return v.Visit(fn.body)
}

func (v *ExprVisitor) VisitUnaryExpr(ctx *ast.UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitNumber(ctx *ast.NumberContext) interface{} {
	text := ctx.NUMBER().GetText()
	if f, err := strconv.ParseFloat(text, 64); err == nil {
		return f
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
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitGroup(ctx *ast.GroupContext) interface{} {
	// group: '(' expression ')'，返回表达式结果
	if expr := ctx.Expression(); expr != nil {
		return v.Visit(expr)
	}
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitParams(ctx *ast.ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ExprVisitor) VisitArguments(ctx *ast.ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
