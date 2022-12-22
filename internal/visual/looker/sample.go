package looker

import (
	"fmt"

	antlr "github.com/antlr/antlr4/runtime/Go/antlr"
	lookml "github.com/datastx/datastx/internal/visual/looker/looker-parser"
)

var lookmlReport = `
include: "/simplified_views/*"
# include: "/demo_vitals/*"

view: realtime_observations {
  sql_table_name: healthcare_demo_live.realtime_observation ;;
  extends: [observation_vitals]
}
`

func test() {
	s := antlr.NewInputStream(lookmlReport)
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
