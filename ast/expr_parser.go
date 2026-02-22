// Code generated from /Users/wepie/Documents/github/go_antlr4/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast // Expr
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ExprParser struct {
	*antlr.BaseParser
}

var ExprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exprParserInit() {
	staticData := &ExprParserStaticData
	staticData.LiteralNames = []string{
		"", "'var'", "'='", "';'", "'return'", "'{'", "'}'", "'if'", "'('",
		"')'", "'else'", "'for'", "'break'", "'continue'", "','", "'function'",
		"'+'", "'-'", "'=='", "'*'", "'/'", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "FUNCTION",
		"PLUS", "MINUS", "EQ", "MULTIPLY", "DIVIDE", "NEQ", "NUMBER", "IDENTIFIER",
		"EMPTY", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "varStatement", "returnStatement", "blockStatement",
		"ifStatement", "forStatement", "breakStatement", "continueStatement",
		"expresstionStatement", "expression", "additionExpression", "multiplicationExpression",
		"primary", "unary", "params", "arguments",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 178, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 4, 0, 36, 8, 0, 11, 0, 12, 0, 37, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 62, 8, 4, 10, 4, 12,
		4, 65, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		76, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 101, 8, 10, 10, 10, 12, 10, 104, 9, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 5, 11, 112, 8, 11, 10, 11, 12, 11, 115, 9, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 123, 8, 12, 10, 12, 12,
		12, 126, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 136, 8, 13, 10, 13, 12, 13, 139, 9, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 153,
		8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 158, 8, 15, 10, 15, 12, 15, 161, 9,
		15, 3, 15, 163, 8, 15, 1, 16, 1, 16, 1, 16, 5, 16, 168, 8, 16, 10, 16,
		12, 16, 171, 9, 16, 1, 16, 3, 16, 174, 8, 16, 3, 16, 176, 8, 16, 1, 16,
		0, 4, 20, 22, 24, 26, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 0, 3, 2, 0, 18, 18, 21, 21, 1, 0, 16, 17, 1, 0, 19, 20,
		182, 0, 35, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 49, 1, 0, 0, 0, 6, 55, 1,
		0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 68, 1, 0, 0, 0, 12, 77, 1, 0, 0, 0, 14,
		85, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 91, 1, 0, 0, 0, 20, 94, 1, 0, 0,
		0, 22, 105, 1, 0, 0, 0, 24, 116, 1, 0, 0, 0, 26, 127, 1, 0, 0, 0, 28, 152,
		1, 0, 0, 0, 30, 162, 1, 0, 0, 0, 32, 175, 1, 0, 0, 0, 34, 36, 3, 2, 1,
		0, 35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38,
		1, 0, 0, 0, 38, 1, 1, 0, 0, 0, 39, 48, 3, 4, 2, 0, 40, 48, 3, 6, 3, 0,
		41, 48, 3, 8, 4, 0, 42, 48, 3, 10, 5, 0, 43, 48, 3, 12, 6, 0, 44, 48, 3,
		14, 7, 0, 45, 48, 3, 16, 8, 0, 46, 48, 3, 18, 9, 0, 47, 39, 1, 0, 0, 0,
		47, 40, 1, 0, 0, 0, 47, 41, 1, 0, 0, 0, 47, 42, 1, 0, 0, 0, 47, 43, 1,
		0, 0, 0, 47, 44, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48,
		3, 1, 0, 0, 0, 49, 50, 5, 1, 0, 0, 50, 51, 5, 23, 0, 0, 51, 52, 5, 2, 0,
		0, 52, 53, 3, 20, 10, 0, 53, 54, 5, 3, 0, 0, 54, 5, 1, 0, 0, 0, 55, 56,
		5, 4, 0, 0, 56, 57, 3, 20, 10, 0, 57, 58, 5, 3, 0, 0, 58, 7, 1, 0, 0, 0,
		59, 63, 5, 5, 0, 0, 60, 62, 3, 2, 1, 0, 61, 60, 1, 0, 0, 0, 62, 65, 1,
		0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65,
		63, 1, 0, 0, 0, 66, 67, 5, 6, 0, 0, 67, 9, 1, 0, 0, 0, 68, 69, 5, 7, 0,
		0, 69, 70, 5, 8, 0, 0, 70, 71, 3, 20, 10, 0, 71, 72, 5, 9, 0, 0, 72, 75,
		3, 8, 4, 0, 73, 74, 5, 10, 0, 0, 74, 76, 3, 8, 4, 0, 75, 73, 1, 0, 0, 0,
		75, 76, 1, 0, 0, 0, 76, 11, 1, 0, 0, 0, 77, 78, 5, 11, 0, 0, 78, 79, 5,
		8, 0, 0, 79, 80, 3, 2, 1, 0, 80, 81, 3, 2, 1, 0, 81, 82, 3, 20, 10, 0,
		82, 83, 5, 9, 0, 0, 83, 84, 3, 8, 4, 0, 84, 13, 1, 0, 0, 0, 85, 86, 5,
		12, 0, 0, 86, 87, 5, 3, 0, 0, 87, 15, 1, 0, 0, 0, 88, 89, 5, 13, 0, 0,
		89, 90, 5, 3, 0, 0, 90, 17, 1, 0, 0, 0, 91, 92, 3, 20, 10, 0, 92, 93, 5,
		3, 0, 0, 93, 19, 1, 0, 0, 0, 94, 95, 6, 10, -1, 0, 95, 96, 3, 22, 11, 0,
		96, 102, 1, 0, 0, 0, 97, 98, 10, 1, 0, 0, 98, 99, 7, 0, 0, 0, 99, 101,
		3, 22, 11, 0, 100, 97, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0,
		0, 0, 102, 103, 1, 0, 0, 0, 103, 21, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0,
		105, 106, 6, 11, -1, 0, 106, 107, 3, 24, 12, 0, 107, 113, 1, 0, 0, 0, 108,
		109, 10, 1, 0, 0, 109, 110, 7, 1, 0, 0, 110, 112, 3, 24, 12, 0, 111, 108,
		1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0,
		0, 0, 114, 23, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 6, 12, -1, 0,
		117, 118, 3, 26, 13, 0, 118, 124, 1, 0, 0, 0, 119, 120, 10, 1, 0, 0, 120,
		121, 7, 2, 0, 0, 121, 123, 3, 26, 13, 0, 122, 119, 1, 0, 0, 0, 123, 126,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 25, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 6, 13, -1, 0, 128, 129, 3, 28, 14,
		0, 129, 137, 1, 0, 0, 0, 130, 131, 10, 1, 0, 0, 131, 132, 5, 8, 0, 0, 132,
		133, 3, 32, 16, 0, 133, 134, 5, 9, 0, 0, 134, 136, 1, 0, 0, 0, 135, 130,
		1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0,
		0, 0, 138, 27, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 153, 5, 22, 0, 0,
		141, 153, 5, 23, 0, 0, 142, 143, 5, 8, 0, 0, 143, 144, 3, 20, 10, 0, 144,
		145, 5, 9, 0, 0, 145, 153, 1, 0, 0, 0, 146, 147, 5, 15, 0, 0, 147, 148,
		5, 8, 0, 0, 148, 149, 3, 30, 15, 0, 149, 150, 5, 9, 0, 0, 150, 151, 3,
		8, 4, 0, 151, 153, 1, 0, 0, 0, 152, 140, 1, 0, 0, 0, 152, 141, 1, 0, 0,
		0, 152, 142, 1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 153, 29, 1, 0, 0, 0, 154,
		159, 5, 23, 0, 0, 155, 156, 5, 14, 0, 0, 156, 158, 5, 23, 0, 0, 157, 155,
		1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0,
		0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 154, 1, 0, 0, 0,
		162, 163, 1, 0, 0, 0, 163, 31, 1, 0, 0, 0, 164, 169, 3, 20, 10, 0, 165,
		166, 5, 14, 0, 0, 166, 168, 3, 20, 10, 0, 167, 165, 1, 0, 0, 0, 168, 171,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 173, 1, 0,
		0, 0, 171, 169, 1, 0, 0, 0, 172, 174, 5, 14, 0, 0, 173, 172, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 164, 1, 0, 0, 0, 175,
		176, 1, 0, 0, 0, 176, 33, 1, 0, 0, 0, 14, 37, 47, 63, 75, 102, 113, 124,
		137, 152, 159, 162, 169, 173, 175,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ExprParserInit initializes any static state used to implement ExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprParserInit() {
	staticData := &ExprParserStaticData
	staticData.once.Do(exprParserInit)
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	ExprParserInit()
	this := new(ExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ExprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

// ExprParser tokens.
const (
	ExprParserEOF           = antlr.TokenEOF
	ExprParserT__0          = 1
	ExprParserT__1          = 2
	ExprParserT__2          = 3
	ExprParserT__3          = 4
	ExprParserT__4          = 5
	ExprParserT__5          = 6
	ExprParserT__6          = 7
	ExprParserT__7          = 8
	ExprParserT__8          = 9
	ExprParserT__9          = 10
	ExprParserT__10         = 11
	ExprParserT__11         = 12
	ExprParserT__12         = 13
	ExprParserT__13         = 14
	ExprParserFUNCTION      = 15
	ExprParserPLUS          = 16
	ExprParserMINUS         = 17
	ExprParserEQ            = 18
	ExprParserMULTIPLY      = 19
	ExprParserDIVIDE        = 20
	ExprParserNEQ           = 21
	ExprParserNUMBER        = 22
	ExprParserIDENTIFIER    = 23
	ExprParserEMPTY         = 24
	ExprParserLINE_COMMENT  = 25
	ExprParserBLOCK_COMMENT = 26
)

// ExprParser rules.
const (
	ExprParserRULE_program                  = 0
	ExprParserRULE_statement                = 1
	ExprParserRULE_varStatement             = 2
	ExprParserRULE_returnStatement          = 3
	ExprParserRULE_blockStatement           = 4
	ExprParserRULE_ifStatement              = 5
	ExprParserRULE_forStatement             = 6
	ExprParserRULE_breakStatement           = 7
	ExprParserRULE_continueStatement        = 8
	ExprParserRULE_expresstionStatement     = 9
	ExprParserRULE_expression               = 10
	ExprParserRULE_additionExpression       = 11
	ExprParserRULE_multiplicationExpression = 12
	ExprParserRULE_primary                  = 13
	ExprParserRULE_unary                    = 14
	ExprParserRULE_params                   = 15
	ExprParserRULE_arguments                = 16
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12630450) != 0) {
		{
			p.SetState(34)
			p.Statement()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarStatement() IVarStatementContext
	ReturnStatement() IReturnStatementContext
	BlockStatement() IBlockStatementContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext
	ExpresstionStatement() IExpresstionStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarStatement() IVarStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ExpresstionStatement() IExpresstionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresstionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresstionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_statement)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.VarStatement()
		}

	case ExprParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.ReturnStatement()
		}

	case ExprParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(41)
			p.BlockStatement()
		}

	case ExprParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(42)
			p.IfStatement()
		}

	case ExprParserT__10:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)
			p.ForStatement()
		}

	case ExprParserT__11:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(44)
			p.BreakStatement()
		}

	case ExprParserT__12:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(45)
			p.ContinueStatement()
		}

	case ExprParserT__7, ExprParserFUNCTION, ExprParserNUMBER, ExprParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(46)
			p.ExpresstionStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarStatementContext is an interface to support dynamic dispatch.
type IVarStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVarStatementContext differentiates from other interfaces.
	IsVarStatementContext()
}

type VarStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarStatementContext() *VarStatementContext {
	var p = new(VarStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_varStatement
	return p
}

func InitEmptyVarStatementContext(p *VarStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_varStatement
}

func (*VarStatementContext) IsVarStatementContext() {}

func NewVarStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarStatementContext {
	var p = new(VarStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_varStatement

	return p
}

func (s *VarStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VarStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *VarStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitVarStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) VarStatement() (localctx IVarStatementContext) {
	localctx = NewVarStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExprParserRULE_varStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(ExprParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(ExprParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(ExprParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.expression(0)
	}
	{
		p.SetState(53)
		p.Match(ExprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExprParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(ExprParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.expression(0)
	}
	{
		p.SetState(57)
		p.Match(ExprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_blockStatement
	return p
}

func InitEmptyBlockStatementContext(p *BlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_blockStatement
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExprParserRULE_blockStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(ExprParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12630450) != 0 {
		{
			p.SetState(60)
			p.Statement()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(ExprParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionContext

	// GetIfBody returns the ifBody rule contexts.
	GetIfBody() IBlockStatementContext

	// GetElseBody returns the elseBody rule contexts.
	GetElseBody() IBlockStatementContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionContext)

	// SetIfBody sets the ifBody rule contexts.
	SetIfBody(IBlockStatementContext)

	// SetElseBody sets the elseBody rule contexts.
	SetElseBody(IBlockStatementContext)

	// Getter signatures
	Expression() IExpressionContext
	AllBlockStatement() []IBlockStatementContext
	BlockStatement(i int) IBlockStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	cond     IExpressionContext
	ifBody   IBlockStatementContext
	elseBody IBlockStatementContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) GetCond() IExpressionContext { return s.cond }

func (s *IfStatementContext) GetIfBody() IBlockStatementContext { return s.ifBody }

func (s *IfStatementContext) GetElseBody() IBlockStatementContext { return s.elseBody }

func (s *IfStatementContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *IfStatementContext) SetIfBody(v IBlockStatementContext) { s.ifBody = v }

func (s *IfStatementContext) SetElseBody(v IBlockStatementContext) { s.elseBody = v }

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllBlockStatement() []IBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockStatementContext); ok {
			tst[i] = t.(IBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) BlockStatement(i int) IBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExprParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(ExprParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(ExprParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)

		var _x = p.expression(0)

		localctx.(*IfStatementContext).cond = _x
	}
	{
		p.SetState(71)
		p.Match(ExprParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)

		var _x = p.BlockStatement()

		localctx.(*IfStatementContext).ifBody = _x
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserT__9 {
		{
			p.SetState(73)
			p.Match(ExprParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)

			var _x = p.BlockStatement()

			localctx.(*IfStatementContext).elseBody = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IStatementContext

	// GetCond returns the cond rule contexts.
	GetCond() IStatementContext

	// GetStep returns the step rule contexts.
	GetStep() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IBlockStatementContext

	// SetInit sets the init rule contexts.
	SetInit(IStatementContext)

	// SetCond sets the cond rule contexts.
	SetCond(IStatementContext)

	// SetStep sets the step rule contexts.
	SetStep(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IBlockStatementContext)

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	Expression() IExpressionContext
	BlockStatement() IBlockStatementContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   IStatementContext
	cond   IStatementContext
	step   IExpressionContext
	body   IBlockStatementContext
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) GetInit() IStatementContext { return s.init }

func (s *ForStatementContext) GetCond() IStatementContext { return s.cond }

func (s *ForStatementContext) GetStep() IExpressionContext { return s.step }

func (s *ForStatementContext) GetBody() IBlockStatementContext { return s.body }

func (s *ForStatementContext) SetInit(v IStatementContext) { s.init = v }

func (s *ForStatementContext) SetCond(v IStatementContext) { s.cond = v }

func (s *ForStatementContext) SetStep(v IExpressionContext) { s.step = v }

func (s *ForStatementContext) SetBody(v IBlockStatementContext) { s.body = v }

func (s *ForStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExprParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(ExprParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(ExprParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)

		var _x = p.Statement()

		localctx.(*ForStatementContext).init = _x
	}
	{
		p.SetState(80)

		var _x = p.Statement()

		localctx.(*ForStatementContext).cond = _x
	}
	{
		p.SetState(81)

		var _x = p.expression(0)

		localctx.(*ForStatementContext).step = _x
	}
	{
		p.SetState(82)
		p.Match(ExprParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)

		var _x = p.BlockStatement()

		localctx.(*ForStatementContext).body = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExprParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(ExprParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(ExprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ExprParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.Match(ExprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresstionStatementContext is an interface to support dynamic dispatch.
type IExpresstionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpresstionStatementContext differentiates from other interfaces.
	IsExpresstionStatementContext()
}

type ExpresstionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresstionStatementContext() *ExpresstionStatementContext {
	var p = new(ExpresstionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expresstionStatement
	return p
}

func InitEmptyExpresstionStatementContext(p *ExpresstionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expresstionStatement
}

func (*ExpresstionStatementContext) IsExpresstionStatementContext() {}

func NewExpresstionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresstionStatementContext {
	var p = new(ExpresstionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expresstionStatement

	return p
}

func (s *ExpresstionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresstionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpresstionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresstionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresstionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExpresstionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ExpresstionStatement() (localctx IExpresstionStatementContext) {
	localctx = NewExpresstionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExprParserRULE_expresstionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.expression(0)
	}
	{
		p.SetState(92)
		p.Match(ExprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqContext struct {
	ExpressionContext
	left  IExpressionContext
	op    antlr.Token
	right IAdditionExpressionContext
}

func NewEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqContext {
	var p = new(EqContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqContext) GetOp() antlr.Token { return s.op }

func (s *EqContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqContext) GetLeft() IExpressionContext { return s.left }

func (s *EqContext) GetRight() IAdditionExpressionContext { return s.right }

func (s *EqContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *EqContext) SetRight(v IAdditionExpressionContext) { s.right = v }

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqContext) AdditionExpression() IAdditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *EqContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExprParserEQ, 0)
}

func (s *EqContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ExprParserNEQ, 0)
}

func (s *EqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitEq(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddInEqContext struct {
	ExpressionContext
}

func NewAddInEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddInEqContext {
	var p = new(AddInEqContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddInEqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddInEqContext) AdditionExpression() IAdditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *AddInEqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAddInEq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ExprParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ExprParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAddInEqContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(95)
		p.additionExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewEqContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*EqContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_expression)
			p.SetState(97)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(98)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*EqContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ExprParserEQ || _la == ExprParserNEQ) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*EqContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(99)

				var _x = p.additionExpression(0)

				localctx.(*EqContext).right = _x
			}

		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditionExpressionContext is an interface to support dynamic dispatch.
type IAdditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAdditionExpressionContext differentiates from other interfaces.
	IsAdditionExpressionContext()
}

type AdditionExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionExpressionContext() *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_additionExpression
	return p
}

func InitEmptyAdditionExpressionContext(p *AdditionExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_additionExpression
}

func (*AdditionExpressionContext) IsAdditionExpressionContext() {}

func NewAdditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_additionExpression

	return p
}

func (s *AdditionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionExpressionContext) CopyAll(ctx *AdditionExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AddContext struct {
	AdditionExpressionContext
	left  IAdditionExpressionContext
	op    antlr.Token
	right IMultiplicationExpressionContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	InitEmptyAdditionExpressionContext(&p.AdditionExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditionExpressionContext))

	return p
}

func (s *AddContext) GetOp() antlr.Token { return s.op }

func (s *AddContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddContext) GetLeft() IAdditionExpressionContext { return s.left }

func (s *AddContext) GetRight() IMultiplicationExpressionContext { return s.right }

func (s *AddContext) SetLeft(v IAdditionExpressionContext) { s.left = v }

func (s *AddContext) SetRight(v IMultiplicationExpressionContext) { s.right = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AdditionExpression() IAdditionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditionExpressionContext)
}

func (s *AddContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *AddContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ExprParserPLUS, 0)
}

func (s *AddContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ExprParserMINUS, 0)
}

func (s *AddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulInAddContext struct {
	AdditionExpressionContext
}

func NewMulInAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulInAddContext {
	var p = new(MulInAddContext)

	InitEmptyAdditionExpressionContext(&p.AdditionExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditionExpressionContext))

	return p
}

func (s *MulInAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulInAddContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *MulInAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitMulInAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) AdditionExpression() (localctx IAdditionExpressionContext) {
	return p.additionExpression(0)
}

func (p *ExprParser) additionExpression(_p int) (localctx IAdditionExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAdditionExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAdditionExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, ExprParserRULE_additionExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewMulInAddContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(106)
		p.multiplicationExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAddContext(p, NewAdditionExpressionContext(p, _parentctx, _parentState))
			localctx.(*AddContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_additionExpression)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(109)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*AddContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ExprParserPLUS || _la == ExprParserMINUS) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*AddContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(110)

				var _x = p.multiplicationExpression(0)

				localctx.(*AddContext).right = _x
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicationExpressionContext is an interface to support dynamic dispatch.
type IMultiplicationExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultiplicationExpressionContext differentiates from other interfaces.
	IsMultiplicationExpressionContext()
}

type MultiplicationExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationExpressionContext() *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_multiplicationExpression
	return p
}

func InitEmptyMultiplicationExpressionContext(p *MultiplicationExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_multiplicationExpression
}

func (*MultiplicationExpressionContext) IsMultiplicationExpressionContext() {}

func NewMultiplicationExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationExpressionContext {
	var p = new(MultiplicationExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_multiplicationExpression

	return p
}

func (s *MultiplicationExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationExpressionContext) CopyAll(ctx *MultiplicationExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MultiplicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PriInMulContext struct {
	MultiplicationExpressionContext
}

func NewPriInMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PriInMulContext {
	var p = new(PriInMulContext)

	InitEmptyMultiplicationExpressionContext(&p.MultiplicationExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicationExpressionContext))

	return p
}

func (s *PriInMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriInMulContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PriInMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitPriInMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulContext struct {
	MultiplicationExpressionContext
	left  IMultiplicationExpressionContext
	op    antlr.Token
	right IPrimaryContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	InitEmptyMultiplicationExpressionContext(&p.MultiplicationExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicationExpressionContext))

	return p
}

func (s *MulContext) GetOp() antlr.Token { return s.op }

func (s *MulContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulContext) GetLeft() IMultiplicationExpressionContext { return s.left }

func (s *MulContext) GetRight() IPrimaryContext { return s.right }

func (s *MulContext) SetLeft(v IMultiplicationExpressionContext) { s.left = v }

func (s *MulContext) SetRight(v IPrimaryContext) { s.right = v }

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) MultiplicationExpression() IMultiplicationExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicationExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExpressionContext)
}

func (s *MulContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *MulContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(ExprParserMULTIPLY, 0)
}

func (s *MulContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ExprParserDIVIDE, 0)
}

func (s *MulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitMul(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) MultiplicationExpression() (localctx IMultiplicationExpressionContext) {
	return p.multiplicationExpression(0)
}

func (p *ExprParser) multiplicationExpression(_p int) (localctx IMultiplicationExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMultiplicationExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultiplicationExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ExprParserRULE_multiplicationExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPriInMulContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(117)
		p.primary(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMulContext(p, NewMultiplicationExpressionContext(p, _parentctx, _parentState))
			localctx.(*MulContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_multiplicationExpression)
			p.SetState(119)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(120)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MulContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ExprParserMULTIPLY || _la == ExprParserDIVIDE) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MulContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(121)

				var _x = p.primary(0)

				localctx.(*MulContext).right = _x
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallExprContext struct {
	PrimaryContext
	left IPrimaryContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *CallExprContext) GetLeft() IPrimaryContext { return s.left }

func (s *CallExprContext) SetLeft(v IPrimaryContext) { s.left = v }

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CallExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	PrimaryContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Primary() (localctx IPrimaryContext) {
	return p.primary(0)
}

func (p *ExprParser) primary(_p int) (localctx IPrimaryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ExprParserRULE_primary, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewUnaryExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(128)
		p.Unary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCallExprContext(p, NewPrimaryContext(p, _parentctx, _parentState))
			localctx.(*CallExprContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_primary)
			p.SetState(130)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(131)
				p.Match(ExprParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(132)
				p.Arguments()
			}
			{
				p.SetState(133)
				p.Match(ExprParserT__8)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) CopyAll(ctx *UnaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupContext struct {
	UnaryContext
}

func NewGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupContext {
	var p = new(GroupContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunInUnaryContext struct {
	UnaryContext
}

func NewFunInUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunInUnaryContext {
	var p = new(FunInUnaryContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *FunInUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunInUnaryContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ExprParserFUNCTION, 0)
}

func (s *FunInUnaryContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunInUnaryContext) BlockStatement() IBlockStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *FunInUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitFunInUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentContext struct {
	UnaryContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, 0)
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	UnaryContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyUnaryContext(&p.UnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ExprParserNUMBER, 0)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExprParserRULE_unary)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExprParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(ExprParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExprParserIDENTIFIER:
		localctx = NewIdentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(ExprParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExprParserT__7:
		localctx = NewGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)
			p.Match(ExprParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.expression(0)
		}
		{
			p.SetState(144)
			p.Match(ExprParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ExprParserFUNCTION:
		localctx = NewFunInUnaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
			p.Match(ExprParserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ExprParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Params()
		}
		{
			p.SetState(149)
			p.Match(ExprParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.BlockStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ExprParserIDENTIFIER)
}

func (s *ParamsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserIDENTIFIER, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserIDENTIFIER {
		{
			p.SetState(154)
			p.Match(ExprParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ExprParserT__13 {
			{
				p.SetState(155)
				p.Match(ExprParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(156)
				p.Match(ExprParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExprParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ExprParserRULE_arguments)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12615936) != 0 {
		{
			p.SetState(164)
			p.expression(0)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(165)
					p.Match(ExprParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(166)
					p.expression(0)
				}

			}
			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserT__13 {
			{
				p.SetState(172)
				p.Match(ExprParserT__13)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 11:
		var t *AdditionExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AdditionExpressionContext)
		}
		return p.AdditionExpression_Sempred(t, predIndex)

	case 12:
		var t *MultiplicationExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicationExpressionContext)
		}
		return p.MultiplicationExpression_Sempred(t, predIndex)

	case 13:
		var t *PrimaryContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryContext)
		}
		return p.Primary_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) AdditionExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) MultiplicationExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ExprParser) Primary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
