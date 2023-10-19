// Code generated from Expr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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
		"", "", "", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "WHITESPACE", "NUMBER", "MUL", "DIV", "ADD", "SUB", "PAROPEN", "PARCLOSE",
	}
	staticData.RuleNames = []string{
		"WHITESPACE", "NUMBER", "MUL", "DIV", "ADD", "SUB", "PAROPEN", "PARCLOSE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 50, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 19, 8, 0, 11,
		0, 12, 0, 20, 1, 0, 1, 0, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 1, 1,
		3, 1, 31, 8, 1, 1, 1, 5, 1, 34, 8, 1, 10, 1, 12, 1, 37, 9, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 0, 0, 8,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 2, 3, 0, 9, 10,
		13, 13, 32, 32, 1, 0, 48, 57, 53, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 18, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0,
		5, 38, 1, 0, 0, 0, 7, 40, 1, 0, 0, 0, 9, 42, 1, 0, 0, 0, 11, 44, 1, 0,
		0, 0, 13, 46, 1, 0, 0, 0, 15, 48, 1, 0, 0, 0, 17, 19, 7, 0, 0, 0, 18, 17,
		1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0,
		21, 22, 1, 0, 0, 0, 22, 23, 6, 0, 0, 0, 23, 2, 1, 0, 0, 0, 24, 26, 7, 1,
		0, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28,
		1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 31, 9, 0, 0, 0, 30, 29, 1, 0, 0, 0,
		30, 31, 1, 0, 0, 0, 31, 35, 1, 0, 0, 0, 32, 34, 7, 1, 0, 0, 33, 32, 1,
		0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36,
		4, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 42, 0, 0, 39, 6, 1, 0, 0,
		0, 40, 41, 5, 47, 0, 0, 41, 8, 1, 0, 0, 0, 42, 43, 5, 43, 0, 0, 43, 10,
		1, 0, 0, 0, 44, 45, 5, 45, 0, 0, 45, 12, 1, 0, 0, 0, 46, 47, 5, 40, 0,
		0, 47, 14, 1, 0, 0, 0, 48, 49, 5, 41, 0, 0, 49, 16, 1, 0, 0, 0, 5, 0, 20,
		27, 30, 35, 1, 6, 0, 0,
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
	ExprLexerWHITESPACE = 1
	ExprLexerNUMBER     = 2
	ExprLexerMUL        = 3
	ExprLexerDIV        = 4
	ExprLexerADD        = 5
	ExprLexerSUB        = 6
	ExprLexerPAROPEN    = 7
	ExprLexerPARCLOSE   = 8
)
