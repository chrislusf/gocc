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
	"code.google.com/p/gocc/io"
	"code.google.com/p/gocc/parser/lr1/items"
	"code.google.com/p/gocc/parser/symbols"
	"path"
	"text/template"
)

func GenGotoTable(outDir string, itemSets *items.ItemSets, sym *symbols.Symbols) {
	tmpl, err := template.New("parser goto table").Parse(gotoTableSrc)
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	tmpl.Execute(wr, getGotoTableData(itemSets, sym))
	io.WriteFile(path.Join(outDir, "parser", "gototable.go"), wr.Bytes())
}

func getGotoTableData(itemSets *items.ItemSets, sym *symbols.Symbols) *gotoTableData {
	data := &gotoTableData{
		NumNTSymbols: sym.NumNTSymbols(),
		Rows:         make([]string, itemSets.Size()),
	}
	for i, set := range itemSets.List() {
		data.Rows[i] = genGotoRow(set, sym)
	}
	return data
}

type gotoTableData struct {
	NumNTSymbols int
	Rows         []string
}

const gotoTableSrc = `
/*
*/
package parser

const numNTSymbols = {{.NumNTSymbols}}
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	{{range $i, $r := .Rows}}gotoRow{ // S{{$i}}
		{{$r}}
	},
	{{end}}
}
`
