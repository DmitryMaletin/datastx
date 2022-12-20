package jinja2

// import (
// 	"fmt"

// 	"github.com/antlr/antlr4/runtime/Go/antlr"
// )

// func main() {
// 	// Create a new input stream from the template string
// 	input := antlr.NewInputStream(templateString)

// 	// Create a new Jinja2 lexer and parser
// 	lexer := jinja.NewJinjaLexer(input)
// 	parser := jinja.NewJinjaParser(antlr.NewCommonTokenStream(lexer, 0))

// 	// Parse the template and get the root node of the AST
// 	tree := parser.Template()

// 	// Print the AST as a string
// 	fmt.Println(tree.ToStringTree(parser.RuleNames, parser))
// }
