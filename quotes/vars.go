package quotes

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
)

var (
	rangesMap enumimpl.DynamicMap = map[string]interface{}{
		"Invalid":                         Invalid,
		constants.DoubleQuoteStringSymbol: Double,
		constants.SingleQuotation:         Single,
		constants.Backtick:                Backtick,
	}

	selfWrap = map[Quote]string{
		Invalid:  constants.EmptyString,
		Double:   constants.DoubleDoubleQuoteStringSymbol,
		Single:   constants.SingleQuotationStartEnd,
		Backtick: constants.CodeQuotationStartEnd,
	}

	otherQuoteCharsMaps = map[byte]QuoteStatus{
		constants.SingleQuoteSymbol: {
			IsQuoteFound: true,
			Found:        Single,
		},
		constants.DoubleQuoteSymbol: {
			IsQuoteFound: true,
			Found:        Double,
		},
		constants.BacktickSymbol: {
			IsQuoteFound: true,
			Found:        Backtick,
		},
	}

	otherQuoteMaps = map[Quote]Quote{
		Single:   Single,
		Double:   Double,
		Backtick: Backtick,
	}

	BasicEnumImpl = rangesMap.BasicByte(
		coredynamic.TypeName(Invalid))
)
