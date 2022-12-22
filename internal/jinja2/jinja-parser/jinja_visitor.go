// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jinja

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by JinjaParser.
type JinjaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JinjaParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by JinjaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by JinjaParser#for_block.
	VisitFor_block(ctx *For_blockContext) interface{}

	// Visit a parse tree produced by JinjaParser#if_block.
	VisitIf_block(ctx *If_blockContext) interface{}

	// Visit a parse tree produced by JinjaParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by JinjaParser#text.
	VisitText(ctx *TextContext) interface{}

}