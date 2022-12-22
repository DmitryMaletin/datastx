// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jinja

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JinjaListener is a complete listener for a parse tree produced by JinjaParser.
type JinjaListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterFor_block is called when entering the for_block production.
	EnterFor_block(c *For_blockContext)

	// EnterIf_block is called when entering the if_block production.
	EnterIf_block(c *If_blockContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitFor_block is called when exiting the for_block production.
	ExitFor_block(c *For_blockContext)

	// ExitIf_block is called when exiting the if_block production.
	ExitIf_block(c *If_blockContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)
}
