// Code generated from /Users/wepie/Documents/github/go_antlr4/Expr.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "'var'", "';'", "'return'", "'{'", "'}'", "'if'", "'('", "')'",
		"'else'", "'for'", "'break'", "'continue'", "','", "'func'", "'+'",
		"'-'", "'=='", "'='", "'*'", "'/'", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "FUNCTION",
		"PLUS", "MINUS", "EQ", "ASSIGN", "MULTIPLY", "DIVIDE", "NEQ", "NUMBER",
		"IDENTIFIER", "EMPTY", "LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "FUNCTION", "PLUS", "MINUS", "EQ",
		"ASSIGN", "MULTIPLY", "DIVIDE", "NEQ", "NUMBER", "IDENTIFIER", "EMPTY",
		"LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 176, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 4, 21, 126, 8, 21, 11, 21, 12, 21, 127,
		1, 21, 1, 21, 4, 21, 132, 8, 21, 11, 21, 12, 21, 133, 3, 21, 136, 8, 21,
		1, 22, 1, 22, 5, 22, 140, 8, 22, 10, 22, 12, 22, 143, 9, 22, 1, 23, 4,
		23, 146, 8, 23, 11, 23, 12, 23, 147, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 5, 24, 156, 8, 24, 10, 24, 12, 24, 159, 9, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 5, 25, 167, 8, 25, 10, 25, 12, 25, 170, 9, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 168, 0, 26, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32,
		32, 2, 0, 10, 10, 13, 13, 182, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 1, 53, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 66,
		1, 0, 0, 0, 9, 68, 1, 0, 0, 0, 11, 70, 1, 0, 0, 0, 13, 73, 1, 0, 0, 0,
		15, 75, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 86, 1,
		0, 0, 0, 23, 92, 1, 0, 0, 0, 25, 101, 1, 0, 0, 0, 27, 103, 1, 0, 0, 0,
		29, 108, 1, 0, 0, 0, 31, 110, 1, 0, 0, 0, 33, 112, 1, 0, 0, 0, 35, 115,
		1, 0, 0, 0, 37, 117, 1, 0, 0, 0, 39, 119, 1, 0, 0, 0, 41, 121, 1, 0, 0,
		0, 43, 125, 1, 0, 0, 0, 45, 137, 1, 0, 0, 0, 47, 145, 1, 0, 0, 0, 49, 151,
		1, 0, 0, 0, 51, 162, 1, 0, 0, 0, 53, 54, 5, 118, 0, 0, 54, 55, 5, 97, 0,
		0, 55, 56, 5, 114, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 59, 0, 0, 58, 4,
		1, 0, 0, 0, 59, 60, 5, 114, 0, 0, 60, 61, 5, 101, 0, 0, 61, 62, 5, 116,
		0, 0, 62, 63, 5, 117, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 110, 0, 0,
		65, 6, 1, 0, 0, 0, 66, 67, 5, 123, 0, 0, 67, 8, 1, 0, 0, 0, 68, 69, 5,
		125, 0, 0, 69, 10, 1, 0, 0, 0, 70, 71, 5, 105, 0, 0, 71, 72, 5, 102, 0,
		0, 72, 12, 1, 0, 0, 0, 73, 74, 5, 40, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76,
		5, 41, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 5, 101, 0, 0, 78, 79, 5, 108,
		0, 0, 79, 80, 5, 115, 0, 0, 80, 81, 5, 101, 0, 0, 81, 18, 1, 0, 0, 0, 82,
		83, 5, 102, 0, 0, 83, 84, 5, 111, 0, 0, 84, 85, 5, 114, 0, 0, 85, 20, 1,
		0, 0, 0, 86, 87, 5, 98, 0, 0, 87, 88, 5, 114, 0, 0, 88, 89, 5, 101, 0,
		0, 89, 90, 5, 97, 0, 0, 90, 91, 5, 107, 0, 0, 91, 22, 1, 0, 0, 0, 92, 93,
		5, 99, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 110, 0, 0, 95, 96, 5, 116,
		0, 0, 96, 97, 5, 105, 0, 0, 97, 98, 5, 110, 0, 0, 98, 99, 5, 117, 0, 0,
		99, 100, 5, 101, 0, 0, 100, 24, 1, 0, 0, 0, 101, 102, 5, 44, 0, 0, 102,
		26, 1, 0, 0, 0, 103, 104, 5, 102, 0, 0, 104, 105, 5, 117, 0, 0, 105, 106,
		5, 110, 0, 0, 106, 107, 5, 99, 0, 0, 107, 28, 1, 0, 0, 0, 108, 109, 5,
		43, 0, 0, 109, 30, 1, 0, 0, 0, 110, 111, 5, 45, 0, 0, 111, 32, 1, 0, 0,
		0, 112, 113, 5, 61, 0, 0, 113, 114, 5, 61, 0, 0, 114, 34, 1, 0, 0, 0, 115,
		116, 5, 61, 0, 0, 116, 36, 1, 0, 0, 0, 117, 118, 5, 42, 0, 0, 118, 38,
		1, 0, 0, 0, 119, 120, 5, 47, 0, 0, 120, 40, 1, 0, 0, 0, 121, 122, 5, 33,
		0, 0, 122, 123, 5, 61, 0, 0, 123, 42, 1, 0, 0, 0, 124, 126, 7, 0, 0, 0,
		125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 135, 1, 0, 0, 0, 129, 131, 5, 46, 0, 0, 130, 132,
		7, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 129, 1, 0, 0, 0,
		135, 136, 1, 0, 0, 0, 136, 44, 1, 0, 0, 0, 137, 141, 7, 1, 0, 0, 138, 140,
		7, 2, 0, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 141, 142, 1, 0, 0, 0, 142, 46, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0,
		144, 146, 7, 3, 0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 150,
		6, 23, 0, 0, 150, 48, 1, 0, 0, 0, 151, 152, 5, 47, 0, 0, 152, 153, 5, 47,
		0, 0, 153, 157, 1, 0, 0, 0, 154, 156, 8, 4, 0, 0, 155, 154, 1, 0, 0, 0,
		156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 6, 24, 0, 0, 161, 50,
		1, 0, 0, 0, 162, 163, 5, 47, 0, 0, 163, 164, 5, 42, 0, 0, 164, 168, 1,
		0, 0, 0, 165, 167, 9, 0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 170, 1, 0, 0,
		0, 168, 169, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170,
		168, 1, 0, 0, 0, 171, 172, 5, 42, 0, 0, 172, 173, 5, 47, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 175, 6, 25, 0, 0, 175, 52, 1, 0, 0, 0, 8, 0, 127, 133,
		135, 141, 147, 157, 168, 1, 6, 0, 0,
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
	ExprLexerT__0          = 1
	ExprLexerT__1          = 2
	ExprLexerT__2          = 3
	ExprLexerT__3          = 4
	ExprLexerT__4          = 5
	ExprLexerT__5          = 6
	ExprLexerT__6          = 7
	ExprLexerT__7          = 8
	ExprLexerT__8          = 9
	ExprLexerT__9          = 10
	ExprLexerT__10         = 11
	ExprLexerT__11         = 12
	ExprLexerT__12         = 13
	ExprLexerFUNCTION      = 14
	ExprLexerPLUS          = 15
	ExprLexerMINUS         = 16
	ExprLexerEQ            = 17
	ExprLexerASSIGN        = 18
	ExprLexerMULTIPLY      = 19
	ExprLexerDIVIDE        = 20
	ExprLexerNEQ           = 21
	ExprLexerNUMBER        = 22
	ExprLexerIDENTIFIER    = 23
	ExprLexerEMPTY         = 24
	ExprLexerLINE_COMMENT  = 25
	ExprLexerBLOCK_COMMENT = 26
)
