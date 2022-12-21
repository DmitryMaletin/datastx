// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // LookML

import (
	"fmt"
	"strconv"
  "sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type LookMLParser struct {
	*antlr.BaseParser
}

var lookmlParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func lookmlParserInit() {
  staticData := &lookmlParserStaticData
  staticData.literalNames = []string{
    "", "'model:'", "'{'", "'}'", "'view:'", "'dimension:'", "'measure:'", 
    "'field:'", "'type:'", "'sql:'", "'description:'", "'group_label:'", 
    "'string'", "'number'", "'boolean'", "'date'", "'datetime'", "'time'", 
    "'duration'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "ID", "STRING", "WS",
  }
  staticData.ruleNames = []string{
    "model", "view", "dimension", "measure", "field", "field_property", 
    "field_type",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 21, 82, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 4, 0, 19, 8, 0, 11, 
	0, 12, 0, 20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 29, 8, 1, 11, 1, 
	12, 1, 30, 1, 1, 4, 1, 34, 8, 1, 11, 1, 12, 1, 35, 1, 1, 1, 1, 1, 2, 1, 
	2, 1, 2, 1, 2, 4, 2, 44, 8, 2, 11, 2, 12, 2, 45, 1, 2, 1, 2, 1, 3, 1, 3, 
	1, 3, 1, 3, 4, 3, 54, 8, 3, 11, 3, 12, 3, 55, 1, 3, 1, 3, 1, 4, 1, 4, 1, 
	4, 1, 4, 4, 4, 64, 8, 4, 11, 4, 12, 4, 65, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 
	1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 78, 8, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 
	0, 2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 12, 18, 83, 0, 14, 1, 0, 0, 0, 2, 24, 
	1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 59, 1, 0, 0, 0, 10, 
	77, 1, 0, 0, 0, 12, 79, 1, 0, 0, 0, 14, 15, 5, 1, 0, 0, 15, 16, 5, 19, 
	0, 0, 16, 18, 5, 2, 0, 0, 17, 19, 3, 2, 1, 0, 18, 17, 1, 0, 0, 0, 19, 20, 
	1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 
	22, 23, 5, 3, 0, 0, 23, 1, 1, 0, 0, 0, 24, 25, 5, 4, 0, 0, 25, 26, 5, 19, 
	0, 0, 26, 28, 5, 2, 0, 0, 27, 29, 3, 4, 2, 0, 28, 27, 1, 0, 0, 0, 29, 30, 
	1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 33, 1, 0, 0, 0, 
	32, 34, 3, 6, 3, 0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1, 
	0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5, 3, 0, 0, 38, 
	3, 1, 0, 0, 0, 39, 40, 5, 5, 0, 0, 40, 41, 5, 19, 0, 0, 41, 43, 5, 2, 0, 
	0, 42, 44, 3, 8, 4, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 43, 
	1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 5, 3, 0, 0, 
	48, 5, 1, 0, 0, 0, 49, 50, 5, 6, 0, 0, 50, 51, 5, 19, 0, 0, 51, 53, 5, 
	2, 0, 0, 52, 54, 3, 8, 4, 0, 53, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 
	53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 5, 3, 0, 
	0, 58, 7, 1, 0, 0, 0, 59, 60, 5, 7, 0, 0, 60, 61, 5, 19, 0, 0, 61, 63, 
	5, 2, 0, 0, 62, 64, 3, 10, 5, 0, 63, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 
	65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 
	3, 0, 0, 68, 9, 1, 0, 0, 0, 69, 70, 5, 8, 0, 0, 70, 78, 3, 12, 6, 0, 71, 
	72, 5, 9, 0, 0, 72, 78, 5, 20, 0, 0, 73, 74, 5, 10, 0, 0, 74, 78, 5, 20, 
	0, 0, 75, 76, 5, 11, 0, 0, 76, 78, 5, 20, 0, 0, 77, 69, 1, 0, 0, 0, 77, 
	71, 1, 0, 0, 0, 77, 73, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 11, 1, 0, 0, 
	0, 79, 80, 7, 0, 0, 0, 80, 13, 1, 0, 0, 0, 7, 20, 30, 35, 45, 55, 65, 77,
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

// LookMLParserInit initializes any static state used to implement LookMLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLookMLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LookMLParserInit() {
  staticData := &lookmlParserStaticData
  staticData.once.Do(lookmlParserInit)
}

// NewLookMLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLookMLParser(input antlr.TokenStream) *LookMLParser {
	LookMLParserInit()
	this := new(LookMLParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &lookmlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}


// LookMLParser tokens.
const (
	LookMLParserEOF = antlr.TokenEOF
	LookMLParserT__0 = 1
	LookMLParserT__1 = 2
	LookMLParserT__2 = 3
	LookMLParserT__3 = 4
	LookMLParserT__4 = 5
	LookMLParserT__5 = 6
	LookMLParserT__6 = 7
	LookMLParserT__7 = 8
	LookMLParserT__8 = 9
	LookMLParserT__9 = 10
	LookMLParserT__10 = 11
	LookMLParserT__11 = 12
	LookMLParserT__12 = 13
	LookMLParserT__13 = 14
	LookMLParserT__14 = 15
	LookMLParserT__15 = 16
	LookMLParserT__16 = 17
	LookMLParserT__17 = 18
	LookMLParserID = 19
	LookMLParserSTRING = 20
	LookMLParserWS = 21
)

// LookMLParser rules.
const (
	LookMLParserRULE_model = 0
	LookMLParserRULE_view = 1
	LookMLParserRULE_dimension = 2
	LookMLParserRULE_measure = 3
	LookMLParserRULE_field = 4
	LookMLParserRULE_field_property = 5
	LookMLParserRULE_field_type = 6
)

// IModelContext is an interface to support dynamic dispatch.
type IModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelContext differentiates from other interfaces.
	IsModelContext()
}

type ModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelContext() *ModelContext {
	var p = new(ModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_model
	return p
}

func (*ModelContext) IsModelContext() {}

func NewModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelContext {
	var p = new(ModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_model

	return p
}

func (s *ModelContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelContext) ID() antlr.TerminalNode {
	return s.GetToken(LookMLParserID, 0)
}

func (s *ModelContext) AllView() []IViewContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IViewContext); ok {
			len++
		}
	}

	tst := make([]IViewContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IViewContext); ok {
			tst[i] = t.(IViewContext)
			i++
		}
	}

	return tst
}

func (s *ModelContext) View(i int) IViewContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IViewContext); ok {
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

	return t.(IViewContext)
}

func (s *ModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ModelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterModel(s)
	}
}

func (s *ModelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitModel(s)
	}
}

func (s *ModelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitModel(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Model() (localctx IModelContext) {
	this := p
	_ = this

	localctx = NewModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LookMLParserRULE_model)
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
	{
		p.SetState(14)
		p.Match(LookMLParserT__0)
	}
	{
		p.SetState(15)
		p.Match(LookMLParserID)
	}
	{
		p.SetState(16)
		p.Match(LookMLParserT__1)
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LookMLParserT__3 {
		{
			p.SetState(17)
			p.View()
		}


		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(22)
		p.Match(LookMLParserT__2)
	}



	return localctx
}


// IViewContext is an interface to support dynamic dispatch.
type IViewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsViewContext differentiates from other interfaces.
	IsViewContext()
}

type ViewContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyViewContext() *ViewContext {
	var p = new(ViewContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_view
	return p
}

func (*ViewContext) IsViewContext() {}

func NewViewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ViewContext {
	var p = new(ViewContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_view

	return p
}

func (s *ViewContext) GetParser() antlr.Parser { return s.parser }

func (s *ViewContext) ID() antlr.TerminalNode {
	return s.GetToken(LookMLParserID, 0)
}

func (s *ViewContext) AllDimension() []IDimensionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDimensionContext); ok {
			len++
		}
	}

	tst := make([]IDimensionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDimensionContext); ok {
			tst[i] = t.(IDimensionContext)
			i++
		}
	}

	return tst
}

func (s *ViewContext) Dimension(i int) IDimensionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensionContext); ok {
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

	return t.(IDimensionContext)
}

func (s *ViewContext) AllMeasure() []IMeasureContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMeasureContext); ok {
			len++
		}
	}

	tst := make([]IMeasureContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMeasureContext); ok {
			tst[i] = t.(IMeasureContext)
			i++
		}
	}

	return tst
}

func (s *ViewContext) Measure(i int) IMeasureContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
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

	return t.(IMeasureContext)
}

func (s *ViewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ViewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ViewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterView(s)
	}
}

func (s *ViewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitView(s)
	}
}

func (s *ViewContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitView(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) View() (localctx IViewContext) {
	this := p
	_ = this

	localctx = NewViewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LookMLParserRULE_view)
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
	{
		p.SetState(24)
		p.Match(LookMLParserT__3)
	}
	{
		p.SetState(25)
		p.Match(LookMLParserID)
	}
	{
		p.SetState(26)
		p.Match(LookMLParserT__1)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LookMLParserT__4 {
		{
			p.SetState(27)
			p.Dimension()
		}


		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LookMLParserT__5 {
		{
			p.SetState(32)
			p.Measure()
		}


		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(LookMLParserT__2)
	}



	return localctx
}


// IDimensionContext is an interface to support dynamic dispatch.
type IDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDimensionContext differentiates from other interfaces.
	IsDimensionContext()
}

type DimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensionContext() *DimensionContext {
	var p = new(DimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_dimension
	return p
}

func (*DimensionContext) IsDimensionContext() {}

func NewDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionContext {
	var p = new(DimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_dimension

	return p
}

func (s *DimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionContext) ID() antlr.TerminalNode {
	return s.GetToken(LookMLParserID, 0)
}

func (s *DimensionContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *DimensionContext) Field(i int) IFieldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *DimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterDimension(s)
	}
}

func (s *DimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitDimension(s)
	}
}

func (s *DimensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitDimension(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Dimension() (localctx IDimensionContext) {
	this := p
	_ = this

	localctx = NewDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LookMLParserRULE_dimension)
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
	{
		p.SetState(39)
		p.Match(LookMLParserT__4)
	}
	{
		p.SetState(40)
		p.Match(LookMLParserID)
	}
	{
		p.SetState(41)
		p.Match(LookMLParserT__1)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LookMLParserT__6 {
		{
			p.SetState(42)
			p.Field()
		}


		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(LookMLParserT__2)
	}



	return localctx
}


// IMeasureContext is an interface to support dynamic dispatch.
type IMeasureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasureContext differentiates from other interfaces.
	IsMeasureContext()
}

type MeasureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasureContext() *MeasureContext {
	var p = new(MeasureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_measure
	return p
}

func (*MeasureContext) IsMeasureContext() {}

func NewMeasureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasureContext {
	var p = new(MeasureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_measure

	return p
}

func (s *MeasureContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasureContext) ID() antlr.TerminalNode {
	return s.GetToken(LookMLParserID, 0)
}

func (s *MeasureContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *MeasureContext) Field(i int) IFieldContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *MeasureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MeasureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterMeasure(s)
	}
}

func (s *MeasureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitMeasure(s)
	}
}

func (s *MeasureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitMeasure(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Measure() (localctx IMeasureContext) {
	this := p
	_ = this

	localctx = NewMeasureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LookMLParserRULE_measure)
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
	{
		p.SetState(49)
		p.Match(LookMLParserT__5)
	}
	{
		p.SetState(50)
		p.Match(LookMLParserID)
	}
	{
		p.SetState(51)
		p.Match(LookMLParserT__1)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LookMLParserT__6 {
		{
			p.SetState(52)
			p.Field()
		}


		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(LookMLParserT__2)
	}



	return localctx
}


// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) ID() antlr.TerminalNode {
	return s.GetToken(LookMLParserID, 0)
}

func (s *FieldContext) AllField_property() []IField_propertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_propertyContext); ok {
			len++
		}
	}

	tst := make([]IField_propertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_propertyContext); ok {
			tst[i] = t.(IField_propertyContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Field_property(i int) IField_propertyContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_propertyContext); ok {
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

	return t.(IField_propertyContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LookMLParserRULE_field)
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
	{
		p.SetState(59)
		p.Match(LookMLParserT__6)
	}
	{
		p.SetState(60)
		p.Match(LookMLParserID)
	}
	{
		p.SetState(61)
		p.Match(LookMLParserT__1)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3840) != 0 {
		{
			p.SetState(62)
			p.Field_property()
		}


		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(LookMLParserT__2)
	}



	return localctx
}


// IField_propertyContext is an interface to support dynamic dispatch.
type IField_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_propertyContext differentiates from other interfaces.
	IsField_propertyContext()
}

type Field_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_propertyContext() *Field_propertyContext {
	var p = new(Field_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_field_property
	return p
}

func (*Field_propertyContext) IsField_propertyContext() {}

func NewField_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_propertyContext {
	var p = new(Field_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_field_property

	return p
}

func (s *Field_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_propertyContext) Field_type() IField_typeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_typeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Field_propertyContext) STRING() antlr.TerminalNode {
	return s.GetToken(LookMLParserSTRING, 0)
}

func (s *Field_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Field_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterField_property(s)
	}
}

func (s *Field_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitField_property(s)
	}
}

func (s *Field_propertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitField_property(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Field_property() (localctx IField_propertyContext) {
	this := p
	_ = this

	localctx = NewField_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LookMLParserRULE_field_property)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LookMLParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(LookMLParserT__7)
		}
		{
			p.SetState(70)
			p.Field_type()
		}


	case LookMLParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(LookMLParserT__8)
		}
		{
			p.SetState(72)
			p.Match(LookMLParserSTRING)
		}


	case LookMLParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(LookMLParserT__9)
		}
		{
			p.SetState(74)
			p.Match(LookMLParserSTRING)
		}


	case LookMLParserT__10:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(LookMLParserT__10)
		}
		{
			p.SetState(76)
			p.Match(LookMLParserSTRING)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LookMLParserRULE_field_type
	return p
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LookMLParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LookMLListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (s *Field_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LookMLVisitor:
		return t.VisitField_type(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LookMLParser) Field_type() (localctx IField_typeContext) {
	this := p
	_ = this

	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LookMLParserRULE_field_type)
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
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 520192) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


