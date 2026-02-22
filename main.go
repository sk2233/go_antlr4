package main

import (
	"ANTLR_test/ast"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	RunCode("return 12*(1+2)+12/2")
	RunCode("var name=2233;" +
		"return name;")
	RunCode("var name=22;" +
		"if(name==22){" +
		"name=11;" +
		"}" +
		"return name;")
	RunCode("var name=22;" +
		"for(var i=0;i==0;i=i+1;){" +
		"name=name+1;" +
		"}" +
		"return name;")
	RunCode("var add=func(num1,num2){\n" +
		"return num1+num2;\n" +
		"}\n" +
		"return add(1+2,2+3);\n")
}

func RunCode(code string) {
	// 输入内容
	input := antlr.NewInputStream(code)
	// 创建词法分析器
	lexer := ast.NewExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// 创建语法分析器
	parser := ast.NewExprParser(stream)
	tree := parser.Program()
	// 创建 visitor 进行遍历输出结果
	v := NewExprVisitor()
	fmt.Println(v.Visit(tree))
}
