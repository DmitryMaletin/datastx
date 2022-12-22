// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jinja

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJinjaListener is a complete listener for a parse tree produced by JinjaParser.
type BaseJinjaListener struct{}

var _ JinjaListener = &BaseJinjaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJinjaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJinjaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJinjaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJinjaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseJinjaListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseJinjaListener) ExitTemplate(ctx *TemplateContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJinjaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJinjaListener) ExitBlock(ctx *BlockContext) {}

// EnterFor_block is called when production for_block is entered.
func (s *BaseJinjaListener) EnterFor_block(ctx *For_blockContext) {}

// ExitFor_block is called when production for_block is exited.
func (s *BaseJinjaListener) ExitFor_block(ctx *For_blockContext) {}

// EnterIf_block is called when production if_block is entered.
func (s *BaseJinjaListener) EnterIf_block(ctx *If_blockContext) {}

// ExitIf_block is called when production if_block is exited.
func (s *BaseJinjaListener) ExitIf_block(ctx *If_blockContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseJinjaListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseJinjaListener) ExitVariable(ctx *VariableContext) {}

// EnterText is called when production text is entered.
func (s *BaseJinjaListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseJinjaListener) ExitText(ctx *TextContext) {}
