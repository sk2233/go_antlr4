package main

import (
	"ANTLR_test/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// 输入内容
	input := antlr.NewInputStream("3*(10+2)+24/3")
	// 创建词法分析器
	lexer := parser.NewExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 创建语法分析器
	p := parser.NewExprParser(stream)
	tree := p.Expr()
	// 创建 visitor 进行遍历输出结果
	v := NewExprVisitor()
	fmt.Println(v.Visit(tree))
}
