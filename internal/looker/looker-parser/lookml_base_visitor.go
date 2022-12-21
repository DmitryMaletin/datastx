// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // LookML

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseLookMLVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLookMLVisitor) VisitModel(ctx *ModelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitView(ctx *ViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitDimension(ctx *DimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitMeasure(ctx *MeasureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitField_property(ctx *Field_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLookMLVisitor) VisitField_type(ctx *Field_typeContext) interface{} {
	return v.VisitChildren(ctx)
}
