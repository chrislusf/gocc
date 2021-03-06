//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package golang

import (
	"bytes"
	"code.google.com/p/gocc/ast"
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/parser/symbols"
	"code.google.com/p/gocc/token"
	"fmt"
	"path"
	"text/template"
)

func GenProductionsTable(pkg, outDir, header string, prods ast.SyntaxProdList, symbols *symbols.Symbols,
	itemsets *items.ItemSets, tokMap *token.TokenMap) {

	fname := path.Join(outDir, "parser", "productionstable.go")
	tmpl, err := template.New("parser productions table").Parse(prodsTabSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	tmpl.Execute(wr, getProdsTab(header, prods, symbols, itemsets, tokMap))
	io.WriteFile(fname, wr.Bytes())
}

func getProdsTab(header string, prods ast.SyntaxProdList, symbols *symbols.Symbols,
	itemsets *items.ItemSets, tokMap *token.TokenMap) *prodsTabData {

	data := &prodsTabData{
		Header:  header,
		ProdTab: make([]prodTabEntry, len(prods)),
	}
	for i, prod := range prods {
		data.ProdTab[i].String = fmt.Sprintf("`%s`", prod.String())
		data.ProdTab[i].Id = prod.Id
		data.ProdTab[i].NTType = symbols.NTType(prod.Id)
		if prod.Body.Symbols[0].SymbolString() == "empty" {
			data.ProdTab[i].NumSymbols = 0
			data.ProdTab[i].ReduceFunc = fmt.Sprintf("return nil, nil")
		} else {
			data.ProdTab[i].NumSymbols = len(prod.Body.Symbols)
			switch {
			case prod.Body.SDT != "":
				data.ProdTab[i].ReduceFunc = fmt.Sprintf("return %s", prod.Body.SDT)
			default:
				data.ProdTab[i].ReduceFunc = fmt.Sprintf("return X[0], nil")
			}
		}
	}

	return data
}

type prodsTabData struct {
	Header  string
	ProdTab []prodTabEntry
}

type prodTabEntry struct {
	String     string
	Id         string
	NTType     int
	NumSymbols int
	ReduceFunc string
}

const prodsTabSrc = `
package parser

{{.Header}}

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	{{range $i, $entry := .ProdTab}}ProdTabEntry{
		String: {{$entry.String}},
		Id: "{{$entry.Id}}",
		NTType: {{$entry.NTType}},
		Index: {{$i}},
		NumSymbols: {{$entry.NumSymbols}},
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			{{$entry.ReduceFunc}}
		},
	},
	{{end}}
}
`
