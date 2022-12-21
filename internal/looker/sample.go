package looker

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	help "github.com/antlr/antlr4/runtime/Go/antlr"
	lookml "github.com/datastx/datastx/internal/looker/looker-parser"
)

func test() {
	s := help.NewInputStream("SELECT * FROM `table`")
	// Create the Lexer
	lexer := lookml.NewLookMLLexer(s)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
