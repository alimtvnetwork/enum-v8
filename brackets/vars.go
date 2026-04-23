package brackets

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	rangesMap enumimpl.DynamicMap = map[string]interface{}{
		"Invalid": Invalid,
		"(":       ParenthesisStart,
		")":       ParenthesisEnd,
		"{":       CurlyStart,
		"}":       CurlyEnd,
		"[":       SquareStart,
		"]":       SquareEnd,
	}

	parenthesisWrap = "()"
	curlyWrap       = "{}"
	squareWrap      = "[]"

	bracketsSelfWrapMap = map[Bracket]string{
		Invalid:          constants.EmptyString,
		ParenthesisStart: parenthesisWrap,
		ParenthesisEnd:   parenthesisWrap,
		CurlyStart:       curlyWrap,
		CurlyEnd:         curlyWrap,
		SquareStart:      squareWrap,
		SquareEnd:        squareWrap,
	}

	categorySelfWrapMap = map[Category]string{
		UnknownCategory: constants.EmptyString,
		Parenthesis:     parenthesisWrap,
		Curly:           curlyWrap,
		Square:          squareWrap,
	}

	startMap = map[Bracket]bool{
		ParenthesisStart: true,
		CurlyStart:       true,
		SquareStart:      true,
	}

	endMap = map[Bracket]bool{
		ParenthesisEnd: true,
		CurlyEnd:       true,
		SquareEnd:      true,
	}

	startCategoryMap = map[Category]Bracket{
		Parenthesis: ParenthesisStart,
		Curly:       CurlyStart,
		Square:      SquareStart,
	}

	endCategoryMap = map[Category]Bracket{
		Parenthesis: ParenthesisEnd,
		Curly:       CurlyEnd,
		Square:      SquareEnd,
	}
	//
	// endMap = map[Bracket]bool{
	// 	ParenthesisEnd: true,
	// 	CurlyEnd:       true,
	// 	SquareEnd:      true,
	// }

	bothBracketsMap = map[Bracket]*BothBrackets{
		Invalid: {
			IsInvalid: true,
		},
		ParenthesisStart: {
			Start:    ParenthesisStart,
			End:      ParenthesisEnd,
			Category: Parenthesis,
		},
		ParenthesisEnd: {
			Start:    ParenthesisStart,
			End:      ParenthesisEnd,
			Category: Parenthesis,
		},
		CurlyStart: {
			Start:    CurlyStart,
			End:      CurlyEnd,
			Category: Curly,
		},
		CurlyEnd: {
			Start:    CurlyStart,
			End:      CurlyEnd,
			Category: Curly,
		},
		SquareStart: {
			Start:    SquareStart,
			End:      SquareEnd,
			Category: Square,
		},
		SquareEnd: {
			Start:    SquareStart,
			End:      SquareEnd,
			Category: Square,
		},
	}

	categoriesRange = [...]string{
		UnknownCategory: "UnknownCategory",
		Parenthesis:     "Parenthesis",
		Curly:           "Curly",
		Square:          "Square",
	}

	otherBracketCharsMaps = map[uint8]BracketStatus{
		constants.ParenthesisStartSymbol: {
			IsBracketFound: true,
			Category:       Parenthesis,
			FoundBracket:   ParenthesisStart,
			OtherBracket:   ParenthesisEnd,
		},
		constants.ParenthesisEndSymbol: {
			IsBracketFound: true,
			Category:       Parenthesis,
			FoundBracket:   ParenthesisEnd,
			OtherBracket:   ParenthesisStart,
		},
		constants.CurlyStartSymbol: {
			IsBracketFound: true,
			Category:       Curly,
			FoundBracket:   CurlyStart,
			OtherBracket:   CurlyEnd,
		},
		constants.CurlyEndSymbol: {
			IsBracketFound: true,
			Category:       Curly,
			FoundBracket:   CurlyEnd,
			OtherBracket:   CurlyStart,
		},
		constants.SquareStartSymbol: {
			IsBracketFound: true,
			Category:       Square,
			FoundBracket:   SquareStart,
			OtherBracket:   SquareEnd,
		},
		constants.SquareEndSymbol: {
			IsBracketFound: true,
			Category:       Square,
			FoundBracket:   SquareEnd,
			OtherBracket:   SquareStart,
		},
	}

	otherBracketMaps = map[Bracket]Bracket{
		ParenthesisStart: ParenthesisEnd,
		ParenthesisEnd:   ParenthesisStart,
		CurlyStart:       CurlyEnd,
		CurlyEnd:         CurlyStart,
		SquareStart:      SquareEnd,
		SquareEnd:        SquareStart,
	}

	BasicEnumImpl = rangesMap.BasicByte(
		coredynamic.TypeName(Invalid))
)
