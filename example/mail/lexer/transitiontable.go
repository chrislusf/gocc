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
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

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
		case r == 33: // ['!','!']
			return 2
		case r == 35: // ['#','#']
			return 2
		case r == 36: // ['$','$']
			return 2
		case r == 37: // ['%','%']
			return 2
		case r == 38: // ['&','&']
			return 2
		case r == 39: // [''',''']
			return 2
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 2
		case r == 45: // ['-','-']
			return 2
		case r == 46: // ['.','.']
			return 4
		case r == 47: // ['/','/']
			return 2
		case 48 <= r && r <= 57: // ['0','9']
			return 2
		case r == 61: // ['=','=']
			return 2
		case r == 63: // ['?','?']
			return 2
		case r == 64: // ['@','@']
			return 5
		case 65 <= r && r <= 90: // ['A','Z']
			return 2
		case r == 94: // ['^','^']
			return 2
		case r == 95: // ['_','_']
			return 2
		case r == 96: // ['`','`']
			return 2
		case 97 <= r && r <= 122: // ['a','z']
			return 2
		case r == 123: // ['{','{']
			return 2
		case r == 124: // ['|','|']
			return 2
		case r == 125: // ['}','}']
			return 2
		case r == 126: // ['~','~']
			return 2
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 2

		}
		return NoState

	},

	// S3
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 6

		default:
			return 7
		}

	},

	// S4
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 8
		case r == 35: // ['#','#']
			return 8
		case r == 36: // ['$','$']
			return 8
		case r == 37: // ['%','%']
			return 8
		case r == 38: // ['&','&']
			return 8
		case r == 39: // [''',''']
			return 8
		case r == 42: // ['*','*']
			return 8
		case r == 43: // ['+','+']
			return 8
		case r == 45: // ['-','-']
			return 8
		case r == 47: // ['/','/']
			return 8
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 61: // ['=','=']
			return 8
		case r == 63: // ['?','?']
			return 8
		case 65 <= r && r <= 90: // ['A','Z']
			return 8
		case r == 94: // ['^','^']
			return 8
		case r == 95: // ['_','_']
			return 8
		case r == 96: // ['`','`']
			return 8
		case 97 <= r && r <= 122: // ['a','z']
			return 8
		case r == 123: // ['{','{']
			return 8
		case r == 124: // ['|','|']
			return 8
		case r == 125: // ['}','}']
			return 8
		case r == 126: // ['~','~']
			return 8
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 8

		}
		return NoState

	},

	// S5
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 9
		case r == 35: // ['#','#']
			return 9
		case r == 36: // ['$','$']
			return 9
		case r == 37: // ['%','%']
			return 9
		case r == 38: // ['&','&']
			return 9
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 9
		case r == 43: // ['+','+']
			return 9
		case r == 45: // ['-','-']
			return 9
		case r == 47: // ['/','/']
			return 9
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		case r == 61: // ['=','=']
			return 9
		case r == 63: // ['?','?']
			return 9
		case 65 <= r && r <= 90: // ['A','Z']
			return 9
		case r == 94: // ['^','^']
			return 9
		case r == 95: // ['_','_']
			return 9
		case r == 96: // ['`','`']
			return 9
		case 97 <= r && r <= 122: // ['a','z']
			return 9
		case r == 123: // ['{','{']
			return 9
		case r == 124: // ['|','|']
			return 9
		case r == 125: // ['}','}']
			return 9
		case r == 126: // ['~','~']
			return 9
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 9

		}
		return NoState

	},

	// S6
	func(r rune) int {
		switch {

		default:
			return 10
		}

	},

	// S7
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 11
		case r == 92: // ['\','\']
			return 12

		default:
			return 7
		}

	},

	// S8
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 8
		case r == 35: // ['#','#']
			return 8
		case r == 36: // ['$','$']
			return 8
		case r == 37: // ['%','%']
			return 8
		case r == 38: // ['&','&']
			return 8
		case r == 39: // [''',''']
			return 8
		case r == 42: // ['*','*']
			return 8
		case r == 43: // ['+','+']
			return 8
		case r == 45: // ['-','-']
			return 8
		case r == 46: // ['.','.']
			return 4
		case r == 47: // ['/','/']
			return 8
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 61: // ['=','=']
			return 8
		case r == 63: // ['?','?']
			return 8
		case r == 64: // ['@','@']
			return 5
		case 65 <= r && r <= 90: // ['A','Z']
			return 8
		case r == 94: // ['^','^']
			return 8
		case r == 95: // ['_','_']
			return 8
		case r == 96: // ['`','`']
			return 8
		case 97 <= r && r <= 122: // ['a','z']
			return 8
		case r == 123: // ['{','{']
			return 8
		case r == 124: // ['|','|']
			return 8
		case r == 125: // ['}','}']
			return 8
		case r == 126: // ['~','~']
			return 8
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 8

		}
		return NoState

	},

	// S9
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 9
		case r == 35: // ['#','#']
			return 9
		case r == 36: // ['$','$']
			return 9
		case r == 37: // ['%','%']
			return 9
		case r == 38: // ['&','&']
			return 9
		case r == 39: // [''',''']
			return 9
		case r == 42: // ['*','*']
			return 9
		case r == 43: // ['+','+']
			return 9
		case r == 45: // ['-','-']
			return 9
		case r == 46: // ['.','.']
			return 13
		case r == 47: // ['/','/']
			return 9
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		case r == 61: // ['=','=']
			return 9
		case r == 63: // ['?','?']
			return 9
		case 65 <= r && r <= 90: // ['A','Z']
			return 9
		case r == 94: // ['^','^']
			return 9
		case r == 95: // ['_','_']
			return 9
		case r == 96: // ['`','`']
			return 9
		case 97 <= r && r <= 122: // ['a','z']
			return 9
		case r == 123: // ['{','{']
			return 9
		case r == 124: // ['|','|']
			return 9
		case r == 125: // ['}','}']
			return 9
		case r == 126: // ['~','~']
			return 9
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 9

		}
		return NoState

	},

	// S10
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 11
		case r == 92: // ['\','\']
			return 12

		default:
			return 7
		}

	},

	// S11
	func(r rune) int {
		switch {
		case r == 64: // ['@','@']
			return 5

		}
		return NoState

	},

	// S12
	func(r rune) int {
		switch {

		default:
			return 10
		}

	},

	// S13
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 35: // ['#','#']
			return 14
		case r == 36: // ['$','$']
			return 14
		case r == 37: // ['%','%']
			return 14
		case r == 38: // ['&','&']
			return 14
		case r == 39: // [''',''']
			return 14
		case r == 42: // ['*','*']
			return 14
		case r == 43: // ['+','+']
			return 14
		case r == 45: // ['-','-']
			return 14
		case r == 47: // ['/','/']
			return 14
		case 48 <= r && r <= 57: // ['0','9']
			return 14
		case r == 61: // ['=','=']
			return 14
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 94: // ['^','^']
			return 14
		case r == 95: // ['_','_']
			return 14
		case r == 96: // ['`','`']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		case r == 123: // ['{','{']
			return 14
		case r == 124: // ['|','|']
			return 14
		case r == 125: // ['}','}']
			return 14
		case r == 126: // ['~','~']
			return 14
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 14

		}
		return NoState

	},

	// S14
	func(r rune) int {
		switch {
		case r == 33: // ['!','!']
			return 14
		case r == 35: // ['#','#']
			return 14
		case r == 36: // ['$','$']
			return 14
		case r == 37: // ['%','%']
			return 14
		case r == 38: // ['&','&']
			return 14
		case r == 39: // [''',''']
			return 14
		case r == 42: // ['*','*']
			return 14
		case r == 43: // ['+','+']
			return 14
		case r == 45: // ['-','-']
			return 14
		case r == 46: // ['.','.']
			return 13
		case r == 47: // ['/','/']
			return 14
		case 48 <= r && r <= 57: // ['0','9']
			return 14
		case r == 61: // ['=','=']
			return 14
		case r == 63: // ['?','?']
			return 14
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 94: // ['^','^']
			return 14
		case r == 95: // ['_','_']
			return 14
		case r == 96: // ['`','`']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		case r == 123: // ['{','{']
			return 14
		case r == 124: // ['|','|']
			return 14
		case r == 125: // ['}','}']
			return 14
		case r == 126: // ['~','~']
			return 14
		case 256 <= r && r <= 1114111: // [\u0100,\U0010ffff]
			return 14

		}
		return NoState

	},
}