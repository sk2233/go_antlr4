package main

import (
	"ANTLR_test/ast"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	RunCode("return 12*(1+2)+12/2;")
	RunCode("var name=2233;\n" +
		"return name;")
	RunCode("var name=22;\n" +
		"if(name==22){\n" +
		"name=11;\n" +
		"}\n" +
		"return name;")
	RunCode("var name=22;\n" +
		"for(var i=0;i==0;i=i+1;){\n" +
		"name=name+1;\n" +
		"}\n" +
		"return name;")
	RunCode("var add=func(num1,num2){\n" +
		"return num1+num2;\n" +
		"};\n" +
		"return add(1+2,2+3);")
}

func RunCode(code string) {
	fmt.Println(strings.Split(code, "\n")[0], "...")
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
