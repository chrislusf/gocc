package astx

import (
	"code.google.com/p/gocc/example/errorrecovery/ast"
	"code.google.com/p/gocc/example/errorrecovery/errors"
	"code.google.com/p/gocc/example/errorrecovery/lexer"
	"code.google.com/p/gocc/example/errorrecovery/parser"
	"fmt"
	"testing"
)

func TestFail(t *testing.T) {
	sml, err := test([]byte("a b ; d e f"))
	if err != nil {
		t.Fail()
	}
	fmt.Print("output: [\n")
	for _, s := range sml {
		switch sym := s.(type) {
		case *errors.Error:
			fmt.Printf("%s\n", sym)
		default:
			fmt.Printf("\t%v\n", sym)
		}
	}
	fmt.Println("]")
}

func test(src []byte) (astree ast.StmtList, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer([]byte(src))
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.StmtList)
	}
	return
}
