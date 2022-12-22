// Code generated from /datastx/internal/jinja2/jinja-parser/Jinja.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jinja

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJinjaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJinjaVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitFor_block(ctx *For_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitIf_block(ctx *If_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJinjaVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}
