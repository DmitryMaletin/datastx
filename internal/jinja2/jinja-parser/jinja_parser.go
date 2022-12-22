// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jinja

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type JinjaParser struct {
	*antlr.BaseParser
}

var jinjaParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func jinjaParserInit() {
  staticData := &jinjaParserStaticData
  staticData.literalNames = []string{
    "", "'{%'", "'for'", "'in'", "'%}'", "'endfor'", "'if'", "'elif'", "'else'", 
    "'endif'", "'{{'", "'}}'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "ID",
  }
  staticData.ruleNames = []string{
    "template", "block", "for_block", "if_block", "variable", "text",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 12, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 1, 1, 1, 1, 1, 
	1, 1, 3, 1, 22, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 31, 
	8, 2, 11, 2, 12, 2, 32, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 
	1, 3, 4, 3, 44, 8, 3, 11, 3, 12, 3, 45, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 
	3, 53, 8, 3, 11, 3, 12, 3, 54, 5, 3, 57, 8, 3, 10, 3, 12, 3, 60, 9, 3, 
	1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 66, 8, 3, 11, 3, 12, 3, 67, 3, 3, 70, 8, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 82, 
	8, 5, 10, 5, 12, 5, 85, 9, 5, 1, 5, 1, 83, 0, 6, 0, 2, 4, 6, 8, 10, 0, 
	1, 2, 0, 1, 1, 10, 10, 91, 0, 13, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 23, 
	1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 79, 1, 0, 0, 0, 12, 
	14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 13, 1, 0, 0, 
	0, 15, 16, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0, 17, 22, 3, 4, 2, 0, 18, 22, 3, 
	6, 3, 0, 19, 22, 3, 8, 4, 0, 20, 22, 3, 10, 5, 0, 21, 17, 1, 0, 0, 0, 21, 
	18, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 20, 1, 0, 0, 0, 22, 3, 1, 0, 0, 
	0, 23, 24, 5, 1, 0, 0, 24, 25, 5, 2, 0, 0, 25, 26, 5, 12, 0, 0, 26, 27, 
	5, 3, 0, 0, 27, 28, 5, 12, 0, 0, 28, 30, 5, 4, 0, 0, 29, 31, 3, 2, 1, 0, 
	30, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 
	0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 5, 1, 0, 0, 35, 36, 5, 5, 0, 0, 36, 
	37, 5, 4, 0, 0, 37, 5, 1, 0, 0, 0, 38, 39, 5, 1, 0, 0, 39, 40, 5, 6, 0, 
	0, 40, 41, 5, 12, 0, 0, 41, 43, 5, 4, 0, 0, 42, 44, 3, 2, 1, 0, 43, 42, 
	1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 
	46, 58, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5, 7, 0, 0, 49, 50, 5, 
	12, 0, 0, 50, 52, 5, 4, 0, 0, 51, 53, 3, 2, 1, 0, 52, 51, 1, 0, 0, 0, 53, 
	54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 
	0, 56, 47, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 
	1, 0, 0, 0, 59, 69, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 5, 1, 0, 0, 
	62, 63, 5, 8, 0, 0, 63, 65, 5, 4, 0, 0, 64, 66, 3, 2, 1, 0, 65, 64, 1, 
	0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 
	70, 1, 0, 0, 0, 69, 61, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 
	0, 71, 72, 5, 1, 0, 0, 72, 73, 5, 9, 0, 0, 73, 74, 5, 4, 0, 0, 74, 7, 1, 
	0, 0, 0, 75, 76, 5, 10, 0, 0, 76, 77, 5, 12, 0, 0, 77, 78, 5, 11, 0, 0, 
	78, 9, 1, 0, 0, 0, 79, 83, 8, 0, 0, 0, 80, 82, 9, 0, 0, 0, 81, 80, 1, 0, 
	0, 0, 82, 85, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 84, 11, 
	1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 9, 15, 21, 32, 45, 54, 58, 67, 69, 83,
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

// JinjaParserInit initializes any static state used to implement JinjaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJinjaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JinjaParserInit() {
  staticData := &jinjaParserStaticData
  staticData.once.Do(jinjaParserInit)
}

// NewJinjaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJinjaParser(input antlr.TokenStream) *JinjaParser {
	JinjaParserInit()
	this := new(JinjaParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &jinjaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Jinja.g4"

	return this
}


// JinjaParser tokens.
const (
	JinjaParserEOF = antlr.TokenEOF
	JinjaParserT__0 = 1
	JinjaParserT__1 = 2
	JinjaParserT__2 = 3
	JinjaParserT__3 = 4
	JinjaParserT__4 = 5
	JinjaParserT__5 = 6
	JinjaParserT__6 = 7
	JinjaParserT__7 = 8
	JinjaParserT__8 = 9
	JinjaParserT__9 = 10
	JinjaParserT__10 = 11
	JinjaParserID = 12
)

// JinjaParser rules.
const (
	JinjaParserRULE_template = 0
	JinjaParserRULE_block = 1
	JinjaParserRULE_for_block = 2
	JinjaParserRULE_if_block = 3
	JinjaParserRULE_variable = 4
	JinjaParserRULE_text = 5
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_template
	return p
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *TemplateContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) Template() (localctx ITemplateContext) {
	this := p
	_ = this

	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JinjaParserRULE_template)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << JinjaParserT__0) | (1 << JinjaParserT__1) | (1 << JinjaParserT__2) | (1 << JinjaParserT__3) | (1 << JinjaParserT__4) | (1 << JinjaParserT__5) | (1 << JinjaParserT__6) | (1 << JinjaParserT__7) | (1 << JinjaParserT__8) | (1 << JinjaParserT__9) | (1 << JinjaParserT__10) | (1 << JinjaParserID))) != 0) {
		{
			p.SetState(12)
			p.Block()
		}


		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) For_block() IFor_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_blockContext)
}

func (s *BlockContext) If_block() IIf_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_blockContext)
}

func (s *BlockContext) Variable() IVariableContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *BlockContext) Text() ITextContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JinjaParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.For_block()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.If_block()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)
			p.Variable()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(20)
			p.Text()
		}

	}


	return localctx
}


// IFor_blockContext is an interface to support dynamic dispatch.
type IFor_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_blockContext differentiates from other interfaces.
	IsFor_blockContext()
}

type For_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_blockContext() *For_blockContext {
	var p = new(For_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_for_block
	return p
}

func (*For_blockContext) IsFor_blockContext() {}

func NewFor_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_blockContext {
	var p = new(For_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_for_block

	return p
}

func (s *For_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *For_blockContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JinjaParserID)
}

func (s *For_blockContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JinjaParserID, i)
}

func (s *For_blockContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *For_blockContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *For_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterFor_block(s)
	}
}

func (s *For_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitFor_block(s)
	}
}

func (s *For_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitFor_block(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) For_block() (localctx IFor_blockContext) {
	this := p
	_ = this

	localctx = NewFor_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JinjaParserRULE_for_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Match(JinjaParserT__0)
	}
	{
		p.SetState(24)
		p.Match(JinjaParserT__1)
	}
	{
		p.SetState(25)
		p.Match(JinjaParserID)
	}
	{
		p.SetState(26)
		p.Match(JinjaParserT__2)
	}
	{
		p.SetState(27)
		p.Match(JinjaParserID)
	}
	{
		p.SetState(28)
		p.Match(JinjaParserT__3)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(29)
					p.Block()
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	{
		p.SetState(34)
		p.Match(JinjaParserT__0)
	}
	{
		p.SetState(35)
		p.Match(JinjaParserT__4)
	}
	{
		p.SetState(36)
		p.Match(JinjaParserT__3)
	}



	return localctx
}


// IIf_blockContext is an interface to support dynamic dispatch.
type IIf_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_blockContext differentiates from other interfaces.
	IsIf_blockContext()
}

type If_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_blockContext() *If_blockContext {
	var p = new(If_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_if_block
	return p
}

func (*If_blockContext) IsIf_blockContext() {}

func NewIf_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_blockContext {
	var p = new(If_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_if_block

	return p
}

func (s *If_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *If_blockContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JinjaParserID)
}

func (s *If_blockContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JinjaParserID, i)
}

func (s *If_blockContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_blockContext) Block(i int) IBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *If_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *If_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterIf_block(s)
	}
}

func (s *If_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitIf_block(s)
	}
}

func (s *If_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitIf_block(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) If_block() (localctx IIf_blockContext) {
	this := p
	_ = this

	localctx = NewIf_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JinjaParserRULE_if_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(JinjaParserT__0)
	}
	{
		p.SetState(39)
		p.Match(JinjaParserT__5)
	}
	{
		p.SetState(40)
		p.Match(JinjaParserID)
	}
	{
		p.SetState(41)
		p.Match(JinjaParserT__3)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(42)
					p.Block()
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(47)
				p.Match(JinjaParserT__0)
			}
			{
				p.SetState(48)
				p.Match(JinjaParserT__6)
			}
			{
				p.SetState(49)
				p.Match(JinjaParserID)
			}
			{
				p.SetState(50)
				p.Match(JinjaParserT__3)
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
						{
							p.SetState(51)
							p.Block()
						}




				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(54)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
			}


		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(61)
			p.Match(JinjaParserT__0)
		}
		{
			p.SetState(62)
			p.Match(JinjaParserT__7)
		}
		{
			p.SetState(63)
			p.Match(JinjaParserT__3)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(64)
						p.Block()
					}




			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}


	}
	{
		p.SetState(71)
		p.Match(JinjaParserT__0)
	}
	{
		p.SetState(72)
		p.Match(JinjaParserT__8)
	}
	{
		p.SetState(73)
		p.Match(JinjaParserT__3)
	}



	return localctx
}


// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(JinjaParserID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JinjaParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(JinjaParserT__9)
	}
	{
		p.SetState(76)
		p.Match(JinjaParserID)
	}
	{
		p.SetState(77)
		p.Match(JinjaParserT__10)
	}



	return localctx
}


// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JinjaParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }
func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *JinjaParser) Text() (localctx ITextContext) {
	this := p
	_ = this

	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JinjaParserRULE_text)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == JinjaParserT__0 || _la == JinjaParserT__9  {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(80)
			p.MatchWildcard()



		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}



	return localctx
}


