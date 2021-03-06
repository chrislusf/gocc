package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{

	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 97: // ['a','a']
			return 2
		case r == 99: // ['c','c']
			return 3

		}
		return NoState
	},

	// S1
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S2
	func(r rune) int {
		switch {

		}
		return NoState
	},

	// S3
	func(r rune) int {
		switch {

		}
		return NoState
	},
}
