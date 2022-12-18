package parser

import (
	"fmt"
	"testing"

	"github.com/datastx/datastx/internal/language/ast"
	"github.com/datastx/datastx/internal/language/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
	 
	let y = 10;
	
	let foobar = 838383;
	`
	program, _, _ := setupTest(t, input)
	countStatements(t, program, 3)
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func TestReturnStatement(t *testing.T) {
	input := `return 5; 
			  return 10; 
			  return 993322;`
	program, _, _ := setupTest(t, input)
	countStatements(t, program, 3)

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifireExpression(t *testing.T) {
	input := `foobar;`
	program, _, _ := setupTest(t, input)
	countStatements(t, program, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}

}

func TestIntegerLiteralExpression(t *testing.T) {
	input := `5;`
	program, _, _ := setupTest(t, input)
	countStatements(t, program, 1)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}

}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input         string
		operator      string
		value         interface{}
		numStatements int
	}{
		{"!5;", "!", 5, 1},
		{"-15;", "-", 15, 1},
		{"!true;", "!", true, 1},
		{"!false;", "!", false, 1},
	}

	for _, tt := range prefixTests {
		program, _, _ := setupTest(t, tt.input)
		countStatements(t, program, tt.numStatements)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp not *ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}

	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input         string
		leftValue     interface{}
		operator      string
		rightValue    interface{}
		numStatements int
	}{
		{"5 + 5;", 5, "+", 5, 1},
		{"5 - 5;", 5, "-", 5, 1},
		{"5 * 5;", 5, "*", 5, 1},
		{"5 / 5;", 5, "/", 5, 1},
		{"5 > 5;", 5, ">", 5, 1},
		{"5 < 5;", 5, "<", 5, 1},
		{"5 == 5;", 5, "==", 5, 1},
		{"5 != 5;", 5, "!=", 5, 1},
		// {"true == true", true, "==", true, 1},
		// {"true != false", true, "!=", false, 1},
		// {"false == false", false, "==", false, 1},
		// {"foobar + barfoo", "foobar", "+", "barfoo", 1},
		// {"foobar - barfoo", "foobar", "-", "barfoo", 1},
		// {"foobar * barfoo", "foobar", "*", "barfoo", 1},
		// {"foobar / barfoo", "foobar", "/", "barfoo", 1},
		// {"foobar > barfoo", "foobar", ">", "barfoo", 1},
	}

	for _, tt := range infixTests {
		program, _, _ := setupTest(t, tt.input)
		countStatements(t, program, tt.numStatements)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}

}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input         string
		expected      string
		numStatements int
	}{
		{"-a * b", "((-a) * b)", 1},
		{"!-a", "(!(-a))", 1},
		{"a + b + c", "((a + b) + c)", 1},
		{"a + b - c", "((a + b) - c)", 1},
		{"a * b * c", "((a * b) * c)", 1},
		{"a * b / c", "((a * b) / c)", 1},
		{"a + b / c", "(a + (b / c))", 1},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)", 1},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)", 2},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))", 1},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))", 1},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))", 1},
		// {"true", "true", 1},
		// {"false", "false", 1},
		// {"3 > 5 == false", "((3 > 5) == false)", 1},
		// {"3 < 5 == true", "((3 < 5) == true)", 1},
		// {"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)", 1},
		// {"(5 + 5) * 2", "((5 + 5) * 2)", 1},
		// {"2 / (5 + 5)", "(2 / (5 + 5))", 1},
		// {"-(5 + 5)", "(-(5 + 5))", 1},
		// {"!(true == true)", "(!(true == true))", 1},
		// {"a + add(b * c) + d", "((a + add((b * c))) + d)", 1},
		// {"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))", "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))", 1},
		// {"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))", 1},
		// {"a * [1, 2, 3, 4][b * c] * d", "((a * ([1, 2, 3, 4][(b * c)])) * d)", 1},
		// {"add(a * b[2], b[1], 2 * [1, 2][1])", "add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))", 1},
	}

	for _, tt := range tests {
		program, _, _ := setupTest(t, tt.input)
		countStatements(t, program, tt.numStatements)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

// Specialized test function for testing the parser

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp not *ast.InfixExpression. got=%T", exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%s", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true

}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	boolExp, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}
	if boolExp.Value != value {
		t.Errorf("boolExp.Value not %t. got=%t", value, boolExp.Value)
		return false
	}
	if boolExp.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("boolExp.TokenLiteral not %t. got=%s", value, boolExp.TokenLiteral())
		return false
	}
	return true
}

// *****Helper functions*******

func setupTest(t *testing.T, input string) (*ast.Program, *lexer.Lexer, *Parser) {
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	checkParserErrors(t, p)
	return program, l, p
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func countStatements(t *testing.T, program *ast.Program, expected int) {
	if len(program.Statements) != expected {
		t.Fatalf("program.Statements does not contain %d statements. got=%d",
			expected, len(program.Statements))
	}
}
