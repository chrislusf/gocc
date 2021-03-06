/*
 */
package parser

const numNTSymbols = 4

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
	gotoRow{ // S0

		-1, // S'
		1,  // RR
		3,  // B
		2,  // A

	},
	gotoRow{ // S1

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
	gotoRow{ // S2

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
	gotoRow{ // S3

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
	gotoRow{ // S4

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
	gotoRow{ // S5

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
	gotoRow{ // S6

		-1, // S'
		-1, // RR
		-1, // B
		-1, // A

	},
}
