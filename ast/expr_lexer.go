// Code generated from /Users/sky/Documents/go/ANTLR_test/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package ast

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ExprLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exprlexerLexerInit() {
	staticData := &ExprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "", "", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "INT", "WS", "MUL", "DIV", "ADD", "SUB",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "INT", "WS", "MUL", "DIV", "ADD", "SUB",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 41, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 4, 2, 23, 8, 2, 11, 2, 12, 2, 24, 1, 3, 4, 3, 28, 8, 3, 11, 3, 12, 3,
		29, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 0, 0, 8,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 2, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 42, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 19, 1, 0, 0, 0,
		5, 22, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 33, 1, 0, 0, 0, 11, 35, 1, 0,
		0, 0, 13, 37, 1, 0, 0, 0, 15, 39, 1, 0, 0, 0, 17, 18, 5, 40, 0, 0, 18,
		2, 1, 0, 0, 0, 19, 20, 5, 41, 0, 0, 20, 4, 1, 0, 0, 0, 21, 23, 7, 0, 0,
		0, 22, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25,
		1, 0, 0, 0, 25, 6, 1, 0, 0, 0, 26, 28, 7, 1, 0, 0, 27, 26, 1, 0, 0, 0,
		28, 29, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 1,
		0, 0, 0, 31, 32, 6, 3, 0, 0, 32, 8, 1, 0, 0, 0, 33, 34, 5, 42, 0, 0, 34,
		10, 1, 0, 0, 0, 35, 36, 5, 47, 0, 0, 36, 12, 1, 0, 0, 0, 37, 38, 5, 43,
		0, 0, 38, 14, 1, 0, 0, 0, 39, 40, 5, 45, 0, 0, 40, 16, 1, 0, 0, 0, 3, 0,
		24, 29, 1, 6, 0, 0,
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

// ExprLexerInit initializes any static state used to implement ExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprLexerInit() {
	staticData := &ExprLexerLexerStaticData
	staticData.once.Do(exprlexerLexerInit)
}

// NewExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExprLexer(input antlr.CharStream) *ExprLexer {
	ExprLexerInit()
	l := new(ExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ExprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0 = 1
	ExprLexerT__1 = 2
	ExprLexerINT  = 3
	ExprLexerWS   = 4
	ExprLexerMUL  = 5
	ExprLexerDIV  = 6
	ExprLexerADD  = 7
	ExprLexerSUB  = 8
)
