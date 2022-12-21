// Code generated from /datastx/internal/looker/looker-parser/LookML.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LookML

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LookMLListener is a complete listener for a parse tree produced by LookMLParser.
type LookMLListener interface {
	antlr.ParseTreeListener

	// EnterModel is called when entering the model production.
	EnterModel(c *ModelContext)

	// EnterView is called when entering the view production.
	EnterView(c *ViewContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterMeasure is called when entering the measure production.
	EnterMeasure(c *MeasureContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterField_property is called when entering the field_property production.
	EnterField_property(c *Field_propertyContext)

	// EnterField_type is called when entering the field_type production.
	EnterField_type(c *Field_typeContext)

	// ExitModel is called when exiting the model production.
	ExitModel(c *ModelContext)

	// ExitView is called when exiting the view production.
	ExitView(c *ViewContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitMeasure is called when exiting the measure production.
	ExitMeasure(c *MeasureContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitField_property is called when exiting the field_property production.
	ExitField_property(c *Field_propertyContext)

	// ExitField_type is called when exiting the field_type production.
	ExitField_type(c *Field_typeContext)
}
