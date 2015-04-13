# Introduction #

[![](https://drone.io/mariusackerman/gocc/status.png)](https://drone.io/mariusackerman/gocc/latest)

Gocc is a compiler kit for Go written in Go.

Gocc generates lexers and parsers or stand-alone DFAs or parsers from a BNF.

Lexers are DFAs, which recognise regular languages. Gocc lexers accept UTF-8 input.

Gocc parsers are PDAs, which recognise LR-1 languages. Optional LR1 conflict handling automatically resolves shift / reduce and reduce / reduce conflicts.

Generating a lexer and parser starts with creating a bnf file. Action expressions embedded in the BNF allows the user to specify semantic actions for syntax productions.

For complex applications the user typically uses an abstract syntax tree (AST) to represent the derivation of the input. The user provides a set of functions to construct the AST, which are called from the action expressions specified in the BNF.

See the [README](https://code.google.com/p/gocc/source/browse/example/bools/README) for an included example.

[User Guide (PDF): Learn You a gocc for Great Good](https://gocc.googlecode.com/git/doc/gocc_user_guide.pdf)
(gocc3 user guide will be published shortly)

# Installation #

  * First download and Install Go From http://golang.org/
  * Setup your GOPATH environment variable.
  * Next in your command line run: go get code.google.com/p/gocc/ (go get will git clone gocc into GOPATH/src/code.google.com/p/gocc and run go install)
  * Alternatively clone the source: https://code.google.com/p/gocc/source/checkout .  Followed by go install code.google.com/p/gocc
  * Finally make sure that the bin folder where the gocc binary is located is in your PATH environment variable.

# Getting Started #

Once installed start by creating your BNF in a package folder.

For example GOPATH/src/foo/bar.bnf:
```
/* Lexical Part */

id : 'a'-'z' {'a'-'z'} ;

!whitespace : ' ' | '\t' | '\n' | '\r' ;

/* Syntax Part */

<< import "foo/ast" >>

Hello:  "hello" id << ast.NewWorld($1) >> ;
```
Next to use gocc, run:
```
cd $GOPATH/src/foo
gocc bar.bnf
```
This will generate a scanner, parser and token package inside GOPATH/src/foo
Following times you might only want to run gocc without the scanner flag, since you might want to start making the scanner your own.  Gocc is after all only a parser generator even if the default scanner is quite useful.

Next create ast.go file at $GOPATH/src/foo/ast with the following contents:
```
package ast

import (
  "foo/token"
)

type Attrib interface {}

type World struct {
  Name string
}

func NewWorld(id Attrib) (*World, error) {
  return &World{string(id.(*token.Token).Lit)}, nil
}

func (this *World) String() string {
  return "hello " + this.Name
}
```

Finally we want to parse a string into the ast, so let us write a test at $GOPATH/src/foo/test/parse\_test.go with the following contents:
```
package test

import (
	"foo/ast"
	"foo/lexer"
	"foo/parser"
	"testing"
)

func TestWorld(t *testing.T) {
	input := []byte(`hello gocc`)
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	st, err := p.Parse(lex)
	if err != nil {
		panic(err)
	}
	w, ok := st.(*ast.World)
	if !ok {
		t.Fatalf("This is not a world")
	}
	if w.Name != `gocc` {
		t.Fatalf("Wrong world %v", w.Name)
	}
}

```

Finally run the test:
```
cd $GOPATH/src/foo/test
go test -v
```

You have now created your first grammar with gocc.
This should now be relatively easy to change into the grammar you actually want to create or an existing LR1 grammar you would like to parse.

# BNF #

The Gocc BNF is specified [here](https://code.google.com/p/gocc/source/browse/spec/gocc2.ebnf)

An example bnf with action expressions can be found [here](https://code.google.com/p/gocc/source/browse/example/example.bnf)

# Action Expressions and AST #

An action expression is specified as "<", "<", goccExpressionList , ">", ">" .
The goccExpressionList is equivalent to a [goExpressionList](http://golang.org/ref/spec#ExpressionList).  This expression list should return an Attrib and an error.
Where Attrib is:
```
type Attrib interface {}
```
Also parsed elements of the corresponding bnf rule can be represented in the expressionList as "$", digit.

Some action expression examples:
```
<< $0, nil >>
<< ast.NewFoo($1) >>
<< ast.NewBar($3, $1) >>
<< ast.TRUE, nil >>
```

Contants, functions, etc. that are returned or called should be programmed by the user in his ast (Abstract Syntax Tree) package.  The ast package requires that you define your own Attrib interface as shown above.
All parameters passed to functions will be of this type.

Some example of functions:
```
func NewFoo(a Attrib) (*Foo, error) { ... }
func NewBar(a, b Attrib) (*Bar, error) { ... }
```

An example of an ast can be found [here](https://code.google.com/p/gocc/source/browse/example/bools/ast/ast.go)