// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // LookML

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"
// A complete Visitor for a parse tree produced by LookMLParser.
type LookMLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LookMLParser#model.
	VisitModel(ctx *ModelContext) interface{}

	// Visit a parse tree produced by LookMLParser#view.
	VisitView(ctx *ViewContext) interface{}

	// Visit a parse tree produced by LookMLParser#dimension.
	VisitDimension(ctx *DimensionContext) interface{}

	// Visit a parse tree produced by LookMLParser#measure.
	VisitMeasure(ctx *MeasureContext) interface{}

	// Visit a parse tree produced by LookMLParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by LookMLParser#field_property.
	VisitField_property(ctx *Field_propertyContext) interface{}

	// Visit a parse tree produced by LookMLParser#field_type.
	VisitField_type(ctx *Field_typeContext) interface{}

}