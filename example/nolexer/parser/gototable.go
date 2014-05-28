/*
 */
package parser

const numNTSymbols = 3

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // Hello
		2,  // Saying

	},
	gotoRow{ // S1

		-1, // S'
		-1, // Hello
		-1, // Saying

	},
	gotoRow{ // S2

		-1, // S'
		-1, // Hello
		-1, // Saying

	},
	gotoRow{ // S3

		-1, // S'
		-1, // Hello
		-1, // Saying

	},
	gotoRow{ // S4

		-1, // S'
		-1, // Hello
		-1, // Saying

	},
	gotoRow{ // S5

		-1, // S'
		-1, // Hello
		-1, // Saying

	},
}
