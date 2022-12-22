// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type JinjaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var jinjalexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func jinjalexerLexerInit() {
  staticData := &jinjalexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.literalNames = []string{
    "", "'{%'", "'for'", "'in'", "'%}'", "'endfor'", "'if'", "'elif'", "'else'", 
    "'endif'", "'{{'", "'}}'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "ID",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "ID",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 12, 77, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 
	1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 
	1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 
	1, 10, 1, 11, 1, 11, 5, 11, 73, 8, 11, 10, 11, 12, 11, 76, 9, 11, 0, 0, 
	12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 
	11, 23, 12, 1, 0, 2, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 
	97, 122, 77, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 
	1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 
	15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 
	0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 28, 1, 0, 0, 0, 5, 32, 1, 0, 0, 
	0, 7, 35, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 45, 1, 0, 0, 0, 13, 48, 1, 
	0, 0, 0, 15, 53, 1, 0, 0, 0, 17, 58, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 
	67, 1, 0, 0, 0, 23, 70, 1, 0, 0, 0, 25, 26, 5, 123, 0, 0, 26, 27, 5, 37, 
	0, 0, 27, 2, 1, 0, 0, 0, 28, 29, 5, 102, 0, 0, 29, 30, 5, 111, 0, 0, 30, 
	31, 5, 114, 0, 0, 31, 4, 1, 0, 0, 0, 32, 33, 5, 105, 0, 0, 33, 34, 5, 110, 
	0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 37, 0, 0, 36, 37, 5, 125, 0, 0, 37, 
	8, 1, 0, 0, 0, 38, 39, 5, 101, 0, 0, 39, 40, 5, 110, 0, 0, 40, 41, 5, 100, 
	0, 0, 41, 42, 5, 102, 0, 0, 42, 43, 5, 111, 0, 0, 43, 44, 5, 114, 0, 0, 
	44, 10, 1, 0, 0, 0, 45, 46, 5, 105, 0, 0, 46, 47, 5, 102, 0, 0, 47, 12, 
	1, 0, 0, 0, 48, 49, 5, 101, 0, 0, 49, 50, 5, 108, 0, 0, 50, 51, 5, 105, 
	0, 0, 51, 52, 5, 102, 0, 0, 52, 14, 1, 0, 0, 0, 53, 54, 5, 101, 0, 0, 54, 
	55, 5, 108, 0, 0, 55, 56, 5, 115, 0, 0, 56, 57, 5, 101, 0, 0, 57, 16, 1, 
	0, 0, 0, 58, 59, 5, 101, 0, 0, 59, 60, 5, 110, 0, 0, 60, 61, 5, 100, 0, 
	0, 61, 62, 5, 105, 0, 0, 62, 63, 5, 102, 0, 0, 63, 18, 1, 0, 0, 0, 64, 
	65, 5, 123, 0, 0, 65, 66, 5, 123, 0, 0, 66, 20, 1, 0, 0, 0, 67, 68, 5, 
	125, 0, 0, 68, 69, 5, 125, 0, 0, 69, 22, 1, 0, 0, 0, 70, 74, 7, 0, 0, 0, 
	71, 73, 7, 1, 0, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 
	0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 24, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 2, 
	0, 74, 0,
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

// JinjaLexerInit initializes any static state used to implement JinjaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJinjaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JinjaLexerInit() {
  staticData := &jinjalexerLexerStaticData
  staticData.once.Do(jinjalexerLexerInit)
}

// NewJinjaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJinjaLexer(input antlr.CharStream) *JinjaLexer {
  JinjaLexerInit()
	l := new(JinjaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &jinjalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Jinja.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JinjaLexer tokens.
const (
	JinjaLexerT__0 = 1
	JinjaLexerT__1 = 2
	JinjaLexerT__2 = 3
	JinjaLexerT__3 = 4
	JinjaLexerT__4 = 5
	JinjaLexerT__5 = 6
	JinjaLexerT__6 = 7
	JinjaLexerT__7 = 8
	JinjaLexerT__8 = 9
	JinjaLexerT__9 = 10
	JinjaLexerT__10 = 11
	JinjaLexerID = 12
)

