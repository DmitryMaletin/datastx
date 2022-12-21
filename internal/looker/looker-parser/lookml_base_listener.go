// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // LookML

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseLookMLListener is a complete listener for a parse tree produced by LookMLParser.
type BaseLookMLListener struct{}

var _ LookMLListener = &BaseLookMLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLookMLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLookMLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLookMLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLookMLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModel is called when production model is entered.
func (s *BaseLookMLListener) EnterModel(ctx *ModelContext) {}

// ExitModel is called when production model is exited.
func (s *BaseLookMLListener) ExitModel(ctx *ModelContext) {}

// EnterView is called when production view is entered.
func (s *BaseLookMLListener) EnterView(ctx *ViewContext) {}

// ExitView is called when production view is exited.
func (s *BaseLookMLListener) ExitView(ctx *ViewContext) {}

// EnterDimension is called when production dimension is entered.
func (s *BaseLookMLListener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *BaseLookMLListener) ExitDimension(ctx *DimensionContext) {}

// EnterMeasure is called when production measure is entered.
func (s *BaseLookMLListener) EnterMeasure(ctx *MeasureContext) {}

// ExitMeasure is called when production measure is exited.
func (s *BaseLookMLListener) ExitMeasure(ctx *MeasureContext) {}

// EnterField is called when production field is entered.
func (s *BaseLookMLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLookMLListener) ExitField(ctx *FieldContext) {}

// EnterField_property is called when production field_property is entered.
func (s *BaseLookMLListener) EnterField_property(ctx *Field_propertyContext) {}

// ExitField_property is called when production field_property is exited.
func (s *BaseLookMLListener) ExitField_property(ctx *Field_propertyContext) {}

// EnterField_type is called when production field_type is entered.
func (s *BaseLookMLListener) EnterField_type(ctx *Field_typeContext) {}

// ExitField_type is called when production field_type is exited.
func (s *BaseLookMLListener) ExitField_type(ctx *Field_typeContext) {}
